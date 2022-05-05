package pagerduty

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEventOrchestrationPathGetRouterPath(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/E-ORC-1/router", eventOrchestrationBaseUrl)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"orchestration_path": {
				"type": "router",
				"parent": {
					"id": "E-ORC-1",
					"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
					"type": "event_orchestration_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Get("E-ORC-1", PathTypeRouter)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPath{
		Type: "router",
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
	}

	if !reflect.DeepEqual(resp.Type, want.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

	if !reflect.DeepEqual(resp.Parent, want.Parent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathGetUnroutedPath(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/E-ORC-1/unrouted", eventOrchestrationBaseUrl)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"orchestration_path": {
				"type": "unrouted",
				"parent": {
					"id": "E-ORC-1",
					"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
					"type": "event_orchestration_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Get("E-ORC-1", PathTypeUnrouted)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPath{
		Type: "unrouted",
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
	}

	if !reflect.DeepEqual(resp.Type, want.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

	if !reflect.DeepEqual(resp.Parent, want.Parent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathGetServicePath(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/services/POOPBUG", eventOrchestrationBaseUrl)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"orchestration_path": {
				"type": "service",
				"parent": {
					"id": "POOPBUG",
					"self": "https://api.pagerduty.com/service/POOPBUG",
					"type": "service_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Get("POOPBUG", PathTypeService)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPath{
		Type: "service",
		Parent: &EventOrchestrationPathReference{
			ID:   "POOPBUG",
			Self: "https://api.pagerduty.com/service/POOPBUG",
			Type: "service_reference",
		},
	}

	if !reflect.DeepEqual(resp.Type, want.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

	if !reflect.DeepEqual(resp.Parent, want.Parent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathRouterPathUpdate(t *testing.T) {
	setup()
	defer teardown()
	// input := &EventOrchestration{Name: "foo", Description: "bar", Team: &EventOrchestrationObject{ID: "P3ZQXDF"}}
	// var id = "abcd"
	// var url = fmt.Sprintf("%s/E-ORC-1/router", eventOrchestrationBaseUrl)

	// mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
	// 	testMethod(t, r, "PUT")
	// 	v := new(EventOrchestration)
	// 	v.Name = "foo"
	// 	v.Description = "bar"
	// 	v.Team = &EventOrchestrationObject{ID: "P3ZQXDF"}
	// 	json.NewDecoder(r.Body).Decode(v)
	// 	if !reflect.DeepEqual(v, input) {
	// 		t.Errorf("Request body = %+v, want %+v", v, input)
	// 	}
	// 	w.Write([]byte(`{"orchestration":{"name": "foo", "description": "bar", "team": {"id": "P3ZQXDF"}, "id": "abcd","routes": 2, "integrations":[{"id":"9c5ff030-12da-4204-a067-25ee61a8df6c","parameters":{"routing_key":"R02","type":"global"}}]}}`))
	// })

	// resp, _, err := client.EventOrchestrations.Update(id, input)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// want := &EventOrchestration{
	// 	Name:        "foo",
	// 	Description: "bar",
	// 	Team:        &EventOrchestrationObject{ID: "P3ZQXDF"},
	// 	ID:          "abcd",
	// 	Routes:      2,
	// 	Integrations: []*EventOrchestrationIntegration{
	// 		{
	// 			ID:         "9c5ff030-12da-4204-a067-25ee61a8df6c",
	// 			Parameters: &EventOrchestrationIntegrationParameters{RoutingKey: "R02", Type: "global"},
	// 		},
	// 	},
	// }

	// if !reflect.DeepEqual(resp, want) {
	// 	t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	// }
}
