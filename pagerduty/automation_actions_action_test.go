package pagerduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAutomationActionsActionTypeScriptGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/actions/01DF4OBNYKW84FS9CCYVYS1MOS", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"action":{"action_data_reference":{"script":"java --version","invocation_command":"sh"},"action_type":"script","action_classification":"diagnostic","creation_time":"2022-12-12T18:51:42.048162Z","id":"01DF4OBNYKW84FS9CCYVYS1MOS","name":"Script Action created by TF","type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Get("01DF4OBNYKW84FS9CCYVYS1MOS")
	if err != nil {
		t.Fatal(err)
	}

	script := "java --version"
	invocation_command := "sh"
	classification := "diagnostic"
	adf := AutomationActionsActionDataReference{
		Script:            &script,
		InvocationCommand: &invocation_command,
	}
	resource_type := "action"
	creation_time := "2022-12-12T18:51:42.048162Z"
	want := &AutomationActionsAction{
		ID:                   "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:                 "Script Action created by TF",
		CreationTime:         &creation_time,
		ActionType:           "script",
		Type:                 &resource_type,
		ActionClassification: &classification,
		ActionDataReference:  adf,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionTypeProcessAutomationGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/actions/01DF4OBNYKW84FS9CCYVYS1MOS", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"action":{"action_data_reference":{"process_automation_job_id":"1519578e-a22a-4340-b58f-08194691e10b"},"action_type":"process_automation","creation_time":"2022-12-12T18:51:42.048162Z","id":"01DF4OBNYKW84FS9CCYVYS1MOS","name":"Action created by TF","privileges":{"permissions":["read"]},"type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Get("01DF4OBNYKW84FS9CCYVYS1MOS")
	if err != nil {
		t.Fatal(err)
	}

	job_id := "1519578e-a22a-4340-b58f-08194691e10b"
	adf := AutomationActionsActionDataReference{
		ProcessAutomationJobId: &job_id,
	}
	permissions_read := "read"
	resource_type := "action"
	creation_time := "2022-12-12T18:51:42.048162Z"
	want := &AutomationActionsAction{
		ID:                  "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:                "Action created by TF",
		CreationTime:        &creation_time,
		ActionType:          "process_automation",
		Type:                &resource_type,
		ActionDataReference: adf,
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionTypeProcessAutomationCreate(t *testing.T) {
	setup()
	defer teardown()

	description := "Description of Action created by TF"
	runner_id := "01DF4O9T1MDPYOUT7SUX9EXZ4R"
	adf_arg := "-arg 123"
	job_id := "1519578e-a22a-4340-b58f-08194691e10b"
	adf := AutomationActionsActionDataReference{
		ProcessAutomationJobId:        &job_id,
		ProcessAutomationJobArguments: &adf_arg,
	}
	input := &AutomationActionsAction{
		Name:                "Action created by TF",
		Description:         &description,
		ActionType:          "process_automation",
		RunnerID:            &runner_id,
		ActionDataReference: adf,
	}

	mux.HandleFunc("/automation_actions/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(AutomationActionsActionPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Action, input) {
			t.Errorf("Request body = %+v, want %+v", v.Action, input)
		}
		w.Write([]byte(`{"action":{"action_data_reference":{"process_automation_job_id":"1519578e-a22a-4340-b58f-08194691e10b","process_automation_job_arguments":"-arg 123"},"action_type":"process_automation","creation_time":"2022-12-12T18:51:42.048162Z","description":"Description of Action created by TF","id":"01DF4OBNYKW84FS9CCYVYS1MOS","last_run":"2022-12-12T18:52:11.937747Z","last_run_by":{"id":"PINL781","type":"user_reference"},"modify_time":"2022-12-12T18:51:42.048162Z","name":"Action created by TF","privileges":{"permissions":["read"]},"runner":"01DF4O9T1MDPYOUT7SUX9EXZ4R","runner_type":"runbook","services":[{"id":"PQWQ0U6","type":"service_reference"}],"teams":[{"id":"PZ31N6S","type":"team_reference"}],"type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	runner_type_runbook := "runbook"
	modify_time := "2022-12-12T18:51:42.048162Z"
	permissions_read := "read"
	resource_type := "action"
	creation_time := "2022-12-12T18:51:42.048162Z"
	want := &AutomationActionsAction{
		ID:           "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:         "Action created by TF",
		Description:  &description,
		CreationTime: &creation_time,
		ActionType:   "process_automation",
		Type:         &resource_type,
		RunnerID:     &runner_id,
		RunnerType:   &runner_type_runbook,
		Teams: []*TeamReference{
			{
				Type: "team_reference",
				ID:   "PZ31N6S",
			},
		},
		Services: []*ServiceReference{
			{
				Type: "service_reference",
				ID:   "PQWQ0U6",
			},
		},
		ActionDataReference: adf,
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
		ModifyTime: &modify_time,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionUpdate(t *testing.T) {
	setup()
	defer teardown()

	description := "Description of Action created by TF"
	runner_id := "01DF4O9T1MDPYOUT7SUX9EXZ4R"
	adf_arg := "-arg 123"
	job_id := "1519578e-a22a-4340-b58f-08194691e10b"
	adf := AutomationActionsActionDataReference{
		ProcessAutomationJobId:        &job_id,
		ProcessAutomationJobArguments: &adf_arg,
	}
	input := &AutomationActionsAction{
		Name:                "Action created by TF",
		Description:         &description,
		ActionType:          "process_automation",
		RunnerID:            &runner_id,
		ActionDataReference: adf,
	}

	var id = "01DF4OBNYKW84FS9CCYVYS1MOS"
	var url = fmt.Sprintf("%s/%s", automationActionsActionBaseUrl, id)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(AutomationActionsActionPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Action, input) {
			t.Errorf("Request body = %+v, want %+v", v.Action, input)
		}
		w.Write([]byte(`{"action":{"action_data_reference":{"process_automation_job_id":"1519578e-a22a-4340-b58f-08194691e10b","process_automation_job_arguments":"-arg 123"},"action_type":"process_automation","creation_time":"2022-12-12T18:51:42.048162Z","description":"Description of Action created by TF","id":"01DF4OBNYKW84FS9CCYVYS1MOS","last_run":"2022-12-12T18:52:11.937747Z","last_run_by":{"id":"PINL781","type":"user_reference"},"modify_time":"2022-12-12T18:51:42.048162Z","name":"Action created by TF","privileges":{"permissions":["read"]},"runner":"01DF4O9T1MDPYOUT7SUX9EXZ4R","runner_type":"runbook","services":[{"id":"PQWQ0U6","type":"service_reference"}],"teams":[{"id":"PZ31N6S","type":"team_reference"}],"type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Update(id, input)
	if err != nil {
		t.Fatal(err)
	}

	runner_type_runbook := "runbook"
	modify_time := "2022-12-12T18:51:42.048162Z"
	permissions_read := "read"
	resource_type := "action"
	creation_time := "2022-12-12T18:51:42.048162Z"
	want := &AutomationActionsAction{
		ID:           id,
		Name:         "Action created by TF",
		Description:  &description,
		CreationTime: &creation_time,
		ActionType:   "process_automation",
		Type:         &resource_type,
		RunnerID:     &runner_id,
		RunnerType:   &runner_type_runbook,
		Teams: []*TeamReference{
			{
				Type: "team_reference",
				ID:   "PZ31N6S",
			},
		},
		Services: []*ServiceReference{
			{
				Type: "service_reference",
				ID:   "PQWQ0U6",
			},
		},
		ActionDataReference: adf,
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
		ModifyTime: &modify_time,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/actions/01DF4OBNYKW84FS9CCYVYS1MOS", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.AutomationActionsAction.Delete("01DF4OBNYKW84FS9CCYVYS1MOS"); err != nil {
		t.Fatal(err)
	}
}

func TestAutomationActionsActionTypeScriptCreate(t *testing.T) {
	setup()
	defer teardown()

	description := "Description of Action created by TF"
	runner_id := "01DF4O9T1MDPYOUT7SUX9EXZ4R"
	invocation_command := "/bin/bash"
	script_data := "java --version"
	adf := AutomationActionsActionDataReference{
		Script:            &script_data,
		InvocationCommand: &invocation_command,
	}
	input := &AutomationActionsAction{
		Name:                "Action created by TF",
		Description:         &description,
		ActionType:          "script",
		RunnerID:            &runner_id,
		ActionDataReference: adf,
	}

	mux.HandleFunc("/automation_actions/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(AutomationActionsActionPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Action, input) {
			t.Errorf("Request body = %+v, want %+v", v.Action, input)
		}
		w.Write([]byte(`{"action":{"action_data_reference":{"script":"java --version","invocation_command":"/bin/bash"},"action_type":"script","creation_time":"2022-12-12T18:51:42.048162Z","description":"Description of Action created by TF","id":"01DF4OBNYKW84FS9CCYVYS1MOS","last_run":"2022-12-12T18:52:11.937747Z","last_run_by":{"id":"PINL781","type":"user_reference"},"modify_time":"2022-12-12T18:51:42.048162Z","name":"Action created by TF","privileges":{"permissions":["read"]},"runner":"01DF4O9T1MDPYOUT7SUX9EXZ4R","runner_type":"sidecar","services":[{"id":"PQWQ0U6","type":"service_reference"}],"teams":[{"id":"PZ31N6S","type":"team_reference"}],"type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	runner_type_sidecar := "sidecar"
	modify_time := "2022-12-12T18:51:42.048162Z"
	permissions_read := "read"
	resource_type := "action"
	creation_time := "2022-12-12T18:51:42.048162Z"
	want := &AutomationActionsAction{
		ID:           "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:         "Action created by TF",
		Description:  &description,
		CreationTime: &creation_time,
		ActionType:   "script",
		Type:         &resource_type,
		RunnerID:     &runner_id,
		RunnerType:   &runner_type_sidecar,
		Teams: []*TeamReference{
			{
				Type: "team_reference",
				ID:   "PZ31N6S",
			},
		},
		Services: []*ServiceReference{
			{
				Type: "service_reference",
				ID:   "PQWQ0U6",
			},
		},
		ActionDataReference: adf,
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
		ModifyTime: &modify_time,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionTeamAssociationCreate(t *testing.T) {
	setup()
	defer teardown()
	actionID := "01DA2MLYN0J5EFC1LKWXUKDDKT"
	teamID := "1"

	mux.HandleFunc(fmt.Sprintf("/automation_actions/actions/%s/teams", actionID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(`{"team":{"id":"1","type":"team_reference"}}`))
	})

	resp, _, err := client.AutomationActionsAction.AssociateToTeam(actionID, teamID)
	if err != nil {
		t.Fatal(err)
	}

	want := &AutomationActionsActionTeamAssociationPayload{
		&TeamReference{
			ID:   teamID,
			Type: "team_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionTeamAssociationDelete(t *testing.T) {
	setup()
	defer teardown()
	actionID := "01DA2MLYN0J5EFC1LKWXUKDDKT"
	teamID := "1"

	mux.HandleFunc(fmt.Sprintf("/automation_actions/actions/%s/teams/%s", actionID, teamID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.AutomationActionsAction.DissociateToTeam(actionID, teamID); err != nil {
		t.Fatal(err)
	}
}

func TestAutomationActionsActionTeamAssociationGet(t *testing.T) {
	setup()
	defer teardown()
	actionID := "01DA2MLYN0J5EFC1LKWXUKDDKT"
	teamID := "1"

	mux.HandleFunc(fmt.Sprintf("/automation_actions/actions/%s/teams/%s", actionID, teamID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"team":{"id":"1","type":"team_reference"}}`))
	})

	resp, _, err := client.AutomationActionsAction.GetAssociationToTeam(actionID, teamID)
	if err != nil {
		t.Fatal(err)
	}

	want := &AutomationActionsActionTeamAssociationPayload{
		&TeamReference{
			ID:   teamID,
			Type: "team_reference",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
