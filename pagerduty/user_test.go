package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestUsersList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"users": [{"id": "P1D3Z4B"}]}`))
	})

	resp, _, err := client.Users.List(&ListUsersOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListUsersResponse{
		Users: []*User{
			{
				ID: "P1D3Z4B",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Name: "foo"}

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.User, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"user": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Users.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &User{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Users.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestUsersGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"user": {"id": "1"}}`))
	})

	resp, _, err := client.Users.Get("1", &GetUserOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &User{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Name: "foo"}

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.User, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"user": {"name": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Users.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &User{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}
