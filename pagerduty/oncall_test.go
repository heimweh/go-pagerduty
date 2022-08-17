package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestOnCallList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oncalls", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"oncalls":[{"escalation_level":2,"start":"2015-03-06T15:28:51-05:00","end":"2015-03-07T15:28:51-05:00"}]}`))
	})

	resp, _, err := client.OnCall.List(&ListOnCallOptions{})
	if err != nil {
		t.Fatal(err)
	}

	start := "2015-03-06T15:28:51-05:00"
	end := "2015-03-07T15:28:51-05:00"
	want := &ListOnCallResponse{
		Oncalls: []*OnCall{
			{
				EscalationLevel: 2,
				Start:           &start,
				End:             &end,
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
