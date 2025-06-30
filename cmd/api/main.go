package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/app/router"
	"github.com/zgack/stocks/app/router/middleware"
	"github.com/zgack/stocks/config"
	"github.com/zgack/stocks/internal/db"
)

func main() {
	c := config.New()
	// TODO: add logger

	db := openDatabase(c.DB.DBPath)
	// TODO: add db migrations

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

	// TODO: add graceful shutdown

	log.Printf("Starting server on %v", c.Server.Port)
	log.Fatal(s.ListenAndServe())
}

func openDatabase(dbURL string) (*pgxpool.Pool) {
	db, err := db.New(dbURL)
	if err != nil {
		log.Panic(err)
	}
	return db
}

// func main() {
//
// 	appConfig := config{
// 		addr: env.GetString("ADDR", ":8080"),
// 		db: dbConfig{
// 			addr: env.GetString("DB_ADDR", "postgresql://root@localhost:26257/defaultdb?sslmode=disable"),
// 		},
//         env: env.GetString("ENV", "development"),
// 	}
//
//     db, ctx, err := db.New(appConfig.db.addr)
//     if err != nil {
//         log.Panic(err)
//     }
//     defer db.Close(ctx)
//     log.Printf("Database connection pool established")
//
//     store := store.NewCockroachStorage(db)
//
// 	app := &application{
// 		config: appConfig,
// 		store:  store,
// 	}
//
// 	mux := app.mount()
//
// 	log.Fatal(app.run(mux))
// }
