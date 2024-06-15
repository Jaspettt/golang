package domain

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Filters struct {
	Offset     int
	Limit      int
	Sort       string
	Order      string
	Filter     string
	FilterType string
}

func ExtractFields(r *http.Request) *Filters {
	var (
		order      = strings.ToUpper(r.URL.Query().Get("order"))
		limit, _   = strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _  = strconv.Atoi(r.URL.Query().Get("offset"))
		sort       = r.URL.Query().Get("sort")
		filter     = "'%%'"
		filterType = "artist"
	)
	validateFilters(r.URL.Query(), &limit, &offset, &sort, &order, &filter, &filterType)
	return &Filters{
		Offset:     offset,
		Limit:      limit,
		Sort:       sort,
		Order:      order,
		Filter:     filter,
		FilterType: filterType,
	}

}
func validateFilters(url url.Values, offset, limit *int, sort, order, filter, filterType *string) {
	validFields := map[string]bool{
		"id":          true,
		"artist":      true,
		"title":       true,
		"releasedate": true,
		"price":       true,
		"rating":      true,
	}
	if *limit <= 0 {
		*limit = 10
	}
	if *offset < 0 {
		*offset = 0
	}
	if *sort == "" {
		*sort = "id"
	}
	if *order != "ASC" && *order != "DESC" {
		*order = "ASC"
	}
	for key, _ := range validFields {
		if url.Get(key) != "" && len(url.Get(key)) < 10 {
			*filterType = fmt.Sprintf("%s::text", key)
			*filter = "'%" + url.Get(key) + "%'"
		}
	}
}
