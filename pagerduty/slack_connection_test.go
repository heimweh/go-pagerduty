package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestSlackConnectionList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 0, "offset": 0, "more": false, "limit": 0, "slack_connections":[{"id": "1"}]}`))
	})

	resp, _, err := client.SlackConnections.List("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &ListSlackConnectionsResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		SlackConnections: []*SlackConnection{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSlackConnectionCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &SlackConnection{SourceID: "1"}
	workspaceID := "1"

	mux.HandleFunc("/integration-slack/workspaces/1/connections", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(SlackConnection)
		v.SourceID = "1"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"slack_connection":{"id": "1", "source_id":"1"}}`))
	})

	resp, _, err := client.SlackConnections.Create(workspaceID, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SlackConnection{
		ID:       "1",
		SourceID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSlackConnectionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"slack_connection":{"source_id": "1", "id":"1"}}`))
	})

	ID := "1"
	workspaceID := "1"
	resp, _, err := client.SlackConnections.Get(workspaceID, ID)

	if err != nil {
		t.Fatal(err)
	}

	want := &SlackConnection{
		SourceID: "1",
		ID:       "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSlackConnectionUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &SlackConnection{
		SourceID: "2",
	}

	mux.HandleFunc("/integration-slack/workspaces/1/connections/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(SlackConnection)
		v.SourceID = "2"

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"slack_connection":{"source_id": "2", "id":"1"}}`))
	})

	workspaceID := "1"
	connID := "1"

	resp, _, err := client.SlackConnections.Update(workspaceID, connID, input)
	if err != nil {
		t.Fatal(err)
	}

	want := &SlackConnection{
		SourceID: "2",
		ID:       "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSlackConnectionDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/integration-slack/workspaces/1/connections/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	workspaceID := "1"
	connID := "1"

	if _, err := client.SlackConnections.Delete(workspaceID, connID); err != nil {
		t.Fatal(err)
	}
}
