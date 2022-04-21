package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestOrchestrationCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &Orchestration{Name: "foo", Description: "bar", Team: &OrchestrationObject{ID: "P3ZQXDF"}}

	mux.HandleFunc("/orchestrations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Orchestration)
		v.Name = "foo"
		v.Description = "bar"
		v.Team = &OrchestrationObject{ID: "P3ZQXDF"}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "1"}}`))
	})

	resp, _, err := client.Orchestrations.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Orchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &OrchestrationObject{ID: "P3ZQXDF"},
		ID:          "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
