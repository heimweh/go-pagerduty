package pagerduty

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCustomFieldSchemaAssignmentCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schema_assignments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"schema_assignment":{"service":{"id":"S1"},"schema":{"id":"FS1"}}}`)

		w.Write([]byte(`
{
    "schema_assignment": {
        "id": "FSA1",
		"type": "schema_assignment",
        "service": {
			"id": "S1"
		},
		"schema": {
			"id": "FS1"
		}
    }
}`))

	})

	resp, _, err := client.CustomFieldSchemaAssignments.Create(&CustomFieldSchemaAssignment{
		Schema:  &CustomFieldSchemaReference{ID: "FS1"},
		Service: &ServiceReference{ID: "S1"},
	})

	if err != nil {
		t.Fatal(err)
	}

	want := &CustomFieldSchemaAssignment{
		ID:      "FSA1",
		Type:    "schema_assignment",
		Schema:  &CustomFieldSchemaReference{ID: "FS1"},
		Service: &ServiceReference{ID: "S1"},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestCustomFieldSchemaAssignmentDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schema_assignments/FSA1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		w.WriteHeader(204)

	})

	resp, err := client.CustomFieldSchemaAssignments.Delete("FSA1")

	if err != nil {
		t.Fatal(err)
	}

	if resp.Response.StatusCode != 204 {
		t.Errorf("unexpected response code. want 204. got %v", resp.Response.StatusCode)
	}
}

func TestFieldAssignmentServiceListForService(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schema_assignments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		testQueryValue(t, r, "service_id", "S1")

		w.Write([]byte(`
{
    "schema_assignments": [
		{
        	"id": "FSA1",
        	"type": "schema_assignment",
			"service": {
				"id": "S1"
			},
			"schema": {
				"id": "FS1"
			}
    	}
	],
	"more": false
}`))
	})

	resp, _, err := client.CustomFieldSchemaAssignments.ListForService("S1", nil)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaAssignmentsResponse{
		SchemaAssignments: []*CustomFieldSchemaAssignment{
			{
				ID:      "FSA1",
				Type:    "schema_assignment",
				Service: &ServiceReference{ID: "S1"},
				Schema:  &CustomFieldSchemaReference{ID: "FS1"},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestFieldAssignmentServiceListForSchema(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schema_assignments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 1)
		testQueryValue(t, r, "schema_id", "FS1")

		w.Write([]byte(`
{
    "schema_assignments": [
		{
        	"id": "FSA1",
        	"type": "schema_assignment",
			"service": {
				"id": "S1"
			},
			"schema": {
				"id": "FS1"
			}
    	}
	],
	"more": false
}`))
	})

	resp, _, err := client.CustomFieldSchemaAssignments.ListForSchema("FS1", nil)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaAssignmentsResponse{
		SchemaAssignments: []*CustomFieldSchemaAssignment{
			{
				ID:      "FSA1",
				Type:    "schema_assignment",
				Service: &ServiceReference{ID: "S1"},
				Schema:  &CustomFieldSchemaReference{ID: "FS1"},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestFieldAssignmentServiceListForSchema_WithLimit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schema_assignments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryCount(t, r, 2)
		testQueryValue(t, r, "schema_id", "FS1")
		testQueryValue(t, r, "limit", "20")

		w.Write([]byte(`
{
    "schema_assignments": [
		{
        	"id": "FSA1",
        	"type": "schema_assignment",
			"service": {
				"id": "S1"
			},
			"schema": {
				"id": "FS1",
				"summary": "some schema title"
			}
    	}
	],
	"limit": 20,
	"offset": 0,
	"more": false
}`))
	})

	resp, _, err := client.CustomFieldSchemaAssignments.ListForSchema("FS1", &ListCustomFieldSchemaAssignmentsOptions{Limit: 20})

	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaAssignmentsResponse{
		SchemaAssignments: []*CustomFieldSchemaAssignment{
			{
				ID:   "FSA1",
				Type: "schema_assignment",
				Service: &ServiceReference{
					ID: "S1",
				},
				Schema: &CustomFieldSchemaReference{
					ID:      "FS1",
					Summary: "some schema title",
				},
			},
		},
		Limit: 20,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestFieldAssignmentServiceListForSchema_TwoPages(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customfields/schema_assignments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValue(t, r, "schema_id", "FS1")
		testQueryMinCount(t, r, 1)
		testQueryMaxCount(t, r, 2)
		offset := r.URL.Query().Get("offset")

		switch offset {
		case "":
			w.Write([]byte(`
{
    "schema_assignments": [
		{
        	"id": "FSA1",
        	"type": "schema_assignment",
			"service": {
				"id": "S1"
			},
			"schema": {
				"id": "FS1"
			}
    	}
	],
	"limit": 1,
	"offset": 0,
	"more": true
}`))
		case "1":
			w.Write([]byte(`
{
    "schema_assignments": [
		{
        	"id": "FSA2",
        	"type": "schema_assignment",
			"service": {
				"id": "S2"
			},
			"schema": {
				"id": "FS1"
			}
    	}
	],
	"limit": 1,
	"offset": 0,
	"more": false
}`))
		default:
			t.Fatalf("Unexpected offset: %v", offset)
		}
	})

	resp, _, err := client.CustomFieldSchemaAssignments.ListForSchema("FS1", nil)

	if err != nil {
		t.Fatal(err)
	}

	want := &ListCustomFieldSchemaAssignmentsResponse{
		SchemaAssignments: []*CustomFieldSchemaAssignment{
			{
				ID:      "FSA1",
				Type:    "schema_assignment",
				Schema:  &CustomFieldSchemaReference{ID: "FS1"},
				Service: &ServiceReference{ID: "S1"},
			},
			{
				ID:      "FSA2",
				Type:    "schema_assignment",
				Schema:  &CustomFieldSchemaReference{ID: "FS1"},
				Service: &ServiceReference{ID: "S2"},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
