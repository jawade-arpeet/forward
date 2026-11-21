package client

import (
	"context"
	"fmt"
	"forward/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClt struct {
	pool *pgxpool.Pool
}

func newPostgresClt(ctx context.Context) (*PostgresClt, error) {
	cfg := config.GetPostgresCfg()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresClt{pool: pool}, nil
}

func (c *PostgresClt) Close() {
	c.pool.Close()
}
