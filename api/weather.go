package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/viper"
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

type WeatherForecast struct {
	City struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
	} `json:"city"`
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt      int `json:"dt"`
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
		Temp    struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		Weather  []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Speed  float64 `json:"speed"`
		Deg    int     `json:"deg"`
		Gust   float64 `json:"gust"`
		Clouds int     `json:"clouds"`
		Rain   float64 `json:"rain"`
		Snow   float64 `json:"snow"`
		Pop    float64 `json:"pop"`
	} `json:"list"`
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

// CurrentWeatherByCity Collects current weather data using a city name
func CurrentWeatherByCity(city string, country string, units string) (CurrentWeather, error) {

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

// CurrentWeatherByZipCode Collects current weather data using a zipcode
func CurrentWeatherByZipCode(zipCode string, country string, units string) (CurrentWeather, error) {

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

// ForecastByCity Collects weather forecast data using a city name
func ForecastByCity(city string, country string, count int, units string) (WeatherForecast, error) {

	// Constructing the api url using a city name
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast/daily?q=%s,%s&cnt=%d&units=%s&APPID=%s", city, country, count, units, apiKey)

	// Http get request to open weather api
	response, err := http.Get(apiUrl)
	if err != nil {
		return WeatherForecast{}, err
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return WeatherForecast{}, err
	}

	// Parsing JSON response to a currentWeather object
	var weatherForecast WeatherForecast
	err = json.Unmarshal(body, &weatherForecast)
	if err != nil {
		return WeatherForecast{}, err
	}

	return weatherForecast, nil
}

// ForecastByZipCode Collects weather forecast data using a zip code
func ForecastByZipCode(zipCode string, country string, count int, units string) (WeatherForecast, error) {

	// Constructing the api url using postal code
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast/daily?zip=%s,%s&cnt=%d&units=%s&APPID=%s", zipCode, country, count, units, apiKey)

	// Http get request to open weather api
	response, err := http.Get(apiUrl)
	if err != nil {
		return WeatherForecast{}, err
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return WeatherForecast{}, err
	}

	// Parsing JSON response to a currentWeather object
	var weatherForecast WeatherForecast
	err = json.Unmarshal(body, &weatherForecast)
	if err != nil {
		return WeatherForecast{}, err
	}

	return weatherForecast, nil
}

// PrintCurrentWeather Prints the current weather data onto the console
func PrintCurrentWeather(weatherData CurrentWeather, detailed bool, units string) {

	// Create a tabwriter instance.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Printf("Here is the current weather for %s, %s: \n", weatherData.Name, weatherData.Sys.Country)
	fmt.Fprintln(w, "Temperature:\t", weatherData.Main.Temp, printTemp(units))
	fmt.Fprintln(w, "Feels like:\t", weatherData.Main.FeelsLike, printTemp(units))
	fmt.Fprintln(w, "Min Temperature:\t", weatherData.Main.TempMin, printTemp(units))
	fmt.Fprintln(w, "Max Temperature:\t", weatherData.Main.TempMax, printTemp(units))
	fmt.Fprintln(w, "Condition:\t", weatherData.Weather[0].Description)
	fmt.Fprintln(w, "Humidity:\t", weatherData.Main.Humidity, "%")

	if detailed {
		fmt.Fprintln(w, "Pressure :\t", weatherData.Main.Pressure, "hPa")
		fmt.Fprintln(w, "Cloudiness :\t", weatherData.Clouds.All, "%")
		fmt.Fprintln(w, "Wind speed :\t", weatherData.Wind.Speed, printSpeed(units))
		fmt.Fprintln(w, "Wind direction :\t", weatherData.Wind.Deg, "%")
		fmt.Fprintln(w, "Wind gust :\t", weatherData.Wind.Gust, printSpeed(units))
		fmt.Fprintln(w, "Sunrise:\t", time.Unix(int64(weatherData.Sys.Sunrise), 0).Format(time.TimeOnly))
		fmt.Fprintln(w, "Sunset:\t", time.Unix(int64(weatherData.Sys.Sunset), 0).Format(time.TimeOnly))
		fmt.Fprintln(w, "Longitude:\t", weatherData.Coordinates.Lon)
		fmt.Fprintln(w, "Latitude:\t", weatherData.Coordinates.Lat)
	}

	fmt.Fprintln(w, "Date & Time of data collection:\t", time.Unix(int64(weatherData.Dt), 0).Format("Mon 02-01-2006 15:04:05"))
	err := w.Flush()
	if err != nil {
		return
	}
}

// PrintWeatherForecast Prints the current weather data onto the console
func PrintWeatherForecast(weatherForecast WeatherForecast, count int, detailed bool, units string) {

	// Create a tabwriter instance.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Printf("Here is the weather forecast for %s, %s (%d day forecast): \n", weatherForecast.City.Name, weatherForecast.City.Country, count)

	for _, day := range weatherForecast.List {
		date := time.Unix(int64(day.Dt), 0).Format("Monday 02 January 2006")
		fmt.Fprintln(w, "----------", date, "----------")
		fmt.Fprintln(w, "Min Temperature:\t", day.Temp.Min, printTemp(units))
		fmt.Fprintln(w, "Max Temperature:\t", day.Temp.Max, printTemp(units))
		fmt.Fprintln(w, "Average Temperature:\t", day.Temp.Day, printTemp(units))
		fmt.Fprintln(w, "Average Temperature Feels like:\t", day.FeelsLike.Day, printTemp(units))
		fmt.Fprintln(w, "Morning Temperature:\t", day.Temp.Morn, printTemp(units))
		fmt.Fprintln(w, "Morning Temperature Feels like:\t", day.FeelsLike.Morn, printTemp(units))
		fmt.Fprintln(w, "Evening Temperature:\t", day.Temp.Eve, printTemp(units))
		fmt.Fprintln(w, "Evening Temperature Feels like:\t", day.FeelsLike.Eve, printTemp(units))
		fmt.Fprintln(w, "Humidity:\t", day.Humidity, "%")
		fmt.Fprintln(w, "Condition:\t", day.Weather[0].Description)

		if detailed {
			fmt.Fprintln(w, "Pressure :\t", day.Pressure, "hPa")
			fmt.Fprintln(w, "Cloudiness\t", day.Clouds, "%")
			fmt.Fprintln(w, "Rain\t", day.Rain, "mm")
			fmt.Fprintln(w, "Clouds\t", day.Clouds, "mm")
			fmt.Fprintln(w, "Wind speed\t", day.Speed, printSpeed(units))
			fmt.Fprintln(w, "Wind direction\t", day.Deg, "%")
			fmt.Fprintln(w, "Wind gust\t", day.Gust, printSpeed(units))
			fmt.Fprintln(w, "Sunrise\t", time.Unix(int64(day.Sunrise), 0).Format(time.TimeOnly))
			fmt.Fprintln(w, "Sunset\t", time.Unix(int64(day.Sunset), 0).Format(time.TimeOnly))
			fmt.Fprintln(w, "Longitude\t", weatherForecast.City.Coord.Lon)
			fmt.Fprintln(w, "Latitude\t", weatherForecast.City.Coord.Lon)

		}
	}
	err := w.Flush()
	if err != nil {
		return
	}
}

// GetUnits Convert units flag value into the correct metric value for the api get request link
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

// printSpeed Print temperature symbols based on unit system
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

// printTemp Print speed symbols based on unit system
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
