package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Grid struct {
	Context  []interface{} `json:"@context"`
	ID       string        `json:"id"`
	Type     string        `json:"type"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		ID                  string `json:"@id"`
		Type                string `json:"@type"`
		Cwa                 string `json:"cwa"`
		ForecastOffice      string `json:"forecastOffice"`
		GridID              string `json:"gridId"`
		GridX               int    `json:"gridX"`
		GridY               int    `json:"gridY"`
		Forecast            string `json:"forecast"`
		ForecastHourly      string `json:"forecastHourly"`
		ForecastGridData    string `json:"forecastGridData"`
		ObservationStations string `json:"observationStations"`
		RelativeLocation    struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Properties struct {
				City     string `json:"city"`
				State    string `json:"state"`
				Distance struct {
					UnitCode string  `json:"unitCode"`
					Value    float64 `json:"value"`
				} `json:"distance"`
				Bearing struct {
					UnitCode string `json:"unitCode"`
					Value    int    `json:"value"`
				} `json:"bearing"`
			} `json:"properties"`
		} `json:"relativeLocation"`
		ForecastZone    string `json:"forecastZone"`
		County          string `json:"county"`
		FireWeatherZone string `json:"fireWeatherZone"`
		TimeZone        string `json:"timeZone"`
		RadarStation    string `json:"radarStation"`
	} `json:"properties"`
}

func getGrid(latitude float32, longitude float32) Grid {
	var grid Grid
	getQuery := fmt.Sprintf("http://api.weather.gov/points/%f,%f", latitude, longitude)
	resp, err := http.Get(getQuery)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &grid)

	if err != nil {
		panic(err)
	}

	return grid
}
