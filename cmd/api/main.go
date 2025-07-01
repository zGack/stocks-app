package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/app/router"
	"github.com/zgack/stocks/app/router/middleware"
	"github.com/zgack/stocks/config"
	"github.com/zgack/stocks/internal/db"
	stocksfetcher "github.com/zgack/stocks/internal/stocksFetcher"
)

func main() {
	c := config.New()
	// TODO: add logger

	db := openDatabase(c.DB.DBPath)

	defer db.Close()

	mux := router.SetupRouter(db)

	middlewaresStack := middleware.CreateStack(
		middleware.InjectDeps(c),
		middleware.CORS(c.Server.CorsOrigins),
		middleware.ContentTypeJSON,
	)

	serverHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Apply middlewares
		middlewaresStack(mux).ServeHTTP(w, r)
	})

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      serverHandler,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	stocksFetcherService := stocksfetcher.NewService(db)

    ctx, cancel := context.WithCancel(context.Background())

	// Fetch initial stock data
	lastPage, err := stocksFetcherService.FetchInitialStocks(context.Background())
	if err != nil {
		log.Fatal("Failed to fetch initial stocks:", err)
	}

    // Init background job to fetch remaining stocks
    go stocksFetcherService.FetchRemainingStocks(ctx, cancel, lastPage)

	// TODO: add graceful shutdown

	log.Printf("Starting server on %v", c.Server.Port)
	log.Fatal(s.ListenAndServe())
}

func openDatabase(dbURL string) *pgxpool.Pool {
	db, err := db.New(dbURL)
	if err != nil {
		log.Panic(err)
	}
	return db
}

