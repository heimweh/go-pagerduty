package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFieldList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryMaxCount(t, r, 1)
		offset := r.URL.Query().Get("offset")

		switch offset {
		case "":
			w.Write([]byte(`{"total": 2, "offset": 0, "more": true, "limit": 1, "fields":[{"id": "1"}]}`))
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "fields":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.CustomFields.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		Fields: []*CustomField{
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

func TestFieldListSecondPage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		offset := r.URL.Query().Get("offset")

		switch offset {
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "fields":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.CustomFields.List(&ListCustomFieldOptions{Offset: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		Fields: []*CustomField{
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestFieldListWithOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		testQueryMinCount(t, r, 1)
		testQueryMaxCount(t, r, 2)
		testQueryValue(t, r, "include[]", "field_options")

		offset := r.URL.Query().Get("offset")
		switch offset {
		case "":
			w.Write([]byte(`{"total": 2, "offset": 0, "more": true, "limit": 1, "fields":[{"id": "1"}]}`))
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "fields":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.CustomFields.List(&ListCustomFieldOptions{
		Includes: []string{"field_options"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		Fields: []*CustomField{
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

func TestFieldGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		testQueryValue(t, r, "include[]", "field_options")
		testBody(t, r, "")

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "display_name": "Environment",
        "datatype": "string",
        "fixed_options": true,
        "field_options": [
            {
                "id": "O1",
                "type": "field_option",
                "data": {
                    "datatype": "string",
                    "value": "abc"
                },
                "disabled": false,
                "created_at": "2021-06-01T21:30:42Z",
                "updated_at": "2021-07-01T21:30:42Z"
            }
        ],
        "created_at": "2021-06-01T21:30:42Z",
        "updated_at": "2021-07-01T21:30:42Z"
    }
}
`))

	})

	resp, _, err := client.CustomFields.Get("F1", &GetCustomFieldOptions{
		Includes: []string{"field_options"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     CustomFieldDataTypeString,
		FixedOptions: true,
		Name:         "environment",
		DisplayName:  "Environment",
		FieldOptions: []*CustomFieldOption{
			{
				ID:   "O1",
				Type: "field_option",
				Data: &CustomFieldOptionData{
					DataType: CustomFieldDataTypeString,
					Value:    "abc",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestFieldCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","datatype":"string","multi_value":false,"fixed_options":false}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "datatype": "string"
    }
}
`))

	})

	resp, _, err := client.CustomFields.Create(&CustomField{
		Name:     "environment",
		DataType: CustomFieldDataTypeString,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomField{
		ID:       "F1",
		Type:     "field",
		DataType: CustomFieldDataTypeString,
		Name:     "environment",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestFieldDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	resp, err := client.CustomFields.Delete("F1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}
}

func TestFieldUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"field":{"id":"F1","name":"environment_edit","datatype":"string","multi_value":false,"fixed_options":false}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment_edit",
        "datatype": "string"
    }
}
`))

	})

	resp, _, err := client.CustomFields.Update("F1", &CustomField{
		ID:       "F1",
		Name:     "environment_edit",
		DataType: CustomFieldDataTypeString,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomField{
		ID:       "F1",
		Type:     "field",
		DataType: CustomFieldDataTypeString,
		Name:     "environment_edit",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
