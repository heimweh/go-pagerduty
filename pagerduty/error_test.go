package pagerduty

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestErrorResponses(t *testing.T) {
	testCases := []struct {
		name string
		body string
		want interface{}
	}{
		{
			name: "error with message",
			body: `{"error": {"message": "Your account is expired and cannot use the API.", "code": 2012}}`,
			want: &errorResponse{
				Error: &Error{
					Code:    2012,
					Message: "Your account is expired and cannot use the API.",
				},
			},
		},

		{
			name: "error with multiple errors",
			body: `{"error": {"errors": ["foo", "bar"], "code": 2001, "message": "Invalid Input Provided"}}`,
			want: &errorResponse{
				Error: &Error{
					Errors:  []interface{}{"foo", "bar"},
					Code:    2001,
					Message: "Invalid Input Provided",
				},
			},
		},

		{
			name: "error with map slice",
			body: `{"error": {"message": "Invalid Schedule", "code": 3001, "errors": {"foo": ["bar"]}}}`,
			want: &errorResponse{
				Error: &Error{
					Errors:  map[string]interface{}{"foo": []interface{}{"bar"}},
					Code:    3001,
					Message: "Invalid Schedule",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := new(errorResponse)

			r := &Response{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(tc.body)))}}

			if err := client.DecodeJSON(r, v); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.want, v) {
				t.Errorf("got \n\n%#v \n\nwant\n\n%#v", tc.want, v)
			}
		})
	}
}
