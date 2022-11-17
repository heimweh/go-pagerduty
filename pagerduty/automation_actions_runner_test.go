package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAutomationActionsRunnerGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/runners/01DA2MLYN0J5EFC1LKWXUKDDKT", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"runner":{"id":"01DA2MLYN0J5EFC1LKWXUKDDKT","name":"us-west-2 prod sidecar runner","type":"runner","creation_time":"2022-10-21T19:42:52.127369Z","runner_type":"sidecar","status":"Configured"}}`))
	})

	resp, _, err := client.AutomationActionsRunner.Get("01DA2MLYN0J5EFC1LKWXUKDDKT")
	if err != nil {
		t.Fatal(err)
	}

	want := &AutomationActionsRunner{
		ID:           "01DA2MLYN0J5EFC1LKWXUKDDKT",
		Name:         "us-west-2 prod sidecar runner",
		CreationTime: "2022-10-21T19:42:52.127369Z",
		RunnerType:   "sidecar",
		Type:         "runner",
		Status:       "Configured",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
