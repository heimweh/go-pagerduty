package pagerduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOrchestrationCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &Orchestration{Name: "foo", Description: "bar", Team: &OrchestrationObject{ID: "P3ZQXDF"}}

	mux.HandleFunc(orchestrationBaseUrl, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Orchestration)
		v.Name = "foo"
		v.Description = "bar"
		v.Team = &OrchestrationObject{ID: "P3ZQXDF"}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd"}}`))
	})

	resp, _, err := client.Orchestrations.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Orchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &OrchestrationObject{ID: "P3ZQXDF"},
		ID:          "abcd",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestOrchestrationGet(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/abcd", orchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd"}}`))
	})

	resp, _, err := client.Orchestrations.Get("abcd")

	if err != nil {
		t.Fatal(err)
	}

	want := &Orchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &OrchestrationObject{ID: "P3ZQXDF"},
		ID:          "abcd",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestOrchestrationUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &Orchestration{Name: "foo", Description: "bar", Team: &OrchestrationObject{ID: "P3ZQXDF"}}
	var id = "abcd"
	var url = fmt.Sprintf("%s/%s", orchestrationBaseUrl, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Orchestration)
		v.Name = "foo"
		v.Description = "bar"
		v.Team = &OrchestrationObject{ID: "P3ZQXDF"}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd"}}`))
	})

	resp, _, err := client.Orchestrations.Update(id, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Orchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &OrchestrationObject{ID: "P3ZQXDF"},
		ID:          "abcd",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestOrchestrationDelete(t *testing.T) {
	setup()
	defer teardown()

	var id = "abcd"
	var url = fmt.Sprintf("%s/%s", orchestrationBaseUrl, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Orchestrations.Delete(id); err != nil {
		t.Fatal(err)
	}
}
