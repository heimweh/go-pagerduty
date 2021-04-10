package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestResponsePlayList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/response_plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 0, "offset": 0, "more": false, "limit": 0, "response_plays":[{"id": "1"}]}`))
	})

	o := &ListResponsePlayOptions{
		From: "foo@email.com",
	}

	resp, _, err := client.ResponsePlays.List(o)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListResponsePlaysResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		ResponsePlays: []*ResponsePlay{
			{
				ID: "1",
			},
		},
	}
	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestResponsePlayCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &ResponsePlay{Name: "foo"}

	mux.HandleFunc("/response_plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(ResponsePlay)
		v.Name = "foo"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"response_play":{"name": "foo", "id":"1"}}`))
	})

	resp, _, err := client.ResponsePlays.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &ResponsePlay{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
func TestResponsePlayGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"response_play":{"name": "foo", "id":"1"}}`))
	})

	ID := "1"
	f := "foo@email.com"

	resp, _, err := client.ResponsePlays.Get(ID, f)

	if err != nil {
		t.Fatal(err)
	}

	want := &ResponsePlay{
		Name:      "foo",
		ID:        "1",
		FromEmail: f,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestResponsePlayUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &ResponsePlay{
		Name: "foo",
	}

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(ResponsePlay)
		v.Name = "foo"

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"response_play":{"name": "foo", "id":"1"}}`))
	})

	resp, _, err := client.ResponsePlays.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &ResponsePlay{
		Name: "foo",
		ID:   "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestResponsePlayDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ID := "1"
	f := "foo@email.com"

	if _, err := client.ResponsePlays.Delete(ID, f); err != nil {
		t.Fatal(err)
	}
}
