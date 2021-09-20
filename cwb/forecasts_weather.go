package cwb

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Data set Ids
	// 36 hour weather forecasts
	F36HW = "F-C0032-001"

	// Townships forecasts
	FTW2DayYilanCounty      = "F-D0047-001"
	FTW7DayYilanCounty      = "F-D0047-003"
	FTW2DayTaoyuanCity      = "F-D0047-005"
	FTW7DayTaoyuanCity      = "F-D0047-007"
	FTW2DayHsinchuCounty    = "F-D0047-009"
	FTW7DayHsinchuCounty    = "F-D0047-011"
	FTW2DayMiaoliCounty     = "F-D0047-013"
	FTW7DayMiaoliCounty     = "F-D0047-015"
	FTW2DayChanghuaCounty   = "F-D0047-017"
	FTW7DayChanghuaCounty   = "F-D0047-019"
	FTW2DayNantouCounty     = "F-D0047-021"
	FTW7DayNantouCounty     = "F-D0047-023"
	FTW2DayYunlinCounty     = "F-D0047-025"
	FTW7DayYunlinCounty     = "F-D0047-027"
	FTW2DayChiayiCounty     = "F-D0047-029"
	FTW7DayChiayiCounty     = "F-D0047-031"
	FTW2DayPingtungCounty   = "F-D0047-033"
	FTW7DayPingtungCounty   = "F-D0047-035"
	FTW2DayTaitungCounty    = "F-D0047-037"
	FTW7DayTaitungCounty    = "F-D0047-039"
	FTW2DayHualienCounty    = "F-D0047-041"
	FTW7DayHualienCounty    = "F-D0047-043"
	FTW2DayPenghuCounty     = "F-D0047-045"
	FTW7DayPenghuCounty     = "F-D0047-047"
	FTW2DayKeelungCity      = "F-D0047-049"
	FTW7DayKeelungCity      = "F-D0047-051"
	FTW2DayHsinchuCity      = "F-D0047-053"
	FTW7DayHsinchuCity      = "F-D0047-055"
	FTW2DayChiayiCity       = "F-D0047-057"
	FTW7DayChiayiCity       = "F-D0047-059"
	FTW2DayTaipeiCity       = "F-D0047-061"
	FTW7DayTaipeiCity       = "F-D0047-063"
	FTW2DayKaohsiungCity    = "F-D0047-065"
	FTW7DayKaohsiungCity    = "F-D0047-067"
	FTW2DayNewTaipeiCity    = "F-D0047-069"
	FTW7DayNewTaipeiCity    = "F-D0047-071"
	FTW2DayTaichungCity     = "F-D0047-073"
	FTW7DayTaichungCity     = "F-D0047-075"
	FTW2DayTainanCity       = "F-D0047-077"
	FTW7DayTainanCity       = "F-D0047-079"
	FTW2DayLienchiangCounty = "F-D0047-081"
	FTW7DayLienchiangCounty = "F-D0047-083"
	FTW2DayKinmenCounty     = "F-D0047-085"
	FTW7DayKinmenCounty     = "F-D0047-087"

	FTW2DayTaiwan = "F-D0047-089"
	FTW7DayTaiwan = "F-D0047-091"

	FTWTaiwan = "F-D0047-093"
)

type Forecast36HourWeather struct {
	Success string `json:"success"`
	Result  Result `json:"result"`
	Records struct {
		DatasetDescription string                 `json:"datasetDescription"`
		Location           []F36HWCountryLocation `json:"location"`
	} `json:"records"`
}

type F36HWCountryLocation struct {
	LocationName   string `json:"locationName"`
	WeatherElement []struct {
		ElementName string `json:"elementName"`
		Time        []struct {
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
			Parameter struct {
				ParameterName  string  `json:"parameterName"`
				ParameterValue *string `json:"parameterValue,omitempty"`
				ParameterUnit  *string `json:"parameterUnit,omitempty"`
			} `json:"parameter"`
		} `json:"time"`
	} `json:"weatherElement"`
}

type F36HWTime struct {
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Parameter Parameter `json:"parameter"`
}

// Get36HourWeather gets 36-hour weather forecasts.
func (s *ForecastsService) Get36HourWeather(ctx context.Context, locationNames, elements []string) (*Forecast36HourWeather, *http.Response, error) {
	forecast := new(Forecast36HourWeather)
	req, err := s.client.Get(ctx, s.generateURL(F36HW, nil, locationNames, elements), forecast)
	if err != nil {
		return nil, nil, err
	}
	return forecast, req, nil
}

type ForecastTownshipsWeather struct {
	Success string `json:"success"`
	Result  Result `json:"result"`
	Records struct {
		Locations []FTWCountryLocation `json:"locations"`
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
	Description    *string   `json:"description,omitempty"`
	ElementName    string    `json:"elementName"`
	ElementMeasure *string   `json:"elementMeasure,omitempty"`
	Time           []FTWTime `json:"time"`
}

type FTWTime struct {
	DataTime     *string           `json:"dataTime,omitempty"`
	StartTime    *string           `json:"startTime,omitempty"`
	EndTime      *string           `json:"endTime,omitempty"`
	ElementValue []FTWElementValue `json:"elementValue,omitempty"`
	Parameter    []Parameter       `json:"parameter,omitempty"`
}

type FTWElementValue struct {
	Value    string `json:"value"`
	Measures string `json:"measures"`
}

// GetTownshipsWeatherByDataId gets townships forecasts by data Id.
// See http://opendata.cwb.gov.tw/datalist for dataId (F-D0047-001 - F-D0047-091).
func (s *ForecastsService) GetTownshipsWeatherByDataId(ctx context.Context, dataId string, locationNames, elements []string) (*ForecastTownshipsWeather, *http.Response, error) {
	forecast := new(ForecastTownshipsWeather)
	req, err := s.client.Get(ctx, s.generateURL(dataId, nil, locationNames, elements), forecast)
	if err != nil {
		return nil, nil, err
	}
	return forecast, req, nil
}

// GetTownshipsWeatherByLocations gets townships forecasts by locationIds and locationNames.
func (s *ForecastsService) GetTownshipsWeatherByLocations(ctx context.Context, locationIds, locationNames, elements []string) (*ForecastTownshipsWeather, *http.Response, error) {
	forecast := new(ForecastTownshipsWeather)
	req, err := s.client.Get(ctx, s.generateURL(FTWTaiwan, locationIds, locationNames, elements), forecast)
	if err != nil {
		return nil, nil, err
	}
	return forecast, req, nil
}

func (s *ForecastsService) generateURL(dataId string, locationIds, locationNames, elements []string) string {
	q := url.Values{}
	if len(locationIds) > 0 {
		q.Set("locationId", strings.Join(locationIds, ","))
	}
	if len(locationNames) > 0 {
		q.Set("locationName", strings.Join(locationNames, ","))
	}
	if len(elements) > 0 {
		q.Set("elementName", strings.Join(elements, ","))
	}
	return s.client.generateURL(dataId, q)
}
