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

func TestDuplicateUsersCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Name: "foo", Email: "foo@bar.com"}

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			v := new(UserPayload)
			json.NewDecoder(r.Body).Decode(v)
			if !reflect.DeepEqual(v.User, input) {
				t.Errorf("Request body = %+v, want %+v", v, input)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":{"errors":["Email has already been taken"],"code":2001,"message":"Invalid Input Provided"}}`))
		} else if r.Method == "GET" {
			qry := r.URL.Query().Get("query")
			if qry != "foo@bar.com" {
				t.Errorf("Request query =%+v, want %+v", qry, "foo@bar.com")
			}
			w.Write([]byte(`{"users": [{"id": "1", "name": "foo", "email": "foo@bar.com"}]}`))
		} else {
			t.Errorf("Request method : %v is neither POST nor GET", r.Method)
		}
	})

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"user": {"id": "1", "name": "foo", "email": "foo@bar.com"}}`))
	})

	resp, _, err := client.Users.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &User{
		Name:  "foo",
		ID:    "1",
		Email: "foo@bar.com",
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

func TestUsersGetLicense(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/license", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"license": {"id": "1", "type": "license"}}`))
	})

	resp, _, err := client.Users.GetLicense("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &License{ID: "1", Type: "license"}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestUsersGetWithLicense(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"user": {"id": "1"}}`))
	})
	mux.HandleFunc("/users/1/license", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"license": {"id": "1", "type": "license"}}`))
	})

	resp, err := client.Users.GetWithLicense("1", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &User{
		ID:      "1",
		License: &LicenseReference{ID: "1", Type: "license_reference"},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}

func TestListAllWithLicenses(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"users": [{"id": "P1D3Z4B", "role": "user"}]}`))
	})
	mux.HandleFunc("/license_allocations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"license_allocations": [
			{
				"user": {"id": "P1D3Z4B", "type": "user_reference"},
				"license": {"id": "P1D3XYZ", "type": "license"}
			}
		]}`))
	})

	resp, err := client.Users.ListAllWithLicenses(&ListUsersOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := []*User{
		{
			ID:      "P1D3Z4B",
			Role:    "user",
			License: &LicenseReference{ID: "P1D3XYZ", Type: "license_reference"},
		},
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

func TestUsersUpdateExistentContactMethod(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactMethod{Address: "foo@bar.com", Type: "email_contact_method"}
	// Counter to ensure that first call to PUT method returns unique contact
	// error, but not the following calls.
	var putReqCount int

	mux.HandleFunc("/users/1/contact_methods/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			v := new(ContactMethodPayload)
			json.NewDecoder(r.Body).Decode(v)
			if !reflect.DeepEqual(v.ContactMethod, input) {
				t.Errorf("Request body = %+v, want %+v", v, input)
			}
			if putReqCount == 0 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error":{"errors":["User Contact method must be unique"],"code":2001,"message":"Invalid Input Provided"}}`))
				putReqCount++
			} else {
				w.Write([]byte(`{"contact_method": { "address": "foo@bar.com", "id": "1", "type": "email_contact_method", "self":"api/users/1/contact_methods/1" }}`))
			}
		} else if r.Method == "DELETE" {
			w.WriteHeader(http.StatusNoContent)
		} else if r.Method == "GET" {
			w.Write([]byte(`{"contact_method": { "address": "foo@bar.com", "id": "1", "type": "email_contact_method", "self":"api/users/1/contact_methods/1" }}`))
		} else {
			t.Errorf("Request method: %v is neither PUT or GET", r.Method)
		}
	})

	mux.HandleFunc("/users/1/contact_methods", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"contact_methods": [{ "address": "foo@bar.com", "id": "1", "type": "email_contact_method", "self":"api/users/1/contact_methods/1" }] }`))
	})

	resp, _, err := client.Users.UpdateContactMethod("1", "1", input)
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

func TestUsersAddDuplicateNotificationRule(t *testing.T) {
	setup()
	defer teardown()

	input := &NotificationRule{
		ContactMethod:       &ContactMethodReference{ID: "c1", Type: "phone_contact_method"},
		StartDelayInMinutes: 0,
		Urgency:             "high",
	}

	mux.HandleFunc("/users/1/notification_rules", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			v := new(NotificationRulePayload)
			json.NewDecoder(r.Body).Decode(v)
			if !reflect.DeepEqual(v.NotificationRule, input) {
				t.Errorf("Request body = %+v, want %+v", v, input)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":{"errors":["Channel Start delay must be unique for a given contact method"],"code":2001,"message":"Invalid Input Provided"}}`))
		} else if r.Method == "GET" {
			w.Write([]byte(`{"notification_rules": [{ "id": "n1", "urgency": "high", "start_delay_in_minutes": 0, "contact_method": {"id": "c1", "type": "phone_contact_method"} }] }`))
		} else {
			t.Errorf("Request method : %v is neither POST nor GET", r.Method)
		}
	})

	mux.HandleFunc("/users/1/notification_rules/n1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"notification_rule": { "id": "n1", "urgency": "high", "start_delay_in_minutes": 0, "contact_method": {"id": "c1", "type": "phone_contact_method"}}}`))
	})

	resp, _, err := client.Users.CreateNotificationRule("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &NotificationRule{
		ContactMethod:       &ContactMethodReference{ID: "c1", Type: "phone_contact_method"},
		StartDelayInMinutes: 0,
		Urgency:             "high",
		ID:                  "n1",
	}

	if !(resp.ContactMethod.ID == want.ContactMethod.ID &&
		resp.ContactMethod.Type == want.ContactMethod.Type &&
		resp.StartDelayInMinutes == want.StartDelayInMinutes &&
		resp.Urgency == want.Urgency &&
		resp.ID == want.ID) {
		t.Errorf("returned %#v; want %#v", resp, want)
	}
}
