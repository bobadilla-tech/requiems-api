package counter

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

// Counter business logic interface.
type Service interface {
	Increment(ctx context.Context, namespace string) (int64, error)
	Get(ctx context.Context, namespace string) (int64, error)
	IncrementBatch(ctx context.Context, namespaces []string) []BatchCounterItem
}

type service struct {
	rdb  *redis.Client
	repo Repository
}

func NewService(rdb *redis.Client, repo Repository) Service {
	return &service{rdb: rdb, repo: repo}
}

// Increment atomically increments the Redis counter and returns the new value.
// Marks the counter as dirty so it will be synced to PostgreSQL on the next cycle.
func (s *service) Increment(ctx context.Context, namespace string) (int64, error) {
	// Fast path: increment in Redis when the key already exists.
	val, missing, err := incrementIfPresent(ctx, s.rdb, namespace)
	if err != nil {
		return 0, err
	}

	if !missing {
		return val, nil
	}

	// Cold-cache bootstrap path: hydrate baseline from PostgreSQL before first increment.
	base, err := s.repo.Get(ctx, namespace)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			base = 0
		} else {
			return 0, err
		}
	}

	return incrementWithBootstrap(ctx, s.rdb, namespace, base)
}

// Get returns the counter value from Redis, falling back to PostgreSQL when the
// key is absent. If a PostgreSQL value is found it is hydrated back into Redis.
func (s *service) Get(ctx context.Context, namespace string) (int64, error) {
	val, err := s.rdb.Get(ctx, redisKey(namespace)).Int64()
	if err == nil {
		return val, nil
	}

	if !errors.Is(err, redis.Nil) {
		return 0, err
	}

	// Fallback to PostgreSQL
	total, err := s.repo.Get(ctx, namespace)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	// Hydrate Redis without clobbering a concurrent increment.
	if err := s.rdb.SetArgs(ctx, redisKey(namespace), total, redis.SetArgs{Mode: "NX"}).Err(); err != nil {
		log.Printf("counter redis hydrate error for %q: %v", namespace, err)
	}

	return total, nil
}

// IncrementBatch increments multiple counters atomically.
// Warm Redis keys are incremented via pipeline (one round-trip);
// cold keys fall back to sequential Increment for PostgreSQL bootstrap.
func (s *service) IncrementBatch(ctx context.Context, namespaces []string) []BatchCounterItem {
	results := make([]BatchCounterItem, len(namespaces))

	type validNS struct {
		idx int
		ns  string
	}
	var valid []validNS
	for i, ns := range namespaces {
		if !namespaceRe.MatchString(ns) {
			results[i] = BatchCounterItem{Namespace: ns, Error: namespaceValidationErrorMessage}
			continue
		}
		valid = append(valid, validNS{idx: i, ns: ns})
	}

	if len(valid) == 0 {
		return results
	}

	// Single MGET to detect which keys are warm in Redis.
	keys := make([]string, len(valid))
	for i, v := range valid {
		keys[i] = redisKey(v.ns)
	}
	vals, mgetErr := s.rdb.MGet(ctx, keys...).Result()

	var warmNS, coldNS []validNS
	if mgetErr == nil {
		for i, v := range valid {
			if vals[i] != nil {
				warmNS = append(warmNS, v)
			} else {
				coldNS = append(coldNS, v)
			}
		}
	} else {
		coldNS = valid
	}

	// Warm path: pipeline INCR in one round-trip.
	if len(warmNS) > 0 {
		pipe := s.rdb.Pipeline()
		cmds := make([]*redis.IntCmd, len(warmNS))
		for i, v := range warmNS {
			cmds[i] = pipe.Incr(ctx, redisKey(v.ns))
		}
		pipe.Exec(ctx) //nolint:errcheck // per-command results captured below
		for i, v := range warmNS {
			val, err := cmds[i].Result()
			if err != nil {
				results[v.idx] = BatchCounterItem{Namespace: v.ns, Error: "failed to increment counter"}
			} else {
				results[v.idx] = BatchCounterItem{Namespace: v.ns, Value: val}
			}
		}
	}

	// Cold path: sequential Increment handles PostgreSQL bootstrap.
	for _, v := range coldNS {
		val, err := s.Increment(ctx, v.ns)
		if err != nil {
			results[v.idx] = BatchCounterItem{Namespace: v.ns, Error: "failed to increment counter"}
		} else {
			results[v.idx] = BatchCounterItem{Namespace: v.ns, Value: val}
		}
	}

	return results
}
