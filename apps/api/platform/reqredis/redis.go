package reqredis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// poolSize is fixed rather than measured, since there is no production
// traffic yet to size against (a single Go replica). Revisit once real
// traffic exists.
const poolSize = 20

func Connect(ctx context.Context, url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)

	if err != nil {
		return nil, fmt.Errorf("parse REDIS_URL: %w", err)
	}

	opts.PoolSize = poolSize

	// Required for the rate limiter/quota middleware's per-call
	// context.WithTimeout to actually bound a slow-but-alive Redis: without
	// this, go-redis ignores a command's context deadline entirely and
	// blocks on its own (much longer) connection-level ReadTimeout instead,
	// silently defeating the "tens of milliseconds" timeout those callers
	// rely on to fail open/fall through instead of hanging.
	opts.ContextTimeoutEnabled = true

	client := redis.NewClient(opts)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		client.Close()

		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return client, nil
}
