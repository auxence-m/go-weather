package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-weather/api"
	"os"
)

var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "display the current weather data for a specific zip code",
	Long: `The zip command will display the current weather data for a specific zip code,
For canadian zip codes only the first three letters are necessary. 
This command also require the use of the country flag for zip codes outside of the usa or else the search works for usa as default.
The latter case could lead an error if the zip code does not exist in the usa`,
	Example: `go-weather zip h1a -c ca 
go-weather zip j4b --country ca --units I --detailed
go-weather zip 75001 -c fr -u S -d
go-weather zip 94040`,
	Run: zipRun,
}

func zipRun(cmd *cobra.Command, args []string) {

	// Make sur a city name is provided
	if len(args) == 0 {
		fmt.Println("You need to provide a zip code")
		os.Exit(1)
	}

	zipCode := args[0]

	// Get currentWeather data for the city
	currentWeather, err := api.GetWeatherByZipCode(zipCode, country, api.GetUnits(units))
	if err != nil {
		fmt.Println("Command error:", err)
		os.Exit(1)
	}

	//Display  weather data
	api.PrintWeatherData(currentWeather, detailed, units)
}

func init() {
	rootCmd.AddCommand(zipCmd)
	zipCmd.Flags().StringVarP(&country, "country", "c", "", "The country where the city is located. Only the country code is required. Example: ca for Canada or fr for France")
	zipCmd.Flags().StringVarP(&units, "units", "u", "M", "Units you want weather data displayed in scientific S(Kelvin, m/s), metrics M(Celsius, m/s) or imperial I(Fahrenheit, mph).")
	zipCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Display a more detailed version of the weather data")
}
