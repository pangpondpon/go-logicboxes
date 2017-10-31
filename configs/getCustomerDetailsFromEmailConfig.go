package configs

import "net/url"

type GetCustomerDetailsFromEmailConfig struct {
	CustomerEmail string
}

func (config GetCustomerDetailsFromEmailConfig) Resource() string {
	return "customers"
}

func (config GetCustomerDetailsFromEmailConfig) Method() string {
	return "details"
}

func (config GetCustomerDetailsFromEmailConfig) ToVariables() (v url.Values) {
	v.Add("username", config.CustomerEmail)

	return
}
