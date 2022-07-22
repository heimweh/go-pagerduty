package pagerduty

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the PagerDuty client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, _ = NewClient(&Config{BaseURL: server.URL, Token: "foo"})
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestClientUserAgentDefault(t *testing.T) {
	client, err := NewClient(&Config{Token: "foo"})
	if err != nil {
		t.Fatal(err)
	}

	if client.Config.UserAgent != defaultUserAgent {
		t.Errorf("got %q, want %q", client.Config.UserAgent, defaultUserAgent)
	}
}

func TestClientUserAgentOverwritten(t *testing.T) {
	newUserAgent := "foo-user-agent"
	client, err := NewClient(&Config{Token: "foo", UserAgent: newUserAgent})
	if err != nil {
		t.Fatal(err)
	}

	if client.Config.UserAgent != newUserAgent {
		t.Errorf("got %q, want %q", client.Config.UserAgent, newUserAgent)
	}
}
