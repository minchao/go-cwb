package cwb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestStationObsElements_GetByName(t *testing.T) {
	elements := StationObsElements{
		{
			ElementName:  "TEMP",
			ElementValue: "5.9",
		},
	}

	element, err := elements.GetByName("TEMP")
	if err != nil {
		t.Errorf("StationObsElements.GetByName returned err: %v", err)
	}

	if element.ElementName != "TEMP" {
		t.Error("StationObsElements.GetByName testdata and restored are not equal")
	}
}

func TestStationObsElements_GetByName_notFound(t *testing.T) {
	elements := StationObsElements{}

	_, err := elements.GetByName("TEMP")
	if err != ErrElementNotFound {
		t.Errorf("StationObsElements.GetByName should return ErrElementNotFound")
	}
}

func TestStationObsService_GetWeather(t *testing.T) {
	setup()
	defer teardown()

	testdata, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/%v.json", StationObsWeatherId))

	mux.HandleFunc(fmt.Sprintf("/api/v1/rest/datastore/%v", StationObsWeatherId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"locationName": "合歡山",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(testdata)
	})

	options := url.Values{}
	options.Set("locationName", "合歡山")

	got, _, err := client.StationObs.GetWeather(context.Background(), options)
	if err != nil {
		t.Errorf("StationObs.GetWeather returned error: %v", err)
	}

	restored, _ := json.Marshal(got)

	areEqual, _ := areEqualJSON(testdata, restored)
	if !areEqual {
		t.Error("StationObs.GetWeather testdata and restored are not equal")
	}
}

func TestStationObsService_GetRainfall(t *testing.T) {
	setup()
	defer teardown()

	testdata, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/%v.json", StationObsRainfallId))

	mux.HandleFunc(fmt.Sprintf("/api/v1/rest/datastore/%v", StationObsRainfallId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"locationName": "合歡山",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(testdata)
	})

	options := url.Values{}
	options.Set("locationName", "合歡山")

	got, _, err := client.StationObs.GetRainfall(context.Background(), options)
	if err != nil {
		t.Errorf("StationObs.GetRainfall returned error: %v", err)
	}

	restored, _ := json.Marshal(got)

	areEqual, _ := areEqualJSON(testdata, restored)
	if !areEqual {
		t.Error("StationObs.GetRainfall testdata and restored are not equal")
	}
}
