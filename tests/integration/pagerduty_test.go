package integration

import (
	"fmt"
	"os"
	"testing"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

var (
	client *pagerduty.Client
	token  string
)

func init() {
	token = os.Getenv("TEST_PAGERDUTY_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "No test token found, please set $TEST_PAGERDUTY_TOKEN")
		os.Exit(0)
	}
}

func setup(t *testing.T) {
	t.Parallel()
	client, _ = pagerduty.NewClient(&pagerduty.Config{Token: os.Getenv("TEST_PAGERDUTY_TOKEN")})
}
