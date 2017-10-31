package logicboxes

import (
	"net/url"
)

type Domain struct {
	CheckAvailability CheckDomainAvailability
}

type CheckDomainAvailability struct {
	DomainName, Tld string
}

func (c CheckDomainAvailability) Call() interface{} {
	return Call("domains", "available", c.toVariables(), true)
}

func (c CheckDomainAvailability) toVariables() url.Values {
	values := url.Values{}
	values.Add("domain-name", c.DomainName)
	values.Add("tlds", c.Tld)
	return values
}
