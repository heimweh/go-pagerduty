package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestBusinessServiceList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/business_services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 0, "offset": 0, "more": false, "limit": 0, "business_services":[{"id": "1"}]}`))
	})

	resp, _, err := client.BusinessServices.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListBusinessServicesResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		BusinessServices: []*BusinessService{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestBusinessServiceCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &BusinessService{Name: "foo"}

	mux.HandleFunc("/business_services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(BusinessService)
		v.Name = "foo"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"business_service":{"name": "foo", "id":"1"}}`))
	})

	resp, _, err := client.BusinessServices.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessService{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
func TestBusinessServiceGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/business_services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"business_service":{"name": "foo", "id":"1"}}`))
	})

	ID := "1"
	resp, _, err := client.BusinessServices.Get(ID)

	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessService{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestBusinessServiceUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &BusinessService{
		Name: "foo",
	}

	mux.HandleFunc("/business_services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(BusinessService)
		v.Name = "foo"

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"business_service":{"name": "foo", "id":"1"}}`))
	})

	resp, _, err := client.BusinessServices.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessService{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestBusinessServiceDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/business_services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.BusinessServices.Delete("1"); err != nil {
		t.Fatal(err)
	}
}
