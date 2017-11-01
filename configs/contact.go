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

// CreateContactConfig is the config of search contact API
type CreateContactConfig struct {
	// Required
	ContactName, Company, Email, AddressLine1, City, Country, ZipCode, PhoneCC, Phone, CustomerID, ContactType string

	// Optional
	AddressLine2, AddressLine3, State, FaxCC, Fax string
	Attrs                                         map[string]string
}

// Resource is use to get resource to call the API
func (config CreateContactConfig) Resource() string {
	return "contacts"
}

// Method is use to get method to call the API
func (config CreateContactConfig) Method() string {
	return "search"
}

// ToVariables is to transform config to url.Values
func (config CreateContactConfig) ToVariables() (v url.Values) {
	v = url.Values{}

	config.addRequiredParameters(&v)
	config.addOptionalParameters(&v)

	return
}

func (config CreateContactConfig) addRequiredParameters(v *url.Values) {
	// Required
	v.Add("name", config.ContactName)
	v.Add("company", config.Company)
	v.Add("email", config.Email)
	v.Add("address-line-1", config.AddressLine1)
	v.Add("city", config.City)
	v.Add("country", config.Country)
	v.Add("zipcode", config.ZipCode)
	v.Add("phone-cc", config.PhoneCC)
	v.Add("phone", config.Phone)
	v.Add("customer-id", config.CustomerID)
	v.Add("type", config.ContactType)
}

func (config CreateContactConfig) addOptionalParameters(v *url.Values) {
	if config.AddressLine2 != "" {
		v.Add("address-line-2", config.AddressLine2)
	}
	if config.AddressLine3 != "" {
		v.Add("address-line-3", config.AddressLine3)
	}
	if config.State != "" {
		v.Add("state", config.State)
	}
	if config.FaxCC != "" {
		v.Add("fax-cc", config.FaxCC)
	}
	if config.Fax != "" {
		v.Add("fax", config.Fax)
	}

	config.addAttrs(v)
}

// addAttrs is used to add the attributes value to the url
func (config CreateContactConfig) addAttrs(v *url.Values) {
	index := 1
	for key, value := range config.Attrs {
		indexString := strconv.Itoa(index)

		v.Add("attr-name"+indexString, key)
		v.Add("attr-value"+indexString, value)

		index++
	}
}

// IsGetMode is to tell API request that this call use GET method
func (config CreateContactConfig) IsGetMode() bool {
	return false
}
