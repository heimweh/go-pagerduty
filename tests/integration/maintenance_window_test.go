package integration

import (
	"testing"
	"time"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func createMaintenanceWindow() (*pagerduty.MaintenanceWindow, *pagerduty.Service, *pagerduty.EscalationPolicy, *pagerduty.User, error) {
	name := randResName()

	sv, es, us, err := installService()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	startTime := time.Now().UTC().Add(24 * time.Hour).String()
	endTime := time.Now().UTC().Add(48 * time.Hour).String()

	resp, _, err := client.MaintenanceWindows.Create(&pagerduty.MaintenanceWindow{
		StartTime:   startTime,
		EndTime:     endTime,
		Description: name,
		Services: []*pagerduty.ServiceReference{
			{
				ID:   sv.ID,
				Type: "service_reference",
			},
		},
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return resp, sv, es, us, nil
}

func TestMaintenanceWindowsList(t *testing.T) {
	setup(t)
	m, s, e, u, err := createMaintenanceWindow()
	if err != nil {
		t.Fatal(err)
	}
	defer client.EscalationPolicies.Delete(e.ID)
	defer client.Users.Delete(u.ID)
	defer client.Services.Delete(s.ID)
	defer client.MaintenanceWindows.Delete(m.ID)

	resp, _, err := client.MaintenanceWindows.List(&pagerduty.ListMaintenanceWindowsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	var found *pagerduty.MaintenanceWindow

	for _, mw := range resp.MaintenanceWindows {
		if mw.ID == m.ID {
			found = mw
		}
	}

	if found == nil {
		t.Fatalf("Could not find: %s", m.ID)
	}

	if m.Description != found.Description {
		t.Fatalf("Expected %s; got %s", m.Description, found.Description)
	}
}
