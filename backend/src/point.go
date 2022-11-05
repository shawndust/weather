package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Point struct {
	Data []struct {
		Latitude            float32 `json:"latitude"`
		Longitude           float32 `json:"longitude"`
		Type                string  `json:"type"`
		Name                string  `json:"name"`
		Number              string  `json:"number"`
		Postal_code         string  `json:"postal_code"`
		Street              string  `json:"street"`
		Confidence          float32 `json:"confidence"`
		Region              string  `json:"region"`
		County              string  `json:"county"`
		Locality            string  `json:"locality"`
		Administrative_area string  `json:"administrative_area"`
		Neighbourhood       string  `json:"neighbourhood"`
		Country             string  `json:"country"`
		Country_code        string  `json:"country_code"`
		Continent           string  `json:"continent"`
		Label               string  `json:"label"`
	}
}

func getPoint(freetext string) Point {
	var point Point
	getQuery := fmt.Sprintf("http://api.positionstack.com/v1/forward?access_key=1e5cbc72a27fdd2a77bb04b38dd79a93&query=%s", freetext)
	resp, err := http.Get(getQuery)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &point)

	if err != nil {
		panic(err)
	}

	return point
}
