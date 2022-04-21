package pagerduty

type OrchestrationService service

type Orchestration struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Team        *OrchestrationObject `json:"team,omitempty"`
	// TODO: Updater, Creator
}

type OrchestrationObject struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type OrchestrationPayload struct {
	Orchestration *Orchestration `json:"orchestration,omitempty"`
}

func (s *OrchestrationService) Create(orchestration *Orchestration) (*Orchestration, *Response, error) {
	u := "/orchestrations"
	v := new(OrchestrationPayload)
	p := &OrchestrationPayload{Orchestration: orchestration}

	resp, err := s.client.newRequestDo("POST", u, nil, p, v)

	if err != nil {
		return nil, nil, err
	}

	return v.Orchestration, resp, nil
}
