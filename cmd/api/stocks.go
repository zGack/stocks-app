package main

import (
	"encoding/json"
	"net/http"

	"github.com/zgack/stocks/internal/env"
	"github.com/zgack/stocks/internal/store"
)

type StockItems struct {
	Items    []store.Stock `json:"items"`
	NextPage string        `json:"next_page"`
}

func (app *application) createStocksHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+env.GetString("STOCKS_API_AUTH_TOKEN", ""))

	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch stocks", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

    var stockItems StockItems

    err = json.NewDecoder(resp.Body).Decode(&stockItems)
    if err != nil {
        http.Error(w, "Failed to read response body", http.StatusInternalServerError)
        return
    }

    ctx := r.Context()

    var stock = &store.Stock{
        Ticker: stockItems.Items[0].Ticker,
        TargetFrom: stockItems.Items[0].TargetFrom,
        TargetTo: stockItems.Items[0].TargetTo,
        Company: stockItems.Items[0].Company,
        Action: stockItems.Items[0].Action,
        Brokerage: stockItems.Items[0].Brokerage,
        RatingFrom: stockItems.Items[0].RatingFrom,
        RatingTo: stockItems.Items[0].RatingTo,
    }


    if err := app.store.Stocks.Create(ctx, stock); err != nil {
        writeJSONError(w, http.StatusInternalServerError, "Failed to create stock in database")
        return
    }

    if err := writeJSON(w, http.StatusCreated, stock); err != nil {
        writeJSONError(w, http.StatusInternalServerError, "Failed to write JSON response")
        return
    }
}
