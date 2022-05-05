package pagerduty

// import (
// 	"fmt"
// )

// TODO: Check omitempty for all structs
type EventOrchestrationPathService service

type EventOrchestrationPath struct {
	Type      string                          `json:"type,omitempty"`
	Parent    *EventOrchestrationPathObject   `json:"parent,omitempty"`
	Sets      []*EventOrchestrationPathSet    `json:"sets,omitempty"`
	CatchAll  *EventOrchestrationPathCatchAll `json:"catch_all,omitempty"`
	CreatedAt string                          `json:"created_at,omitempty"`
	CreatedBy *EventOrchestrationPathObject   `json:"created_by,omitempty"`
	UpdatedAt string                          `json:"updated_at,omitempty"`
	UpdatedBy *EventOrchestrationPathObject   `json:"updated_by,omitempty"`
	Version   string                          `json:"version,omitempty"`
}

type EventOrchestrationPathObject struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Self string `json:"self,omitempty"`
}

type EventOrchestrationPathSet struct {
	ID    string                        `json:"id,omitempty"`
	Rules []*EventOrchestrationPathRule `json:"rules,omitempty"`
}

type EventOrchestrationPathRule struct {
	ID         string                                 `json:"id,omitempty"`
	Label      string                                 `json:"label,omitempty"`
	Conditions []*EventOrchestrationPathRuleCondition `json:"conditions,omitempty"`
	Actions    *EventOrchestrationPathRuleAction      `json:"actions,omitempty"`
	Disabled   bool                                   `json:"disabled,omitempty"`
}

type EventOrchestrationPathRuleCondition struct {
	Expression string `json:"expression,omitempty"`
}

type EventOrchestrationPathAction struct {
	Suppress                   bool                                               `json:"suppress,omitempty"`
	Suspend                    int                                                `json:"suspend,omitempty"`
	Priority                   string                                             `json:"priority,omitempty"`
	Annotate                   string                                             `json:"annotate,omitempty"`
	PagerdutyAutomationActions []*EventOrchestrationPathPagerdutyAutomationAction `json:"pagerduty_automation_actions,omitempty"`
	AutomationActions          []*EventOrchestrationPathAutomationAction          `json:"automation_actions,omitempty"`
	Severity                   string                                             `json:"severity,omitempty"`
	EventAction                string                                             `json:"event_action,omitempty"`
	Variables                  []*EventOrchestrationPathActionVariables           `json:"variables,omitempty"`
	Extractions                []*EventOrchestrationPathActionExtractions         `json:"extractions,omitempty"`
}

type EventOrchestrationPathRuleAction struct {
	*EventOrchestrationPathAction
	RouteTo  string `json:"route_to,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

type EventOrchestrationPathPagerdutyAutomationAction struct {
	ActionId string `json:"action_id,omitempty"`
}

type EventOrchestrationPathAutomationAction struct {
	Name       string                                          `json:"name,omitempty"`
	Url        string                                          `json:"url,omitempty"`
	AutoSend   bool                                            `json:"auto_send,omitempty"`
	Headers    []*EventOrchestrationPathAutomationActionObject `json:"headers,omitempty"`
	Parameters []*EventOrchestrationPathAutomationActionObject `json:"parameters,omitempty"`
}

type EventOrchestrationPathAutomationActionObject struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type EventOrchestrationPathActionVariables struct {
	Name  string `json:"name,omitempty"`
	Path  string `json:"path,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type EventOrchestrationPathActionExtractions struct {
	Target   string `json:"target,omitempty"`
	Template string `json:"template,omitempty"`
}

type EventOrchestrationPathCatchAll struct {
	Actions *EventOrchestrationPathAction `json:"actions,omitempty"`
}
