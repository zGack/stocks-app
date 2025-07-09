package stocksfetcher

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zgack/stocks/app/stock"
)

type StockAPIRequest struct {
	apiURL    string
	authToken string
}

type StocksAPIResponse struct {
	Items    []stock.Stock `json:"items"`
	NextPage string        `json:"next_page"`
}

var ratingScore = map[string]int{
	"strong-buy":          15,
	"buy":                 12,
	"top pick":            12,
	"outperform":          10,
	"outperformer":        10,
	"overweight":          10,
	"market outperform":   10,
	"moderate buy":        9,
	"positive":            9,
	"speculative buy":     8,
	"sector outperform":   8,
	"equal weight":        5,
	"in-line":             5,
	"hold":                5,
	"neutral":             5,
	"market perform":      4,
	"sector perform":      4,
	"reduce":              -5,
	"underweight":         -7,
	"underperform":        -8,
	"sector underperform": -8,
	"sell":                -10,
	"strong sell":         -12,
	"unchanged":           0,
	"":                    0,
}

var actionScore = map[string]int{
	"upgraded by":       15,
	"target raised by":  12,
	"initiated by":      10,
	"target set by":     7,
	"reiterated by":     5,
	"target lowered by": -10,
	"downgraded by":     -15,
}

func parsePrice(stockPrice string) float64 {
	if stockPrice == "" {
		return 0.0
	}

    // Remove comma
    stockPrice = strings.ReplaceAll(stockPrice, ",", "")

	parsedPrice, err := strconv.ParseFloat(strings.TrimPrefix(stockPrice, "$"), 64)
	if err != nil {
        fmt.Printf("Error parsing stock price '%s': %v\n", stockPrice, err)
		return 0.0
	}
	return parsedPrice
}

func getStockScore(stock stock.Stock) float64 {
	from := parsePrice(stock.TargetFrom)
	to := parsePrice(stock.TargetTo)

	change := to - from
	percent := (change / from) * 100
	score := 0.0

	if to > from {
		score += percent
	}

    ratingTo := strings.ToLower(stock.RatingTo)
    ratingFrom := strings.ToLower(stock.RatingFrom)

	score += float64(ratingScore[ratingTo]) - float64(ratingScore[ratingFrom])
	score += float64(actionScore[stock.Action])

	return score
}
