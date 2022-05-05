package pagerduty

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEventOrchestrationPathRouterPathGet(t *testing.T) {
	setup()
	defer teardown()

	var url = UrlBuilder("E-ORC-1", PathTypeRouter)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"type": "router",
			"parent": {
				"id": "E-ORC-1",
				"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
				"type": "event_orchestration_reference"
			},
			"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1/router",
			"sets": [
				{
					"id": "start",
					"rules": [
						{ "actions": {}, "conditions": [], "id": "rule-1", "label": null }
					]
				}
			],
			"catch_all": { "actions": {} },
			"created_at": "2022-03-22T16:32:20Z",
			"created_by": null,
			"updated_at": "2022-03-22T16:32:20Z",
			"updated_by": {
				"id": "POVFTKB",
				"self": "https://api.pagerduty.com/users/POVFTKB",
				"type": "user_reference"
			},
			"version": "new_version_1"
		}
		`))
	})

	resp, _, err := client.EventOrchestrationPaths.Get("E-ORC-1", PathTypeRouter)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf(" \n\n%#v", resp)

	want := &EventOrchestrationPath{
		Type: "router",
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
		Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1/router",
		Sets: []*EventOrchestrationPathSet{
			{
				ID: "start",
				Rules: []*EventOrchestrationPathRule{
					{
						ID:         "rule-1",
						Label:      "A first routing rule",
						Conditions: []*EventOrchestrationPathRuleCondition{},
						Actions:    &EventOrchestrationPathRuleActions{},
					},
				},
			},
		},
		CatchAll:  &EventOrchestrationPathCatchAll{},
		CreatedAt: "2022-03-22T16:32:20Z",
		CreatedBy: nil,
		UpdatedAt: "2022-03-22T16:32:20Z",
		UpdatedBy: &EventOrchestrationPathReference{
			ID:   "POVFTKB",
			Type: "user_reference",
			Self: "https://api.pagerduty.com/users/POVFTKB",
		},
		Version: "new_version_1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathUpdate(t *testing.T) {
	setup()
	defer teardown()
	// input := &EventOrchestration{Name: "foo", Description: "bar", Team: &EventOrchestrationObject{ID: "P3ZQXDF"}}
	// var id = "abcd"
	// var url = fmt.Sprintf("%s/%s", eventOrchestrationBaseUrl, id)

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
