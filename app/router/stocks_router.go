package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zgack/stocks/app/stock"
)

type StockHandler struct {
	DB *pgxpool.Pool
}

func NewStockHandler(db *pgxpool.Pool) *StockHandler {
	return &StockHandler{
		DB: db,
	}
}

func (h *StockHandler) SetupRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /", h.GetAllStocksHandler)
}

func (h *StockHandler) GetAllStocksHandler(w http.ResponseWriter, r *http.Request) {
    stock.GetAllStocksHandler(h.DB, w, r)
}

func SetupStockRouter(db *pgxpool.Pool) *http.ServeMux {
    stockRouter := http.NewServeMux()

    handler := NewStockHandler(db)
    handler.SetupRoutes(stockRouter)
    return stockRouter
}
