package integration

import (
	"fmt"
	"testing"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func installAddon() (*pagerduty.Addon, error) {
	name := randResName()
	addon := &pagerduty.Addon{Name: name, Src: fmt.Sprintf("https://%s", name)}
	resp, _, err := client.Addons.Install(addon)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func TestAddonsList(t *testing.T) {
	setup(t)
	res, err := installAddon()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Addons.Delete(res.ID)

	resp, _, err := client.Addons.List(&pagerduty.ListAddonsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	var found *pagerduty.Addon

	for _, r := range resp.Addons {
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

func TestAddonsGet(t *testing.T) {
	setup(t)
	res, err := installAddon()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Addons.Delete(res.ID)

	resp, _, err := client.Addons.Get(res.ID)
	if err != nil {
		t.Fatal(err)
	}

	if res.Name != resp.Name {
		t.Fatalf("Expected %s; got %s", res.Name, resp.Name)
	}
}
