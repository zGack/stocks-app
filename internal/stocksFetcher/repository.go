package stocksfetcher

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StocksFetcherRepo interface {
	BulkInsert(context.Context, [][]any) error
}

type cockroachStockRepo struct {
	db *pgxpool.Pool
}

func NewCockroachStockRepo(db *pgxpool.Pool) StocksFetcherRepo {
    return &cockroachStockRepo{
        db: db,
    }
}

func (r *cockroachStockRepo) BulkInsert(ctx context.Context, rows [][]any) error {
	copyCount, err := r.db.CopyFrom(
		ctx,
		pgx.Identifier{"stock"},
		[]string{"ticker", "target_from", "target_to", "company", "action", "brokerage", "rating_from", "rating_to", "time", "stock_score"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}

    // TODO: replace with logger
	fmt.Printf("Inserted %d rows into stock table\n", copyCount)

	return nil
}
