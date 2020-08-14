package cwb

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestDatasetService_GetIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/rest/dataset", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"dataid":["F-A0021-001","F-C0032-001","O-A0017-001"]}`))
	})

	want := &Dataset{
		Dataid: []string{
			"F-A0021-001",
			"F-C0032-001",
			"O-A0017-001",
		},
	}

	got, _, err := client.Dataset.GetIds(context.Background())
	if err != nil {
		t.Errorf("Dataset.GetIds returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dataset.GetIds is %+v, want %+v", got, want)
	}
}

func TestDatasetService_GetData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/rest/dataset/F-C0032-001", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
  "cwbCategory": "1",
  "datanamefixid": "2",
  "title": "3",
  "description": "4",
  "fieldDescription": "5",
  "type": "6",
  "license": "7",
  "licenseURL": "8",
  "cost": "9",
  "costURL": "10",
  "costLaw": "11",
  "organization": "12",
  "organizationContactName": "13",
  "organizationContactPhone": "14",
  "organizationContactEmail": "15",
  "publisher": "16",
  "publisherContactName": "17",
  "publisherContactPhone": "18",
  "publisherContactEmail": "19",
  "accrualPeriodicity": "20",
  "issued": "21",
  "modified": "22",
  "spatial": "23",
  "language": "24",
  "landingPage": "25",
  "notes": "26"
}`))
	})

	want := &Data{
		CwbCategory:              "1",
		Datanamefixid:            "2",
		Title:                    "3",
		Description:              "4",
		FieldDescription:         "5",
		Type:                     "6",
		License:                  "7",
		LicenseURL:               "8",
		Cost:                     "9",
		CostURL:                  "10",
		CostLaw:                  "11",
		Organization:             "12",
		OrganizationContactName:  "13",
		OrganizationContactPhone: "14",
		OrganizationContactEmail: "15",
		Publisher:                "16",
		PublisherContactName:     "17",
		PublisherContactPhone:    "18",
		PublisherContactEmail:    "19",
		AccrualPeriodicity:       "20",
		Issued:                   "21",
		Modified:                 "22",
		Spatial:                  "23",
		Language:                 "24",
		LandingPage:              "25",
		Notes:                    "26",
	}

	got, _, err := client.Dataset.GetData(context.Background(), "F-C0032-001")
	if err != nil {
		t.Errorf("Dataset.GetData returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dataset.GetData is %+v, want %+v", got, want)
	}
}
