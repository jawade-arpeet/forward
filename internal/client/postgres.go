package client

import (
	"context"
	"errors"
	"fmt"
	"forward/internal/config"
	"forward/internal/errs"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var errPgOperationFailed = errors.New("postgres operation failed")

type PostgresClient struct {
	pool *pgxpool.Pool
}

func newPostgresClient(ctx context.Context) (*PostgresClient, error) {
	cfg := config.GetPostgresConfig()

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
		return nil, fmt.Errorf("failed to create postgres client: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresClient{pool: pool}, nil
}

func (c *PostgresClient) Close() {
	c.pool.Close()
}

func (c *PostgresClient) wrapErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%w: %w", errs.ErrPgNoRows, err)
	}

	pgErr, ok := errors.AsType[*pgconn.PgError](err)
	if !ok {
		return fmt.Errorf("%w: %w", errPgOperationFailed, err)
	}

	var sentinel error
	switch pgErr.Code {
	case pgerrcode.UniqueViolation:
		sentinel = errs.ErrPgUniqueViolation
	case pgerrcode.ForeignKeyViolation:
		sentinel = errs.ErrPgForeignKeyViolation
	case pgerrcode.NotNullViolation:
		sentinel = errs.ErrPgNotNullViolation
	case pgerrcode.CheckViolation:
		sentinel = errs.ErrPgCheckViolation
	default:
		return fmt.Errorf("%w: %w", errPgOperationFailed, err)
	}

	return fmt.Errorf("%w: %w", sentinel, err)
}

func (c *PostgresClient) QueryOne[T any](
	ctx context.Context,
	query string,
	args pgx.NamedArgs,
) (*T, error) {
	row, err := c.pool.Query(ctx, query, args)
	if err != nil {
		return nil, c.wrapErr(err)
	}

	defer row.Close()

	var result T
	if err := pgxscan.ScanOne(&result, row); err != nil {
		return nil, c.wrapErr(err)
	}

	return &result, nil
}
