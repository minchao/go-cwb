package cwb

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	// tide forecasts 1 month
	Tide1MonthId = "F-A0021-001"
)

type TideForecastsService service

type TideForecast1MonthOptions struct {
	Limit        int    `url:"limit,omitempty"`
	Offset       int    `url:"offset,omitempty"`
	LocationName string `url:"locationName,omitempty"` // see https://opendata.cwa.gov.tw/opendatadoc/MMC/A0021-001.pdf
	ElementName  string `url:"elementName,omitempty"`
	Sort         string `url:"sort,omitempty"`
	StartTime    string `url:"startTime,omitempty"`
	DataTime     string `url:"dataTime,omitempty"`
	TimeForm     string `url:"timeForm,omitempty"`
	TimeTo       string `url:"timeTo,omitempty"`
}

type TideForecast1MonthResponse struct {
	Success string `json:"success"`
	Result  Result `json:"result"`
	Records struct {
		Dataid   string                 `json:"dataid"`
		Note     string                 `json:"note"`
		Location []TideForecastLocation `json:"location"`
	} `json:"records"`
}

type TideForecastLocation struct {
	LocationName string `json:"locationName"`
	StationId    string `json:"stationId"`
	ValidTime    []struct {
		StartTime      string                       `json:"startTime"`
		EndTime        string                       `json:"endTime"`
		WeatherElement []TideForecastWeatherElement `json:"weatherElement"`
	} `json:"validTime"`
}

type TideForecastWeatherElement struct {
	ElementName  string             `json:"elementName"`
	ElementValue *string            `json:"elementValue,omitempty"`
	Time         []TideForecastTime `json:"time,omitempty"`
}

type TideForecastTime struct {
	DataTime  string                  `json:"dataTime"`
	Parameter []TideForecastParameter `json:"parameter"`
}

type TideForecastParameter struct {
	ParameterName    string  `json:"parameterName"`
	ParameterValue   *string `json:"parameterValue,omitempty"`
	ParameterMeasure *string `json:"parameterMeasure,omitempty"`
}

// Get1MonthTide gets 1 month tide forecasts.
// see https://opendata.cwa.gov.tw/dist/opendata-swagger.html#/%E9%A0%90%E5%A0%B1/get_v1_rest_datastore_F_A0021_001
func (s *TideForecastsService) Get1MonthTide(ctx context.Context, options TideForecast1MonthOptions) (*TideForecast1MonthResponse, *http.Response, error) {
	values, _ := query.Values(options)
	forecast := new(TideForecast1MonthResponse)
	req, err := s.client.Get(ctx, s.client.generateURL(Tide1MonthId, values), forecast)
	if err != nil {
		return nil, nil, err
	}
	return forecast, req, nil
}
