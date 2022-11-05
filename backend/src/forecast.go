package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Forecast struct {
	Context  []interface{} `json:"@context"`
	Type     string        `json:"type"`
	Geometry struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Updated           time.Time `json:"updated"`
		Units             string    `json:"units"`
		ForecastGenerator string    `json:"forecastGenerator"`
		GeneratedAt       time.Time `json:"generatedAt"`
		UpdateTime        time.Time `json:"updateTime"`
		ValidTimes        string    `json:"validTimes"`
		Elevation         struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		Periods []struct {
			Number           int         `json:"number"`
			Name             string      `json:"name"`
			StartTime        string      `json:"startTime"`
			EndTime          string      `json:"endTime"`
			IsDaytime        bool        `json:"isDaytime"`
			Temperature      int         `json:"temperature"`
			TemperatureUnit  string      `json:"temperatureUnit"`
			TemperatureTrend interface{} `json:"temperatureTrend"`
			WindSpeed        string      `json:"windSpeed"`
			WindDirection    string      `json:"windDirection"`
			Icon             string      `json:"icon"`
			ShortForecast    string      `json:"shortForecast"`
			DetailedForecast string      `json:"detailedForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

func getForecast(office string, gridX int, gridY int) string {
	var forecast Forecast

	getQuery := fmt.Sprintf("http://api.weather.gov/gridpoints/%s/%d,%d/forecast", office, gridX, gridY)
	resp, err := http.Get(getQuery)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &forecast)

	if err != nil {
		panic(err)
	}

	theCurrentForecast := forecast.Properties.Periods[0].DetailedForecast
	return theCurrentForecast
}
