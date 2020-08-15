package cwb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTideForecastService_Get1MonthTide(t *testing.T) {
	// Arrange
	setup()
	defer teardown()

	testdata, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/%s.json", Tide1MonthId))

	mux.HandleFunc(fmt.Sprintf("/api/v1/rest/datastore/%s", Tide1MonthId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"locationName": "宜蘭縣南澳鄉",
			"sort":         "validTime",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(testdata)
	})

	options := TideForecast1MonthOptions{
		LocationName: "宜蘭縣南澳鄉",
		Sort:         "validTime",
	}

	// Act
	got, _, err := client.TideForecasts.Get1MonthTide(context.Background(), options)

	// Assert
	if err != nil {
		t.Errorf("TideForecasts.Get1MonthTide returned error: %v", err)
	}

	restored, _ := json.Marshal(got)
	areEqual, _ := areEqualJSON(testdata, restored)
	if !areEqual {
		t.Error("TideForecasts.Get1MonthTide testdata and restored are not equal")
	}
}
