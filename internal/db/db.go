package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(addr string) (*pgxpool.Pool, error) {
    connConfig, err := pgxpool.ParseConfig(addr)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    poolConn, err := pgxpool.NewWithConfig(ctx, connConfig)
    if err != nil {
        return nil, err
    }
    return poolConn, nil
}
