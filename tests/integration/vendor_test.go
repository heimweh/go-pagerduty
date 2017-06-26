package integration

import (
	"testing"

	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
)

func TestVendorsList(t *testing.T) {
	setup(t)

	resp, _, err := client.Vendors.List(&pagerduty.ListVendorsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Vendors) == 0 {
		t.Fatal("expected at least one vendor to be found")
	}
}
