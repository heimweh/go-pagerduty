package pagerduty

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
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

func testHeader(t *testing.T, r *http.Request, key, value string) {
	v := r.Header.Get(key)
	if value != v {
		t.Errorf("unexpected header for key %s.\n\n%s want\n\n%s", key, v, value)
	}
}

func testBody(t *testing.T, r *http.Request, expectedBody string) {
	b := new(bytes.Buffer)
	b.ReadFrom(r.Body)
	bodyStr := strings.TrimSpace(b.String())
	if bodyStr != expectedBody {
		t.Errorf("unexpected body.\n\n%v want\n\n%v", bodyStr, expectedBody)
	}
}

func testQueryValue(t *testing.T, r *http.Request, wantKey string, wantValue string) {
	if wantValue == "" {
		if r.URL.Query().Get(wantKey) == "" {
			t.Errorf("Request missing query param: %v, was %v", wantKey, r.URL.Query().Encode())
		} else if got := r.URL.Query().Get(wantKey); got != wantValue {
			t.Errorf("Request unexpected query param value for %v: %v, want %v", wantKey, wantValue, got)
		}
	}
}

func testQueryMinCount(t *testing.T, r *http.Request, minCount int) {
	if l := len(r.URL.Query()); l < minCount {
		t.Errorf("Request contained unexpected number of query params: %v, want at least %v", l, minCount)
	}
}

func testQueryMaxCount(t *testing.T, r *http.Request, maxCount int) {
	if l := len(r.URL.Query()); l > maxCount {
		t.Errorf("Request contained unexpected number of query params: %v, want at most %v", l, maxCount)
	}
}

func testQueryCount(t *testing.T, r *http.Request, count int) {
	if l := len(r.URL.Query()); l != count {
		t.Errorf("Request contained unexpected number of query params: %v, want exactly %v", l, count)
	}
}
