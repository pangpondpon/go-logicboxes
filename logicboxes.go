package logicboxes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var LogicboxesInstance Logicboxes

// Main type of logicboxes
type Logicboxes struct {
	UserID, APIKey string
	TestMode       bool

	Domain  Domain
	Contact Contact
}

func NewLogicboxes(userId, apiKey string, testMode bool) Logicboxes {
	logicboxes := Logicboxes{
		UserID:   userId,
		APIKey:   apiKey,
		TestMode: testMode,
		Domain:   Domain{},
		Contact:  Contact{},
	}

	LogicboxesInstance = logicboxes

	return logicboxes
}

func Call(resource, method string, variables url.Values, getMode bool) interface{} {
	url := fmt.Sprintf("%s%s%s%s", BaseUrl(), ResourceMethodUrl(resource, method), CredentialUrl(), VariablesUrl(variables))

	var response *http.Response
	var err error

	if getMode {
		response, err = http.Get(url)
	} else {
		response, err = http.Post(url, "text/plain", bytes.NewBuffer([]byte{}))
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return response
	}

	// Transform body byte to json
	var result interface{}
	err = json.Unmarshal(bodyBytes, &result)

	return result
}

func BaseUrl() string {
	if LogicboxesInstance.TestMode {
		return "https://test.httpapi.com/api/"
	}

	return "https://httpapi.com/api/"
}

func ResourceMethodUrl(resource, method string) string {
	return resource + "/" + method + ".json?"
}

func CredentialUrl() string {
	return fmt.Sprintf("auth-userid=%s&api-key=%s&", LogicboxesInstance.UserID, LogicboxesInstance.APIKey)
}

func VariablesUrl(variables url.Values) string {
	return variables.Encode()
}
