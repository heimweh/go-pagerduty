package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentWorkflowList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryMaxCount(t, r, 1)
		offset := r.URL.Query().Get("offset")

		switch offset {
		case "":
			w.Write([]byte(`{"total": 2, "offset": 0, "more": true, "limit": 1, "incident_workflows":[{"id": "1"}]}`))
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "incident_workflows":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.IncidentWorkflows.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		IncidentWorkflows: []*IncidentWorkflow{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowList_SecondPage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		offset := r.URL.Query().Get("offset")

		switch offset {
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "incident_workflows":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.IncidentWorkflows.List(&ListIncidentWorkflowOptions{Offset: 1})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		IncidentWorkflows: []*IncidentWorkflow{
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowList_Limit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		testQueryValue(t, r, "limit", "42")

		w.Write([]byte(`{"total": 1, "offset": 0, "more": false, "limit": 42, "incident_workflows":[{"id": "2"}]}`))

	})

	resp, _, err := client.IncidentWorkflows.List(&ListIncidentWorkflowOptions{Limit: 42})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowResponse{
		Total:  1,
		Offset: 0,
		More:   false,
		Limit:  42,
		IncidentWorkflows: []*IncidentWorkflow{
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowList_WithOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		testQueryMinCount(t, r, 1)
		testQueryMaxCount(t, r, 2)
		testQueryValue(t, r, "include[]", "steps")

		offset := r.URL.Query().Get("offset")
		switch offset {
		case "":
			w.Write([]byte(`{"total": 2, "offset": 0, "more": true, "limit": 1, "incident_workflows":[{"id": "1"}]}`))
		case "1":
			w.Write([]byte(`{"total": 2, "offset": 1, "more": false, "limit": 1, "incident_workflows":[{"id": "2"}]}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}

	})

	resp, _, err := client.IncidentWorkflows.List(&ListIncidentWorkflowOptions{
		Includes: []string{"steps"},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListIncidentWorkflowResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		IncidentWorkflows: []*IncidentWorkflow{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/IW1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testBody(t, r, "")

		w.Write([]byte(`
{
  "incident_workflow": {
    "id": "TO38234",
    "type": "incident_workflow",
    "name": "Example Workflow",
    "description": "This workflow serves as an example",
    "self": "https://api.pagerduty.com/incident_workflows/TO38234",
    "html_url": "https://subdomain.pagerduty.com/flex-workflows/workflows/TO38234/edit",
    "created_at": "2022-06-07T00:01:55Z",
    "last_started_at": "2022-06-07T00:01:55Z",
    "team": {
        "type": "team_reference",
        "id": "T1"
    },
    "steps": [
      {
        "id": "32OIHWEJ",
        "type": "step",
        "name": "Example Step",
        "description": "An example workflow step",
        "action_configuration": {
            "action_id": "example/action/v1",
            "description": "Description of the example action",
            "inputs": [
                {
                    "name": "Example input",
                    "parameter_type": "text",
                    "value": "{{ example-value }}"
                }
            ],
            "outputs": [
                {
                    "name": "Example output",
                    "reference_name": "example-output",
                    "parameter_type": "text"
                }
            ]
        }
      },
      {
        "id": "D3IT0D3",
        "type": "step",
        "name": "Subsequent Step",
        "description": "A subsequent step in this workflow",
        "action_configuration": {
            "action_id": "example/action/v1",
            "description": "Description of the example action",
            "inputs": [
                {
                    "name": "Example input",
                    "parameter_type": "text",
                    "value": "{{ example-value }}"
                }
            ],
            "inline_steps_inputs": [
                {
                    "name": "Example inline_steps_input",
                    "value": {
                        "steps": [
                            {
                                "name": "Inline Step 1",
                                "action_configuration": {
                                    "action_id": "example/action/v1",
                                    "inputs": [
                                        {
                                            "name": "Example input",
                                            "value": "B"
                                        }
                                    ]
                                }
                            }
                        ]
                    }
                }
            ],
            "outputs": [
                {
                    "name": "Example output",
                    "reference_name": "example-output",
                    "parameter_type": "text"
                }
            ]
        }
      }
    ]
  }
}
`))

	})

	resp, _, err := client.IncidentWorkflows.Get("IW1")
	if err != nil {
		t.Fatal(err)
	}

	workflowDesc := "This workflow serves as an example"
	firstStepDesc := "An example workflow step"
	secondStepDesc := "A subsequent step in this workflow"
	actionDesc := "Description of the example action"

	want := &IncidentWorkflow{
		ID:          "TO38234",
		Type:        "incident_workflow",
		Name:        "Example Workflow",
		Description: &workflowDesc,
		Self:        "https://api.pagerduty.com/incident_workflows/TO38234",
		Team: &TeamReference{
			Type: "team_reference",
			ID:   "T1",
		},
		Steps: []*IncidentWorkflowStep{
			{
				ID:          "32OIHWEJ",
				Type:        "step",
				Name:        "Example Step",
				Description: &firstStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID:    "example/action/v1",
					Description: &actionDesc,
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
				},
			}, {
				ID:          "D3IT0D3",
				Type:        "step",
				Name:        "Subsequent Step",
				Description: &secondStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID:    "example/action/v1",
					Description: &actionDesc,
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
					InlineStepsInputs: []*IncidentWorkflowActionInlineStepsInput{
						{
							Name: "Example inline_steps_input",
							Value: &IncidentWorkflowActionInlineStepsInputValue{
								Steps: []*IncidentWorkflowActionInlineStep{
									{
										Name: "Inline Step 1",
										Configuration: &IncidentWorkflowActionConfiguration{
											ActionID: "example/action/v1",
											Inputs: []*IncidentWorkflowActionInput{
												{
													Name:  "Example input",
													Value: "B",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"incident_workflow":{"name":"Example Workflow","description":"This workflow serves as an example","steps":[{"name":"Example Step","description":"An example workflow step","action_configuration":{"action_id":"example/action/v1","inputs":[{"name":"Example input","value":"{{ example-value }}"}]}},{"name":"Subsequent Step","description":"A subsequent step in this workflow","action_configuration":{"action_id":"example/action/v1","inputs":[{"name":"Example input","value":"{{ example-value }}"}],"inline_steps_inputs":[{"name":"Example inline_steps_input","value":{"steps":[{"name":"Inline Step 1","action_configuration":{"action_id":"example/action/v1","inputs":[{"name":"Example input","value":"B"}]}}]}}]}}]}}`)

		w.Write([]byte(`
{
  "incident_workflow": {
    "id": "TO38234",
    "type": "incident_workflow",
    "name": "Example Workflow",
    "description": "This workflow serves as an example",
    "self": "https://api.pagerduty.com/incident_workflows/TO38234",
    "html_url": "https://subdomain.pagerduty.com/flex-workflows/workflows/TO38234/edit",
    "created_at": "2022-06-07T00:01:55Z",
    "last_started_at": "2022-06-07T00:01:55Z",
    "steps": [
      {
        "id": "32OIHWEJ",
        "type": "step",
        "name": "Example Step",
        "description": "An example workflow step",
        "action_configuration": {
            "action_id": "example/action/v1",
            "description": "Description of the example action",
            "inputs": [
                {
                    "name": "Example input",
                    "parameter_type": "text",
                    "value": "{{ example-value }}"
                }
            ],
            "outputs": [
                {
                    "name": "Example output",
                    "reference_name": "example-output",
                    "parameter_type": "text"
                }
            ]
        }
      },
      {
        "id": "D3IT0D3",
        "type": "step",
        "name": "Subsequent Step",
        "description": "A subsequent step in this workflow",
        "action_configuration": {
            "action_id": "example/action/v1",
            "description": "Description of the example action",
            "inputs": [
                {
                    "name": "Example input",
                    "parameter_type": "text",
                    "value": "{{ example-value }}"
                }
            ],
            "inline_steps_inputs": [
                {
                    "name": "Example inline_steps_input",
                    "value": {
                        "steps": [
                            {
                                "name": "Inline Step 1",
                                "action_configuration": {
                                    "action_id": "example/action/v1",
                                    "inputs": [
                                        {
                                            "name": "Example input",
                                            "value": "B"
                                        }
                                    ]
                                }
                            }
                        ]
                    }
                }
            ],
            "outputs": [
                {
                    "name": "Example output",
                    "reference_name": "example-output",
                    "parameter_type": "text"
                }
            ]
        }
      }
    ]
  }
}
`))

	})

	workflowDesc := "This workflow serves as an example"
	firstStepDesc := "An example workflow step"
	secondStepDesc := "A subsequent step in this workflow"
	actionDesc := "Description of the example action"

	resp, _, err := client.IncidentWorkflows.Create(&IncidentWorkflow{
		Name:        "Example Workflow",
		Description: &workflowDesc,
		Steps: []*IncidentWorkflowStep{
			{
				Name:        "Example Step",
				Description: &firstStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID: "example/action/v1",
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
				},
			}, {
				Name:        "Subsequent Step",
				Description: &secondStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID: "example/action/v1",
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
					InlineStepsInputs: []*IncidentWorkflowActionInlineStepsInput{
						{
							Name: "Example inline_steps_input",
							Value: &IncidentWorkflowActionInlineStepsInputValue{
								Steps: []*IncidentWorkflowActionInlineStep{
									{
										Name: "Inline Step 1",
										Configuration: &IncidentWorkflowActionConfiguration{
											ActionID: "example/action/v1",
											Inputs: []*IncidentWorkflowActionInput{
												{
													Name:  "Example input",
													Value: "B",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentWorkflow{
		ID:          "TO38234",
		Type:        "incident_workflow",
		Name:        "Example Workflow",
		Description: &workflowDesc,
		Self:        "https://api.pagerduty.com/incident_workflows/TO38234",
		Steps: []*IncidentWorkflowStep{
			{
				ID:          "32OIHWEJ",
				Type:        "step",
				Name:        "Example Step",
				Description: &firstStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID:    "example/action/v1",
					Description: &actionDesc,
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
				},
			}, {
				ID:          "D3IT0D3",
				Type:        "step",
				Name:        "Subsequent Step",
				Description: &secondStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID:    "example/action/v1",
					Description: &actionDesc,
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
					InlineStepsInputs: []*IncidentWorkflowActionInlineStepsInput{
						{
							Name: "Example inline_steps_input",
							Value: &IncidentWorkflowActionInlineStepsInputValue{
								Steps: []*IncidentWorkflowActionInlineStep{
									{
										Name: "Inline Step 1",
										Configuration: &IncidentWorkflowActionConfiguration{
											ActionID: "example/action/v1",
											Inputs: []*IncidentWorkflowActionInput{
												{
													Name:  "Example input",
													Value: "B",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestIncidentWorkflowDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/IW1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(200)
	})

	resp, err := client.IncidentWorkflows.Delete("IW1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 200 {
		t.Errorf("unexpected response code. want 200. got %v", resp.Response.StatusCode)
	}
}

func TestIncidentWorkflowUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/incident_workflows/IW1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"incident_workflow":{"description":"Updated description","steps":[{"id":"32OIHWEJ"},{"id":"D3IT0D3","name":"Subsequent Step Updated Name","action_configuration":{"action_id":"example/action/v1","inputs":[{"name":"Example input","value":"{{ example-value }}"}],"inline_steps_inputs":[{"name":"Example inline_steps_input","value":{"steps":[{"name":"Inline Step 1 Updated Name","action_configuration":{"action_id":"example/action/v1","inputs":[{"name":"Example input","value":"B"}]}}]}}]}}]}}`)

		w.Write([]byte(`
{
  "incident_workflow": {
    "id": "TO38234",
    "type": "incident_workflow",
    "name": "Example Workflow",
    "description": "Updated description",
    "self": "https://api.pagerduty.com/incident_workflows/TO38234",
    "html_url": "https://subdomain.pagerduty.com/flex-workflows/workflows/TO38234/edit",
    "created_at": "2022-06-07T00:01:55Z",
    "last_started_at": "2022-06-07T00:01:55Z",
    "steps": [
      {
        "id": "32OIHWEJ",
        "type": "step",
        "name": "Example Step",
        "description": "An example workflow step",
        "action_configuration": {
            "action_id": "example/action/v1",
            "description": "Description of the example action",
            "inputs": [
                {
                    "name": "Example input",
                    "parameter_type": "text",
                    "value": "{{ example-value }}"
                }
            ],
            "outputs": [
                {
                    "name": "Example output",
                    "reference_name": "example-output",
                    "parameter_type": "text"
                }
            ]
        }
      },
      {
        "id": "D3IT0D3",
        "type": "step",
        "name": "Subsequent Step Updated Name",
        "description": "A subsequent step in this workflow",
        "action_configuration": {
            "action_id": "example/action/v1",
            "description": "Description of the example action",
            "inputs": [
                {
                    "name": "Example input",
                    "parameter_type": "text",
                    "value": "{{ example-value }}"
                }
            ],
            "inline_steps_inputs": [
                {
                    "name": "Example inline_steps_input",
                    "value": {
                        "steps": [
                            {
                                "name": "Inline Step 1 Updated Name",
                                "action_configuration": {
                                    "action_id": "example/action/v1",
                                    "inputs": [
                                        {
                                            "name": "Example input",
                                            "value": "B"
                                        }
                                    ]
                                }
                            }
                        ]
                    }
                }
            ],
            "outputs": [
                {
                    "name": "Example output",
                    "reference_name": "example-output",
                    "parameter_type": "text"
                }
            ]
        }
      }
    ]
  }
}
`))

	})

	updatedWorkflowDesc := "Updated description"
	firstStepDesc := "An example workflow step"
	secondStepDesc := "A subsequent step in this workflow"
	actionDesc := "Description of the example action"

	resp, _, err := client.IncidentWorkflows.Update("IW1", &IncidentWorkflow{
		Description: &updatedWorkflowDesc,
		Steps: []*IncidentWorkflowStep{
			{
				ID: "32OIHWEJ",
			},
			{
				ID:   "D3IT0D3",
				Name: "Subsequent Step Updated Name",
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID: "example/action/v1",
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
					InlineStepsInputs: []*IncidentWorkflowActionInlineStepsInput{
						{
							Name: "Example inline_steps_input",
							Value: &IncidentWorkflowActionInlineStepsInputValue{
								Steps: []*IncidentWorkflowActionInlineStep{
									{
										Name: "Inline Step 1 Updated Name",
										Configuration: &IncidentWorkflowActionConfiguration{
											ActionID: "example/action/v1",
											Inputs: []*IncidentWorkflowActionInput{
												{
													Name:  "Example input",
													Value: "B",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &IncidentWorkflow{
		ID:          "TO38234",
		Type:        "incident_workflow",
		Name:        "Example Workflow",
		Description: &updatedWorkflowDesc,
		Self:        "https://api.pagerduty.com/incident_workflows/TO38234",
		Steps: []*IncidentWorkflowStep{
			{
				ID:          "32OIHWEJ",
				Type:        "step",
				Name:        "Example Step",
				Description: &firstStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID:    "example/action/v1",
					Description: &actionDesc,
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
				},
			}, {
				ID:          "D3IT0D3",
				Type:        "step",
				Name:        "Subsequent Step Updated Name",
				Description: &secondStepDesc,
				Configuration: &IncidentWorkflowActionConfiguration{
					ActionID:    "example/action/v1",
					Description: &actionDesc,
					Inputs: []*IncidentWorkflowActionInput{
						{
							Name:  "Example input",
							Value: "{{ example-value }}",
						},
					},
					InlineStepsInputs: []*IncidentWorkflowActionInlineStepsInput{
						{
							Name: "Example inline_steps_input",
							Value: &IncidentWorkflowActionInlineStepsInputValue{
								Steps: []*IncidentWorkflowActionInlineStep{
									{
										Name: "Inline Step 1 Updated Name",
										Configuration: &IncidentWorkflowActionConfiguration{
											ActionID: "example/action/v1",
											Inputs: []*IncidentWorkflowActionInput{
												{
													Name:  "Example input",
													Value: "B",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
