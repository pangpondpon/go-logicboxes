package logicboxes

import "net/url"

// VariablesURL method use to change url.Values to the url string
func VariablesURL(variables url.Values) string {
	return variables.Encode()
}

// ResourceMethodURL method use to concat resource and method
func ResourceMethodURL(resource, method string) string {
	return resource + "/" + method + ".json?"
}
