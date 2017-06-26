package integration

import (
	"testing"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func createEscalationPolicy() (*pagerduty.EscalationPolicy, *pagerduty.User, error) {
	name := randResName()

	user, err := createUser()
	if err != nil {
		return nil, nil, err
	}

	ep := &pagerduty.EscalationPolicy{
		Name: name,
		EscalationRules: []*pagerduty.EscalationRule{
			{
				EscalationDelayInMinutes: 1,
				Targets: []*pagerduty.EscalationTargetReference{
					{
						ID:   user.ID,
						Type: "user_reference",
					},
				},
			},
		},
	}

	resp, _, err := client.EscalationPolicies.Create(ep)
	if err != nil {
		return nil, nil, err
	}

	return resp, user, nil
}

func TestEscalationPoliciesList(t *testing.T) {
	setup(t)
	res, user, err := createEscalationPolicy()
	if err != nil {
		t.Fatal(err)
	}
	defer client.EscalationPolicies.Delete(res.ID)
	defer client.Users.Delete(user.ID)

	resp, _, err := client.EscalationPolicies.List(&pagerduty.ListEscalationPoliciesOptions{})
	if err != nil {
		t.Fatal(err)
	}

	var found *pagerduty.EscalationPolicy

	for _, r := range resp.EscalationPolicies {
		if res.ID == r.ID {
			found = r
		}
	}

	if found == nil {
		t.Fatalf("Could not find: %s", res.ID)
	}

	if res.Name != found.Name {
		t.Fatalf("Expected %s; got %s", res.Name, found.Name)
	}
}

func TestEscalationPoliciesGet(t *testing.T) {
	setup(t)
	res, user, err := createEscalationPolicy()
	if err != nil {
		t.Fatal(err)
	}
	defer client.EscalationPolicies.Delete(res.ID)
	defer client.Users.Delete(user.ID)

	resp, _, err := client.EscalationPolicies.Get(res.ID, &pagerduty.GetEscalationPolicyOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if res.Name != resp.Name {
		t.Fatalf("Expected %s; got %s", res.Name, resp.Name)
	}
}
