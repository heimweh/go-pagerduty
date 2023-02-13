package pagerduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCustomFieldSchemaFieldConfigurationList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 0)

		w.Write([]byte(`{"field_configurations":[{"id": "1"}]}`))

	})

	resp, _, err := client.CustomFieldSchemas.ListFieldConfigurations("S1", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaFieldConfigurationsResponse{
		FieldConfigurations: []*CustomFieldSchemaFieldConfiguration{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationListWithOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValue(t, r, "include[]", "fields")

		w.Write([]byte(`{"field_configurations":[{"id": "1"}]}`))

	})

	resp, _, err := client.CustomFieldSchemas.ListFieldConfigurations("S1", &ListCustomFieldSchemaConfigurationsOptions{
		Includes: []string{"fields"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaFieldConfigurationsResponse{
		FieldConfigurations: []*CustomFieldSchemaFieldConfiguration{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations/C1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 0)

		w.Write([]byte(`
{
	"field_configuration": {
		"id": "C1",
		"type": "field_configuration",
		"field": {
			"id": "F1",
			"type": "field_reference"
		},
		"required": false,
		"created_at": "2021-06-01T21:30:42Z",
		"updated_at": "2021-06-01T21:30:42Z"
	}
}`))

	})

	resp, _, err := client.CustomFieldSchemas.GetFieldConfiguration("S1", "C1", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: false,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationGetWithOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations/C1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		q := r.URL.Query()

		includes := q.Get("include[]")
		if includes != "fields" {
			t.Errorf("request URL (%v) did not include fields", r.URL)
		}

		w.Write([]byte(`
{
	"field_configuration": {
		"id": "C1",
		"type": "field_configuration",
		"field": {
			"id": "F1",
			"type": "field",
			"name": "environment",
			"datatype": "string",
			"multi_value": false,
			"fixed_options": false,
			"created_at": "2021-06-01T21:30:42Z",
			"updated_at": "2021-07-01T21:30:42Z"
		},
		"required": false,
		"created_at": "2021-06-01T21:30:42Z",
		"updated_at": "2021-06-01T21:30:42Z"
	}
}`))

	})

	resp, _, err := client.CustomFieldSchemas.GetFieldConfiguration("S1", "C1", &GetCustomFieldSchemaConfigurationsOptions{
		Includes: []string{"fields"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: false,
		Field: &CustomField{
			ID:           "F1",
			Type:         "field",
			Name:         "environment",
			DataType:     CustomFieldDataTypeString,
			MultiValue:   false,
			FixedOptions: false,
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations/C1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	resp, err := client.CustomFieldSchemas.DeleteFieldConfiguration("S1", "C1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":false,"field":{"id":"F1","multi_value":false,"fixed_options":false}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": false
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Field: &CustomField{
			ID: "F1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: false,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultString(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":false},"default_value":{"datatype":"string","multi_value":false,"value":"prod"}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "string",
			"multi_value": false,
			"value": "prod"
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID: "F1",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeString,
			MultiValue: false,
			Value:      "prod",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeString,
			MultiValue: false,
			Value:      "prod",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultStringArray(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":false},"default_value":{"datatype":"string","multi_value":true,"value":["dev","prod"]}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "string",
			"multi_value": true,
			"value": [
				"dev",
				"prod"
			]
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID: "F1",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeString,
			MultiValue: true,
			Value:      []string{"dev", "prod"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeString,
			MultiValue: true,
			Value:      []interface{}{"dev", "prod"},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultInt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":false},"default_value":{"datatype":"integer","multi_value":false,"value":42}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "integer",
			"multi_value": false,
			"value": 42
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID: "F1",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeInt,
			MultiValue: false,
			Value:      42,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeInt,
			MultiValue: false,
			Value:      int64(42),
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultIntArray(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":false},"default_value":{"datatype":"integer","multi_value":true,"value":[42,78]}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "integer",
			"multi_value": true,
			"value": [
				42,
				78
			]
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID: "F1",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeInt,
			MultiValue: true,
			Value:      []int{42, 78},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeInt,
			MultiValue: true,
			Value:      []interface{}{int64(42), int64(78)},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultFloat(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":false},"default_value":{"datatype":"float","multi_value":false,"value":3.14}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "float",
			"multi_value": false,
			"value": 3.14
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID: "F1",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeFloat,
			MultiValue: false,
			Value:      3.14,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeFloat,
			MultiValue: false,
			Value:      3.14,
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultFloatArray(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":false},"default_value":{"datatype":"float","multi_value":true,"value":[3.14,5.78]}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "float",
			"multi_value": true,
			"value": [
				3.14,
				5.78
			]
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID: "F1",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeFloat,
			MultiValue: true,
			Value:      []float64{3.14, 5.78},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeFloat,
			MultiValue: true,
			Value:      []interface{}{3.14, 5.78},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationCreateWithDefaultFieldOption(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"field_configuration":{"required":true,"field":{"id":"F1","multi_value":false,"fixed_options":true},"default_value":{"datatype":"field_option","multi_value":false,"value":{"id":"FO1","type":"field_option_reference"}}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": true,
		"default_value": {
			"datatype": "field_option",
			"multi_value": false,
			"value": {
				"type": "field_option_reference",
				"id": "FO1"
			}
		}
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.CreateFieldConfiguration("S1", &CustomFieldSchemaFieldConfiguration{
		Required: true,
		Field: &CustomField{
			ID:           "F1",
			FixedOptions: true, // this isn't actually important but makes the test body above more sensible
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeFieldOption,
			MultiValue: false,
			Value:      "FO1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: true,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
		DefaultValue: &CustomFieldDefaultValue{
			DataType:   CustomFieldDataTypeFieldOption,
			MultiValue: false,
			Value:      "FO1",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaFieldConfigurationUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schemas/S1/field_configurations/C1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"field_configuration":{"required":false,"field":{"id":"F1","multi_value":false,"fixed_options":false}}}`)

		w.Write([]byte(`
{
    "field_configuration": {
        "id": "C1",
		"type": "field_configuration",
        "field" : {
			"id": "F1",
			"type": "field_reference"
		},
        "required": false
    }
}
`))
	})

	resp, _, err := client.CustomFieldSchemas.UpdateFieldConfiguration("S1", "C1", &CustomFieldSchemaFieldConfiguration{
		Required: false,
		Field: &CustomField{
			ID: "F1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaFieldConfiguration{
		ID:       "C1",
		Type:     "field_configuration",
		Required: false,
		Field: &CustomField{
			ID:   "F1",
			Type: "field_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

type testFieldDefaultValueUnmarshalJSON_ErrorCases struct {
	multiValue bool
	datatype   CustomFieldDataType
	input      string
}

func TestFieldDefaultValueUnmarshalJSON_ErrorCases(t *testing.T) {
	data := []testFieldDefaultValueUnmarshalJSON_ErrorCases{
		{
			multiValue: true,
			datatype:   CustomFieldDataTypeInt,
			input:      "42",
		},
		{
			multiValue: false,
			datatype:   CustomFieldDataTypeInt,
			input:      "[42]",
		},
		{
			multiValue: true,
			datatype:   CustomFieldDataTypeInt,
			input:      "[42,false]",
		},
		{
			multiValue: false,
			datatype:   CustomFieldDataTypeInt,
			input:      "false",
		},
	}

	for _, v := range data {
		js := fmt.Sprintf(`
{
	"datatype":"%s",
	"multi_value":%v,
	"value":%s
}`, v.datatype.String(), v.multiValue, v.input)
		var def CustomFieldDefaultValue
		err := json.Unmarshal([]byte(js), &def)
		if err == nil {
			t.Errorf("Was expecting an error parsing\n%s\nand did not receive one", js)
		}
	}

}
