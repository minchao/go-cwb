package cwb

import (
	"context"
	"net/http"
	"net/url"
)

const (
	// Data set Ids
	// Weather observation data
	StationObsWeatherId = "O-A0001-001"

	// Rainfall observation data
	StationObsRainfallId = "O-A0002-001"
)

type StationObsService service

type StationObsResponse struct {
	Success string `json:"success"`
	Result  Result `json:"result"`
	Records struct {
		Location []StationObsLocation `json:"location"`
	} `json:"records"`
}

type StationObsWeather StationObsResponse

type StationObsLocation struct {
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	LocationName string `json:"locationName"`
	StationId    string `json:"stationId"`
	Time         struct {
		ObsTime string `json:"obsTime"`
	} `json:"time"`
	WeatherElement []StationObsElement `json:"weatherElement"`
	Parameter      []Parameter         `json:"parameter"`
}

type StationObsElement struct {
	ElementName  string `json:"elementName"`
	ElementValue string `json:"elementValue"`
}

// GetWeatherObs gets weather observation data.
func (s *StationObsService) GetWeather(ctx context.Context, options url.Values) (*StationObsWeather, *http.Response, error) {
	obs := new(StationObsWeather)
	req, err := s.client.Get(ctx, s.client.generateURL(StationObsWeatherId, options), obs)
	if err != nil {
		return nil, nil, err
	}
	return obs, req, nil
}

type StationObsRainfall StationObsResponse

// GetRainfall gets rainfall observation data.
func (s *StationObsService) GetRainfall(ctx context.Context, options url.Values) (*StationObsRainfall, *http.Response, error) {
	obs := new(StationObsRainfall)
	req, err := s.client.Get(ctx, s.client.generateURL(StationObsRainfallId, options), obs)
	if err != nil {
		return nil, nil, err
	}
	return obs, req, nil
}
