package client

import "context"

type Client struct {
	Postgres *PostgresClt
	Redis    *RedisClt
}

func New(ctx context.Context) (*Client, error) {
	pg, err := newPostgresClt(ctx)
	if err != nil {
		return nil, err
	}

	rds, err := newRedisClt(ctx)
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
