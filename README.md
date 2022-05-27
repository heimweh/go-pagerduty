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

## Contributing
1. Fork it ( https://github.com/heimweh/go-pagerduty/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

### Testing

Run all unit tests with `make test`

Run a specific subset of unit test by name using `make test TESTARGS="-v -run TestTeams"` which will run all test functions with "TestTeams" in their name while `-v` enables verbose output.