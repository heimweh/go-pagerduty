package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCustomFieldDeleteFieldOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1/field_options/O1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	resp, err := client.CustomFields.DeleteFieldOption("F1", "O1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}

}

func TestCustomFieldCreateFieldOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1/field_options", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_option":{"data":{"datatype":"string","value":"dev"}}}`)
		w.Write([]byte(`
{
    "field_option": {
        "id": "O1",
        "type": "field_option",
        "data": {
			"datatype": "string",
			"value": "dev"
		}
	}
}
`))
	})

	resp, _, err := client.CustomFields.CreateFieldOption("F1", &CustomFieldOption{
		Data: &CustomFieldOptionData{
			DataType: CustomFieldDataTypeString,
			Value:    "dev",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldOption{
		ID:   "O1",
		Type: "field_option",
		Data: &CustomFieldOptionData{
			DataType: CustomFieldDataTypeString,
			Value:    "dev",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldOptionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1/field_options/FO1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.Write([]byte(`
{
    "field_option": {
        "id": "FO1",
        "type": "field_option",
		"data": {
            "datatype": "string",
            "value": "abc"
        },
        "created_at": "2021-06-01T21:30:42Z",
        "updated_at": "2021-07-01T21:30:42Z"
    }
}
`))

	})

	resp, _, err := client.CustomFields.GetFieldOption("F1", "FO1")
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldOption{
		ID:   "FO1",
		Type: "field_option",
		Data: &CustomFieldOptionData{
			DataType: CustomFieldDataTypeString,
			Value:    "abc",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldOptionList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1/field_options", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`
{
  "field_options": [
    {
      "data": {
        "datatype": "string",
        "value": "Chicago"
      },
      "id": "FO1",
      "type": "field_option"
    },
    {
      "data": {
        "datatype": "string",
        "value": "San Francisco"
      },
      "id": "FO2",
      "type": "field_option"
    }
  ]
}
`))

	})

	resp, _, err := client.CustomFields.ListFieldOptions("F1")
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldOptionsResponse{
		FieldOptions: []*CustomFieldOption{
			{
				ID:   "FO1",
				Type: "field_option",
				Data: &CustomFieldOptionData{
					DataType: CustomFieldDataTypeString,
					Value:    "Chicago",
				},
			},
			{
				ID:   "FO2",
				Type: "field_option",
				Data: &CustomFieldOptionData{
					DataType: CustomFieldDataTypeString,
					Value:    "San Francisco",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldOptionUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/fields/F1/field_options/FO1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"field_option":{"id":"FO1","data":{"datatype":"string","value":"value-edit"}}}`)

		w.Write([]byte(`
{
    "field_option": {
        "id": "FO1",
        "type": "field_option",
        "data": {
        	"datatype": "string",
        	"value": "value-edit"
      	}
    }
}
`))

	})

	resp, _, err := client.CustomFields.UpdateFieldOption("F1", "FO1", &CustomFieldOption{
		ID: "FO1",
		Data: &CustomFieldOptionData{
			DataType: CustomFieldDataTypeString,
			Value:    "value-edit",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldOption{
		ID:   "FO1",
		Type: "field_option",
		Data: &CustomFieldOptionData{
			DataType: CustomFieldDataTypeString,
			Value:    "value-edit",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
