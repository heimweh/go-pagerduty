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
		v := new(UserPayload)
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
		v := new(UserPayload)
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

func TestUsersAddContactMethod(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactMethod{Address: "foo@bar.com", Type: "email_contact_method"}

	mux.HandleFunc("/users/1/contact_methods", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(ContactMethodPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.ContactMethod, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"contact_method": { "address": "foo@bar.com", "id": "1", "type": "email_contact_method" }}`))
	})

	resp, _, err := client.Users.CreateContactMethod("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &ContactMethod{
		ID:      "1",
		Type:    "email_contact_method",
		Address: "foo@bar.com",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersAddDuplicateContactMethod(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactMethod{Address: "foo@bar.com", Type: "email_contact_method"}

	mux.HandleFunc("/users/1/contact_methods", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			v := new(ContactMethodPayload)
			json.NewDecoder(r.Body).Decode(v)
			if !reflect.DeepEqual(v.ContactMethod, input) {
				t.Errorf("Request body = %+v, want %+v", v, input)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":{"errors":["User Contact method must be unique"],"code":2001,"message":"Invalid Input Provided"}}`))
		} else if r.Method == "GET" {
			w.Write([]byte(`{"contact_methods": [{ "address": "foo@bar.com", "id": "1", "type": "email_contact_method", "self":"api/users/1/contact_methods/1" }] }`))
		} else {
			t.Errorf("Request method : %v is neither POST nor GET", r.Method)
		}
	})

	mux.HandleFunc("/users/1/contact_methods/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"contact_method": { "address": "foo@bar.com", "id": "1", "type": "email_contact_method", "self":"api/users/1/contact_methods/1" }}`))
	})

	resp, _, err := client.Users.CreateContactMethod("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &ContactMethod{
		ID:      "1",
		Type:    "email_contact_method",
		Address: "foo@bar.com",
		Self:    "api/users/1/contact_methods/1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersUpdateContactMethod(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactMethod{Address: "foo@bar.com", Type: "email_contact_method"}

	mux.HandleFunc("/users/1/contact_methods/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(ContactMethodPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.ContactMethod, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"contact_method": { "address": "foo@bar.com", "id": "1", "type": "email_contact_method" }}`))
	})

	resp, _, err := client.Users.UpdateContactMethod("1", "1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &ContactMethod{
		ID:      "1",
		Type:    "email_contact_method",
		Address: "foo@bar.com",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersDeleteContactMethod(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/contact_methods/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Users.DeleteContactMethod("1", "1"); err != nil {
		t.Fatal(err)
	}
}
