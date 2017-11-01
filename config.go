package logicboxes

import "net/url"

// Config is the interface that use as an abstract type to use with Call() method
type Config interface {
	Resource() string
	Method() string
	ToVariables() (v url.Values)
	IsGetMode() bool
}
