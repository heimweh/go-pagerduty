package pagerduty

import (
	"fmt"
)

type EventOrchestrationService service

type EventOrchestration struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Team        *EventOrchestrationObject `json:"team,omitempty"`
	// TODO: add Integrations, Routes, Updater, Creator + expand tests to verify these props
}

type EventOrchestrationObject struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type EventOrchestrationPayload struct {
	Orchestration *EventOrchestration `json:"orchestration,omitempty"`
}

var eventOrchestrationBaseUrl = "/event_orchestrations"

func (s *EventOrchestrationService) Create(orchestration *EventOrchestration) (*EventOrchestration, *Response, error) {
	v := new(EventOrchestrationPayload)
	p := &EventOrchestrationPayload{Orchestration: orchestration}

	resp, err := s.client.newRequestDo("POST", eventOrchestrationBaseUrl, nil, p, v)

	if err != nil {
		return nil, nil, err
	}

	return v.Orchestration, resp, nil
}

func (s *EventOrchestrationService) Get(ID string) (*EventOrchestration, *Response, error) {
	u := fmt.Sprintf("%s/%s", eventOrchestrationBaseUrl, ID)
	v := new(EventOrchestrationPayload)
	p := &EventOrchestrationPayload{}

	resp, err := s.client.newRequestDo("GET", u, nil, p, v)
	if err != nil {
		return nil, nil, err
	}

	return v.Orchestration, resp, nil
}

func (s *EventOrchestrationService) Update(ID string, orchestration *EventOrchestration) (*EventOrchestration, *Response, error) {
	u := fmt.Sprintf("%s/%s", eventOrchestrationBaseUrl, ID)
	v := new(EventOrchestrationPayload)
	p := &EventOrchestrationPayload{Orchestration: orchestration}

	resp, err := s.client.newRequestDo("PUT", u, nil, p, v)
	if err != nil {
		return nil, nil, err
	}

	return v.Orchestration, resp, nil
}

func (s *EventOrchestrationService) Delete(ID string) (*Response, error) {
	u := fmt.Sprintf("%s/%s", eventOrchestrationBaseUrl, ID)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}
