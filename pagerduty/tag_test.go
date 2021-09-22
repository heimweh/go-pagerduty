package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTagsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"tags": [{"id": "1"}]}`))
	})

	o := new(ListTagsOptions)

	resp, _, err := client.Tags.List(o)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListTagsResponse{
		Tags: []*Tag{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTagsCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Tag{
		Label: "foo",
		Type:  "tag",
	}

	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Tag)
		v.Label = "foo"
		v.Type = "tag"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"tag": {"label": "foo", "id": "1"}}`))
	})

	resp, _, err := client.Tags.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Tag{
		Label: "foo",
		ID:    "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestTagsDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Tags.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestTagsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"tag": {"id": "1","label": "foo"}}`))
	})

	resp, _, err := client.Tags.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Tag{
		ID:    "1",
		Label: "foo",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// assign add test
func TestTagsAssignAdd(t *testing.T) {
	setup()
	defer teardown()

	input := &TagAssignments{
		Add: []*TagAssignment{
			{
				Type:  "tag_reference",
				TagID: "1",
			},
			{
				Type:  "tag",
				Label: "NewTag",
			},
		},
	}

	mux.HandleFunc("/users/1/change_tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(TagAssignments)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
	})
	// this endpoint only returns  an "ok" in the body. no point in testing for it.
	if _, err := client.Tags.Assign("users", "1", input); err != nil {
		t.Fatal(err)
	}
}

// assign remove test
func TestTagsAssignRemove(t *testing.T) {
	setup()
	defer teardown()

	input := &TagAssignments{
		Remove: []*TagAssignment{
			{
				Type:  "tag_reference",
				TagID: "1",
			},
		},
	}

	mux.HandleFunc("/users/1/change_tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(TagAssignments)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
	})
	// this endpoint only returns  an "ok" in the body. no point in testing for it.
	if _, err := client.Tags.Assign("users", "1", input); err != nil {
		t.Fatal(err)
	}
}
