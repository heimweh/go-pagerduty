package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPriorityList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/priorities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 1, "offset": 0, "more": false, "limit": 25, "priorities":[{"id": "1"}]}`))
	})

	resp, _, err := client.Priorities.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListPrioritiesResponse{
		Total:  1,
		Offset: 0,
		More:   false,
		Limit:  25,
		Priorities: []*Priority{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
func TestPriorityListFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/priorities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusForbidden)
	})

	if _, _, err := client.Priorities.List(); err == nil {
		t.Fatal("expected error; got nil")
	}
}
