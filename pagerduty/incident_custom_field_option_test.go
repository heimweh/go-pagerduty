package pagerduty

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentCustomFieldDeleteFieldOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1/field_options/O1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	resp, err := client.IncidentCustomFields.DeleteFieldOptionContext(context.Background(), "F1", "O1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}

}

func TestIncidentCustomFieldCreateFieldOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1/field_options", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_option":{"data":{"data_type":"string","value":"dev"}}}`)
		w.Write([]byte(`
{
    "field_option": {
        "id": "O1",
        "type": "field_option",
        "data": {
			"data_type": "string",
			"value": "dev"
		}
	}
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateFieldOptionContext(context.Background(), "F1", &IncidentCustomFieldOption{
		Data: &IncidentCustomFieldOptionData{
			DataType: IncidentCustomFieldDataTypeString,
			Value:    "dev",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomFieldOption{
		ID:   "O1",
		Type: "field_option",
		Data: &IncidentCustomFieldOptionData{
			DataType: IncidentCustomFieldDataTypeString,
			Value:    "dev",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldOptionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1/field_options", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.Write([]byte(`
{
    "field_options": [{
        "id": "FO1",
        "type": "field_option",
		"data": {
            "data_type": "string",
            "value": "abc"
        },
        "created_at": "2021-06-01T21:30:42Z",
        "updated_at": "2021-07-01T21:30:42Z"
    }]
}
`))

	})

	resp, _, err := client.IncidentCustomFields.GetFieldOptionContext(context.Background(), "F1", "FO1")
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomFieldOption{
		ID:   "FO1",
		Type: "field_option",
		Data: &IncidentCustomFieldOptionData{
			DataType: IncidentCustomFieldDataTypeString,
			Value:    "abc",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldOptionGetNotExisting(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1/field_options", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.Write([]byte(`
{
    "field_options": [{
        "id": "FO1",
        "type": "field_option",
		"data": {
            "data_type": "string",
            "value": "abc"
        },
        "created_at": "2021-06-01T21:30:42Z",
        "updated_at": "2021-07-01T21:30:42Z"
    }]
}
`))

	})

	_, _, err := client.IncidentCustomFields.GetFieldOptionContext(context.Background(), "F1", "FO2")
	if err != nil {
		if err.Error() != "no field option with ID FO2 under field F1 can be found" {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		return
	}

	t.Errorf("was not expecting to receive a field option")
}

func TestIncidentCustomFieldOptionList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1/field_options", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`
{
  "field_options": [
    {
      "data": {
        "data_type": "string",
        "value": "Chicago"
      },
      "id": "FO1",
      "type": "field_option"
    },
    {
      "data": {
        "data_type": "string",
        "value": "San Francisco"
      },
      "id": "FO2",
      "type": "field_option"
    }
  ]
}
`))

	})

	resp, _, err := client.IncidentCustomFields.ListFieldOptionsContext(context.Background(), "F1")
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentCustomFieldOptionsResponse{
		FieldOptions: []*IncidentCustomFieldOption{
			{
				ID:   "FO1",
				Type: "field_option",
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "Chicago",
				},
			},
			{
				ID:   "FO2",
				Type: "field_option",
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "San Francisco",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldOptionUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1/field_options/FO1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"field_option":{"id":"FO1","data":{"data_type":"string","value":"value-edit"}}}`)

		w.Write([]byte(`
{
    "field_option": {
        "id": "FO1",
        "type": "field_option",
        "data": {
        	"data_type": "string",
        	"value": "value-edit"
      	}
    }
}
`))

	})

	resp, _, err := client.IncidentCustomFields.UpdateFieldOptionContext(context.Background(), "F1", "FO1", &IncidentCustomFieldOption{
		ID: "FO1",
		Data: &IncidentCustomFieldOptionData{
			DataType: IncidentCustomFieldDataTypeString,
			Value:    "value-edit",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomFieldOption{
		ID:   "FO1",
		Type: "field_option",
		Data: &IncidentCustomFieldOptionData{
			DataType: IncidentCustomFieldDataTypeString,
			Value:    "value-edit",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
