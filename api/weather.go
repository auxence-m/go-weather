package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

var apiKey string

type CurrentWeather struct {
	Coordinates struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	Main struct {
		Temp        float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		TempMin     float64 `json:"temp_min"`
		TempMax     float64 `json:"temp_max"`
		Pressure    int     `json:"pressure"`
		Humidity    int     `json:"humidity"`
		SeaLevel    int     `json:"sea_level"`
		GroundLevel int     `json:"grnd_level"`
	} `json:"main"`

	Visibility int `json:"visibility"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Rain struct {
		Precipitation float64 `json:"1h"`
	} `json:"rain"`

	Snow struct {
		Precipitation float64 `json:"1h"`
	} `json:"snow"`

	Dt int `json:"dt"`

	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`

	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Code     int    `json:"cod"`
	Message  string `json:"message"`
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("An error occured while reading configuration file: %s", err)
	}

	apiKey = viper.GetString("OPEN_WEATHER_MAP_API_KEY")
}

func GetWeatherByCity(city string, country string, units string) (CurrentWeather, error) {

	// Constructing the api url using city name
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=%s&APPID=%s", city, country, units, apiKey)

	// Http get request to open weather api
	response, err := http.Get(apiUrl)
	if err != nil {
		return CurrentWeather{}, err
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return CurrentWeather{}, err
	}

	// Parsing JSON response to a currentWeather object
	var currentWeather CurrentWeather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return CurrentWeather{}, err
	}

	// Handle specific bad request errors
	if currentWeather.Code != 200 {
		return CurrentWeather{}, errors.New(currentWeather.Message)
	}

	return currentWeather, nil
}

func GetWeatherByZip(zipCode string, country string, units string) (CurrentWeather, error) {

	// Constructing the api url using postal code
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?zip=%s,%s&units=%s&APPID=%s", zipCode, country, units, apiKey)

	// Http get request to open weather api
	response, err := http.Get(apiUrl)
	if err != nil {
		return CurrentWeather{}, err
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return CurrentWeather{}, err
	}

	// Parsing JSON response to a currentWeather object
	var currentWeather CurrentWeather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return CurrentWeather{}, err
	}

	return currentWeather, nil
}

func init() {
	initConfig()
}
