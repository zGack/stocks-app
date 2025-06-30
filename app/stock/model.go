package stock

import (
	"net/http"
	"strconv"
	"time"
)

type Stock struct {
	ID         int64     `json:"id"`
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

type StockQueryFilters struct {
	Limit      int    `json:"limit" validate:"gte=1,lte=20"`
	Offset     int    `json:"offset" validate:"gte=0"`
	SortDir    string `json:"sort_dir" validate:"oneof=asc desc"`
	SortBy     string `json:"sort_by"`
	SearchTerm string `json:"search_term" validate:"omitempty,alphanumunicode,max=50"` // optional search term for filtering stocks
	SearchBy   string `json:"search_by" validate:"omitempty,oneof=company"`            // optional field to search by
}

func (f StockQueryFilters) Parse(r *http.Request) (StockQueryFilters, error) {
	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return f, nil
		}

		f.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return f, nil
		}

		f.Offset = o
	}

	sortDir := qs.Get("sort_dir")
	if sortDir != "" {
		f.SortDir = sortDir
	}

	sortBy := qs.Get("sort_by")
	if sortBy != "" {
		f.SortBy = sortBy
	}

    searchTerm := qs.Get("search_term")
    if searchTerm != "" {
        f.SearchTerm = searchTerm
    }

    searchBy := qs.Get("search_by")
    if searchBy != "" {
        f.SearchBy = searchBy
    }

	return f, nil
}
