package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestMaintenanceWindowsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/maintenance_windows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"maintenance_windows": [{"id": "1"}]}`))
	})

	resp, _, err := client.MaintenanceWindows.List(&ListMaintenanceWindowsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListMaintenanceWindowsResponse{
		MaintenanceWindows: []*MaintenanceWindow{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestMaintenanceWindowsCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &MaintenanceWindow{Description: "foo"}

	mux.HandleFunc("/maintenance_windows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(MaintenanceWindowPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.MaintenanceWindow, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"maintenance_window": {"description": "foo", "id": "1"}}`))
	})

	resp, _, err := client.MaintenanceWindows.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &MaintenanceWindow{
		Description: "foo",
		ID:          "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestMaintenanceWindowsDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/maintenance_windows/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.MaintenanceWindows.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestMaintenanceWindowsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/maintenance_windows/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"maintenance_window": {"id": "1"}}`))
	})

	resp, _, err := client.MaintenanceWindows.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &MaintenanceWindow{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestMaintenanceWindowsUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &MaintenanceWindow{
		Description: "foo",
	}

	mux.HandleFunc("/maintenance_windows/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.Write([]byte(`{"maintenance_window": {"description": "foo", "id": "1"}}`))
	})

	resp, _, err := client.MaintenanceWindows.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &MaintenanceWindow{
		Description: "foo",
		ID:          "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
