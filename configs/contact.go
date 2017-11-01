package configs

import (
	"net/url"
	"strconv"
)

// SearchContactConfig is the config of search contact API
type SearchContactConfig struct {
	CustomerID string
	PageNo     int
	PerPage    int
}

// Resource is use to get resource to call the API
func (config SearchContactConfig) Resource() string {
	return "contacts"
}

// Method is use to get method to call the API
func (config SearchContactConfig) Method() string {
	return "search"
}

// ToVariables is to transform config to url.Values
func (config SearchContactConfig) ToVariables() (v url.Values) {
	v.Add("customer-id", config.CustomerID)
	v.Add("no-of-records", strconv.Itoa(config.PerPage))
	v.Add("page-no", strconv.Itoa(config.PageNo))

	return
}

// IsGetMode is to tell API request that this call use GET method
func (config SearchContactConfig) IsGetMode() bool {
	return true
}
