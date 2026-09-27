package client

import (
	"context"
	"fmt"
	"forward/internal/config"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rds *redis.Client
}

func newRedisClient(ctx context.Context) (*RedisClient, error) {
	cfg := config.GetRedisConfig()

	dsn := fmt.Sprintf(
		"redis://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}

	rds := redis.NewClient(opts)

	if err := rds.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &RedisClient{rds: rds}, nil
}

func (c *RedisClient) Close() {
	c.rds.Close()
}
