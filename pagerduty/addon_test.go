package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestAddonsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/addons", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"addons": [{"name": "Internal Status Page"}]}`))
	})

	addons, _, err := client.Addons.List(&ListAddonsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListAddonsResponse{
		Addons: []*Addon{
			{
				Name: "Internal Status Page",
			},
		},
	}

	if !reflect.DeepEqual(addons, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", addons, want)
	}
}

func TestAddonsInstall(t *testing.T) {
	setup()
	defer teardown()

	input := &Addon{
		Name: "Internal Status Page",
	}

	mux.HandleFunc("/addons", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		v := new(Addon)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Addon, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		w.Write([]byte(`{"addon": {"name": "Internal Status Page", "id": "1"}}`))
	})

	addon, _, err := client.Addons.Install(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Addon{
		Name: "Internal Status Page",
		ID:   "1",
	}

	if !reflect.DeepEqual(addon, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", addon, want)
	}
}

func TestAddonsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/addons/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"addon": {"id": "1"}}`))
	})

	addon, _, err := client.Addons.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Addon{
		ID: "1",
	}

	if !reflect.DeepEqual(addon, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", addon, want)
	}
}

func TestAddonsUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Addon{
		Name: "Internal Status Page",
	}

	mux.HandleFunc("/addons/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Addon)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Addon, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		w.Write([]byte(`{"addon": {"name": "Internal Status Page", "id": "1"}}`))
	})

	addon, _, err := client.Addons.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Addon{
		Name: "Internal Status Page",
		ID:   "1",
	}

	if !reflect.DeepEqual(addon, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", addon, want)
	}
}

func TestAddonsDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/addons/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Addons.Delete("1"); err != nil {
		t.Fatal(err)
	}
}
