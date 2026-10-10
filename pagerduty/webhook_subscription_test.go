package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestWebhookSubscriptionList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/webhook_subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 0, "offset": 0, "more": false, "limit": 0, "webhook_subscriptions":[{"id": "1"}]}`))
	})

	resp, _, err := client.WebhookSubscriptions.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListWebhookSubscriptionsResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		WebhookSubscriptions: []*WebhookSubscription{
			{
				ID: "1",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestWebhookSubscriptionActivation(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		initialActive bool
		active        bool
	}{
		{"create inactive", http.MethodPost, true, false},
		{"create active", http.MethodPost, true, true},
		{"deactivate", http.MethodPut, true, false},
		{"reactivate", http.MethodPut, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			defer teardown()

			storedActive := tt.initialActive
			path := "/webhook_subscriptions"
			if tt.method == http.MethodPut {
				path += "/1"
			}
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, tt.method)
				var payload struct {
					WebhookSubscription struct {
						Active *bool `json:"active"`
					} `json:"webhook_subscription"`
				}
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					t.Errorf("decode request: %v", err)
					http.Error(w, "invalid request", http.StatusBadRequest)
					return
				}
				// PagerDuty defaults new subscriptions to active and preserves
				// the current state when an update omits the active field.
				if payload.WebhookSubscription.Active != nil {
					storedActive = *payload.WebhookSubscription.Active
				}
				json.NewEncoder(w).Encode(WebhookSubscriptionPayload{
					WebhookSubscription: &WebhookSubscription{
						ID:     "1",
						Active: storedActive,
					},
				})
			})

			input := &WebhookSubscription{Active: tt.active}
			var resp *WebhookSubscription
			var err error
			if tt.method == http.MethodPost {
				resp, _, err = client.WebhookSubscriptions.Create(input)
			} else {
				resp, _, err = client.WebhookSubscriptions.Update("1", input)
			}
			if err != nil {
				t.Fatal(err)
			}
			if resp.Active != tt.active {
				t.Errorf("subscription active = %v, want %v", resp.Active, tt.active)
			}
		})
	}
}

func TestWebhookSubscriptionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/webhook_subscriptions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"webhook_subscription":{"source_id": "1", "id":"1"}}`))
	})

	ID := "1"
	resp, _, err := client.WebhookSubscriptions.Get(ID)

	if err != nil {
		t.Fatal(err)
	}

	want := &WebhookSubscription{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestWebhookSubscriptionDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/webhook_subscriptions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ID := "1"

	if _, err := client.WebhookSubscriptions.Delete(ID); err != nil {
		t.Fatal(err)
	}
}
