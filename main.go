package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func main() {
	apiKey := "86112171bb3c767fcb25a57b118dffe8"
	city := "London"

	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		type WeatherData struct {
			Main struct {
				Temp      float64 `json:"temp"`
				FeelsLike float64 `json:"feels_like"`
				Humidity  int     `json:"humidity"`
				Pressure  int     `json:"pressure"`
			} `json:"main"`
			Wind struct {
				Speed float64 `json:"speed"`
				Deg   int     `json:"deg"`
			} `json:"wind"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
		}

		var data WeatherData

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}

		temp := data.Main.Temp
		feelsLike := data.Main.FeelsLike
		humidity := data.Main.Humidity
		pressure := data.Main.Pressure
		windSpeed := data.Wind.Speed
		windDeg := data.Wind.Deg
		weather := data.Weather[0].Description

		tempC := temp - 273.15
		feelsLikeC := feelsLike - 273.15

		fmt.Printf("The weather in %s is %s.\n", city, weather)
		fmt.Printf("The temperature is %.2f °C, feels like %.2f °C.\n", tempC, feelsLikeC)
		fmt.Printf("The humidity is %d%%, the pressure is %d hPa, and the wind is %.2f m/s at %d degrees.\n", humidity, pressure, windSpeed, windDeg)
	} else {
		fmt.Printf("Something went wrong. The API call returned status code %d.\n", resp.StatusCode)
	}
}
