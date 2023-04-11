package pagerduty

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLicensesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"licenses": [{"id": "P1D3Z4B"}]}`))
	})

	resp, _, err := client.Licenses.List()
	if err != nil {
		t.Fatal(err)
	}

	want := []*License{
		{
			ID: "P1D3Z4B",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestListAllocations(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/license_allocations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"license_allocations": [
			{
				"user": {"id": "P1D3Z4B", "type": "user_reference"},
				"license": {"id": "P1D3XYZ", "type": "license"},
				"allocated_at": "2021-06-01T21:30:42Z"
			}
		]}`))
	})

	resp, _, err := client.Licenses.ListAllocations(&ListLicenseAllocationsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListLicenseAllocationsResponse{
		LicenseAllocations: []*LicenseAllocation{
			{
				License:     &License{ID: "P1D3XYZ", Type: "license"},
				User:        &UserReference{ID: "P1D3Z4B", Type: "user_reference"},
				AllocatedAt: "2021-06-01T21:30:42Z",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestListAllAllocations(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/license_allocations", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		more := true
		userID := "P1D3Z4B"
		if _, exists := params["offset"]; exists {
			more = false
			userID = "P1D3Z4A"
		}

		testMethod(t, r, "GET")
		w.Write([]byte(fmt.Sprintf(`{
			"license_allocations": [
				{
					"user": {"id": "%s", "type": "user_reference"},
					"license": {"id": "P1D3XYZ", "type": "license"},
					"allocated_at": "2021-06-01T21:30:42Z"
				}
			],
			"limit": 1,
			"more": %t,
			"total": 2
		}`, userID, more)))
	})

	resp, err := client.Licenses.ListAllAllocations(&ListLicenseAllocationsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := []*LicenseAllocation{
		{
			License:     &License{ID: "P1D3XYZ", Type: "license"},
			User:        &UserReference{ID: "P1D3Z4B", Type: "user_reference"},
			AllocatedAt: "2021-06-01T21:30:42Z",
		},
		{
			License:     &License{ID: "P1D3XYZ", Type: "license"},
			User:        &UserReference{ID: "P1D3Z4A", Type: "user_reference"},
			AllocatedAt: "2021-06-01T21:30:42Z",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}
