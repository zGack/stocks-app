package stocksfetcher

import "github.com/zgack/stocks/app/stock"

type StockApiRequest struct {
	apiUrl    string
	authToken string
}

type StocksAPIResponse struct {
	Items    []stock.Stock `json:"items"`
	NextPage string        `json:"next_page"`
}
