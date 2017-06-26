# go-pagerduty

## Installation
```bash
go get github.com/heimweh/go-pagerduty
```

## Example usage
```go
func main() {
  client, err := pagerduty.NewClient(&Config{Token: "foo"})
  if err != nil {
    panic(err)
  }

  // List all users
  resp, raw, err := client.Users.List(&pagerduty.ListUsersOptions{})
  if err != nil {
    panic(err)
  }

  for _, user := range resp.Users {
    fmt.Println(user.Name)
  }

  // All calls returns the raw *http.Response for further inspection
  fmt.Println(raw.StatusCode)
}
```

## Contributing
1. Fork it ( https://github.com/heimweh/go-pagerduty/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
