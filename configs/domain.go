package configs

import (
	"net/url"
	"strconv"
)

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

// ValidateContactFor2ndLevelUKDomainConfig is the config of validate 2nd level uk domain API
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
	AddToQueryIfValueIsNotEmptyString(&v, config.AddressLine2, "address-line-2")
	AddToQueryIfValueIsNotEmptyString(&v, config.AddressLine3, "address-line-3")
	AddToQueryIfValueIsNotEmptyString(&v, config.PhoneCC, "phone-cc")
	AddToQueryIfValueIsNotEmptyString(&v, config.Phone, "phone")

	return
}

// IsGetMode is to tell API request that this call use GET method
func (config ValidateContactFor2ndLevelUKDomainConfig) IsGetMode() bool {
	return false
}

// RegisterDomainConfig is the config of register domain API
type RegisterDomainConfig struct {
	// Required
	DomainName       string   `json:"domain_name" binding:"required"`
	Years            int      `json:"years" binding:"required"`
	NS               []string `json:"ns" binding:"required"`
	CustomerID       string   `json:"customer_id" binding:"required"`
	RegContactID     string   `json:"reg_contact_id" binding:"required"`
	AdminContactID   string   `json:"admin_contact_id" binding:"required"`
	TechContactID    string   `json:"tech_contact_id" binding:"required"`
	BillingContactID string   `json:"billing_contact_id" binding:"required"`
	InvoiceOption    string   `json:"invoice_option" binding:"required"`

	// Optional
	PurchasePrivacy bool              `json:"purchase_privacy"`
	ProtectPrivacy  bool              `json:"protect_privacy"`
	Attrs           map[string]string `json:"attrs"`
}

// Resource is use to get resource to call the API
func (config RegisterDomainConfig) Resource() string {
	return "domains"
}

// Method is use to get method to call the API
func (config RegisterDomainConfig) Method() string {
	return "register"
}

// ToVariables is to transform config to url.Values
func (config RegisterDomainConfig) ToVariables() (v url.Values) {
	v = url.Values{}

	config.addRequiredToURL(&v)
	config.addOptionalToURL(&v)

	return
}

func (config RegisterDomainConfig) addRequiredToURL(v *url.Values) {
	v.Add("domain-name", config.DomainName)
	v.Add("years", strconv.Itoa(config.Years))

	for _, ns := range config.NS {
		v.Add("ns", ns)
	}

	v.Add("customer-id", config.CustomerID)
	v.Add("reg-contact-id", config.RegContactID)
	v.Add("admin-contact-id", config.AdminContactID)
	v.Add("tech-contact-id", config.TechContactID)
	v.Add("billing-contact-id", config.BillingContactID)
	v.Add("invoice-option", config.InvoiceOption)
}

func (config RegisterDomainConfig) addOptionalToURL(v *url.Values) {
	if config.PurchasePrivacy {
		v.Add("purchase-privacy", "true")
	} else {
		v.Add("purchase-privacy", "false")
	}

	if config.ProtectPrivacy {
		v.Add("protect-privacy", "true")
	} else {
		v.Add("protect-privacy", "false")
	}

	config.addAttrsToURL(v)
}

func (config RegisterDomainConfig) addAttrsToURL(v *url.Values) {
	index := 1
	for key, value := range config.Attrs {
		indexString := strconv.Itoa(index)

		v.Add("attr-name"+indexString, key)
		v.Add("attr-value"+indexString, value)

		index++
	}
}

// IsGetMode is to tell API request that this call use GET method
func (config RegisterDomainConfig) IsGetMode() bool {
	return false
}
