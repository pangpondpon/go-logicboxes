package logicboxes

import "net/url"

func VariablesUrl(variables url.Values) string {
	return variables.Encode()
}

func ResourceMethodUrl(resource, method string) string {
	return resource + "/" + method + ".json?"
}