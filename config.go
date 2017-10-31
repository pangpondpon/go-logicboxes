package logicboxes

import "net/url"

type Config interface {
	Resource() string
	Method() string
	ToVariables() (v url.Values)
}
