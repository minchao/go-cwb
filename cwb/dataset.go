package cwb

import (
	"context"
	"net/http"
)

type DatasetService service

type Dataset struct {
	Dataid []string `json:"dataid"`
}

type Data struct {
	CwbCategory              string `json:"cwbCategory"`
	Datanamefixid            string `json:"datanamefixid"`
	Title                    string `json:"title"`
	Description              string `json:"description"`
	FieldDescription         string `json:"fieldDescription"`
	Type                     string `json:"type"`
	License                  string `json:"license"`
	LicenseURL               string `json:"licenseURL"`
	Cost                     string `json:"cost"`
	CostURL                  string `json:"costURL"`
	CostLaw                  string `json:"costLaw"`
	Organization             string `json:"organization"`
	OrganizationContactName  string `json:"organizationContactName"`
	OrganizationContactPhone string `json:"organizationContactPhone"`
	OrganizationContactEmail string `json:"organizationContactEmail"`
	Publisher                string `json:"publisher"`
	PublisherContactName     string `json:"publisherContactName"`
	PublisherContactPhone    string `json:"publisherContactPhone"`
	PublisherContactEmail    string `json:"publisherContactEmail"`
	AccrualPeriodicity       string `json:"accrualPeriodicity"`
	TemporalCoverageFrom     string `json:"temporalCoverageFrom,omitempty"`
	TemporalCoverageTo       string `json:"temporalCoverageTo,omitempty"`
	Issued                   string `json:"issued"`
	Modified                 string `json:"modified"`
	Spatial                  string `json:"spatial"`
	Language                 string `json:"language"`
	LandingPage              string `json:"landingPage"`
	Notes                    string `json:"notes"`
}

func (s *DatasetService) GetIds(ctx context.Context) (*Dataset, *http.Response, error) {
	dataset := new(Dataset)
	req, err := s.client.Get(ctx, "api/v1/rest/dataset", dataset)
	if err != nil {
		return nil, nil, err
	}
	return dataset, req, nil
}

func (s *DatasetService) GetData(ctx context.Context, id string) (*Data, *http.Response, error) {
	data := new(Data)
	req, err := s.client.Get(ctx, "api/v1/rest/dataset/"+id, data)
	if err != nil {
		return nil, nil, err
	}
	return data, req, nil
}
