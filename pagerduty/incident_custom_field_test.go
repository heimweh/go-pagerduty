package pagerduty

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentCustomFieldList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 0)

		w.Write([]byte(`{"fields":[{"id": "1"}, {"id": "2"}]}`))
	})

	resp, _, err := client.IncidentCustomFields.ListContext(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentCustomFieldResponse{
		Fields: []*IncidentCustomField{
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

func TestIncidentCustomFieldListWithOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		testQueryValue(t, r, "include[]", "field_options")

		w.Write([]byte(`
{
	"fields":[
		{
			"id": "1",
			"field_options": [
            	{
                	"id": "O1"
	            }
    	    ]
		},
		{"id": "2"}
	]
}`))

	})

	resp, _, err := client.IncidentCustomFields.ListContext(context.Background(), &ListIncidentCustomFieldOptions{
		Includes: []string{"field_options"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentCustomFieldResponse{
		Fields: []*IncidentCustomField{
			{
				ID: "1",
				FieldOptions: []*IncidentCustomFieldOption{
					{
						ID: "O1",
					},
				},
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

func TestIncidentCustomFieldGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1", func(w http.ResponseWriter, r *http.Request) {
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
        "data_type": "string",
		"field_type": "single_value_fixed",
        "fixed_options": true,
        "field_options": [
            {
                "id": "O1",
                "type": "field_option",
                "data": {
                    "data_type": "string",
                    "value": "abc"
                },
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

	resp, _, err := client.IncidentCustomFields.GetContext(context.Background(), "F1", &GetIncidentCustomFieldOptions{
		Includes: []string{"field_options"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:          "F1",
		Type:        "field",
		DataType:    IncidentCustomFieldDataTypeString,
		FieldType:   IncidentCustomFieldFieldTypeSingleValueFixed,
		Name:        "environment",
		DisplayName: "Environment",
		FieldOptions: []*IncidentCustomFieldOption{
			{
				ID:   "O1",
				Type: "field_option",
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "abc",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"string","field_type":"single_value"}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "string",
		"field_type": "single_value"
    }
}
`))

	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:      "environment",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeSingleValue,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:        "F1",
		Type:      "field",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeSingleValue,
		Name:      "environment",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateFixedOptionsWithStringOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"string","field_type":"single_value_fixed","field_options":[{"data":{"data_type":"string","value":"Chicago"}}]}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "string",
		"field_type": "single_value_fixed",
		"field_options": [
			{
                "id": "FO1",
                "type": "field_option",
                "data": {
                    "data_type": "string",
                    "value": "Chicago"
                }
            }
		]
    }
}
`))

	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:      "environment",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeSingleValueFixed,
		FieldOptions: []*IncidentCustomFieldOption{
			{
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "Chicago",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:        "F1",
		Type:      "field",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeSingleValueFixed,
		Name:      "environment",
		FieldOptions: []*IncidentCustomFieldOption{
			{
				ID:   "FO1",
				Type: "field_option",
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "Chicago",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateMultiFixedOptionsWithStringOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"string","field_type":"multi_value_fixed","field_options":[{"data":{"data_type":"string","value":"Chicago"}}]}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "string",
		"field_type": "multi_value_fixed",
		"field_options": [
			{
                "id": "FO1",
                "type": "field_option",
                "data": {
                    "data_type": "string",
                    "value": "Chicago"
                }
            }
		]
    }
}
`))

	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:      "environment",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeMultiValueFixed,
		FieldOptions: []*IncidentCustomFieldOption{
			{
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "Chicago",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:        "F1",
		Type:      "field",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeMultiValueFixed,
		Name:      "environment",
		FieldOptions: []*IncidentCustomFieldOption{
			{
				ID:   "FO1",
				Type: "field_option",
				Data: &IncidentCustomFieldOptionData{
					DataType: IncidentCustomFieldDataTypeString,
					Value:    "Chicago",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateWithDefaultString(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"string","field_type":"single_value","default_value":"prod"}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "string",
		"field_type": "single_value",
		"default_value": "prod"
    }
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:         "environment",
		DataType:     IncidentCustomFieldDataTypeString,
		FieldType:    IncidentCustomFieldFieldTypeSingleValue,
		DefaultValue: "prod",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     IncidentCustomFieldDataTypeString,
		FieldType:    IncidentCustomFieldFieldTypeSingleValue,
		Name:         "environment",
		DefaultValue: "prod",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateWithDefaultStringArray(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"string","field_type":"multi_value","default_value":["dev","prod"]}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "string",
		"field_type": "multi_value",
		"default_value": ["dev","prod"]
    }
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:         "environment",
		DataType:     IncidentCustomFieldDataTypeString,
		FieldType:    IncidentCustomFieldFieldTypeMultiValue,
		DefaultValue: []string{"dev", "prod"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     IncidentCustomFieldDataTypeString,
		FieldType:    IncidentCustomFieldFieldTypeMultiValue,
		Name:         "environment",
		DefaultValue: []interface{}{"dev", "prod"},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateWithDefaultInt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"integer","field_type":"single_value","default_value":42}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "integer",
		"field_type": "single_value",
		"default_value": 42
    }
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:         "environment",
		DataType:     IncidentCustomFieldDataTypeInt,
		FieldType:    IncidentCustomFieldFieldTypeSingleValue,
		DefaultValue: 42,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     IncidentCustomFieldDataTypeInt,
		FieldType:    IncidentCustomFieldFieldTypeSingleValue,
		Name:         "environment",
		DefaultValue: int64(42),
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateWithDefaultIntArray(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"integer","field_type":"multi_value","default_value":[42,56]}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "integer",
		"field_type": "multi_value",
		"default_value": [42,56]
    }
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:         "environment",
		DataType:     IncidentCustomFieldDataTypeInt,
		FieldType:    IncidentCustomFieldFieldTypeMultiValue,
		DefaultValue: []int{42, 56},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     IncidentCustomFieldDataTypeInt,
		FieldType:    IncidentCustomFieldFieldTypeMultiValue,
		Name:         "environment",
		DefaultValue: []interface{}{int64(42), int64(56)},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateWithDefaultFloat(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"float","field_type":"single_value","default_value":3.14}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "float",
		"field_type": "single_value",
		"default_value": 3.14
    }
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:         "environment",
		DataType:     IncidentCustomFieldDataTypeFloat,
		FieldType:    IncidentCustomFieldFieldTypeSingleValue,
		DefaultValue: 3.14,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     IncidentCustomFieldDataTypeFloat,
		FieldType:    IncidentCustomFieldFieldTypeSingleValue,
		Name:         "environment",
		DefaultValue: 3.14,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldCreateWithDefaultFloatArray(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field":{"name":"environment","data_type":"float","field_type":"multi_value","default_value":[3.14,5.76]}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment",
        "data_type": "float",
		"field_type": "multi_value",
		"default_value": [3.14, 5.76]
    }
}
`))
	})

	resp, _, err := client.IncidentCustomFields.CreateContext(context.Background(), &IncidentCustomField{
		Name:         "environment",
		DataType:     IncidentCustomFieldDataTypeFloat,
		FieldType:    IncidentCustomFieldFieldTypeMultiValue,
		DefaultValue: []float64{3.14, 5.76},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:           "F1",
		Type:         "field",
		DataType:     IncidentCustomFieldDataTypeFloat,
		FieldType:    IncidentCustomFieldFieldTypeMultiValue,
		Name:         "environment",
		DefaultValue: []interface{}{3.14, 5.76},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentCustomFieldDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	resp, err := client.IncidentCustomFields.DeleteContext(context.Background(), "F1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}
}

func TestIncidentCustomFieldUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incidents/custom_fields/F1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"field":{"id":"F1","name":"environment_edit","data_type":"string","field_type":"single_value"}}`)

		w.Write([]byte(`
{
    "field": {
        "id": "F1",
        "type": "field",
        "name": "environment_edit",
        "data_type": "string",
		"field_type": "single_value"
    }
}
`))

	})

	resp, _, err := client.IncidentCustomFields.UpdateContext(context.Background(), "F1", &IncidentCustomField{
		ID:        "F1",
		Name:      "environment_edit",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeSingleValue,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentCustomField{
		ID:        "F1",
		Type:      "field",
		DataType:  IncidentCustomFieldDataTypeString,
		FieldType: IncidentCustomFieldFieldTypeSingleValue,
		Name:      "environment_edit",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
