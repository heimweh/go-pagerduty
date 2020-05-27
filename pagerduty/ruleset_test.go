package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestRulesetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/rulesets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 1, "offset": 0, "more": false, "limit": 25, "rulesets":[{"id": "1"}]}`))
	})

	rulesets, _, err := client.Rulesets.List()
	if err != nil {
		t.Fatal(err)
	}

	want := []*Ruleset{
		{
			ID: "1",
		},
	}

	if !reflect.DeepEqual(rulesets, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", rulesets, want)
	}
}

func TestRulesetCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &Ruleset{Name: "foo"}

	mux.HandleFunc("/rulesets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Ruleset)
		v.Name = "foo"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"ruleset":{"name": "foo", "id":"1"}}`))
	})

	resp, _, err := client.Rulesets.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Ruleset{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
func TestRulesetGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/rulesets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"ruleset":{"name": "foo", "id":"1"}}`))
	})

	ID := "1"
	resp, _, err := client.Rulesets.Get(ID)

	if err != nil {
		t.Fatal(err)
	}

	want := &Ruleset{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestRulesetUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &Ruleset{
		Name: "foo",
	}

	mux.HandleFunc("/rulesets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Ruleset)
		v.Name = "foo"

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"ruleset":{"name": "foo", "id":"1"}}`))
	})

	resp, _, err := client.Rulesets.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Ruleset{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestRulesetDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/rulesets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Rulesets.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

// RulesListTest
func TestRulesetRulesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/rulesets/1/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 1, "offset": 0, "more": false, "limit": 25, "rules":[{"id": "1"}]}`))
	})

	rulesetID := "1"
	resp, _, err := client.Rulesets.ListRules(rulesetID)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListRulesetRulesResponse{
		Total:  1,
		Offset: 0,
		More:   false,
		Limit:  25,
		Rules: []*RulesetRule{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// CreateRuleTest
func TestRulesetRuleCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &RulesetRule{}
	ra := RuleActions{}
	input.Actions = &ra

	mux.HandleFunc("/rulesets/1/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(RulesetRule)
		v.Actions = &RuleActions{}

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"rule":{"id": "1"}}`))
	})
	rulesetID := "1"
	resp, _, err := client.Rulesets.CreateRule(rulesetID, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &RulesetRule{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// UpdateRuleTest
func TestRulesetRuleUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &RulesetRule{}
	ra := RuleActions{}
	input.Actions = &ra

	mux.HandleFunc("/rulesets/1/rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(RulesetRule)
		v.Actions = &RuleActions{}

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"rule":{"id": "1"}}`))
	})
	rulesetID := "1"
	ruleID := "1"

	resp, _, err := client.Rulesets.UpdateRule(rulesetID, ruleID, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &RulesetRule{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// GetRuleTest
func TestRulesetRuleGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/rulesets/1/rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"rule":{"id": "1"}}`))
	})
	rulesetID := "1"
	ruleID := "1"

	resp, _, err := client.Rulesets.GetRule(rulesetID, ruleID)
	if err != nil {
		t.Fatal(err)
	}

	want := &RulesetRule{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// DeleteRuleTest
func TestRulesetRuleDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/rulesets/1/rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Rulesets.DeleteRule("1", "1"); err != nil {
		t.Fatal(err)
	}
}
