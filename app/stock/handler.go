package stock

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/config"
	contextkeys "github.com/zgack/stocks/pkg/contextKeys"
	"github.com/zgack/stocks/pkg/router"
	"github.com/zgack/stocks/pkg/validator"
)

func GetAllStocksHandler(db *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
    stocksBuffer := make(chan []Stock) // buffer to hold stocks

	service, err := getServiceFromContext(db, r)
	if err != nil {
		router.RespondWithError(w, http.StatusInternalServerError, "service initialization failed", err)
		return
	}

	filters := StockQueryFilters{
		Limit:   20,     // default limit
		Offset:  0,      // default offset
		SortBy:  "time", // default sort by
		SortDir: "desc", // default sort dir
	}

	filters, err = filters.Parse(r)
	if err != nil {
		router.RespondWithError(w, http.StatusBadRequest, "invalid query parameters", err)
		return
	}

	fmt.Printf("Filters: %+v\n", filters)

    dbStocks, err := service.GetAllStocks(r.Context(), filters)
	if err != nil {
		router.RespondWithError(w, http.StatusInternalServerError, "failed to fecth all stocks from db", err)
		return
	}

    fmt.Printf("Fetched %d stocks from the database\n", len(dbStocks))

    // FIXME: cover case when background job is still running, maybe i can use wait group to wait for the background job to finish
    // FIXME: not strong condition, this could be triggered when offset is greater than the number of stocks in the database
    bStocks, ok := <-stocksBuffer
    if len(dbStocks) == 0 && ok {
        // chan still open, that means the background job is still running, so retry the request
        // GetAllStocksHandler(db, w, r)
    }

	if len(dbStocks) == filters.Limit {
		router.RespondWithJSON(w, http.StatusOK, dbStocks)
		return
	}

	// TODO: validate the filters struct
	// if err := Validate.Struct(filters); err != nil {
	//     router.RespondWithError(w, http.StatusBadRequest, "invalid query parameters", err)
	//     return
	// }

	// try to fetch stocks from the database first
	//    for this we need to limit and offset

	// if the database is empty, fetch from the API

	// fetch 2 first pages from the API
	// store the results in the database
	// create a background job to fetch the remaining stocks
	//    fetch 5 mores pages
	//    flush the stocks in the database
	//    if the next_page is empty, stop fetching
	//    otherwise, continue fetching

	// I suppose this only occurs once, when the database is empty
	apiStocks, lastPage, err := service.FetchStockPages(r.Context(), 2, "")
	if err != nil {
		router.RespondWithError(w, http.StatusInternalServerError, "failed to fecth all page stocks", err)
		return
	}

	err = service.InsertStocks(r.Context(), apiStocks)
	if err != nil {
		router.RespondWithError(w, http.StatusInternalServerError, "failed to perform bulk stocks insert", err)
		return
	}

	// background job to get the remaining stocks
	go func() {
		ctx := context.Background()
		fmt.Println("Starting background job to fetch remaining stocks from the API")
        defer close(stocksBuffer)
		// defer cancelCtx()
		for {
			// get stocks by batches of 10 e.g: 100 stocks records
			asyncStocks, nextPage, err := service.FetchStockPages(ctx, 10, lastPage)
            // TODO: add score to the stocks

			if err != nil {
				fmt.Printf("failed to fetch stocks from API: %v", err)
				return
			}

			// flush the stocks batch in the database
			err = service.InsertStocks(ctx, asyncStocks)
			if err != nil {
				fmt.Printf("failed to perform bulk stocks insertion: %v", err)
				return
			}

            // send the stocks to the buffer
            stocksBuffer <- asyncStocks

			if nextPage == "" {
				break
			}
			lastPage = nextPage
		}
		fmt.Println("Finished fetching all stocks from the API and inserting them into the database")
	}()

	router.RespondWithJSON(w, http.StatusOK, apiStocks)
}

func getServiceFromContext(db *pgxpool.Pool, r *http.Request) (*Service, error) {
	c, err := validator.ExtractAndValidateContext[*config.Conf](r.Context(), contextkeys.CtxKeyConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to extract config from context: %w", err)
	}
	return NewService(c, db), nil
}
