package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
    // repository for stocks
    Stocks interface {
        Create(context.Context, *Stock) error
    }
}

//
func NewCockroachStorage(conn *pgx.Conn) Storage {
    return Storage{
        Stocks: &StocksStore{conn},
    }
}
