package main

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getWeatherForecast(locationFreetext string) string {
	point := getPoint(locationFreetext)

	grid := getGrid(point.Data[0].Latitude, point.Data[0].Longitude)

	officeURL := strings.Split(grid.Properties.ForecastOffice, "/")
	office := officeURL[len(officeURL)-1]

	forecast := getForecast(office, grid.Properties.GridX, grid.Properties.GridY)

	return forecast
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
	}))
	r.GET("/forecast", func(c *gin.Context) {
		locationToBeForecast := c.Request.URL.Query().Get("freetext")
		forecast := getWeatherForecast(locationToBeForecast)
		c.JSON(200, gin.H{
			"forecast": forecast,
		})
	})
	r.Run(":8090")
}
