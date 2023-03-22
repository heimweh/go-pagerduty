package pagerduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEventOrchestrationIntegrationList(t *testing.T) {
	setup()
	defer teardown()

	oId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	url := fmt.Sprintf("%s/%s/integrations", eventOrchestrationBaseUrl, oId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"integrations": [
				{
					"id": "b2c92c66-bb25-48a3-b1cb-f38b26837277",
					"label": "My Orchestration Default Integration",
					"parameters": {
						"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
						"type": "global"
					}
				},
				{
					"id": "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
					"label": "My Integration 2",
					"parameters": {
						"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX2",
						"type": "global"
					}
				},
				{
					"id": "8e61060f-fa66-4c21-a858-9fd997c28ccc",
					"label": "My Integration 3",
					"parameters": {
						"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX3",
						"type": "global"
					}
				}
			],
			"total": 3
		}`))
	})

	resp, _, err := client.EventOrchestrationIntegrations.List(oId)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListEventOrchestrationIntegrationsResponse{
		Integrations: []*EventOrchestrationIntegration{
			{
				ID:    "b2c92c66-bb25-48a3-b1cb-f38b26837277",
				Label: "My Orchestration Default Integration",
				Parameters: &EventOrchestrationIntegrationParameters{
					RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
					Type:       "global",
				},
			},
			{
				ID:    "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
				Label: "My Integration 2",
				Parameters: &EventOrchestrationIntegrationParameters{
					RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX2",
					Type:       "global",
				},
			},
			{
				ID:    "8e61060f-fa66-4c21-a858-9fd997c28ccc",
				Label: "My Integration 3",
				Parameters: &EventOrchestrationIntegrationParameters{
					RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX3",
					Type:       "global",
				},
			},
		},
		Total: 3,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationIntegrationCreate(t *testing.T) {
	setup()
	defer teardown()
	oId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	input := &EventOrchestrationIntegration{
		Label: "My Integration",
	}
	url := fmt.Sprintf("%s/%s/integrations", eventOrchestrationBaseUrl, oId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EventOrchestrationIntegrationPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Integration, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"integration": {
				"id": "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
				"label": "My Integration",
				"parameters": {
					"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"type": "global"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationIntegrations.Create(oId, input)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationIntegration{
		ID:    "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
		Label: "My Integration",
		Parameters: &EventOrchestrationIntegrationParameters{
			RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Type:       "global",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationIntegrationGet(t *testing.T) {
	setup()
	defer teardown()

	oId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	id := "b8a1f84c-418c-417c-90bb-597bf5ca7ffc"
	url := fmt.Sprintf("%s/%s/integrations/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{
			"integration": {
				"id": "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
				"label": "My Integration",
				"parameters": {
					"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"type": "global"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationIntegrations.Get(oId, id)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationIntegration{
		ID:    "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
		Label: "My Integration",
		Parameters: &EventOrchestrationIntegrationParameters{
			RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Type:       "global",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationIntegrationUpdate(t *testing.T) {
	setup()
	defer teardown()
	oId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	id := "b8a1f84c-418c-417c-90bb-597bf5ca7ffc"
	input := &EventOrchestrationIntegration{
		Label: "My Updated Integration",
	}
	url := fmt.Sprintf("%s/%s/integrations/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(EventOrchestrationIntegrationPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Integration, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"integration": {
				"id": "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
				"label": "My Updated Integration",
				"parameters": {
					"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"type": "global"
				}
			}
		}`))
	})

	resp, _, err := client.EventOrchestrationIntegrations.Update(oId, id, input)

	if err != nil {
		t.Fatal(err)
	}

	want := &EventOrchestrationIntegration{
		ID:    "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
		Label: "My Updated Integration",
		Parameters: &EventOrchestrationIntegrationParameters{
			RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Type:       "global",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestEventOrchestrationIntegrationDelete(t *testing.T) {
	setup()
	defer teardown()

	oId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	id := "b8a1f84c-418c-417c-90bb-597bf5ca7ffc"
	url := fmt.Sprintf("%s/%s/integrations/%s", eventOrchestrationBaseUrl, oId, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.EventOrchestrationIntegrations.Delete(oId, id); err != nil {
		t.Fatal(err)
	}
}

func TestEventOrchestrationIntegrationMigrate(t *testing.T) {
	setup()
	defer teardown()
	soId := "a64f9c87-6adc-4f89-a64c-2fdd8cba4639"
	doId := "a99a20bd-7699-4722-8bef-b48c2fcac116"
	id := "b8a1f84c-418c-417c-90bb-597bf5ca7ffc"
	input := &EventOrchestrationIntegrationMigrationPayload{
		SourceType:    "orchestration",
		SourceId:      soId,
		IntegrationId: id,
	}
	url := fmt.Sprintf("%s/%s/integrations/migration", eventOrchestrationBaseUrl, doId)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(EventOrchestrationIntegrationMigrationPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{
			"integrations": [
				{
					"id": "b2c92c66-bb25-48a3-b1cb-f38b26837277",
					"label": "My Orchestration Default Integration",
					"parameters": {
						"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
						"type": "global"
					}
				},
				{
					"id": "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
					"label": "Source Orchestration Integraton",
					"parameters": {
						"routing_key": "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX2",
						"type": "global"
					}
				}
			],
			"total": 2
		}`))
	})

	resp, _, err := client.EventOrchestrationIntegrations.MigrateFromOrchestration(doId, soId, id)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListEventOrchestrationIntegrationsResponse{
		Integrations: []*EventOrchestrationIntegration{
			{
				ID:    "b2c92c66-bb25-48a3-b1cb-f38b26837277",
				Label: "My Orchestration Default Integration",
				Parameters: &EventOrchestrationIntegrationParameters{
					RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
					Type:       "global",
				},
			},
			{
				ID:    "b8a1f84c-418c-417c-90bb-597bf5ca7ffc",
				Label: "Source Orchestration Integraton",
				Parameters: &EventOrchestrationIntegrationParameters{
					RoutingKey: "R0XXXXXXXXXXXXXXXXXXXXXXXXXXXXX2",
					Type:       "global",
				},
			},
		},
		Total: 2,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
