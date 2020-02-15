package pagerduty

import "fmt"

// RulesetService handles the communication with rulesets
// related methods of the PagerDuty API.
type RulesetService service

// Ruleset represents a ruleset.
type Ruleset struct {
	ID          string        `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Type        string        `json:"type,omitempty"`
	RoutingKeys []interface{} `json:"routing_keys,omitempty"`
	Team        string        `json:"team,omitempty"`
	Updater     string        `json:"updater,omitempty"`
	Creator     string        `json:"creator,omitempty"`
}

// RulesetPayload represents payload with a ruleset object
type RulesetPayload struct {
	Ruleset *Ruleset `json:"ruleset,omitempty"`
}

// ListRulesetsResponse represents a list response of rulesets.
type ListRulesetsResponse struct {
	Total    int        `json:"total,omitempty"`
	Rulesets []*Ruleset `json:"rulesets,omitempty"`
	Offset   int        `json:"offset,omitempty"`
	More     bool       `json:"more,omitempty"`
	Limit    int        `json:"limit,omitempty"`
}

// RulesetRule represents a Ruleset rule
type RulesetRule struct {
	ID                 string         `json:"id,omitempty"`
	Position           int            `json:"position,omitempty"`
	Disabled           bool           `json:"disabled,omitempty"`
	Conditions         RuleConditions `json:"conditions,omitempty"`
	AdvancedConditions []interface{}  `json:"advanced_conditions,omitempty"`
	Actions            []*RuleAction  `json:"actions,omitempty"`
}

// RulesetRulePayload represents a payload for ruleset rules
type RulesetRulePayload struct {
	Rule *RulesetRule `json:"rule,omitempty"`
}

// RuleConditions represents the conditions field for a Ruleset
type RuleConditions struct {
	Operator          string              `json:"operator,omitempty"`
	RuleSubconditions []*RuleSubcondition `json:"subconditions,omitempty"`
}

// RuleSubcondition represents a subcondition of a ruleset condition
type RuleSubcondition struct {
	Operator   string                `json:"operator,omitempty"`
	Parameters []*ConditionParameter `json:"parameters,omitempty"`
}

// ConditionParameter represents  parameters in a rule condition
type ConditionParameter struct {
	Path  string `json:"path,omitempty"`
	Value string `json:"value,omitempty"`
}

// ListRulesetRulesResponse represents a list of rules in a ruleset
type ListRulesetRulesResponse struct {
	Total  int            `json:"total,omitempty"`
	Rules  []*RulesetRule `json:"rules,omitempty"`
	Offset int            `json:"offset,omitempty"`
	More   bool           `json:"more,omitempty"`
	Limit  int            `json:"limit,omitempty"`
}

// // RuleAdvancedCondition represents advanced conditions for rules
// type RuleAdvancedCondition struct {

// }

// RuleAction represents a rule action
type RuleAction struct {
	Type       string            `json:"type,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// List lists existing rulesets.
func (s *RulesetService) List() (*ListRulesetsResponse, *Response, error) {
	u := "/rulesets"
	v := new(ListRulesetsResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new ruleset.
func (s *RulesetService) Create(ruleset *Ruleset) (*RulesetPayload, *Response, error) {
	u := "/rulesets"
	v := new(RulesetPayload)
	p := RulesetPayload{Ruleset: ruleset}

	resp, err := s.client.newRequestDo("POST", u, nil, p, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete deletes an existing ruleset.
func (s *RulesetService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/rulesets/%s", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}

// Update updates an existing ruleset.
func (s *RulesetService) Update(id string, ruleset *Ruleset) (*RulesetPayload, *Response, error) {
	u := fmt.Sprintf("/rulesets/%s", id)
	v := new(RulesetPayload)
	p := RulesetPayload{Ruleset: ruleset}

	resp, err := s.client.newRequestDo("PUT", u, nil, p, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// ListRules Lists Event Rules for Ruleset
func (s *RulesetService) ListRules(rulesetID string) (*ListRulesetRulesResponse, *Response, error) {
	u := fmt.Sprintf("/rulesets/%s/rules", rulesetID)
	v := new(ListRulesetRulesResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// CreateRule for Ruleset
func (s *RulesetService) CreateRule(rulesetID string, rule *RulesetRule) (*RulesetRulePayload, *Response, error) {
	u := fmt.Sprintf("/rulesets/%s/rules", rulesetID)
	v := new(RulesetRulePayload)
	p := RulesetRulePayload{Rule: rule}

	resp, err := s.client.newRequestDo("POST", u, nil, p, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// GetRule for Ruleset
func (s *RulesetService) GetRule(rulesetID, ruleID string) (*RulesetRulePayload, *Response, error) {
	u := fmt.Sprintf("/rulesets/%s/rules/%s", rulesetID, ruleID)
	v := new(RulesetRulePayload)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// UpdateRule for Ruleset
func (s *RulesetService) UpdateRule(rulesetID, ruleID string, rule *RulesetRule) (*RulesetRulePayload, *Response, error) {
	u := fmt.Sprintf("/rulesets/%s/rules/%s", rulesetID, ruleID)
	v := new(RulesetRulePayload)
	p := RulesetRulePayload{Rule: rule}

	resp, err := s.client.newRequestDo("PUT", u, nil, p, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// DeleteRule deletes an existing rule from the ruleset.
func (s *RulesetService) DeleteRule(rulesetID, ruleID string) (*Response, error) {
	u := fmt.Sprintf("/rulesets/%s/rules/%s", rulesetID, ruleID)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}
