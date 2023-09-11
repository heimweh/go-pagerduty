package pagerduty

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEventOrchestrationPathGlobalGet(t *testing.T) {
	setup()
	defer teardown()

	var url = fmt.Sprintf("%s/E-ORC-1/global", eventOrchestrationBaseUrl)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"orchestration_path": {
				"catch_all": {
					"actions": {}
				},
				"created_at": "2022-07-13T20:44:58Z",
				"created_by": {
					"id": "P8B9WR7",
					"self": "https://api.pagerduty.com/users/P8B9WR7",
					"type": "user_reference"
				},
				"parent": {
					"id": "E-ORC-1",
					"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
					"type": "event_orchestration_reference"
				},
				"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1/global",
				"sets": [
					{
						"id": "start",
						"rules": []
					}
				],
				"type": "global",
				"updated_at": "2022-12-15T13:57:08Z",
				"updated_by": {
					"id": "P8B9WR8",
					"self": "https://api.pagerduty.com/users/P8B9WR8",
					"type": "user_reference"
				},
				"version": "Abcd.1234"
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Get("E-ORC-1", PathTypeGlobal)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPath{
		CatchAll: &EventOrchestrationPathCatchAll{
			Actions: &EventOrchestrationPathRuleActions{},
		},
		CreatedAt: "2022-07-13T20:44:58Z",
		CreatedBy: &EventOrchestrationPathReference{
			ID:   "P8B9WR7",
			Self: "https://api.pagerduty.com/users/P8B9WR7",
			Type: "user_reference",
		},
		Parent: &EventOrchestrationPathReference{
			ID:   "E-ORC-1",
			Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
			Type: "event_orchestration_reference",
		},
		Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1/global",
		Sets: []*EventOrchestrationPathSet{
			{
				ID:    "start",
				Rules: []*EventOrchestrationPathRule{},
			},
		},
		Type:      "global",
		UpdatedAt: "2022-12-15T13:57:08Z",
		UpdatedBy: &EventOrchestrationPathReference{
			ID:   "P8B9WR8",
			Self: "https://api.pagerduty.com/users/P8B9WR8",
			Type: "user_reference",
		},
		Version: "Abcd.1234",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathRouterGet(t *testing.T) {
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

func TestEventOrchestrationPathUnroutedGet(t *testing.T) {
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

func TestEventOrchestrationPathServiceGet(t *testing.T) {
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

	resp, _, err := client.EventOrchestrationPaths.GetServiceActiveStatusContext(context.Background(), "POOPBUG")

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

func TestEventOrchestrationPathGlobalUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &EventOrchestrationPath{
		CatchAll: &EventOrchestrationPathCatchAll{
			Actions: &EventOrchestrationPathRuleActions{
				Suppress: true,
			},
		},
		Sets: []*EventOrchestrationPathSet{
			{
				ID: "start",
				Rules: []*EventOrchestrationPathRule{
					{
						Actions: &EventOrchestrationPathRuleActions{
							DropEvent: true,
						},
						Conditions: []*EventOrchestrationPathRuleCondition{
							{
								Expression: "event.summary matches part '[TEST]'",
							},
						},
						ID:    "240790f7",
						Label: "drop test events",
					},
					{
						Actions: &EventOrchestrationPathRuleActions{
							AutomationActions: []*EventOrchestrationPathAutomationAction{
								{
									AutoSend: true,
									Headers: []*EventOrchestrationPathAutomationActionObject{
										{
											Key:   "x-header-1",
											Value: "h-one",
										},
										{
											Key:   "x-header-2",
											Value: "h-two",
										},
									},
									Name: "test webhook",
									Parameters: []*EventOrchestrationPathAutomationActionObject{
										{
											Key:   "hostname",
											Value: "{{variables.hostname}}",
										},
										{
											Key:   "info",
											Value: "{{event.summary}}",
										},
									},
									Url: "https://test.com",
								},
							},
							EventAction: "trigger",
							Extractions: []*EventOrchestrationPathActionExtractions{
								{
									Target:   "event.summary",
									Template: "{{event.summary}}, hostname: {{variables.hostname}}",
								},
							},
							Priority: "PCMUB6F",
							RouteTo:  "7589a1b9",
							Severity: "warning",
							Variables: []*EventOrchestrationPathActionVariables{
								{
									Name:  "hostname",
									Path:  "event.source",
									Type:  "regex",
									Value: ".*",
								},
							},
						},
						Conditions: []*EventOrchestrationPathRuleCondition{
							{
								Expression: "now in Mon,Tue,Wed,Thu,Fri 08:00:00 to 18:00:00 America/Los_Angeles",
							},
						},
						ID:    "4ad2c1be",
						Label: "business hours",
					},
					{
						Actions: &EventOrchestrationPathRuleActions{
							IncidentCustomFieldActions: []*EventOrchestrationPathIncidentCustomFieldAction{
								{
									ID:    "PN1C4A2",
									Value: "{{event.timestamp}}",
								},
							},
						},
						Label: "Set Impact Start custom field from event",
						ID:   "yu3bv02m",
						Conditions: []*EventOrchestrationPathRuleCondition{},
					},
				},
			},
			{
				ID: "7589a1b9",
				Rules: []*EventOrchestrationPathRule{
					{
						Actions: &EventOrchestrationPathRuleActions{
							Annotate: "received more than 10 events over 5 minutes",
						},
						Conditions: []*EventOrchestrationPathRuleCondition{
							{
								Expression: "trigger_count over 5 minute > 10",
							},
						},
						ID:    "fed68019",
						Label: "too many events",
					},
				},
			},
		},
	}

	var url = fmt.Sprintf("%s/E-ORC-1/global", eventOrchestrationBaseUrl)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationPathPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.OrchestrationPath, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"orchestration_path": {
				"catch_all": {"actions": {"suppress": true}},
				"created_at": "2022-07-13T20:44:58Z",
				"created_by": {"id": "P8B9WR7", "self": "https://api.pagerduty.com/users/P8B9WR7", "type": "user_reference"},
				"parent": {
					"id": "E-ORC-1",
					"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
					"type": "event_orchestration_reference"
				},
				"self": "https://api.pagerduty.com/event_orchestrations/E-ORC-1/global",
				"sets": [
					{
						"id": "start",
						"rules": [
							{
								"actions": {"drop_event": true},
								"conditions": [{"expression": "event.summary matches part '[TEST]'"}],
								"id": "240790f7",
								"label": "drop test events"
							},
							{
								"actions": {
									"automation_actions": [
										{
											"auto_send": true,
											"headers": [{"key": "x-header-1", "value": "h-one"}, {"key": "x-header-2","value": "h-two"}],
											"name": "test webhook",
											"parameters": [{"key": "hostname", "value": "{{variables.hostname}}"}, {"key": "info", "value": "{{event.summary}}"}],
											"url": "https://test.com"
										}
									],
									"event_action": "trigger",
									"extractions": [
										{
											"regex": null, "source": null, "target": "event.summary", "template": "{{event.summary}}, hostname: {{variables.hostname}}"
										}
									],
									"priority": "PCMUB6F",
									"route_to": "7589a1b9",
									"severity": "warning",
									"variables": [{"name": "hostname", "path": "event.source", "type": "regex", "value": ".*"}]
								},
								"conditions": [{"expression": "now in Mon,Tue,Wed,Thu,Fri 08:00:00 to 18:00:00 America/Los_Angeles"}],
								"id": "4ad2c1be",
								"label": "business hours"
							},
							{
								"label": "Set Impact Start custom field from event",
								"id": "yu3bv02m",
								"conditions": [],
								"actions": {
								"incident_custom_field_updates": [
									{
									"id": "PN1C4A2",
									"value": "{{event.timestamp}}"
									}
								]
								}
							}
						]
					},
					{
						"id": "7589a1b9",
						"rules": [
							{
								"actions": {"annotate": "received more than 10 events over 5 minutes"},
								"conditions": [{"expression": "trigger_count over 5 minute > 10"}],
								"id": "fed68019",
								"label": "too many events"
							}
						]
					}
				],
				"type": "global",
				"updated_at": "2023-01-26T16:03:55Z",
				"updated_by": {"id": "P8B9WR8", "self": "https://api.pagerduty.com/users/P8B9WR8", "type": "user_reference"},
				"version": "AbCdE.5"
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationPaths.Update("E-ORC-1", PathTypeGlobal, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathPayload{
		OrchestrationPath: &EventOrchestrationPath{
			CatchAll:  input.CatchAll,
			CreatedAt: "2022-07-13T20:44:58Z",
			CreatedBy: &EventOrchestrationPathReference{
				ID:   "P8B9WR7",
				Self: "https://api.pagerduty.com/users/P8B9WR7",
				Type: "user_reference",
			},
			Parent: &EventOrchestrationPathReference{
				ID:   "E-ORC-1",
				Self: "https://api.pagerduty.com/event_orchestrations/E-ORC-1",
				Type: "event_orchestration_reference",
			},
			Self:      "https://api.pagerduty.com/event_orchestrations/E-ORC-1/global",
			Sets:      input.Sets,
			Type:      "global",
			UpdatedAt: "2023-01-26T16:03:55Z",
			UpdatedBy: &EventOrchestrationPathReference{
				ID:   "P8B9WR8",
				Self: "https://api.pagerduty.com/users/P8B9WR8",
				Type: "user_reference",
			},
			Version: "AbCdE.5",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathRouterUpdate(t *testing.T) {
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
		v := new(EventOrchestrationPathPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.OrchestrationPath, input) {
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

func TestEventOrchestrationPathUnroutedUpdate(t *testing.T) {
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
		v := new(EventOrchestrationPathPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.OrchestrationPath, input) {
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
			{
				Feature:     "variables",
				FeatureType: "actions",
				Message:     "Message 1",
				RuleId:      "abcd001",
				WarningType: "forbidden_feature",
			},
			{
				Feature:     "extractions",
				FeatureType: "actions",
				Message:     "Message 2",
				RuleId:      "abcd002",
				WarningType: "forbidden_feature",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationPathServiceUpdate(t *testing.T) {
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
		v := new(EventOrchestrationPathPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.OrchestrationPath, input) {
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

	resp, _, err := client.EventOrchestrationPaths.UpdateServiceActiveStatusContext(context.Background(), "P3ZQXDF", input.Active)
	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationPathServiceActiveStatus{Active: false}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
