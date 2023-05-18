package pagerduty

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIncidentCustomFieldDataTypeIsKnown(t *testing.T) {
	for k, v := range incidentCustomFieldDataTypeToString {
		if v == "unknown" {
			if k.IsKnown() {
				t.Errorf("'unknown' data type should not be known")
			}
		} else if !k.IsKnown() {
			t.Errorf("'%s' data type should be known", v)
		}
	}
}

type incidentCustomFieldDataTypeWrapper struct {
	DataType IncidentCustomFieldDataType `json:"data_type"`
}

func TestIncidentCustomFieldDataTypeMarshalJSON(t *testing.T) {
	for k, v := range incidentCustomFieldDataTypeToString {
		o := incidentCustomFieldDataTypeWrapper{DataType: k}
		b, _ := json.Marshal(o)
		s := string(b)
		exp := fmt.Sprintf(`{"data_type":"%s"}`, v)
		if s != exp {
			t.Errorf(`%s was not marshalled correctly. want:\n%s\ngot:\n%s`, v, exp, s)
		}
	}
}

func TestIncidentCustomFieldDataTypeUnmarshalJSON(t *testing.T) {
	for k, v := range incidentCustomFieldDataTypeToString {
		js := fmt.Sprintf(`{"data_type":"%s"}`, v)
		var o incidentCustomFieldDataTypeWrapper
		err := json.Unmarshal([]byte(js), &o)
		if err != nil {
			t.Errorf("Error when unmarhsalling %s", js)
		}
		if o.DataType != k {
			t.Errorf(`%s was not unmarshalled correctly. want:\n%s\ngot:\n%s`, js, k, o.DataType)
		}
	}
}

func TestIncidentCustomFieldDataTypeUnmarshalJSON_Error(t *testing.T) {
	js := `{"data_type":1234}`
	var o incidentCustomFieldDataTypeWrapper
	err := json.Unmarshal([]byte(js), &o)
	if err == nil {
		t.Errorf("Unmarshalling %s should have produced an error, but didn't.", js)
	}
}
