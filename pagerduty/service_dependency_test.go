package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestServiceDependencyGetBusinessServiceDependencies(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/service_dependencies/business_services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"relationships":[{"type": "service_dependency", "supporting_service": {"type":"business_service_reference","id":"1"}, "dependent_service": {"type":"technical_service_reference","id":"1"}, "id":"1"}]}`))
	})

	serveID := "1"
	serveType := "business_service"
	resp, _, err := client.ServiceDependencies.GetServiceDependenciesForType(serveID, serveType)
	if err != nil {
		t.Fatal(err)
	}
	want := &ListServiceDependencies{
		Relationships: []*ServiceDependency{
			{
				Type: "service_dependency",
				ID:   "1",
				SupportingService: &ServiceObj{
					ID:   "1",
					Type: "business_service_reference",
				},
				DependentService: &ServiceObj{
					ID:   "1",
					Type: "technical_service_reference",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
func TestServiceDependencyGetTechnicalServiceDependencies(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/service_dependencies/technical_services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"relationships":[{"type": "service_dependency", "supporting_service": {"type":"service","id":"1"}, "dependent_service": {"type":"technical_service_reference","id":"1"}, "id":"1"}]}`))
	})

	serveID := "1"
	serveType := "service"
	resp, _, err := client.ServiceDependencies.GetServiceDependenciesForType(serveID, serveType)
	if err != nil {
		t.Fatal(err)
	}
	want := &ListServiceDependencies{
		Relationships: []*ServiceDependency{
			{
				Type: "service_dependency",
				ID:   "1",
				SupportingService: &ServiceObj{
					ID:   "1",
					Type: "service",
				},
				DependentService: &ServiceObj{
					ID:   "1",
					Type: "technical_service_reference",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
func TestServiceDependencyGetServiceDependencies_WrongType(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/service_dependencies/technical_services/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"relationships":[{"type": "service_dependency", "supporting_service": {"type":"service","id":"1"}, "dependent_service": {"type":"technical_service_reference","id":"1"}, "id":"1"}]}`))
	})

	serveID := "1"
	serveType := "foo"
	_, _, err := client.ServiceDependencies.GetServiceDependenciesForType(serveID, serveType)
	if err == nil {
		t.Fatal("this should have been an error")
	}
}
func TestServiceDependencyAssociate(t *testing.T) {
	setup()
	defer teardown()
	input := &ListServiceDependencies{}

	mux.HandleFunc("/service_dependencies/associate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(ListServiceDependencies)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"relationships":[{"type": "service_dependency", "supporting_service": {"type":"business_service_reference","id":"1"}, "dependent_service": {"type":"technical_service_reference","id":"1"}, "id":"1"}]}`))
	})

	resp, _, err := client.ServiceDependencies.AssociateServiceDependencies(input)
	if err != nil {
		t.Fatal(err)
	}
	want := &ListServiceDependencies{
		Relationships: []*ServiceDependency{
			{
				Type: "service_dependency",
				ID:   "1",
				SupportingService: &ServiceObj{
					ID:   "1",
					Type: "business_service_reference",
				},
				DependentService: &ServiceObj{
					ID:   "1",
					Type: "technical_service_reference",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestBusinessServiceDisassociateServiceDependencies(t *testing.T) {
	setup()
	defer teardown()
	input := &ListServiceDependencies{
		Relationships: []*ServiceDependency{
			{
				Type: "service_dependency",
				SupportingService: &ServiceObj{
					ID:   "foo123",
					Type: "business_service_reference",
				},
				DependentService: &ServiceObj{
					ID:   "bar123",
					Type: "technical_service_reference",
				},
			},
		},
	}

	mux.HandleFunc("/service_dependencies/disassociate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(ListServiceDependencies)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"relationships":[{"type": "service_dependency", "supporting_service": {"type":"business_service_reference","id":"1"}, "dependent_service": {"type":"technical_service_reference","id":"1"}, "id":"1"}]}`))

	})

	resp, _, err := client.ServiceDependencies.DisassociateServiceDependencies(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &ListServiceDependencies{
		Relationships: []*ServiceDependency{
			{
				Type: "service_dependency",
				ID:   "1",
				SupportingService: &ServiceObj{
					ID:   "1",
					Type: "business_service_reference",
				},
				DependentService: &ServiceObj{
					ID:   "1",
					Type: "technical_service_reference",
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}

}
