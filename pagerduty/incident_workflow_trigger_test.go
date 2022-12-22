package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentWorkflowTriggerList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		testQueryMaxCount(t, r, 1)
		pageToken := r.URL.Query().Get("page_token")

		switch pageToken {
		case "":
			w.Write([]byte(`{"next_page_token":"abc", "triggers":[{"id": "1"}]}`))
		case "abc":
			w.Write([]byte(`{"next_page_token":"def", "triggers":[{"id": "2"}]}`))
		case "def":
			w.Write([]byte(`{"next_page_token":null, "triggers":[{"id": "3"}]}`))
		default:
			t.Fatalf("Unexpected pageToken: %v", pageToken)
		}

	})

	resp, _, err := client.IncidentWorkflowTriggers.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowTriggerResponse{
		Limit: 0,
		Triggers: []*IncidentWorkflowTrigger{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
			{
				ID: "3",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowTriggerList_SecondPage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		testQueryCount(t, r, 1)
		pageToken := r.URL.Query().Get("page_token")

		switch pageToken {
		case "def":
			w.Write([]byte(`{"next_page_token":null, "triggers":[{"id": "3"}]}`))
		default:
			t.Fatalf("Unexpected pageToken: %v", pageToken)
		}

	})

	resp, _, err := client.IncidentWorkflowTriggers.List(&ListIncidentWorkflowTriggerOptions{PageToken: "def"})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowTriggerResponse{
		Limit: 0,
		Triggers: []*IncidentWorkflowTrigger{
			{
				ID: "3",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowTriggerList_Limit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		testQueryCount(t, r, 1)
		testQueryValue(t, r, "limit", "42")

		w.Write([]byte(`{"limit": 42, "triggers":[{"id": "2"}]}`))

	})

	resp, _, err := client.IncidentWorkflowTriggers.List(&ListIncidentWorkflowTriggerOptions{Limit: 42})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowTriggerResponse{
		Limit: 42,
		Triggers: []*IncidentWorkflowTrigger{
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowTriggerList_WithOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")

		testQueryMinCount(t, r, 1)
		testQueryMaxCount(t, r, 2)
		testQueryValue(t, r, "trigger_type", "manual")

		pageToken := r.URL.Query().Get("page_token")
		switch pageToken {
		case "":
			w.Write([]byte(`{"next_page_token":"abc", "triggers":[{"id": "1"}]}`))
		case "abc":
			w.Write([]byte(`{"triggers":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected pageToken: %v", pageToken)
		}

	})

	resp, _, err := client.IncidentWorkflowTriggers.List(&ListIncidentWorkflowTriggerOptions{
		TriggerType: IncidentWorkflowTriggerTypeManual,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowTriggerResponse{
		Triggers: []*IncidentWorkflowTrigger{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowTriggerGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers/IWT1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		testBody(t, r, "")

		w.Write([]byte(`
{
   "trigger":{
      "id":"abc-123",
      "type":"workflow_trigger",
      "trigger_type_name":"Manual Incident Trigger",
      "trigger_type":"manual",
      "trigger_url":"https://api.pagerduty.com/incident_workflows/triggers/abc-123/start",
      "self":"https://api.pagerduty.com/incident_workflows/triggers/abc-123",
      "workflow":{
         "id":"TO38234",
         "type":"workflow",
         "name":"Example Workflow",
         "description":"This workflow serves as an example",
         "self":"https://api.pagerduty.com/incident_workflows/TO38234",
         "created_at":"2022-06-07T00:01:55Z"
      },
      "services":[
         {
            "id":"PIJ90N7",
            "summary":"My Application Service",
            "type":"service",
            "self":"https://api.pagerduty.com/services/PIJ90N7",
            "html_url":"https://subdomain.pagerduty.com/services/PIJ90N7",
            "name":"My Application Service",
            "created_at":"2015-11-06T11:12:51-05:00",
            "status":"active"
         }
      ],
      "condition":"incident.priority matches 'P1'",
      "permissions":{
         "restricted":true,
         "team_id":"PDEJ7MP"
      }
   }
}
`))

	})

	resp, _, err := client.IncidentWorkflowTriggers.Get("IWT1")
	if err != nil {
		t.Fatal(err)
	}

	workflowDesc := "This workflow serves as an example"
	cond := "incident.priority matches 'P1'"

	want := &IncidentWorkflowTrigger{
		ID:          "abc-123",
		Type:        "workflow_trigger",
		TriggerType: IncidentWorkflowTriggerTypeManual,
		Workflow: &IncidentWorkflow{
			ID:          "TO38234",
			Type:        "workflow",
			Name:        "Example Workflow",
			Description: &workflowDesc,
			Self:        "https://api.pagerduty.com/incident_workflows/TO38234",
		},
		Services: []*ServiceReference{
			{
				ID:      "PIJ90N7",
				Summary: "My Application Service",
				Type:    "service",
				Self:    "https://api.pagerduty.com/services/PIJ90N7",
				HTMLURL: "https://subdomain.pagerduty.com/services/PIJ90N7",
			},
		},
		Condition: &cond,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowTriggerCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		testBody(t, r, `{"trigger_type":"manual","workflow":{"id":"TO38234"},"services":[{"id":"PIJ90N7"}],"condition":"incident.priority matches 'P1'"}`)

		w.Write([]byte(`
{
   "trigger":{
      "id":"abc-123",
      "type":"workflow_trigger",
      "trigger_type_name":"Manual Incident Trigger",
      "trigger_type":"manual",
      "trigger_url":"https://api.pagerduty.com/incident_workflows/triggers/abc-123/start",
      "self":"https://api.pagerduty.com/incident_workflows/triggers/abc-123",
      "workflow":{
         "id":"TO38234",
         "type":"workflow",
         "name":"Example Workflow",
         "description":"This workflow serves as an example",
         "self":"https://api.pagerduty.com/incident_workflows/TO38234",
         "created_at":"2022-06-07T00:01:55Z"
      },
      "services":[
         {
            "id":"PIJ90N7",
            "summary":"My Application Service",
            "type":"service",
            "self":"https://api.pagerduty.com/services/PIJ90N7",
            "html_url":"https://subdomain.pagerduty.com/services/PIJ90N7",
            "name":"My Application Service",
            "created_at":"2015-11-06T11:12:51-05:00",
            "status":"active"
         }
      ],
      "is_subscribed_to_all_services": true,
      "condition":"incident.priority matches 'P1'"
   }
}
`))

	})

	cond := "incident.priority matches 'P1'"

	resp, _, err := client.IncidentWorkflowTriggers.Create(&IncidentWorkflowTrigger{
		TriggerType: IncidentWorkflowTriggerTypeManual,
		Workflow: &IncidentWorkflow{
			ID: "TO38234",
		},
		Services: []*ServiceReference{
			{
				ID: "PIJ90N7",
			},
		},
		Condition: &cond,
	})
	if err != nil {
		t.Fatal(err)
	}

	workflowDesc := "This workflow serves as an example"

	want := &IncidentWorkflowTrigger{
		ID:          "abc-123",
		Type:        "workflow_trigger",
		TriggerType: IncidentWorkflowTriggerTypeManual,
		Workflow: &IncidentWorkflow{
			ID:          "TO38234",
			Type:        "workflow",
			Name:        "Example Workflow",
			Description: &workflowDesc,
			Self:        "https://api.pagerduty.com/incident_workflows/TO38234",
		},
		Services: []*ServiceReference{
			{
				ID:      "PIJ90N7",
				Summary: "My Application Service",
				Type:    "service",
				Self:    "https://api.pagerduty.com/services/PIJ90N7",
				HTMLURL: "https://subdomain.pagerduty.com/services/PIJ90N7",
			},
		},
		SubscribedToAllServices: true,
		Condition:               &cond,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowTriggerDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers/IWT1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		w.WriteHeader(200)
	})

	resp, err := client.IncidentWorkflowTriggers.Delete("IWT1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 200 {
		t.Errorf("unexpected response code. want 200. got %v", resp.Response.StatusCode)
	}
}

func TestIncidentWorkflowTriggerUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/triggers/IW1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "x-early-access", "incident-workflows-early-access")
		testBody(t, r, `{"services":[{"id":"PIJ90N7"}],"condition":"incident.priority matches 'P1'"}`)

		w.Write([]byte(`
{
   "trigger":{
      "id":"abc-123",
      "type":"workflow_trigger",
      "trigger_type_name":"Manual Incident Trigger",
      "trigger_type":"manual",
      "trigger_url":"https://api.pagerduty.com/incident_workflows/triggers/abc-123/start",
      "self":"https://api.pagerduty.com/incident_workflows/triggers/abc-123",
      "workflow":{
         "id":"TO38234",
         "type":"workflow",
         "name":"Example Workflow",
         "description":"This workflow serves as an example",
         "self":"https://api.pagerduty.com/incident_workflows/TO38234",
         "created_at":"2022-06-07T00:01:55Z"
      },
      "workflow_id":"xyz-123",
      "workflow_name":"High Priority Incident",
      "services":[
         {
            "id":"PIJ90N7",
            "summary":"My Application Service",
            "type":"service",
            "self":"https://api.pagerduty.com/services/PIJ90N7",
            "html_url":"https://subdomain.pagerduty.com/services/PIJ90N7",
            "name":"My Application Service",
            "created_at":"2015-11-06T11:12:51-05:00",
            "status":"active"
         }
      ],
      "condition":"incident.priority matches 'P1'"
   }
}
`))

	})

	workflowDesc := "This workflow serves as an example"
	cond := "incident.priority matches 'P1'"

	resp, _, err := client.IncidentWorkflowTriggers.Update("IW1", &IncidentWorkflowTrigger{
		Condition: &cond,
		Services: []*ServiceReference{
			{
				ID: "PIJ90N7",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentWorkflowTrigger{
		ID:          "abc-123",
		Type:        "workflow_trigger",
		TriggerType: IncidentWorkflowTriggerTypeManual,
		Workflow: &IncidentWorkflow{
			ID:          "TO38234",
			Type:        "workflow",
			Name:        "Example Workflow",
			Description: &workflowDesc,
			Self:        "https://api.pagerduty.com/incident_workflows/TO38234",
		},
		Services: []*ServiceReference{
			{
				ID:      "PIJ90N7",
				Summary: "My Application Service",
				Type:    "service",
				Self:    "https://api.pagerduty.com/services/PIJ90N7",
				HTMLURL: "https://subdomain.pagerduty.com/services/PIJ90N7",
			},
		},
		Condition: &cond,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
