package logicboxes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pangpondpon/go-logicboxes/configs"
)

// APIResponse type is the type that return from the API call
type APIResponse map[string]interface{}

// Logicboxes type is the main type that user must initilize
// and use this type as a base to call the API
type Logicboxes struct {
	UserID, APIKey string
	TestMode       bool
}

// CallByConfig allow user to use a custom Config type to call the API
func (logicboxes Logicboxes) CallByConfig(config Config) APIResponse {
	return logicboxes.Call(config.Resource(), config.Method(), config.ToVariables(), logicboxes.TestMode)
}

// Call allow user to call the API without using Config type
func (logicboxes Logicboxes) Call(resource, method string, variables url.Values, testMode bool) map[string]interface{} {
	u := fmt.Sprintf("%s%s%s%s", logicboxes.BaseURL(), ResourceMethodURL(resource, method), logicboxes.CredentialURL(), VariablesURL(variables))

	var response *http.Response
	var err error

	if testMode {
		response, err = http.Get(u)
	} else {
		response, err = http.Post(u, "text/plain", bytes.NewBuffer([]byte{}))
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil
	}

	// Transform body byte to json
	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)

	return result
}

// BaseURL method is for getting the base url to call the API
func (logicboxes Logicboxes) BaseURL() string {
	if logicboxes.TestMode {
		return "https://test.httpapi.com/api/"
	}

	return "https://httpapi.com/api/"
}

// CredentialURL method is for getting credential URL to call the API
func (logicboxes Logicboxes) CredentialURL() string {
	return fmt.Sprintf("auth-userid=%s&api-key=%s&", logicboxes.UserID, logicboxes.APIKey)
}

// NewLogicboxes function can be use to initilize the Logicboxes type instance
func NewLogicboxes(userID, apiKey string, testMode bool) Logicboxes {
	logicboxes := Logicboxes{
		UserID:   userID,
		APIKey:   apiKey,
		TestMode: testMode,
	}

	return logicboxes
}

// CheckDomainAvailability method is use to call check domain availability API
// The official API page is located at https://manage.logicboxes.com/kb/answer/764
func (logicboxes Logicboxes) CheckDomainAvailability(config configs.CheckDomainAvailabilityConfig) APIResponse {
	return logicboxes.CallByConfig(config)
}

// GetCustomerDetailsFromEmail method is use to call get customer details using customer email
// The official API page is located at https://manage.logicboxes.com/kb/answer/874
func (logicboxes Logicboxes) GetCustomerDetailsFromEmail(config configs.GetCustomerDetailsFromEmailConfig) APIResponse {
	return logicboxes.CallByConfig(config)
}

// SearchContact method is use to call search contact API
// The official API page is located at https://manage.logicboxes.com/kb/answer/793
func (logicboxes Logicboxes) SearchContact(config configs.SearchContactConfig) APIResponse {
	return logicboxes.CallByConfig(config)
}
