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

type APIResponse map[string]interface{}

// Main type of logicboxes
type Logicboxes struct {
	UserID, APIKey string
	TestMode       bool
}

func (logicboxes Logicboxes) CallByConfig(config Config) APIResponse {
	return logicboxes.Call(config.Resource(), config.Method(), config.ToVariables(), logicboxes.TestMode)
}

func (logicboxes Logicboxes) Call(resource, method string, variables url.Values, testMode bool) map[string]interface{} {
	u := fmt.Sprintf("%s%s%s%s", logicboxes.BaseUrl(), ResourceMethodUrl(resource, method), logicboxes.CredentialUrl(), VariablesUrl(variables))

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

func (logicboxes Logicboxes) BaseUrl() string {
	if logicboxes.TestMode {
		return "https://test.httpapi.com/api/"
	}

	return "https://httpapi.com/api/"
}

func (logicboxes Logicboxes) CredentialUrl() string {
	return fmt.Sprintf("auth-userid=%s&api-key=%s&", logicboxes.UserID, logicboxes.APIKey)
}

func NewLogicboxes(userId, apiKey string, testMode bool) Logicboxes {
	logicboxes := Logicboxes{
		UserID:   userId,
		APIKey:   apiKey,
		TestMode: testMode,
	}

	return logicboxes
}

func (logicboxes Logicboxes) CheckDomainAvailability(config configs.CheckDomainAvailabilityConfig) APIResponse {
	return logicboxes.CallByConfig(config)
}

func (logicboxes Logicboxes) GetCustomerDetailsFromEmail(config configs.GetCustomerDetailsFromEmailConfig) APIResponse {
	return logicboxes.CallByConfig(config)
}

func (logicboxes Logicboxes) SearchContact(config configs.SearchContactConfig) APIResponse {
	return logicboxes.CallByConfig(config)
}
