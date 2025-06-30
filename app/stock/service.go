package stock

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/config"
)

type Service struct {
	DB    *pgxpool.Pool
	store StockRepo
}

func NewService(c *config.Conf, db *pgxpool.Pool) *Service {
	return &Service{
		DB:    db,
		store: NewCockroachStockRepo(db),
	}
}

func (s *Service) GetAllStocks(ctx context.Context, filters StockQueryFilters) ([]*Stock, error) {
    var stocks []*Stock

    if filters.SearchTerm != "" && filters.SearchBy != "" {
	    stocks, err := s.store.GetAllWithQuery(ctx, filters)
        if err != nil {
            return nil, err
        }
	    return stocks, nil
    }

    stocks, err := s.store.GetAll(ctx, filters)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

