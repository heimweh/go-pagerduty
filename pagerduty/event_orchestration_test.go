package pagerduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEventOrchestrationCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestration{Name: "foo", Description: "bar", Team: &EventOrchestrationObject{ID: "P3ZQXDF"}}

	mux.HandleFunc(eventOrchestrationBaseUrl, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EventOrchestration)
		v.Name = "foo"
		v.Description = "bar"
		v.Team = &EventOrchestrationObject{ID: "P3ZQXDF"}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd","routes": 0, "integrations":[{"id":"9c5ff030-12da-4204-a067-25ee61a8df6c","parameters":{"routing_key":"R02","type":"global"}}]}}`))
	})

	resp, _, err := client.EventOrchestrations.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &EventOrchestrationObject{ID: "P3ZQXDF"},
		ID:          "abcd",
		Routes:      0,
		Integrations: []*EventOrchestrationIntegration{
			{
				ID:         "9c5ff030-12da-4204-a067-25ee61a8df6c",
				Parameters: &EventOrchestrationIntegrationParameters{RoutingKey: "R02", Type: "global"},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationGet(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/abcd", eventOrchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd","routes": 2, "integrations":[{"id":"9c5ff030-12da-4204-a067-25ee61a8df6c","parameters":{"routing_key":"R02","type":"global"}}]}}`))
	})

	resp, _, err := client.EventOrchestrations.Get("abcd")

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &EventOrchestrationObject{ID: "P3ZQXDF"},
		ID:          "abcd",
		Routes:      2,
		Integrations: []*EventOrchestrationIntegration{
			{
				ID:         "9c5ff030-12da-4204-a067-25ee61a8df6c",
				Parameters: &EventOrchestrationIntegrationParameters{RoutingKey: "R02", Type: "global"},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestration{Name: "foo", Description: "bar", Team: &EventOrchestrationObject{ID: "P3ZQXDF"}}
	var id = "abcd"
	var url = fmt.Sprintf("%s/%s", eventOrchestrationBaseUrl, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestration)
		v.Name = "foo"
		v.Description = "bar"
		v.Team = &EventOrchestrationObject{ID: "P3ZQXDF"}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd","routes": 2, "integrations":[{"id":"9c5ff030-12da-4204-a067-25ee61a8df6c","parameters":{"routing_key":"R02","type":"global"}}]}}`))
	})

	resp, _, err := client.EventOrchestrations.Update(id, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestration{
		Name:        "foo",
		Description: "bar",
		Team:        &EventOrchestrationObject{ID: "P3ZQXDF"},
		ID:          "abcd",
		Routes:      2,
		Integrations: []*EventOrchestrationIntegration{
			{
				ID:         "9c5ff030-12da-4204-a067-25ee61a8df6c",
				Parameters: &EventOrchestrationIntegrationParameters{RoutingKey: "R02", Type: "global"},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationDelete(t *testing.T) {
	setup()
	defer teardown()

	var id = "abcd"
	var url = fmt.Sprintf("%s/%s", eventOrchestrationBaseUrl, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.EventOrchestrations.Delete(id); err != nil {
		t.Fatal(err)
	}
}
