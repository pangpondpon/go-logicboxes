package configs

import "net/url"

type CheckDomainAvailabilityConfig struct {
	DomainName, Tld string
}

func (config CheckDomainAvailabilityConfig) Resource() string {
	return "domains"
}

func (config CheckDomainAvailabilityConfig) Method() string {
	return "available"
}

func (config CheckDomainAvailabilityConfig) ToVariables() (v url.Values) {
	v.Add("domain-name", config.DomainName)
	v.Add("tlds", config.Tld)

	return
}
