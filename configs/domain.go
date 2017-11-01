package configs

import "net/url"

// CheckDomainAvailabilityConfig is the config of check domain availability API
type CheckDomainAvailabilityConfig struct {
	DomainName, Tld string
}

// Resource is use to get resource to call the API
func (config CheckDomainAvailabilityConfig) Resource() string {
	return "domains"
}

// Method is use to get method to call the API
func (config CheckDomainAvailabilityConfig) Method() string {
	return "available"
}

// ToVariables is to transform config to url.Values
func (config CheckDomainAvailabilityConfig) ToVariables() (v url.Values) {
	v = url.Values{}
	v.Add("domain-name", config.DomainName)
	v.Add("tlds", config.Tld)

	return
}

// IsGetMode is to tell API request that this call use GET method
func (config CheckDomainAvailabilityConfig) IsGetMode() bool {
	return true
}
