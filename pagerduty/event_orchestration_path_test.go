package pagerduty

import (
	"encoding/json"
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

func TestEventOrchestrationPathGetServiceActiveStatus(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/services/POOPBUG/active", eventOrchestrationBaseUrl)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
      "active": false
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.GetServiceActiveStatus("POOPBUG")

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathServiceActiveStatus{
		Active: false,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathRouterPathUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestrationPath{
		Type: "router",
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
	}

	var url = fmt.Sprintf("%s/E-ORC-1/router", eventOrchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationPath)
		v.Type = "router"
		v.Parent = &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"orchestration_path": { "type": "router", "parent": { "id": "E-ORC-1", "self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1", "type": "event_orchestration_reference" }, "sets": [ { "id": "start", "rules": [ { "actions": { "route_to": "P3ZQXDF" }, "conditions": [ { "expression": "event.summary matches part 'orca'" }, { "expression": "event.summary matches part 'humpback'" } ], "id": "E-ORC-RULE-1"}]}]}}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Update("E-ORC-1", PathTypeRouter, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathPayload{
		OrchestrationPath: &EventOrchestrationPath{
			Type: "router",
			Parent: &EventOrchestrationPathReference{
				ID:   "E-ORC-1",
				Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
				Type: "event_orchestration_reference",
			},
			Sets: []*EventOrchestrationPathSet{
				{
					ID: "start",
					Rules: []*EventOrchestrationPathRule{
						{
							Actions: &EventOrchestrationPathRuleActions{
								RouteTo: "P3ZQXDF",
							},
							Conditions: []*EventOrchestrationPathRuleCondition{
								{
									Expression: "event.summary matches part 'orca'",
								},
								{
									Expression: "event.summary matches part 'humpback'",
								},
							},
							ID: "E-ORC-RULE-1",
						},
					},
				},
			},
		},
		Warnings: nil,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathUnroutedPathUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestrationPath{
		Type: "unrouted",
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
	}

	var url = fmt.Sprintf("%s/E-ORC-1/unrouted", eventOrchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationPath)
		v.Type = "unrouted"
		v.Parent = &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"orchestration_path": { "type": "unrouted", "parent": { "id": "E-ORC-1", "self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1", "type": "event_orchestration_reference" }, "sets": [ { "id": "start", "rules": [ { "actions": { "route_to": "P3ZQXDF" }, "conditions": [ { "expression": "event.summary matches part 'orca'" }, { "expression": "event.summary matches part 'humpback'" } ], "id": "E-ORC-RULE-1"}]}]},
			"warnings": [
				{"feature": "variables", "feature_type": "actions", "message": "Message 1", "rule_id": "abcd001", "warning_type": "forbidden_feature"},
				{"feature": "extractions", "feature_type": "actions", "message": "Message 2", "rule_id": "abcd002", "warning_type": "forbidden_feature"}
			]
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Update("E-ORC-1", PathTypeUnrouted, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathPayload{
		OrchestrationPath: &EventOrchestrationPath{
			Type: "unrouted",
			Parent: &EventOrchestrationPathReference{
				ID:   "E-ORC-1",
				Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
				Type: "event_orchestration_reference",
			},
			Sets: []*EventOrchestrationPathSet{
				{
					ID: "start",
					Rules: []*EventOrchestrationPathRule{
						{
							Actions: &EventOrchestrationPathRuleActions{
								RouteTo: "P3ZQXDF",
							},
							Conditions: []*EventOrchestrationPathRuleCondition{
								{
									Expression: "event.summary matches part 'orca'",
								},
								{
									Expression: "event.summary matches part 'humpback'",
								},
							},
							ID: "E-ORC-RULE-1",
						},
					},
				},
			},
		},
		Warnings: []*EventOrchestrationPathWarning{
			&EventOrchestrationPathWarning{
				Feature: "variables",
				FeatureType: "actions",
				Message: "Message 1",
				RuleId: "abcd001",
				WarningType: "forbidden_feature",
			},
			&EventOrchestrationPathWarning{
				Feature: "extractions",
				FeatureType: "actions",
				Message: "Message 2",
				RuleId: "abcd002",
				WarningType: "forbidden_feature",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathServiceActiveStatusUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestrationPathServiceActiveStatus{Active: false}

	var url = fmt.Sprintf("%s/services/P3ZQXDF/active", eventOrchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationPathServiceActiveStatus)
		v.Active = false
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"active":false}`))
	})

	resp, _, err := client.EventOrchestrationPaths.UpdateServiceActiveStatus("P3ZQXDF", input.Active)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathServiceActiveStatus{Active: false}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathServicePathUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestrationPath{
		Type: "service",
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
	}

	var url = fmt.Sprintf("%s/services/P3ZQXDF", eventOrchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationPath)
		v.Type = "service"
		v.Parent = &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"orchestration_path": { "type": "service", "parent": { "id": "E-ORC-1", "self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1", "type": "event_orchestration_reference" }, "sets": [ { "id": "start", "rules": [ { "actions": { "route_to": "P3ZQXDF" }, "conditions": [ { "expression": "event.summary matches part 'orca'" }, { "expression": "event.summary matches part 'humpback'" } ], "id": "E-ORC-RULE-1"}]}]},
			"warnings": []
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Update("P3ZQXDF", PathTypeService, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathPayload{
		OrchestrationPath: &EventOrchestrationPath{
			Type: "service",
			Parent: &EventOrchestrationPathReference{
				ID:   "E-ORC-1",
				Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
				Type: "event_orchestration_reference",
			},
			Sets: []*EventOrchestrationPathSet{
				{
					ID: "start",
					Rules: []*EventOrchestrationPathRule{
						{
							Actions: &EventOrchestrationPathRuleActions{
								RouteTo: "P3ZQXDF",
							},
							Conditions: []*EventOrchestrationPathRuleCondition{
								{
									Expression: "event.summary matches part 'orca'",
								},
								{
									Expression: "event.summary matches part 'humpback'",
								},
							},
							ID: "E-ORC-RULE-1",
						},
					},
				},
			},
		},
		Warnings: []*EventOrchestrationPathWarning{},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
