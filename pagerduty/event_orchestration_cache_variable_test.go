package pagerduty

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGlobalEventOrchestrationCacheVariableList(t *testing.T) {
	setup()
	defer teardown()

	oId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	oType := "global"
	url := fmt.Sprintf("%s/%s/cache_variables/", eventOrchestrationBaseUrl, oId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"cache_variables": [
				{
					"id": "45be0a94-55b2-4691-b285-7a14478f4fe2",
					"name": "example_1",
					"conditions": [],
					"configuration": {
						"type": "trigger_event_count",
						"ttl_seconds": 30
					},
					"created_at": "2024-02-12T14:44:58Z",
					"created_by": {
						"id": "P8B9WR7",
						"self": "https://api.pagerduty.com/users/P8B9WR7",
						"type": "user_reference"
					},
					"updated_at": "2024-02-13T12:41:43Z",
					"updated_by": {
						"id": "P8B9WR8",
						"self": "https://api.pagerduty.com/users/P8B9WR8",
						"type": "user_reference"
					}
				},
				{
					"id": "b5e7eef2-d2f8-46ce-80cc-ad7c3efa6d0a",
					"name": "example_2",
					"conditions": [{"expression": "event.summary matches part '[TEST]'"}],
					"configuration": {
						"type": "recent_value",
						"source": "event.source",
						"regex": ".*"
					},
					"created_at": "2024-02-12T14:44:58Z",
					"created_by": {
						"id": "P8B9WR7",
						"self": "https://api.pagerduty.com/users/P8B9WR7",
						"type": "user_reference"
					},
					"updated_at": "2024-02-13T12:41:43Z",
					"updated_by": {
						"id": "P8B9WR8",
						"self": "https://api.pagerduty.com/users/P8B9WR8",
						"type": "user_reference"
					}
				},
				{
					"id": "fffd20a2-8240-4df3-b996-964968464a3f",
					"name": "example_3",
					"configuration": {
						"type": "external_data",
						"data_type": "boolean",
						"ttl_seconds": 1200
					},
					"created_at": "2025-02-07T18:50:58Z",
					"created_by": {
						"id": "P24PCRC",
						"self": "https://api.pagerduty.com/users/P24PCRC",
						"type": "user_reference"
					},
					"data_endpoint": "https://api.pagerduty.com/event_orchestrations/a64f9c87-6adc-4f89-a64c-2fdd8cba4639/cache_variables/fffd20a2-8240-4df3-b996-964968464a3f/data",
					"updated_at": "2025-02-07T20:00:43Z",
					"updated_by": {
						"id": "P23PCRC",
						"self": "https://api.pagerduty.com/users/P23PCRC",
						"type": "user_reference"
					}
				}
			],
			"total": 3
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.List(context.Background(), oType, oId)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListEventOrchestrationCacheVariablesResponse{
		CacheVariables: []*EventOrchestrationCacheVariable{
			{
				ID:         "45be0a94-55b2-4691-b285-7a14478f4fe2",
				Name:       "example_1",
				Conditions: []*EventOrchestrationCacheVariableCondition{},
				Configuration: &EventOrchestrationCacheVariableConfiguration{
					Type:       "trigger_event_count",
					TTLSeconds: 30,
				},
				CreatedAt: "2024-02-12T14:44:58Z",
				CreatedBy: &UserReference{
					ID:   "P8B9WR7",
					Self: "https://api.pagerduty.com/users/P8B9WR7",
					Type: "user_reference",
				},
				UpdatedAt: "2024-02-13T12:41:43Z",
				UpdatedBy: &UserReference{
					ID:   "P8B9WR8",
					Self: "https://api.pagerduty.com/users/P8B9WR8",
					Type: "user_reference",
				},
			},
			{
				ID:   "b5e7eef2-d2f8-46ce-80cc-ad7c3efa6d0a",
				Name: "example_2",
				Conditions: []*EventOrchestrationCacheVariableCondition{
					{
						Expression: "event.summary matches part '[TEST]'",
					},
				},
				Configuration: &EventOrchestrationCacheVariableConfiguration{
					Type:   "recent_value",
					Source: "event.source",
					Regex:  ".*",
				},
				CreatedAt: "2024-02-12T14:44:58Z",
				CreatedBy: &UserReference{
					ID:   "P8B9WR7",
					Self: "https://api.pagerduty.com/users/P8B9WR7",
					Type: "user_reference",
				},
				UpdatedAt: "2024-02-13T12:41:43Z",
				UpdatedBy: &UserReference{
					ID:   "P8B9WR8",
					Self: "https://api.pagerduty.com/users/P8B9WR8",
					Type: "user_reference",
				},
			},
			{
				ID:   "fffd20a2-8240-4df3-b996-964968464a3f",
				Name: "example_3",
				Configuration: &EventOrchestrationCacheVariableConfiguration{
					Type:   "external_data",
					DataType: "boolean",
					TTLSeconds:  1200,
				},
				CreatedAt: "2025-02-07T18:50:58Z",
				CreatedBy: &UserReference{
					ID:   "P24PCRC",
					Self: "https://api.pagerduty.com/users/P24PCRC",
					Type: "user_reference",
				},
				DataEndpoint: "https://api.pagerduty.com/event_orchestrations/a64f9c87-6adc-4f89-a64c-2fdd8cba4639/cache_variables/fffd20a2-8240-4df3-b996-964968464a3f/data",
				UpdatedAt: "2025-02-07T20:00:43Z",
				UpdatedBy: &UserReference{
					ID:   "P23PCRC",
					Self: "https://api.pagerduty.com/users/P23PCRC",
					Type: "user_reference",
				},
			},
		},
		Total: 3,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestGlobalOrchestrationCacheVariableCreate(t *testing.T) {
	setup()
	defer teardown()
	oId := "9ad1fdb7-5f69-4e36-9a5e-0da1f3290dda"
	oType := "global"
	input := &EventOrchestrationCacheVariable{
		Name: "create_example",
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type:   "recent_value",
			Source: "event.custom_details.error",
			Regex:  "[0-9]+",
		},
	}
	url := fmt.Sprintf("%s/%s/cache_variables/", eventOrchestrationBaseUrl, oId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EventOrchestrationCacheVariablePayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.CacheVariable, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"cache_variable": {
				"id": "43f061e0-e92b-49f0-91a0-be79404319fc",
				"name": "create_example",
				"configuration": {
					"type": "recent_value",
					"source": "event.custom_details.error",
					"regex": "[0-9]+"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.Create(context.Background(), oType, oId, input)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationCacheVariable{
		ID:   "43f061e0-e92b-49f0-91a0-be79404319fc",
		Name: "create_example",
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type:   "recent_value",
			Source: "event.custom_details.error",
			Regex:  "[0-9]+",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestGlobalOrchestrationCacheVariableGet(t *testing.T) {
	setup()
	defer teardown()

	oId := "d054a57d-5933-44ce-8173-8be0bbacdfff"
	oType := "global"
	id := "9aa13ae3-81f3-4456-9abc-79233555fc3f"
	url := fmt.Sprintf("%s/%s/cache_variables/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"cache_variable": {
				"id": "9aa13ae3-81f3-4456-9abc-79233555fc3f",
				"name": "get_example",
				"conditions": [{
					"expression": "event.source matches part 'test'"
				}],
				"configuration": {
					"type": "trigger_event_count",
					"ttl_seconds": 70
				},
				"created_at": "2024-02-12T14:44:58Z",
				"created_by": {
					"id": "P8B9WR7",
					"self": "https://api.pagerduty.com/users/P8B9WR7",
					"type": "user_reference"
				},
				"updated_at": "2024-02-13T12:41:43Z",
				"updated_by": {
					"id": "P8B9WR8",
					"self": "https://api.pagerduty.com/users/P8B9WR8",
					"type": "user_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.Get(context.Background(), oType, oId, id)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationCacheVariable{
		ID:   "9aa13ae3-81f3-4456-9abc-79233555fc3f",
		Name: "get_example",
		Conditions: []*EventOrchestrationCacheVariableCondition{
			{
				Expression: "event.source matches part 'test'",
			},
		},
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type:       "trigger_event_count",
			TTLSeconds: 70,
		},
		CreatedAt: "2024-02-12T14:44:58Z",
		CreatedBy: &UserReference{
			ID:   "P8B9WR7",
			Self: "https://api.pagerduty.com/users/P8B9WR7",
			Type: "user_reference",
		},
		UpdatedAt: "2024-02-13T12:41:43Z",
		UpdatedBy: &UserReference{
			ID:   "P8B9WR8",
			Self: "https://api.pagerduty.com/users/P8B9WR8",
			Type: "user_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestGlobalOrchestrationCacheVariableUpdate(t *testing.T) {
	setup()
	defer teardown()
	oId := "31112383-c14b-4464-b8fc-660f2508f43b"
	oType := "global"
	id := "460d9629-269a-4d26-8b6b-b96d61102436"
	input := &EventOrchestrationCacheVariable{
		Disabled: true,
	}
	url := fmt.Sprintf("%s/%s/cache_variables/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationCacheVariablePayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.CacheVariable, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"cache_variable": {
				"id": "460d9629-269a-4d26-8b6b-b96d61102436",
				"name": "updated_example",
				"configuration": {
					"type": "trigger_event_count",
					"ttl_seconds": 10
				},
				"disabled": true,
				"created_at": "2024-02-12T14:44:58Z",
				"created_by": {
					"id": "P8B9WR7",
					"self": "https://api.pagerduty.com/users/P8B9WR7",
					"type": "user_reference"
				},
				"updated_at": "2024-02-13T12:41:43Z",
				"updated_by": {
					"id": "P8B9WR8",
					"self": "https://api.pagerduty.com/users/P8B9WR8",
					"type": "user_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.Update(context.Background(), oType, oId, id, input)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationCacheVariable{
		ID:   "460d9629-269a-4d26-8b6b-b96d61102436",
		Name: "updated_example",
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type:       "trigger_event_count",
			TTLSeconds: 10,
		},
		Disabled:  true,
		CreatedAt: "2024-02-12T14:44:58Z",
		CreatedBy: &UserReference{
			ID:   "P8B9WR7",
			Self: "https://api.pagerduty.com/users/P8B9WR7",
			Type: "user_reference",
		},
		UpdatedAt: "2024-02-13T12:41:43Z",
		UpdatedBy: &UserReference{
			ID:   "P8B9WR8",
			Self: "https://api.pagerduty.com/users/P8B9WR8",
			Type: "user_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestGlobalOrchestrationCacheVariableDelete(t *testing.T) {
	setup()
	defer teardown()

	oId := "1e51d2d3-1384-46f1-a7bb-920b8b114f80"
	oType := "global"
	id := "58cab484-ed93-4c46-8188-8f02b0cf3d9f"
	url := fmt.Sprintf("%s/%s/cache_variables/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.EventOrchestrationCacheVariables.Delete(context.Background(), oType, oId, id); err != nil {
		t.Fatal(err)
	}
}

func TestServiceEventOrchestrationCacheVariableList(t *testing.T) {
	setup()
	defer teardown()

	oId := "P3ZQXDF"
	oType := "service"
	url := fmt.Sprintf("%s/services/%s/cache_variables/", eventOrchestrationBaseUrl, oId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"cache_variables": [
				{
					"id": "45be0a94-55b2-4691-b285-7a14478f4fe2",
					"name": "example_1",
					"conditions": [],
					"configuration": {
						"type": "trigger_event_count",
						"ttl_seconds": 30
					},
					"created_at": "2024-02-12T14:44:58Z",
					"created_by": {
						"id": "P8B9WR7",
						"self": "https://api.pagerduty.com/users/P8B9WR7",
						"type": "user_reference"
					},
					"updated_at": "2024-02-13T12:41:43Z",
					"updated_by": {
						"id": "P8B9WR8",
						"self": "https://api.pagerduty.com/users/P8B9WR8",
						"type": "user_reference"
					}
				},
				{
					"id": "b5e7eef2-d2f8-46ce-80cc-ad7c3efa6d0a",
					"name": "example_2",
					"conditions": [{"expression": "event.summary matches part '[TEST]'"}],
					"configuration": {
						"type": "recent_value",
						"source": "event.source",
						"regex": ".*"
					},
					"created_at": "2024-02-12T14:44:58Z",
					"created_by": {
						"id": "P8B9WR7",
						"self": "https://api.pagerduty.com/users/P8B9WR7",
						"type": "user_reference"
					},
					"updated_at": "2024-02-13T12:41:43Z",
					"updated_by": {
						"id": "P8B9WR8",
						"self": "https://api.pagerduty.com/users/P8B9WR8",
						"type": "user_reference"
					}
				},
				{
					"id": "fffd20a2-8240-4df3-b996-964968464a3f",
					"name": "example_3",
					"configuration": {
						"type": "external_data",
						"data_type": "boolean",
						"ttl_seconds": 1200
					},
					"created_at": "2025-02-07T18:50:58Z",
					"created_by": {
						"id": "P24PCRC",
						"self": "https://api.pagerduty.com/users/P24PCRC",
						"type": "user_reference"
					},
					"data_endpoint": "https://api.pagerduty.com/event_orchestrations/a64f9c87-6adc-4f89-a64c-2fdd8cba4639/cache_variables/fffd20a2-8240-4df3-b996-964968464a3f/data",
					"updated_at": "2025-02-07T20:00:43Z",
					"updated_by": {
						"id": "P23PCRC",
						"self": "https://api.pagerduty.com/users/P23PCRC",
						"type": "user_reference"
					}
				}
			],
			"total": 3
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.List(context.Background(), oType, oId)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListEventOrchestrationCacheVariablesResponse{
		CacheVariables: []*EventOrchestrationCacheVariable{
			{
				ID:         "45be0a94-55b2-4691-b285-7a14478f4fe2",
				Name:       "example_1",
				Conditions: []*EventOrchestrationCacheVariableCondition{},
				Configuration: &EventOrchestrationCacheVariableConfiguration{
					Type:       "trigger_event_count",
					TTLSeconds: 30,
				},
				CreatedAt: "2024-02-12T14:44:58Z",
				CreatedBy: &UserReference{
					ID:   "P8B9WR7",
					Self: "https://api.pagerduty.com/users/P8B9WR7",
					Type: "user_reference",
				},
				UpdatedAt: "2024-02-13T12:41:43Z",
				UpdatedBy: &UserReference{
					ID:   "P8B9WR8",
					Self: "https://api.pagerduty.com/users/P8B9WR8",
					Type: "user_reference",
				},
			},
			{
				ID:   "b5e7eef2-d2f8-46ce-80cc-ad7c3efa6d0a",
				Name: "example_2",
				Conditions: []*EventOrchestrationCacheVariableCondition{
					{
						Expression: "event.summary matches part '[TEST]'",
					},
				},
				Configuration: &EventOrchestrationCacheVariableConfiguration{
					Type:   "recent_value",
					Source: "event.source",
					Regex:  ".*",
				},
				CreatedAt: "2024-02-12T14:44:58Z",
				CreatedBy: &UserReference{
					ID:   "P8B9WR7",
					Self: "https://api.pagerduty.com/users/P8B9WR7",
					Type: "user_reference",
				},
				UpdatedAt: "2024-02-13T12:41:43Z",
				UpdatedBy: &UserReference{
					ID:   "P8B9WR8",
					Self: "https://api.pagerduty.com/users/P8B9WR8",
					Type: "user_reference",
				},
			},
			{
				ID:   "fffd20a2-8240-4df3-b996-964968464a3f",
				Name: "example_3",
				Configuration: &EventOrchestrationCacheVariableConfiguration{
					Type:   "external_data",
					DataType: "boolean",
					TTLSeconds:  1200,
				},
				CreatedAt: "2025-02-07T18:50:58Z",
				CreatedBy: &UserReference{
					ID:   "P24PCRC",
					Self: "https://api.pagerduty.com/users/P24PCRC",
					Type: "user_reference",
				},
				DataEndpoint: "https://api.pagerduty.com/event_orchestrations/a64f9c87-6adc-4f89-a64c-2fdd8cba4639/cache_variables/fffd20a2-8240-4df3-b996-964968464a3f/data",
				UpdatedAt: "2025-02-07T20:00:43Z",
				UpdatedBy: &UserReference{
					ID:   "P23PCRC",
					Self: "https://api.pagerduty.com/users/P23PCRC",
					Type: "user_reference",
				},
			},
		},
		Total: 3,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServiceOrchestrationCacheVariableCreate(t *testing.T) {
	setup()
	defer teardown()
	oId := "P3ZQXDF"
	oType := "service"
	input := &EventOrchestrationCacheVariable{
		Name: "create_example",
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type: "external_data",
			DataType: "string",
			TTLSeconds: 900,
		},
	}
	url := fmt.Sprintf("%s/services/%s/cache_variables/", eventOrchestrationBaseUrl, oId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EventOrchestrationCacheVariablePayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.CacheVariable, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"cache_variable": {
				"id": "43f061e0-e92b-49f0-91a0-be79404319fc",
				"name": "create_example",
				"configuration": {
					"type": "external_data",
					"data_type": "string",
					"ttl_seconds": 900
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.Create(context.Background(), oType, oId, input)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationCacheVariable{
		ID:   "43f061e0-e92b-49f0-91a0-be79404319fc",
		Name: "create_example",
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type: "external_data",
			DataType: "string",
			TTLSeconds: 900,
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServiceOrchestrationCacheVariableGet(t *testing.T) {
	setup()
	defer teardown()

	oId := "P3ZQXDF"
	oType := "service"
	id := "9aa13ae3-81f3-4456-9abc-79233555fc3f"
	url := fmt.Sprintf("%s/services/%s/cache_variables/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"cache_variable": {
				"id": "9aa13ae3-81f3-4456-9abc-79233555fc3f",
				"name": "get_example",
				"conditions": [{
					"expression": "event.source matches part 'test'"
				}],
				"configuration": {
					"type": "trigger_event_count",
					"ttl_seconds": 70
				},
				"created_at": "2024-02-12T14:44:58Z",
				"created_by": {
					"id": "P8B9WR7",
					"self": "https://api.pagerduty.com/users/P8B9WR7",
					"type": "user_reference"
				},
				"updated_at": "2024-02-13T12:41:43Z",
				"updated_by": {
					"id": "P8B9WR8",
					"self": "https://api.pagerduty.com/users/P8B9WR8",
					"type": "user_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.Get(context.Background(), oType, oId, id)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationCacheVariable{
		ID:   "9aa13ae3-81f3-4456-9abc-79233555fc3f",
		Name: "get_example",
		Conditions: []*EventOrchestrationCacheVariableCondition{
			{
				Expression: "event.source matches part 'test'",
			},
		},
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type:       "trigger_event_count",
			TTLSeconds: 70,
		},
		CreatedAt: "2024-02-12T14:44:58Z",
		CreatedBy: &UserReference{
			ID:   "P8B9WR7",
			Self: "https://api.pagerduty.com/users/P8B9WR7",
			Type: "user_reference",
		},
		UpdatedAt: "2024-02-13T12:41:43Z",
		UpdatedBy: &UserReference{
			ID:   "P8B9WR8",
			Self: "https://api.pagerduty.com/users/P8B9WR8",
			Type: "user_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServiceOrchestrationCacheVariableUpdate(t *testing.T) {
	setup()
	defer teardown()
	oId := "P3ZQXDF"
	oType := "service"
	id := "460d9629-269a-4d26-8b6b-b96d61102436"
	input := &EventOrchestrationCacheVariable{
		Disabled: true,
	}
	url := fmt.Sprintf("%s/services/%s/cache_variables/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationCacheVariablePayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.CacheVariable, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"cache_variable": {
				"id": "460d9629-269a-4d26-8b6b-b96d61102436",
				"name": "updated_example",
				"configuration": {
					"type": "trigger_event_count",
					"ttl_seconds": 10
				},
				"disabled": true,
				"created_at": "2024-02-12T14:44:58Z",
				"created_by": {
					"id": "P8B9WR7",
					"self": "https://api.pagerduty.com/users/P8B9WR7",
					"type": "user_reference"
				},
				"updated_at": "2024-02-13T12:41:43Z",
				"updated_by": {
					"id": "P8B9WR8",
					"self": "https://api.pagerduty.com/users/P8B9WR8",
					"type": "user_reference"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationCacheVariables.Update(context.Background(), oType, oId, id, input)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationCacheVariable{
		ID:   "460d9629-269a-4d26-8b6b-b96d61102436",
		Name: "updated_example",
		Configuration: &EventOrchestrationCacheVariableConfiguration{
			Type:       "trigger_event_count",
			TTLSeconds: 10,
		},
		Disabled:  true,
		CreatedAt: "2024-02-12T14:44:58Z",
		CreatedBy: &UserReference{
			ID:   "P8B9WR7",
			Self: "https://api.pagerduty.com/users/P8B9WR7",
			Type: "user_reference",
		},
		UpdatedAt: "2024-02-13T12:41:43Z",
		UpdatedBy: &UserReference{
			ID:   "P8B9WR8",
			Self: "https://api.pagerduty.com/users/P8B9WR8",
			Type: "user_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestServiceOrchestrationCacheVariableDelete(t *testing.T) {
	setup()
	defer teardown()

	oId := "P3ZQXDF"
	oType := "service"
	id := "58cab484-ed93-4c46-8188-8f02b0cf3d9f"
	url := fmt.Sprintf("%s/services/%s/cache_variables/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.EventOrchestrationCacheVariables.Delete(context.Background(), oType, oId, id); err != nil {
		t.Fatal(err)
	}
}
