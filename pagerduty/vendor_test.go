package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestVendorsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/vendors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"vendors": [{"id": "1"}]}`))
	})

	resp, _, err := client.Vendors.List(&ListVendorsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListVendorsResponse{
		Vendors: []*Vendor{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}
func TestVendorsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/vendors/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"vendor": {"id": "1"}}`))
	})

	resp, _, err := client.Vendors.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Vendor{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}
