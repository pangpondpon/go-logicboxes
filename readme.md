# go-logicboxes
This library has a purpose to help you talk with Logicboxes API with ease.

Status: In Progress

## Important
This project does not intent to implement all the API call that Logicboxes provided, I'll keep adding the one that I use in my project.

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