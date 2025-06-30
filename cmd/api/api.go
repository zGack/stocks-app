package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zgack/stocks/internal/store"
)

type application struct {
	config configOld
	store  store.Storage
}

type configOld struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/stocks", app.createStocksHandler)

	return mux
}

// var stocks = []Stock{}

func (app *application) run(mux *http.ServeMux) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return server.ListenAndServe()
}
