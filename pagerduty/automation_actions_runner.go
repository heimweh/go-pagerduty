package pagerduty

import "fmt"

// AutomationActionsRunner handles the communication with schedule
// related methods of the PagerDuty API.
type AutomationActionsRunnerService service

type AutomationActionsRunner struct {
	ID           string                       `json:"id"`
	Name         string                       `json:"name"`
	Summary      string                       `json:"summary,omitempty"`
	Type         string                       `json:"type"`
	Description  string                       `json:"description,omitempty"`
	CreationTime string                       `json:"creation_time"`
	RunnerType   string                       `json:"runner_type,omitempty"`
	Status       string                       `json:"status"`
	Teams        []*TeamReference             `json:"teams,omitempty"`
	Privileges   *AutomationActionsPriviliges `json:"privileges,omitempty"`
}

type AutomationActionsPriviliges struct {
	Permissions []*string `json:"permissions,omitempty"`
}

type AutomationActionsRunnerPayload struct {
	Runner *AutomationActionsRunner `json:"runner,omitempty"`
}

// Get retrieves information about a service.
func (s *AutomationActionsRunnerService) Get(id string) (*AutomationActionsRunner, *Response, error) {
	u := fmt.Sprintf("/automation_actions/runners/%s", id)
	v := new(AutomationActionsRunnerPayload)
	o := RequestOptions{
		Type:  "header",
		Label: "X-EARLY-ACCESS",
		Value: "automation-actions-early-access",
	}

	resp, err := s.client.newRequestDoOptions("GET", u, nil, nil, &v, o)
	if err != nil {
		return nil, nil, err
	}

	return v.Runner, resp, nil
}
