package stock

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockRepo interface {
	Create(context.Context, *Stock) error
	GetAll(context.Context, StockQueryFilters) ([]*Stock, error)
	GetAllWithQuery(context.Context, StockQueryFilters) ([]*Stock, error)
}

type cockroachStockRepo struct {
	db *pgxpool.Pool
}

func NewCockroachStockRepo(db *pgxpool.Pool) StockRepo {
	return &cockroachStockRepo{db: db}
}

func (r *cockroachStockRepo) Create(ctx context.Context, stock *Stock) error {
	query := `
        INSERT INTO stock (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	return r.db.QueryRow(ctx,
		query,
		stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company, stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo,
	).Scan(&stock.ID)
}

func (r *cockroachStockRepo) GetAll(ctx context.Context, filters StockQueryFilters) ([]*Stock, error) {
	query := fmt.Sprintf("SELECT * FROM stock ORDER BY %s %s LIMIT $1 OFFSET $2", pgx.Identifier{filters.SortBy}.Sanitize(), filters.SortDir)

	rows, err := r.db.Query(ctx, query, filters.Limit, filters.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stocks []*Stock
	for rows.Next() {
		var stock Stock
		if err := rows.Scan(&stock.ID, &stock.Ticker, &stock.TargetFrom, &stock.TargetTo,
			&stock.Company, &stock.Action, &stock.Brokerage, &stock.RatingFrom, &stock.RatingTo, &stock.Time, &stock.StockScore); err != nil {
			return nil, err
		}
		stocks = append(stocks, &stock)
	}
	return stocks, nil
}

func (r *cockroachStockRepo) GetAllWithQuery(ctx context.Context, filters StockQueryFilters) ([]*Stock, error) {
	searchTerm := fmt.Sprintf("%%%s%%", strings.ToLower(filters.SearchTerm))
	query := fmt.Sprintf("SELECT * FROM stock WHERE LOWER(company) LIKE $1 ORDER BY %s %s LIMIT $2 OFFSET $3", pgx.Identifier{filters.SortBy}.Sanitize(), filters.SortDir)
	fmt.Printf("Executing query: %s\n", query)
	fmt.Printf("Search term: %s\n", searchTerm)
	fmt.Printf("Filters: %+v\n", filters)

	rows, err := r.db.Query(ctx, query, searchTerm, filters.Limit, filters.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stocks []*Stock
	for rows.Next() {
		var stock Stock
		if err := rows.Scan(&stock.ID, &stock.Ticker, &stock.TargetFrom, &stock.TargetTo,
			&stock.Company, &stock.Action, &stock.Brokerage, &stock.RatingFrom, &stock.RatingTo, &stock.Time, &stock.StockScore); err != nil {
			return nil, err
		}
		stocks = append(stocks, &stock)
	}
	return stocks, nil
}
