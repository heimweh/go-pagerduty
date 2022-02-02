# go-pagerduty
PagerDuty API client in Go, primarily used by the [PagerDuty](https://github.com/nordcloud/terraform-provider-pagerduty) provider in Terraform.

This project is a fork of [heimweh/go-pagerduty](https://github.com/heimweh/go-pagerduty).

[![GoDoc](https://godoc.org/github.com/nordcloud/go-pagerduty?status.svg)](http://godoc.org/github.com/nordcloud/go-pagerduty/pagerduty)

## Installation
```bash
go get github.com/nordcloud/go-pagerduty/pagerduty
```

## Example usage
```go
package main

import (
	"fmt"
	"os"

	"github.com/nordcloud/go-pagerduty/pagerduty"
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
	fmt.Println(raw.StatusCode)
}
```

## Contributing
1. Fork it (https://github.com/nordcloud/go-pagerduty/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
