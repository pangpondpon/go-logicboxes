package logicboxes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pangpondpon/go-logicboxes/configs"
)

// APIResponse type is the type that return from the API call
type APIResponse interface{}

// Logicboxes type is the main type that user must initilize
// and use this type as a base to call the API
type Logicboxes struct {
	UserID, APIKey string
	TestMode       bool
	Log            bool
}

// CallByConfig allow user to use a custom Config type to call the API
func (logicboxes Logicboxes) CallByConfig(config Config) (APIResponse, error) {
	return logicboxes.Call(config.Resource(), config.Method(), config.ToVariables(), config.IsGetMode())
}

// Call allow user to call the API without using Config type
func (logicboxes Logicboxes) Call(resource, method string, variables url.Values, getMode bool) (APIResponse, error) {
	u := fmt.Sprintf("%s%s%s%s", logicboxes.BaseURL(), ResourceMethodURL(resource, method), logicboxes.CredentialURL(), VariablesURL(variables))

	var response *http.Response
	var err error

	if getMode {
		response, err = http.Get(u)
	} else {
		response, err = http.Post(u, "text/plain", bytes.NewBuffer([]byte{}))
	}

	if response == nil {
		return nil, errors.New("Can't connect to Logicboxes, please try again")
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	// Transform body byte to json
	var result interface{}
	err = json.Unmarshal(bodyBytes, &result)

	if logicboxes.Log {
		logicboxes.logRequestResponse(u, getMode, result)
	}

	return result, nil
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

// logRequestResponse will log the URL, HTTP Method and the result to the terminal
// It's good for debugging
func (logicboxes Logicboxes) logRequestResponse(u string, getMode bool, result APIResponse) {
	fmt.Println("\n[LOGICBOXES-debug] " + logicboxes.getHTTPMethodString(getMode) + "   " + u)
	fmt.Println("Response:")
	fmt.Println(result)
	fmt.Println("")
}

// getHTTPMethodString return the full string name of method
// Only GET and POST will be returned from this method
func (logicboxes Logicboxes) getHTTPMethodString(getMode bool) string {
	if getMode {
		return "GET"
	}

	return "POST"
}

// NewLogicboxes function can be use to initilize the Logicboxes type instance
func NewLogicboxes(userID, apiKey string, testMode bool) Logicboxes {
	logicboxes := Logicboxes{
		UserID:   userID,
		APIKey:   apiKey,
		TestMode: testMode,
		Log:      false,
	}

	return logicboxes
}

// NewLogicboxesWithLogging function can be use to initilize the Logicboxes type instance
// It will set Log property to true, so you'll see all the outgoing API call
func NewLogicboxesWithLogging(userID, apiKey string, testMode bool) Logicboxes {
	logicboxes := Logicboxes{
		UserID:   userID,
		APIKey:   apiKey,
		TestMode: testMode,
		Log:      true,
	}

	return logicboxes
}

// CheckDomainAvailability method is use to call check domain availability API
// The official API page is located at https://manage.logicboxes.com/kb/answer/764
func (logicboxes Logicboxes) CheckDomainAvailability(config configs.CheckDomainAvailabilityConfig) (APIResponse, error) {
	return logicboxes.CallByConfig(config)
}

// GetCustomerDetailsFromEmail method is use to call get customer details using customer email
// The official API page is located at https://manage.logicboxes.com/kb/answer/874
func (logicboxes Logicboxes) GetCustomerDetailsFromEmail(config configs.GetCustomerDetailsFromEmailConfig) (APIResponse, error) {
	return logicboxes.CallByConfig(config)
}

// SearchContact method is use to call search contact API
// The official API page is located at https://manage.logicboxes.com/kb/answer/793
func (logicboxes Logicboxes) SearchContact(config configs.SearchContactConfig) (APIResponse, error) {
	return logicboxes.CallByConfig(config)
}

// CreateContact method is use to call create contact API
// The official API page is located at https://manage.netearthone.com/kb/answer/790
func (logicboxes Logicboxes) CreateContact(config configs.CreateContactConfig) (APIResponse, error) {
	return logicboxes.CallByConfig(config)
}

// ValidateContactFor2ndLevelUKDomainName method is use to call Validating Contact Information for a 2nd Level .UK Domain Name API
// The official API page is located at https://manage.netearthone.com/kb/answer/2187
func (logicboxes Logicboxes) ValidateContactFor2ndLevelUKDomainName(config configs.ValidateContactFor2ndLevelUKDomainConfig) (APIResponse, error) {
	return logicboxes.CallByConfig(config)
}

// RegisterDomain method is use to call Register Domain API
// The official API page is located at https://manage.netearthone.com/kb/answer/752
func (logicboxes Logicboxes) RegisterDomain(config configs.RegisterDomainConfig) (APIResponse, error) {
	return logicboxes.CallByConfig(config)
}
