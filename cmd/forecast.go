package cmd

import (
	"fmt"
	"go-weather/api"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var count int

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get the weather forecast",
}

var forecastCityCmd = &cobra.Command{
	Use:   "city [city_name]",
	Args:  cobra.ExactArgs(1), // Make sure a city name is provided
	Short: "Get the daily weather forecast for a city",
	Long: `The city command will display the weather forecast for a specific city for the next 7 days,
If a city name exist in multiple countries, the command will only get the weather forecast for the country the open weather api sets as default`,
	Example: `go-weather forecast city london 
go-weather forecast city london --country ca --units I --detailed
go-weather forecast city madrid --count 4 --units S --detailed
go-weather forecast city montreal -c ca -u S -d
go-weather forecast city new-york --units I`,
	Run: forecastCityRun,
}

var forecastZipCmd = &cobra.Command{
	Use:   "zip [zip_code]",
	Args:  cobra.ExactArgs(1),
	Short: "Get the daily weather forecast for a specific zip code",
	Long: `The zip command will display the weather forecast for a specific zip code for the next 7 days,
For canadian zip codes only the first three letters are necessary. 
This command also require the use of the country flag for zip codes outside of the usa or else the search works for usa as default.
The latter case could lead an error if the zip code does not exist in the usa`,
	Example: `go-weather forecast zip h1a -c ca 
go-weather forecast zip j4b --country ca --units I --detailed
go-weather forecast zip 75001 -c fr -u S -d
go-weather forecast zip 94040`,
	Run: forecastZipRun,
}

func forecastCityRun(cmd *cobra.Command, args []string) {

	// For cities that contains two substrings, replace the "-" character "+" character
	city := strings.ReplaceAll(args[0], "-", "+")

	// Check for count value
	if (count < 1) || (count > 16) {
		fmt.Println("Error: the number of days must be between 1 and 16")
		os.Exit(1)
	}

	// Get currentWeather data for the city
	weatherForecast, err := api.ForecastByCity(city, country, count, api.GetUnits(units))
	if err != nil {
		fmt.Println("Command error:", err)
		os.Exit(1)
	}

	// Display current weather
	api.PrintWeatherForecast(weatherForecast, count, detailed, units)
}

func forecastZipRun(cmd *cobra.Command, args []string) {

	zipCode := args[0]

	// Check for count value
	if (count < 1) || (count > 16) {
		fmt.Println("Error: the number of days must be between 1 and 16")
		os.Exit(1)
	}

	// Get currentWeather data for the city
	weatherForecast, err := api.ForecastByZipCode(zipCode, country, count, api.GetUnits(units))
	if err != nil {
		fmt.Println("Command error:", err)
		os.Exit(1)
	}

	// Display current weather
	api.PrintWeatherForecast(weatherForecast, count, detailed, units)
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.AddCommand(forecastCityCmd, forecastZipCmd)

	forecastCmd.PersistentFlags().StringVarP(&country, "country", "c", "", "The country where the city is located. Only the country code is required. Example: ca for Canada or fr for France")
	forecastCmd.PersistentFlags().StringVarP(&units, "units", "u", "M", "Units you want weather data displayed in scientific S(Kelvin, m/s), metrics M(Celsius, m/s) or imperial I(Fahrenheit, mph).")
	forecastCmd.PersistentFlags().IntVarP(&count, "count", "n", 7, "The number of days of daily weather forecast for a city. Maximum of 16 and a minimum of 1")
	forecastCmd.PersistentFlags().BoolVarP(&detailed, "detailed", "d", false, "Display a more detailed version of the weather data")
}
