package cmd

import (
	"fmt"
	"go-weather/api"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var country string
var units string
var detailed bool

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current weather",
}

var currentCityCmd = &cobra.Command{
	Use:   "city [city_name]",
	Args:  cobra.ExactArgs(1), // Make sure a city name is provided
	Short: "Gets the current weather data for a city",
	Long: `The city command will display the current weather data for a specific city,
If a city name exist in multiple countries, the command will only get the weather data for the country the open weather api sets as default`,
	Example: `go-weather current city london 
go-weather current city london --country ca --units I --detailed
go-weather current city montreal -c ca -u S -d
go-weather current city new-york --units I`,
	Run: currentCityRun,
}

var currentZipCmd = &cobra.Command{
	Use:   "zip [zip_code]",
	Args:  cobra.ExactArgs(1), // Make sure a city name is provided
	Short: "Get the current weather data for a specific zip code",
	Long: `The zip command will display the current weather data for a specific zip code,
For canadian zip codes only the first three letters are necessary. 
This command also require the use of the country flag for zip codes outside of the usa or else the search works for usa as default.
The latter case could lead an error if the zip code does not exist in the usa`,
	Example: `go-weather zip h1a -c ca 
go-weather current zip j4b --country ca --units I --detailed
go-weather current zip 75001 -c fr -u S -d
go-weather current zip 94040`,
	Run: currentZipRun,
}

func currentCityRun(cmd *cobra.Command, args []string) {

	// For cities that contains two substrings, replace the "-" character "+" character
	city := strings.ReplaceAll(args[0], "-", "+")

	// Get currentWeather data for the city
	currentWeather, err := api.CurrentWeatherByCity(city, country, api.GetUnits(units))
	if err != nil {
		fmt.Println("Command error:", err)
		os.Exit(1)
	}

	// Display current weather
	api.PrintCurrentWeather(currentWeather, detailed, units)
}

func currentZipRun(cmd *cobra.Command, args []string) {

	zipCode := args[0]

	// Get currentWeather data for the city
	currentWeather, err := api.CurrentWeatherByZipCode(zipCode, country, api.GetUnits(units))
	if err != nil {
		fmt.Println("Command error:", err)
		os.Exit(1)
	}

	// Display current weather
	api.PrintCurrentWeather(currentWeather, detailed, units)
}

func init() {
	rootCmd.AddCommand(currentCmd)
	currentCmd.AddCommand(currentCityCmd, currentZipCmd)

	currentCityCmd.Flags().StringVarP(&country, "country", "c", "", "The country where the city is located. Only the country code is required. Example: ca for Canada or fr for France")
	currentCityCmd.Flags().StringVarP(&units, "units", "u", "M", "Units you want weather data displayed in scientific S(Kelvin, m/s), metrics M(Celsius, m/s) or imperial I(Fahrenheit, mph).")
	currentCityCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Display a more detailed version of the weather data")

	currentZipCmd.Flags().StringVarP(&country, "country", "c", "", "The country where the city is located. Only the country code is required. Example: ca for Canada or fr for France")
	currentZipCmd.Flags().StringVarP(&units, "units", "u", "M", "Units you want weather data displayed in scientific S(Kelvin, m/s), metrics M(Celsius, m/s) or imperial I(Fahrenheit, mph).")
	currentZipCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Display a more detailed version of the weather data")
}
