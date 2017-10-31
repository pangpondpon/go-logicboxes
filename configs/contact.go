package configs

import (
	"strconv"
	"net/url"
)

type SearchContactConfig struct {
	CustomerID string
	PageNo     int
	PerPage    int
}

func (config SearchContactConfig) Resource() string {
	return "contacts"
}

func (config SearchContactConfig) Method() string {
	return "search"
}

func (config SearchContactConfig) ToVariables() (v url.Values) {
	v.Add("customer-id", config.CustomerID)
	v.Add("no-of-records", strconv.Itoa(config.PerPage))
	v.Add("page-no", strconv.Itoa(config.PageNo))

	return
}
