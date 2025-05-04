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

func TestWebhookSubscriptionCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &WebhookSubscription{}

	mux.HandleFunc("/webhook_subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(WebhookSubscriptionPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.WebhookSubscription, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"webhook_subscription":{"id": "1"}}`))
	})

	resp, _, err := client.WebhookSubscriptions.Create(input)
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

func TestWebhookSubscriptionCreateNonActive(t *testing.T) {
	setup()
	defer teardown()

	input := &WebhookSubscription{Active: false}

	mux.HandleFunc("/webhook_subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(WebhookSubscriptionPayload)
		json.NewDecoder(r.Body).Decode(v)
		encoded, _ := json.Marshal(v)
		expectedMarshalValue := `{"webhook_subscription":{"active":false,"delivery_method":{"custom_headers":null},"filter":{}}}`
		if string(encoded) != expectedMarshalValue {
			t.Errorf("Marshalled body = %s, want %s", string(encoded), expectedMarshalValue)
		}
		if !reflect.DeepEqual(v.WebhookSubscription, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"webhook_subscription":{"id": "1"}}`))
	})

	resp, _, err := client.WebhookSubscriptions.Create(input)
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

func TestWebhookSubscriptionUpdate(t *testing.T) {
	setup()
	defer teardown()
	input := &WebhookSubscription{
		ID: "2",
	}

	mux.HandleFunc("/webhook_subscriptions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(WebhookSubscriptionPayload)

		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.WebhookSubscription, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"webhook_subscription":{"id":"1"}}`))
	})

	ID := "1"

	resp, _, err := client.WebhookSubscriptions.Update(ID, input)
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
