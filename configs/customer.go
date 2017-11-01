package configs

import "net/url"

// GetCustomerDetailsFromEmailConfig is the config to get customer details
type GetCustomerDetailsFromEmailConfig struct {
	CustomerEmail string
}

// Resource is use to get resource to call the API
func (config GetCustomerDetailsFromEmailConfig) Resource() string {
	return "customers"
}

// Method is use to get method to call the API
func (config GetCustomerDetailsFromEmailConfig) Method() string {
	return "details"
}

// ToVariables is to transform config to url.Values
func (config GetCustomerDetailsFromEmailConfig) ToVariables() (v url.Values) {
	v.Add("username", config.CustomerEmail)

	return
}

// IsGetMode is to tell API request that this call use GET method
func (config GetCustomerDetailsFromEmailConfig) IsGetMode() bool {
	return true
}
