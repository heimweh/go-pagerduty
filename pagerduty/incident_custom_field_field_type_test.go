package pagerduty

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIncidentCustomFieldFieldTypeIsKnown(t *testing.T) {
	for k, v := range incidentCustomFieldFieldTypeToString {
		if v == "unknown" {
			if k.IsKnown() {
				t.Errorf("'unknown' data type should not be known")
			}
		} else if !k.IsKnown() {
			t.Errorf("'%s' data type should be known", v)
		}
	}
}

type incidentCustomFieldFieldTypeWrapper struct {
	FieldType IncidentCustomFieldFieldType `json:"field_type"`
}

func TestIncidentCustomFieldFieldTypeMarshalJSON(t *testing.T) {
	for k, v := range incidentCustomFieldFieldTypeToString {
		o := incidentCustomFieldFieldTypeWrapper{FieldType: k}
		b, _ := json.Marshal(o)
		s := string(b)
		exp := fmt.Sprintf(`{"field_type":"%s"}`, v)
		if s != exp {
			t.Errorf(`%s was not marshalled correctly. want:\n%s\ngot:\n%s`, v, exp, s)
		}
	}
}

func TestIncidentCustomFieldFieldTypeUnmarshalJSON(t *testing.T) {
	for k, v := range incidentCustomFieldFieldTypeToString {
		js := fmt.Sprintf(`{"field_type":"%s"}`, v)
		var o incidentCustomFieldFieldTypeWrapper
		err := json.Unmarshal([]byte(js), &o)
		if err != nil {
			t.Errorf("Error when unmarhsalling %s", js)
		}
		if o.FieldType != k {
			t.Errorf(`%s was not unmarshalled correctly. want:\n%s\ngot:\n%s`, js, k, o.FieldType)
		}
	}
}

func TestIncidentCustomFieldFieldTypeUnmarshalJSON_Error(t *testing.T) {
	js := `{"field_type":1234}`
	var o incidentCustomFieldFieldTypeWrapper
	err := json.Unmarshal([]byte(js), &o)
	if err == nil {
		t.Errorf("Unmarshalling %s should have produced an error, but didn't.", js)
	}
}
