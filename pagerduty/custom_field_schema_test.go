package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCustomFieldSchemaList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		offset := r.URL.Query().Get("offset")
		switch offset {
		case "":
			w.Write([]byte(`{"total": 2, "offset": 0, "more": true, "limit": 1, "schemas":[{"id": "1"}]}`))
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "schemas":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.CustomFieldSchemas.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		Schemas: []*CustomFieldSchema{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaListSecondPage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		offset := r.URL.Query().Get("offset")
		switch offset {
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "schemas":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.CustomFieldSchemas.List(&ListCustomFieldSchemaOptions{Offset: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		Schemas: []*CustomFieldSchema{
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testBody(t, r, "")

		w.Write([]byte(`
{
    "schema": {
        "id": "S1",
        "type": "field_schema",
        "title": "Some title",
		"description": "Some description",
        "created_at": "2021-06-01T21:30:42Z",
        "updated_at": "2021-07-01T21:30:42Z"
    }
}
`))

	})

	resp, _, err := client.CustomFieldSchemas.Get("S1", nil)
	if err != nil {
		t.Fatal(err)
	}

	desc := "Some description"

	want := &CustomFieldSchema{
		ID:          "S1",
		Type:        "field_schema",
		Title:       "Some title",
		Description: &desc,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"schema":{"title":"Some title","description":"Some description"}}`)

		w.Write([]byte(`
{
    "schema": {
        "id": "S1",
        "type": "field_schema",
        "title": "Some title",
        "description": "Some description"
    }
}
`))

	})

	desc := "Some description"

	resp, _, err := client.CustomFieldSchemas.Create(&CustomFieldSchema{
		Title:       "Some title",
		Description: &desc,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchema{
		ID:          "S1",
		Type:        "field_schema",
		Title:       "Some title",
		Description: &desc,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"schema":{"id":"S1","title":"Some new title","description":"Some new description"}}`)

		w.Write([]byte(`
{
    "schema": {
        "id": "S1",
        "type": "field_schema",
        "title": "Some new title",
        "description": "Some new description"
    }
}
`))

	})

	desc := "Some new description"

	resp, _, err := client.CustomFieldSchemas.Update("S1", &CustomFieldSchema{
		ID:          "S1",
		Title:       "Some new title",
		Description: &desc,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchema{
		ID:          "S1",
		Type:        "field_schema",
		Title:       "Some new title",
		Description: &desc,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	resp, err := client.CustomFieldSchemas.Delete("S1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}
}
