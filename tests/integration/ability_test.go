package integration

import "testing"

func TestAbilitiesList(t *testing.T) {
	setup(t)
	if _, _, err := client.Abilities.List(); err != nil {
		t.Fatal(err)
	}
}

func TestAbilitiesSSO(t *testing.T) {
	setup(t)
	if _, err := client.Abilities.Test("sso"); err != nil {
		t.Fatal(err)
	}
}
