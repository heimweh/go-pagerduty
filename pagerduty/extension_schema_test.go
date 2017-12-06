package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestExtensionSchemasList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/extension_schemas", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"extension_schemas": [{"id": "1"}]}`))
	})

	resp, _, err := client.ExtensionSchemas.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListExtensionSchemasResponse{
		ExtensionSchemas: []*ExtensionSchema{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestExtensionSchemasGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/extension_schemas/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"extension_schema": {"id": "1"}}`))
	})

	resp, _, err := client.ExtensionSchemas.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &ExtensionSchema{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
