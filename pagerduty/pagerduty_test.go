package pagerduty

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestRetryURL(t *testing.T) {

	setup()
	defer teardown()

	timesCalled := 0
	expectedURL := "/members?offset=100"

	options := GetMembersOptions{
		Offset: 100,
	}

	mux.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		timesCalled++
		testMethod(t, r, "GET")
		url := r.URL.String()
		if url != expectedURL {
			t.Fatalf("Request url: %v, want %v", url, expectedURL)
		}

		if timesCalled > 1 {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Ratelimit-Reset", "1")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error":{"code":"2020", "message":"Rate Limit Exceeded"}}`))

	})

	_, err := client.newRequestDoOptionsContext(context.Background(), http.MethodGet, "/members", options, nil, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	timesCalled = 0
	_, err = client.newRequestDoContext(context.Background(), http.MethodGet, "/members", options, nil, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestHandleRatelimitErrorWithRatelimitHeaders(t *testing.T) {
	setup()
	defer teardown()

	count := 0
	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		if count > 0 {
			w.Write([]byte(`{"teams": [{"id": "1"}]}`))
			return
		}

		// Expected response ref. https://developer.pagerduty.com/docs/72d3b724589e3-rest-api-rate-limits#reaching-the-limit
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("ratelimit-limit", "960")
		w.Header().Add("ratelimit-remaining", "0")
		w.Header().Add("ratelimit-reset", "1")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error":{"message":"Rate Limit Exceeded","code":2020}}`))
		count++
	})

	resp, _, err := client.Teams.List(&ListTeamsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListTeamsResponse{
		Teams: []*Team{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

}

func TestHandleRatelimitErrorNoRatelimitHeaders(t *testing.T) {
	setup()
	defer teardown()

	count := 0
	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		if count > 0 {
			w.Write([]byte(`{"teams": [{"id": "1"}]}`))
			return
		}

		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`
      <html>
      <head><title>429 Too Many Requests</title></head>
      <body>
      <center><h1>429 Too Many Requests</h1></center>
      <hr><center>nginx</center>
      </body>
      </html>`))
		count++
	})

	resp, _, err := client.Teams.List(&ListTeamsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	want := &ListTeamsResponse{
		Teams: []*Team{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

}
