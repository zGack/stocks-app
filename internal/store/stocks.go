package store

import (
	"context"
	"fmt"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
)

// Model for the stocks table
type Stock struct {
	ID         int64  `json:"id"`
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

// stocks cockroach repository
type StocksStore struct {
	conn *pgx.Conn
}

func (s *StocksStore) Create(ctx context.Context, stock *Stock) error {
	query := `
	   INSERT INTO stock (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to)
	   VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
    // TODO: parse time correctly

	return crdbpgx.ExecuteTx(ctx, s.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		err := tx.QueryRow(ctx,
			query,
			stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company, stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo,
		).Scan(&stock.ID)
        fmt.Println("Stock created with ID:", err)
		return err
	})

	// query := `
	//    INSERT INTO stocks (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time)
	//    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, ticker`
	//
	//    err := s.db.QueryRowContext(ctx, query,
	//    ).Scan(&stock.ID)
	//
	//    if err != nil {
	//        return nil
	//    }
	//
	// return nil
}
