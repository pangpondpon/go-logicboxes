# go-logicboxes
This library has a purpose to help you talk with Logicboxes API with ease.

Status: In Progress

## Important
1. This project does not intent to implement all the API call that Logicboxes provided, I'll keep adding the one that I use in my project.
2. The library does not parse the response into custom response, it will be the pure response from Logicboxes in JSON format. If you want to transform it, please use your own code to transform the response after get it from this library. 

## Installation
I didn't publish this package to Go package directory yet. So if you want to use this package you'll need to run git clone to your go src and manually import it.
TODO: Publish this package to Go package directory.

## Usage
```go
package main

import (
	"github.com/pangpondpon/go-logicboxes"
)

// Live mode
lb := logicboxes.NewLogicboxes("USER_ID", "API_KEY", false)

// Test mode
lb = logicboxes.NewLogicboxes("USER_ID", "API_KEY", true)

// Check domain availability
config := configs.CheckDomainAvailabilityConfig{
	DomainName: "xxx"
	Tld: "com"
}

response := lb.CheckDomainAvailability(config)

fmt.Println(response)
```