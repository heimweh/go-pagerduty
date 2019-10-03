package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestEventRuleList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/event_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"external_id": "1", "object_version": "objVersion", "format_version": "2", "rules":[{"id": "1"}]}`))
	})

	resp, _, err := client.EventRules.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListEventRulesResponse{
		ExternalID:    "1",
		ObjectVersion: "objVersion",
		FormatVersion: 2,
		EventRules: []*EventRule{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventRuleCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventRule{Actions: []interface{}{[]interface{}{"route", "P5DTL0K"}}, Condition: []interface{}{"and", []interface{}{"contains", []interface{}{"path", "payload", "source"}, "website"}}}

	mux.HandleFunc("/event_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EventRule)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"actions":[["route","P5DTL0K"]],"condition": ["and",["contains",["path","payload","source"],"website"]]}`))
	})

	resp, _, err := client.EventRules.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventRule{
		Actions:           []interface{}{[]interface{}{"route", "P5DTL0K"}},
		Condition:         []interface{}{"and", []interface{}{"contains", []interface{}{"path", "payload", "source"}, "website"}},
		CatchAll:          false,
		AdvancedCondition: []interface{}(nil),
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventRuleDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/event_rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.EventRules.Delete("1"); err != nil {
		t.Fatal(err)
	}
}
