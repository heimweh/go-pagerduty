package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestBusinessServiceSubscriberList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/business_services/1/subscribers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"total": 0, "offset": 0, "more": false, "limit": 0, "subscribers":[{"subscriber_id": "1", "subscriber_type": "team"}]}`))
	})

	resp, _, err := client.BusinessServiceSubscribers.List("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &ListBusinessServiceSubscribersResponse{
		Total:  0,
		Offset: 0,
		More:   false,
		Limit:  0,
		BusinessServiceSubscribers: []*BusinessServiceSubscriber{
			{
				ID: "1",
				Type: "team",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestBusinessServiceSubscriberCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &BusinessServiceSubscriber{ID: "foo", Type: "team"}
	businessServiceID := "1"

	mux.HandleFunc("/business_services/1/subscribers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(BusinessServiceSubscriber)
		v.ID = "foo"
		v.Type = "team"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"account_is_subscribed": true}`))
	})
	// this endpoint only returns  an "ok" in the body. no point in testing for it.
	if _, err := client.BusinessServiceSubscribers.Create(businessServiceID, input); err != nil {
		t.Fatal(err)
	}
}

func TestBusinessServiceSubscriberDelete(t *testing.T) {
	setup()
	defer teardown()

	businessServiceID := "1"
	subscriber := &BusinessServiceSubscriber{ID: "foo", Type: "team"}

	mux.HandleFunc("/business_services/1/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(BusinessServiceSubscriber)
		v.ID = "foo"
		v.Type = "team"
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, subscriber) {
			t.Errorf("Request body = %+v, want %+v", v, subscriber)
		}
		w.Write([]byte(`{"deleted_count": 1, "unauthorized_count": 1, "non_existent_count": 0}`))
	})

	// this endpoint only returns  an "ok" in the body. no point in testing for it.
	if _, err := client.BusinessServiceSubscribers.Delete(businessServiceID, subscriber); err != nil {
		t.Fatal(err)
	}
}
