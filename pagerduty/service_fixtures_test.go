package pagerduty

const (
	validListServicesJSON = `{
  "services": [
    {
      "id": "PIJ90N7",
      "summary": "My Application Service",
      "type": "service",
      "self": "https://api.pagerduty.com/services/PIJ90N7",
      "html_url": "https://subdomain.pagerduty.com/services/PIJ90N7",
      "name": "My Application Service",
      "description": null,
      "auto_resolve_timeout": 14400,
      "acknowledgement_timeout": 600,
      "created_at": "2015-11-06T11:12:51-05:00",
      "status": "active",
      "last_incident_timestamp": null,
      "alert_creation": "create_alerts_and_incidents",
      "alert_grouping": "intelligent",
      "alert_grouping_timeout": null,
      "integrations": [
        {
          "id": "PQ12345",
          "type": "generic_email_inbound_integration_reference",
          "summary": "Email Integration",
          "self": "https://api.pagerduty.com/services/PIJ90N7/integrations/PQ12345",
          "html_url": "https://subdomain.pagerduty.com/services/PIJ90N7/integrations/PQ12345"
        }
      ],
      "escalation_policy": {
        "id": "PT20YPA",
        "type": "escalation_policy_reference",
        "summary": "Another Escalation Policy",
        "self": "https://api.pagerduty.com/escalation_policies/PT20YPA",
        "html_url": "https://subdomain.pagerduty.com/escalation_policies/PT20YPA"
      },
      "teams": [
        {
          "id": "PQ9K7I8",
          "type": "team_reference",
          "summary": "Engineering",
          "self": "https://api.pagerduty.com/teams/PQ9K7I8",
          "html_url": "https://subdomain.pagerduty.com/teams/PQ9K7I8"
        }
      ],
      "incident_urgency_rule": {
        "type": "use_support_hours",
        "during_support_hours": {
          "type": "constant",
          "urgency": "high"
        },
        "outside_support_hours": {
          "type": "constant",
          "urgency": "low"
        }
      },
      "support_hours": {
        "type": "fixed_time_per_day",
        "time_zone": "America/Lima",
        "start_time": "09:00:00",
        "end_time": "17:00:00",
        "days_of_week": [
          1,
          2,
          3,
          4,
          5
        ]
      },
      "scheduled_actions": [
        {
          "type": "urgency_change",
          "at": {
            "type": "named_time",
            "name": "support_hours_start"
          },
          "to_urgency": "high"
        }
      ]
    }
  ],
  "limit": 25,
  "offset": 0,
  "total": null,
  "more": false
}`
)

var (
	defaultTestServiceAcknowledgementTimeout = 600
	defaultAutoResolveTimeout                = 14400

	validListServicesResponse = &ListServicesResponse{
		Services: []*Service{
			&Service{
				AcknowledgementTimeout: &defaultTestServiceAcknowledgementTimeout,
				Addons:                 nil,
				AlertCreation:          "create_alerts_and_incidents",
				AlertGrouping:          "intelligent",
				AlertGroupingTimeout:   nil,
				AutoResolveTimeout:     &defaultAutoResolveTimeout,
				CreatedAt:              "2015-11-06T11:12:51-05:00",
				Description:            "",
				EscalationPolicy: &EscalationPolicyReference{
					HTMLURL: "https://subdomain.pagerduty.com/escalation_policies/PT20YPA",
					ID:      "PT20YPA",
					Self:    "https://api.pagerduty.com/escalation_policies/PT20YPA",
					Summary: "Another Escalation Policy",
					Type:    "escalation_policy_reference",
				},
				HTMLURL: "https://subdomain.pagerduty.com/services/PIJ90N7",
				ID:      "PIJ90N7",
				IncidentUrgencyRule: &IncidentUrgencyRule{
					DuringSupportHours: &IncidentUrgencyType{
						Type:    "constant",
						Urgency: "high",
					},
					OutsideSupportHours: &IncidentUrgencyType{
						Type:    "constant",
						Urgency: "low",
					},
					Type:    "use_support_hours",
					Urgency: "",
				},
				Integrations: []*IntegrationReference{
					{
						ID:      "PQ12345",
						Type:    "generic_email_inbound_integration_reference",
						Summary: "Email Integration",
						Self:    "https://api.pagerduty.com/services/PIJ90N7/integrations/PQ12345",
						HTMLURL: "https://subdomain.pagerduty.com/services/PIJ90N7/integrations/PQ12345",
					},
				},
				LastIncidentTimestamp: "",
				Name:                  "My Application Service",
				ScheduledActions: []*ScheduledAction{
					{
						Type: "urgency_change",
						At: &At{
							Type: "named_time",
							Name: "support_hours_start",
						},
						ToUrgency: "high",
					},
				},
				Self:    "https://api.pagerduty.com/services/PIJ90N7",
				Status:  "active",
				Summary: "My Application Service",
				SupportHours: &SupportHours{
					DaysOfWeek: []int{1, 2, 3, 4, 5},
					EndTime:    "17:00:00",
					StartTime:  "09:00:00",
					TimeZone:   "America/Lima",
					Type:       "fixed_time_per_day",
				},
				Teams: []*TeamReference{
					{
						ID:      "PQ9K7I8",
						Type:    "team_reference",
						Summary: "Engineering",
						Self:    "https://api.pagerduty.com/teams/PQ9K7I8",
						HTMLURL: "https://subdomain.pagerduty.com/teams/PQ9K7I8",
					},
				},
				Type: "service",
			},
		},
		Limit: 25,
	}
)
