package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestExtensionsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/extensions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"extensions": [{"id": "1"}]}`))
	})

	resp, _, err := client.Extensions.List(&ListExtensionsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListExtensionsResponse{
		Extensions: []*Extension{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestExtensionsCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Extension{Name: "foo"}

	mux.HandleFunc("/extensions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(ExtensionPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Extension, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"extension": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Extensions.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Extension{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestExtensionsDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/extensions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Extensions.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestExtensionsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/extensions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"extension": {"id": "1"}}`))
	})

	resp, _, err := client.Extensions.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Extension{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestExtensionsUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Extension{
		Name: "foo",
	}

	mux.HandleFunc("/extensions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.Write([]byte(`{"extension": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Extensions.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Extension{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
