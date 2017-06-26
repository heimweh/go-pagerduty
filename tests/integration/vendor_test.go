package integration

import (
	"testing"

	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
)

func TestVendorsList(t *testing.T) {
	setup(t)

	var offset int
	var vendors []*pagerduty.Vendor

	for {
		resp, _, err := client.Vendors.List(&pagerduty.ListVendorsOptions{Pagination: &pagerduty.Pagination{Offset: offset}})
		if err != nil {
			t.Fatal(err)
		}

		vendors = append(vendors, resp.Vendors...)

		if !resp.More {
			break
		}

		offset += 25
	}

	if len(vendors) == 0 {
		t.Fatal("expected at least one vendor to be found")
	}
}
