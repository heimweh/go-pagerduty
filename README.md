# go-pagerduty
PagerDuty API client in Go, primarily used by the [PagerDuty](https://github.com/PagerDuty/terraform-provider-pagerduty) provider in Terraform.


[![GoDoc](https://godoc.org/github.com/heimweh/go-pagerduty?status.svg)](http://godoc.org/github.com/heimweh/go-pagerduty/pagerduty)
[![Build
Status](https://travis-ci.org/heimweh/go-pagerduty.svg?branch=master)](https://travis-ci.org/heimweh/go-pagerduty)


## Installation
```bash
go get github.com/heimweh/go-pagerduty/pagerduty
```

## Example usage
1. Add something like the following to a go project:

```go
package main

import (
	"fmt"
	"os"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func main() {
	client, err := pagerduty.NewClient(&pagerduty.Config{Token: os.Getenv("PAGERDUTY_TOKEN")})
	if err != nil {
		panic(err)
	}

	resp, raw, err := client.Users.List(&pagerduty.ListUsersOptions{})
	if err != nil {
		panic(err)
	}

	for _, user := range resp.Users {
		fmt.Println(user.Name)
	}

	// All calls returns the raw *http.Response for further inspection.
	fmt.Println(raw.Response.StatusCode)
}
```

2. run:
```bash
$ PAGERDUTY_TOKEN=<SECRET> go run <PATH/TO/PROJECT/WITH/ABOVE/CODE>/main.go
```

## Caching support

Since some of the APIs implemented into this library doesn't offer a query mechanism for querying specific resources by their attributes, each time an implementation on the side of the Terraform Provider relies on that kind of logic, what it is done is to list all the resources of an specific entity and the lookup is executed in memory. Therefore, this leads to an inefficient use of the APIs, on top of that for use cases with a big amount of resources this repetitive API calls for lists of resources definitions start to pile up with the form of time consumption performance penalties that are nowadays causing uncomfortable experience for the Terraform Provider users.

### APIs Currently supporting caching on `go-pagerduty` library

* Abilities
* Contact Methods
* Notification Rules
* Team Members
* Users

### Caching mechanisms available

* In memory.
* MongoDB.

### To activate caching support

| Environment Variable       | Example Value                                                                      | Description                                                                                                                                  |
| -------------------------- | ---------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| TF_PAGERDUTY_CACHE         | memory                                                                             | Activate **In Memory** cache.                                                                                                                |
| TF_PAGERDUTY_CACHE         | `mongodb+srv://[mongouser]:[mongopass]@[mongodbname].[mongosubdomain].mongodb.net` | Activate MongoDB cache.                                                                                                                      |
| TF_PAGERDUTY_CACHE_MAX_AGE | 30s                                                                                | Only applicable for MongoDB cache. Time in seconds for cached data to become staled. Default value `10s`.                                    |
| TF_PAGERDUTY_CACHE_PREFILL | 1                                                                                  | Only applicable for MongoDB cache. Indicates to pre-fill data in cache for *Abilities*, *Users*, *Contact Methods* and *Notification Rules*. |

## Contributing
1. Fork it ( https://github.com/heimweh/go-pagerduty/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

### Testing

Run all unit tests with `make test`

Run a specific subset of unit test by name using `make test TESTARGS="-v -run TestTeams"` which will run all test functions with "TestTeams" in their name while `-v` enables verbose output.

### Environment Variables to test specific feature sets

| Environment Variable    | Example Value | Feature Set                                             |
| ----------------------- | ------------- | ------------------------------------------------------- |
| TF_PAGERDUTY_TEST_CACHE | 1             | Indicates to execute test sets that make use of caching |
