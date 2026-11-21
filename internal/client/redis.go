package client

import (
	"context"
	"fmt"
	"forward/internal/config"

	"github.com/redis/go-redis/v9"
)

type RedisClt struct {
	rds *redis.Client
}

func newRedisClt(ctx context.Context) (*RedisClt, error) {
	cfg := config.GetRedisCfg()

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

	return &RedisClt{rds: rds}, nil
}

func (c *RedisClt) Close() {
	c.rds.Close()
}
