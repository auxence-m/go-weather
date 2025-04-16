package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

var apiKey string

type WeatherData struct {
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

// Reads config file and sets the api key
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

// GetWeatherByCity Collects weather data using a city name
func GetWeatherByCity(city string, country string, units string) (WeatherData, error) {

	// Constructing the api url using city name
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=%s&APPID=%s", city, country, units, apiKey)

	// Http get request to open weather api
	response, err := http.Get(apiUrl)
	if err != nil {
		return WeatherData{}, err
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return WeatherData{}, err
	}

	// Parsing JSON response to a currentWeather object
	var currentWeather WeatherData
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return WeatherData{}, err
	}

	// Handle specific bad request errors
	if currentWeather.Code != 200 {
		return WeatherData{}, errors.New(currentWeather.Message)
	}

	return currentWeather, nil
}

// GetWeatherByZipCode Collects weather data using a zipcode
func GetWeatherByZipCode(zipCode string, country string, units string) (WeatherData, error) {

	// Constructing the api url using postal code
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?zip=%s,%s&units=%s&APPID=%s", zipCode, country, units, apiKey)

	// Http get request to open weather api
	response, err := http.Get(apiUrl)
	if err != nil {
		return WeatherData{}, err
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return WeatherData{}, err
	}

	// Parsing JSON response to a currentWeather object
	var currentWeather WeatherData
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return WeatherData{}, err
	}

	return currentWeather, nil
}

func PrintWeatherData(currentWeather WeatherData, detailed bool, units string) {

	// Create a tabwriter instance.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Println("Here is the current Condition data.")
	fmt.Fprintln(w, "City:\t", currentWeather.Name, currentWeather.Sys.Country)
	fmt.Fprintln(w, "Temperature:\t", currentWeather.Main.Temp, printTemp(units))
	fmt.Fprintln(w, "Feels like:\t", currentWeather.Main.FeelsLike, printTemp(units))
	fmt.Fprintln(w, "Min Temperature:\t", currentWeather.Main.TempMin, printTemp(units))
	fmt.Fprintln(w, "Max Temperature:\t", currentWeather.Main.TempMax, printTemp(units))
	fmt.Fprintln(w, "Condition:\t", currentWeather.Weather[0].Description)

	if detailed {
		fmt.Fprintln(w, "Humidity:\t", currentWeather.Main.Humidity, "%")
		fmt.Fprintln(w, "Pressure :\t", currentWeather.Main.Pressure, "hPa")
		fmt.Fprintln(w, "Cloudiness :\t", currentWeather.Clouds.All, "%")
		fmt.Fprintln(w, "Wind speed :\t", currentWeather.Wind.Speed, printSpeed(units))
		fmt.Fprintln(w, "Wind direction :\t", currentWeather.Wind.Deg, "%")
		fmt.Fprintln(w, "Wind gust :\t", currentWeather.Wind.Gust, printSpeed(units))
		fmt.Fprintln(w, "Sunrise:\t", time.Unix(int64(currentWeather.Sys.Sunrise), 0).Format(time.TimeOnly))
		fmt.Fprintln(w, "Sunset:\t", time.Unix(int64(currentWeather.Sys.Sunset), 0).Format(time.TimeOnly))
		fmt.Fprintln(w, "Longitude:\t", currentWeather.Coordinates.Lon)
		fmt.Fprintln(w, "Latitude:\t", currentWeather.Coordinates.Lat)
	}

	fmt.Fprintln(w, "Time of data collection:\t", time.Unix(int64(currentWeather.Dt), 0).Format(time.DateTime))

	err := w.Flush()
	if err != nil {
		return
	}
}

// GetUnits Convert units flag value into the correct metric value
func GetUnits(flag string) string {
	switch flag {
	case "S":
		return "standard"
	case "I":
		return "imperial"
	default:
		return "metric"
	}
}

// Print temperature symbols based on unit system
func printSpeed(units string) string {
	switch units {
	case "S":
		return "m/s"
	case "I":
		return "mph"
	default:
		return "m/s"
	}
}

// Print speed symbols based on unit system
func printTemp(units string) string {
	switch units {
	case "S":
		return "K"
	case "I":
		return "°F"
	default:
		return "°C"
	}
}

func init() {
	initConfig()
}
