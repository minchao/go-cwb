package cwb

import (
	"context"
	"net/http"
)

type Forecast36HourWeather struct {
	Success string `json:"success"`
	Result  Result `json:"result"`
	Records struct {
		DatasetDescription string                `json:"datasetDescription"`
		Location           []F36HCountryLocation `json:"location"`
	} `json:"records"`
}

type F36HCountryLocation struct {
	LocationName   string `json:"locationName"`
	WeatherElement []struct {
		ElementName string `json:"elementName"`
		Time        []struct {
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
			Parameter struct {
				ParameterName  string  `json:"paramterName"` // Typo in CWB API
				ParameterValue *string `json:"parameterValue,omitempty"`
				ParameterUnit  *string `json:"parameterUnit,omitempty"`
			} `json:"parameter"`
		} `json:"time"`
	} `json:"weatherElement"`
}

type F36HTime struct {
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Parameter Parameter `json:"parameter"`
}

// GetForecasts gets 36 hour weather forecasts.
func (s *ForecastsService) Get36HourWeather(ctx context.Context) (*Forecast36HourWeather, *http.Response, error) {
	forecast := new(Forecast36HourWeather)
	req, err := s.client.Get(ctx, "api/v1/rest/datastore/F-C0032-001", forecast)
	if err != nil {
		return nil, nil, err
	}
	return forecast, req, nil
}

type ForecastTownshipsWeather struct {
	Success string `json:"success"`
	Result  Result `json:"result"`
	Records struct {
		ContentDescription string               `json:"contentDescription"`
		Locations          []FTWCountryLocation `json:"locations"`
	} `json:"records"`
}

type FTWCountryLocation struct {
	DatasetDescription string               `json:"datasetDescription"`
	LocationsName      *string              `json:"locationsName,omitempty"`
	Dataid             string               `json:"dataid"`
	Location           []FTWDatasetLocation `json:"location"`
}

type FTWDatasetLocation struct {
	LocationName   string              `json:"locationName"`
	Geocode        string              `json:"geocode"`
	Lat            string              `json:"lat"`
	Lon            string              `json:"lon"`
	WeatherElement []FTWWeatherElement `json:"weatherElement"`
}

type FTWWeatherElement struct {
	ElementName    string    `json:"elementName"`
	ElementMeasure *string   `json:"elementMeasure,omitempty"`
	Time           []FTWTime `json:"time"`
}

type FTWTime struct {
	DataTime     *string     `json:"dataTime,omitempty"`
	StartTime    *string     `json:"startTime,omitempty"`
	EndTime      *string     `json:"endTime,omitempty"`
	ElementValue *string     `json:"elementValue,omitempty"`
	Parameter    []Parameter `json:"parameter,omitempty"`
}

// GetTownshipsWeatherByCity gets townships forecasts by data Id.
// See http://opendata.cwb.gov.tw/datalist for dataId (F-D0047-001 - F-D0047-091).
func (s *ForecastsService) GetTownshipsWeatherByDataId(ctx context.Context, dataId string) (*ForecastTownshipsWeather, *http.Response, error) {
	forecast := new(ForecastTownshipsWeather)
	req, err := s.client.Get(ctx, "api/v1/rest/datastore/"+dataId, forecast)
	if err != nil {
		return nil, nil, err
	}
	return forecast, req, nil
}
