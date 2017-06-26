package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTeamsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"teams": [{"id": "1"}]}`))
	})

	resp, _, err := client.Teams.List(&ListTeamsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListTeamsResponse{
		Teams: []*Team{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTeamsCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Team{
		Name: "foo",
	}

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Team)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Team, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"team": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Teams.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Team{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTeamsDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Teams.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestTeamsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"team": {"id": "1"}}`))
	})

	resp, _, err := client.Teams.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Team{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTeamsUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Team{
		Name: "foo",
	}

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Team)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Team, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"team": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Teams.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Team{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTeamsAddUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams/1/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	if _, err := client.Teams.AddUser("1", "1"); err != nil {
		t.Fatal(err)
	}
}

func TestTeamsRemoveUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams/1/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	if _, err := client.Teams.RemoveUser("1", "1"); err != nil {
		t.Fatal(err)
	}
}
