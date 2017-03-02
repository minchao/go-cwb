package cwb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestForecastsService_Get36HourWeather(t *testing.T) {
	setup()
	defer teardown()

	testdata, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/%v.json", F36HW))

	mux.HandleFunc(fmt.Sprintf("/api/v1/rest/datastore/%v", F36HW), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"locationName": "宜蘭縣",
			"elementName":  "Wx,PoP,CI,MinT,MaxT",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(testdata)
	})

	got, _, err := client.Forecasts.Get36HourWeather(context.Background(),
		[]string{"宜蘭縣"},
		[]string{"Wx,PoP,CI,MinT,MaxT"})
	if err != nil {
		t.Errorf("Forecasts.Get36HourWeather returned error: %v", err)
	}

	restored, _ := json.Marshal(got)
	areEqual, err := areEqualJSON(testdata, restored)
	if !areEqual {
		t.Error("Forecasts.Get36HourWeather testdata and restored are not equal")
	}
}

func TestForecastsService_GetTownshipsWeatherByDataId(t *testing.T) {
	setup()
	defer teardown()

	testdata, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/%v.json", FTW2DayYilanCounty))

	mux.HandleFunc(fmt.Sprintf("/api/v1/rest/datastore/%v", FTW2DayYilanCounty), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"locationName": "羅東鎮",
			"elementName":  "Wx,PoP,AT,T,CI,RH,WeatherDescription,PoP6h,Wind,Td",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(testdata)
	})

	got, _, err := client.Forecasts.GetTownshipsWeatherByDataId(context.Background(),
		FTW2DayYilanCounty,
		[]string{"羅東鎮"},
		[]string{"Wx,PoP,AT,T,CI,RH,WeatherDescription,PoP6h,Wind,Td"})
	if err != nil {
		t.Errorf("Forecasts.GetTownshipsWeatherByDataId returned error: %v", err)
	}

	restored, _ := json.Marshal(got)
	areEqual, err := areEqualJSON(testdata, restored)
	if !areEqual {
		t.Error("Forecasts.GetTownshipsWeatherByDataId testdata and restored are not equal")
	}
}

func TestForecastsService_GetTownshipsWeatherByLocations(t *testing.T) {
	setup()
	defer teardown()

	testdata, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/%v.json", FTWTaiwan))

	mux.HandleFunc(fmt.Sprintf("/api/v1/rest/datastore/%v", FTWTaiwan), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"locationId":   "F-D0047-033,F-D0047-031",
			"locationName": "恆春鎮,阿里山鄉",
			"elementName":  "Wx,PoP,AT,T,CI,RH,WeatherDescription,PoP6h,Wind,Td",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(testdata)
	})

	got, _, err := client.Forecasts.GetTownshipsWeatherByLocations(context.Background(),
		[]string{"F-D0047-033,F-D0047-031"},
		[]string{"恆春鎮,阿里山鄉"},
		[]string{"Wx,PoP,AT,T,CI,RH,WeatherDescription,PoP6h,Wind,Td"})
	if err != nil {
		t.Errorf("Forecasts.GetTownshipsWeatherByLocations returned error: %v", err)
	}

	restored, _ := json.Marshal(got)
	areEqual, err := areEqualJSON(testdata, restored)
	if !areEqual {
		t.Error("Forecasts.GetTownshipsWeatherByLocations testdata and restored are not equal")
	}
}
