package pagerduty

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIncidentWorkflowTriggerTypeIsKnown(t *testing.T) {
	for k, v := range incidentWorkflowTriggerTypeToString {
		if v == "unknown" {
			if k.IsKnown() {
				t.Errorf("'unknown' data type should not be known")
			}
		} else if !k.IsKnown() {
			t.Errorf("'%s' data type should be known", v)
		}
	}
}

type incidentWorkflowTriggerTypeWrapper struct {
	Type IncidentWorkflowTriggerType `json:"trigger_type"`
}

func TestIncidentWorkflowTriggerTypeMarshalJSON(t *testing.T) {
	for k, v := range incidentWorkflowTriggerTypeToString {
		o := incidentWorkflowTriggerTypeWrapper{Type: k}
		b, _ := json.Marshal(o)
		s := string(b)
		exp := fmt.Sprintf(`{"trigger_type":"%s"}`, v)
		if s != exp {
			t.Errorf(`%s was not marshalled correctly. want:\n%s\ngot:\n%s`, v, exp, s)
		}
	}
}

func TestIncidentWorkflowTriggerTypeUnmarshalJSON(t *testing.T) {
	for k, v := range incidentWorkflowTriggerTypeToString {
		js := fmt.Sprintf(`{"trigger_type":"%s"}`, v)
		var o incidentWorkflowTriggerTypeWrapper
		err := json.Unmarshal([]byte(js), &o)
		if err != nil {
			t.Errorf("Error when unmarhsalling %s", js)
		}
		if o.Type != k {
			t.Errorf(`%s was not unmarshalled correctly. want:\n%s\ngot:\n%s`, js, k, o.Type)
		}
	}
}

func TestIncidentWorkflowTriggerTypeUnmarshalJSON_Error(t *testing.T) {
	js := `{"trigger_type":1234}`
	var o incidentWorkflowTriggerTypeWrapper
	err := json.Unmarshal([]byte(js), &o)
	if err == nil {
		t.Errorf("Unmarshalling %s should have produced an error, but didn't.", js)
	}
}
