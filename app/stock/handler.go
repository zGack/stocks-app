package stock

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/config"
	contextkeys "github.com/zgack/stocks/pkg/contextKeys"
	"github.com/zgack/stocks/pkg/router"
	"github.com/zgack/stocks/pkg/validator"
)

func GetAllStocksHandler(db *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {

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

    stocks, err := service.GetAllStocks(r.Context(), filters)
	if err != nil {
		router.RespondWithError(w, http.StatusInternalServerError, "failed to fecth all stocks from db", err)
		return
	}

    fmt.Printf("Fetched %d stocks from the database\n", len(stocks))

	router.RespondWithJSON(w, http.StatusOK, stocks)
}

func getServiceFromContext(db *pgxpool.Pool, r *http.Request) (*Service, error) {
	c, err := validator.ExtractAndValidateContext[*config.Conf](r.Context(), contextkeys.CtxKeyConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to extract config from context: %w", err)
	}
	return NewService(c, db), nil
}
