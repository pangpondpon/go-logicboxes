package configs

import "net/url"

// AddToQueryIfValueIsNotEmptyString add the value to query if not empty
func AddToQueryIfValueIsNotEmptyString(v *url.Values, value string, key string) {
	if value != "" {
		v.Add(key, value)
	}
}
