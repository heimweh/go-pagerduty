package integration

import (
	"testing"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func installService() (*pagerduty.Service, *pagerduty.EscalationPolicy, *pagerduty.User, error) {
	name := randResName()
	es, user, err := createEscalationPolicy()
	if err != nil {
		return nil, nil, nil, err
	}

	resp, _, err := client.Services.Create(&pagerduty.Service{
		Name: name,
		EscalationPolicy: &pagerduty.EscalationPolicyReference{
			ID:   es.ID,
			Type: "escalation_policy_reference",
		},
	})
	if err != nil {
		return nil, nil, nil, err
	}
	return resp, es, user, nil
}

func TestServicesList(t *testing.T) {
	setup(t)
	service, es, user, err := installService()
	if err != nil {
		t.Fatal(err)
	}
	defer client.EscalationPolicies.Delete(es.ID)
	defer client.Users.Delete(user.ID)
	defer client.Services.Delete(service.ID)

	resp, _, err := client.Services.List(&pagerduty.ListServicesOptions{})
	if err != nil {
		t.Fatal(err)
	}

	var found *pagerduty.Service

	for _, r := range resp.Services {
		if service.ID == r.ID {
			found = r
		}
	}

	if found == nil {
		t.Fatalf("Could not find: %s", service.ID)
	}

	if service.Name != found.Name {
		t.Fatalf("Expected %s; got %s", service.Name, found.Name)
	}
}
