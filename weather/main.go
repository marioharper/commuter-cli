package weather

import (
	"log"
	"strconv"
	"time"

	forecast "github.com/mlbright/forecast/v2"
)

// supplied via -ldflags in Makefile
var DARK_SKY_API_KEY = ""

type WeatherInfo struct {
	Summary           string
	PrecipProbability float64
	Temp              float64
	Wind              float64
}

func GetInfo(idealTime int, AMPM string, latitude float64, longitude float64) WeatherInfo {

	var info WeatherInfo
	if AMPM == "PM" {
		idealTime += 12
	}

	lat := strconv.FormatFloat(latitude, 'f', 6, 64)
	lng := strconv.FormatFloat(longitude, 'f', 6, 64)

	//Get forecast for commute starting point
	f, err := forecast.Get(DARK_SKY_API_KEY, lat, lng, "now", forecast.US)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range f.Hourly.Data {
		var weatherHour = strconv.FormatFloat(value.Time, 'f', 0, 64)
		weatherHourInt, err := strconv.ParseInt(weatherHour, 10, 64)
		hr, _, _ := time.Unix(weatherHourInt, 0).Clock()

		if err != nil {
			panic(err)
		}

		if hr == idealTime {
			info.Summary = value.Summary
			info.Temp = value.Temperature
			info.PrecipProbability = (value.PrecipProbability * 100)
			info.Wind = value.WindSpeed
			break
		}
	}

	return info
}