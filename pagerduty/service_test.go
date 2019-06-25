package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestServicesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(validListServicesJSON))
	})

	resp, _, err := client.Services.List(&ListServicesOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp, validListServicesResponse) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, validListServicesResponse)
	}
}

func TestServicesCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Service{
		Name: "foo",
	}

	mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Service)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Service, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"service": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Services.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Service{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServicesDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Services.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestServicesGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"service": {"id": "1"}}`))
	})

	resp, _, err := client.Services.Get("1", &GetServiceOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &Service{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServicesUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Service{
		Name: "foo",
	}

	mux.HandleFunc("/services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Service)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Service, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"service": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Services.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Service{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServicesCreateIntegration(t *testing.T) {
	setup()
	defer teardown()

	input := &Integration{
		Name: "foo",
		ID:   "1",
	}

	mux.HandleFunc("/services/1/integrations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Integration)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Integration, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"integration": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Services.CreateIntegration("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Integration{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServicesUpdateIntegration(t *testing.T) {
	setup()
	defer teardown()

	input := &Integration{
		Name: "foo",
	}

	mux.HandleFunc("/services/1/integrations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Integration)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Integration, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"integration": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Services.UpdateIntegration("1", "1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Integration{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServicesGetIntegration(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/services/1/integrations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"integration": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Services.GetIntegration("1", "1", &GetIntegrationOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &Integration{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServicesDeleteIntegration(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/1/integrations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Services.DeleteIntegration("1", "1"); err != nil {
		t.Fatal(err)
	}
}
