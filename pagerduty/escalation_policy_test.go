package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestEscalationPoliciesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"escalation_policies": [{"id": "1"}]}`))
	})

	resp, _, err := client.EscalationPolicies.List(&ListEscalationPoliciesOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListEscalationPoliciesResponse{
		EscalationPolicies: []*EscalationPolicy{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEscalationPoliciesCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &EscalationPolicy{Name: "foo"}

	mux.HandleFunc("/escalation_policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EscalationPolicy)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.EscalationPolicy, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"escalation_policy": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.EscalationPolicies.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EscalationPolicy{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEscalationPoliciesDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.EscalationPolicies.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestEscalationPoliciesGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"escalation_policy": {"id": "1"}}`))
	})

	resp, _, err := client.EscalationPolicies.Get("1", &GetEscalationPolicyOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &EscalationPolicy{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEscalationPoliciesUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &EscalationPolicy{
		Name: "foo",
	}

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.Write([]byte(`{"escalation_policy": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.EscalationPolicies.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EscalationPolicy{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
