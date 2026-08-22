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

	client := redis.NewClient(opts)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		client.Close()

		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return client, nil
}
