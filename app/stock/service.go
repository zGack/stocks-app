package stock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/config"
	"github.com/zgack/stocks/internal/env"
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

func (s *Service) InsertStocks(ctx context.Context, stocks []Stock) error {
	var rows [][]any
	for _, stock := range stocks {
		row := []any{
			stock.Ticker, stock.TargetFrom, stock.TargetTo,
			stock.Company, stock.Action, stock.Brokerage,
			stock.RatingFrom, stock.RatingTo, stock.Time,
		}
		rows = append(rows, row)
	}

	err := s.store.BulkInsert(ctx, rows)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FetchStockPage(ctx context.Context, nextPage string) ([]Stock, string, error) {
	apiURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

	if nextPage != "" {
		apiURL += fmt.Sprintf("?next_page=%s", nextPage)
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+env.GetString("STOCKS_API_AUTH_TOKEN", ""))

	if err != nil {
		return nil, "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var stockItems StockAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&stockItems)
	if err != nil {
		return nil, "", err
	}

	return stockItems.Items, stockItems.NextPage, nil
}

func (s *Service) FetchStockPages(ctx context.Context, numOfPages int, nextPage string) ([]Stock, string, error) {
	var allStocks []Stock

	for range numOfPages {

		apiStocks, newNextPage, err := s.FetchStockPage(ctx, nextPage)
		if err != nil {
			return nil, "", err
		}

		allStocks = append(allStocks, apiStocks...)

		nextPage = newNextPage
		if nextPage == "" {
			break
		}
	}

	return allStocks, nextPage, nil
}
