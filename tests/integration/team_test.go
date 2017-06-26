package integration

import (
	"testing"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func createTeam() (*pagerduty.Team, error) {
	name := randResName()

	resp, _, err := client.Teams.Create(&pagerduty.Team{Name: name})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TestTeamsList(t *testing.T) {
	setup(t)
	res, err := createTeam()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Teams.Delete(res.ID)

	resp, _, err := client.Teams.List(&pagerduty.ListTeamsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	var found *pagerduty.Team

	for _, r := range resp.Teams {
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

func TestTeamsGet(t *testing.T) {
	setup(t)
	res, err := createTeam()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Teams.Delete(res.ID)

	resp, _, err := client.Teams.Get(res.ID)
	if err != nil {
		t.Fatal(err)
	}

	if res.Name != resp.Name {
		t.Fatalf("Expected %s; got %s", res.Name, resp.Name)
	}
}
