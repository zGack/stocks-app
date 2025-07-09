package stocksfetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/app/stock"
	"github.com/zgack/stocks/internal/env"
)

type Service struct {
	DB     *pgxpool.Pool
	store  StocksFetcherRepo
	apiURL string
	once   sync.Once
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		DB:    db,
		store: NewCockroachStockRepo(db),
	}
}

func (s *Service) InsertStocks(ctx context.Context, stocks []stock.Stock) error {
	var rows [][]any
	for _, stock := range stocks {
        stock.StockScore = getStockScore(stock)

		row := []any{
			stock.Ticker, stock.TargetFrom, stock.TargetTo,
			stock.Company, stock.Action, stock.Brokerage,
			stock.RatingFrom, stock.RatingTo, stock.Time,
            stock.StockScore,
		}
		rows = append(rows, row)
	}

	err := s.store.BulkInsert(ctx, rows)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FetchInitialStocks(ctx context.Context) (string, error) {
    shouldPopulate := env.GetBool("STOCKS_API_POPULATE", false)
    if !shouldPopulate {
        return "", nil
    }

	stocks, nextPage, err := s.FetchStockPages(ctx, 2, "") // initially fetch 2 pages
	if err != nil {
		return "", err
	}

	err = s.InsertStocks(ctx, stocks)
	if err != nil {
		return "", err
	}

	return nextPage, err
}

func (s *Service) FetchRemainingStocks(ctx context.Context, cancel context.CancelFunc, lastPage string) {
    defer cancel() // ensure the context is cancelled when done

    shouldPopulate := env.GetBool("STOCKS_API_POPULATE", false)
    if !shouldPopulate {
        return
    }

    s.once.Do(func() {
		fmt.Println("Starting background job to fetch remaining stocks from the API")
		for {
			// get stocks by batches of 10 e.g: 100 stocks records
			asyncStocks, nextPage, err := s.FetchStockPages(ctx, 10, lastPage)
            // TODO: add score to the stocks

			if err != nil {
				fmt.Printf("failed to fetch stocks from API: %v", err)
				return
			}

			// flush the stocks batch in the database
			err = s.InsertStocks(ctx, asyncStocks)
			if err != nil {
				fmt.Printf("failed to perform bulk stocks insertion: %v", err)
				return
			}


			if nextPage == "" {
				break
			}
			lastPage = nextPage
		}
		fmt.Println("Finished fetching all stocks from the API and inserting them into the database")
    })
}

func (s *Service) FetchStockPage(ctx context.Context, nextPage string) ([]stock.Stock, string, error) {
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

	var stockItems StocksAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&stockItems)
	if err != nil {
		return nil, "", err
	}

	return stockItems.Items, stockItems.NextPage, nil
}

func (s *Service) FetchStockPages(ctx context.Context, numOfPages int, nextPage string) ([]stock.Stock, string, error) {
	var allStocks []stock.Stock

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
