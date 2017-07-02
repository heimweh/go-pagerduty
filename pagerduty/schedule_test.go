package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestSchedulesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"schedules": [{"id": "1"}]}`))
	})

	resp, _, err := client.Schedules.List(&ListSchedulesOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListSchedulesResponse{
		Schedules: []*Schedule{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSchedulesCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Schedule{
		Name: "foo",
	}

	mux.HandleFunc("/schedules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Schedule)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Schedule, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"schedule": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Schedules.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Schedule{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSchedulesDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Schedules.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestSchedulesGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"schedule": {"id": "1"}}`))
	})

	resp, _, err := client.Schedules.Get("1", &GetScheduleOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &Schedule{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSchedulesUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Schedule{
		Name: "foo",
	}

	mux.HandleFunc("/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Schedule)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Schedule, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"schedule": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Schedules.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Schedule{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSchedulesListOverrides(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/1/overrides", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"overrides": [{"id": "1"}]}`))
	})

	resp, _, err := client.Schedules.ListOverrides("1", &ListOverridesOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListOverridesResponse{
		Overrides: []*Override{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSchedulesCreateOverride(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/1/overrides", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(`{"override": {"id": "1", "user": { "id": "1" }}}`))
	})

	resp, _, err := client.Schedules.CreateOverride("1", &Override{User: &UserReference{ID: "1"}})
	if err != nil {
		t.Fatal(err)
	}

	want := &Override{
		ID: "1",
		User: &UserReference{
			ID: "1",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestSchedulesDeleteOverride(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/1/overrides/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Schedules.DeleteOverride("1", "1"); err != nil {
		t.Fatal(err)
	}
}

func TestSchedulesListOnCalls(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/1/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"users": [{"id": "1"}]}`))
	})

	want := &ListOnCallsResponse{
		Users: []*User{
			{
				ID: "1",
			},
		},
	}

	resp, _, err := client.Schedules.ListOnCalls("1", &ListOnCallsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
