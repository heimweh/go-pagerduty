package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAbilitiesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/abilities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"abilities": ["sso"]}`))
	})

	abilities, _, err := client.Abilities.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListAbilitiesResponse{
		Abilities: []string{"sso"},
	}

	if !reflect.DeepEqual(abilities, want) {
		t.Errorf("returned %#v; want %#v", abilities, want)
	}
}

func TestAbilitiesListFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/abilities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusForbidden)
	})

	if _, _, err := client.Abilities.List(); err == nil {
		t.Fatal("expected error; got nil")
	}
}

func TestAbilitiesTestAbility(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/abilities/sso", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Abilities.Test("sso"); err != nil {
		t.Fatal(err)
	}
}

func TestAbilitiesTestAbilityFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/abilities/sso", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusForbidden)
	})

	if _, err := client.Abilities.Test("sso"); err == nil {
		t.Fatal("expected error; got nil")
	}
}
