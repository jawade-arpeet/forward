package client

import "context"

type Client struct {
	Postgres *PostgresClient
	Redis    *RedisClient
}

func New(ctx context.Context) (*Client, error) {
	pg, err := newPostgresClient(ctx)
	if err != nil {
		return nil, err
	}

	rds, err := newRedisClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		Postgres: pg,
		Redis:    rds,
	}, nil
}

func (c *Client) Close() {
	c.Postgres.Close()
	c.Redis.Close()
}
