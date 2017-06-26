package integration

import (
	"fmt"
	"testing"

	"github.com/heimweh/go-pagerduty/pagerduty"
)

func createUser() (*pagerduty.User, error) {
	name := randResName()

	resp, _, err := client.Users.Create(&pagerduty.User{Name: name, Email: fmt.Sprintf("%s@bar.com", name)})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TestUsersList(t *testing.T) {
	setup(t)
	res, err := createUser()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Users.Delete(res.ID)

	resp, _, err := client.Users.List(&pagerduty.ListUsersOptions{})
	if err != nil {
		t.Fatal(err)
	}

	var found *pagerduty.User

	for _, r := range resp.Users {
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

func TestUsersGet(t *testing.T) {
	setup(t)
	res, err := createUser()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Users.Delete(res.ID)

	resp, _, err := client.Users.Get(res.ID, &pagerduty.GetUserOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if res.Name != resp.Name {
		t.Fatalf("Expected %s; got %s", res.Name, resp.Name)
	}
}
