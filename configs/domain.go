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

// ValidateContactFor2ndLevelUKDomainConfig is the config of check domain availability API
type ValidateContactFor2ndLevelUKDomainConfig struct {
	// Required
	DomainName   string `json:"domain_name" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Company      string `json:"company" binding:"required"`
	Email        string `json:"email" binding:"required"`
	AddressLine1 string `json:"address_line_1" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	ZipCode      string `json:"zipcode" binding:"required"`
	Country      string `json:"country" binding:"required"`

	// Optional
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	PhoneCC      string `json:"phone_cc"`
	Phone        string `json:"phone"`
}

// Resource is use to get resource to call the API
func (config ValidateContactFor2ndLevelUKDomainConfig) Resource() string {
	return "domains/uk"
}

// Method is use to get method to call the API
func (config ValidateContactFor2ndLevelUKDomainConfig) Method() string {
	return "available"
}

// ToVariables is to transform config to url.Values
func (config ValidateContactFor2ndLevelUKDomainConfig) ToVariables() (v url.Values) {
	v = url.Values{}

	// Required
	v.Add("domain-name", config.DomainName)
	v.Add("name", config.Name)
	v.Add("company", config.Company)
	v.Add("email", config.Email)
	v.Add("address-line-1", config.AddressLine1)
	v.Add("city", config.City)
	v.Add("state", config.State)
	v.Add("zipcode", config.ZipCode)
	v.Add("country", config.Country)

	// Optional
	v.Add("address-line-2", config.AddressLine2)
	v.Add("address-line-3", config.AddressLine3)
	v.Add("phone-cc", config.PhoneCC)
	v.Add("phone", config.Phone)

	return
}

// IsGetMode is to tell API request that this call use GET method
func (config ValidateContactFor2ndLevelUKDomainConfig) IsGetMode() bool {
	return false
}
