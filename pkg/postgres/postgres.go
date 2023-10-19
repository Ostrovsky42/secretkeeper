package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func OpenDB(dsn string) (*pgxpool.Pool, error) {
	poolConfig, err := newPoolConfig(dsn)
	if err != nil {
		return nil, err
	}
	conn, err := newConnection(poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func newPoolConfig(dsn string) (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}

func newConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
