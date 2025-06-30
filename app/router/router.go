package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/app/health"
	"github.com/zgack/stocks/pkg/router"
)

func SetupRouter(db *pgxpool.Pool) *http.ServeMux {
	routerMux := http.NewServeMux()

	routerMux.HandleFunc("GET /v1/health", health.HealthHandler)

	stockRouter := SetupStockRouter(db)
	router.Mount(routerMux, "/v1/stocks", stockRouter)

	return routerMux
}
