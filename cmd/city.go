package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-weather/api"
	"os"
)

var country string
var units string
var detailed bool

var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "display the current weather data for a city",
	Long: `The city command will display the current weather data for a specific city,
If a city name exist in multiple counties, the command will only get the weather data for the country the open weather api sets as default`,
	Example: `go-weather city london 
go-weather city london --country ca -units I --detailed
go-weather city montreal -c ca -u S -d`,
	Run: cityRun,
}

func cityRun(cmd *cobra.Command, args []string) {

	// Make sur a city name is provided
	if len(args) == 0 {
		fmt.Println("You need to provide a city name")
		os.Exit(1)
	}

	city := args[0]

	// Get currentWeather data for the city
	currentWeather, err := api.GetWeatherByCity(city, country, api.GetUnits(units))
	if err != nil {
		fmt.Println("Command error:", err)
		os.Exit(1)
	}

	//Display  weather data
	api.PrintWeatherData(currentWeather, detailed, units)
}

func init() {
	rootCmd.AddCommand(cityCmd)
	cityCmd.Flags().StringVarP(&country, "country", "c", "", "The country where the city is located. Only the country code is required. Example: ca for Canada or fr for France")
	cityCmd.Flags().StringVarP(&units, "units", "u", "M", "Units you want weather data displayed in scientific S(Kelvin, m/s), metrics M(Celsius, m/s) or imperial I(Fahrenheit, mph).")
	cityCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Display a more detailed version of the weather data")
}
