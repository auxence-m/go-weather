## go-weather: A Weather CLI App in Go

`go-weather` is a simple CLI tool built with Golang that fetches current weather or weather forecasts (up to 16 days) for a specific city or zip code. It uses the [spf13/cobra](https://github.com/spf13/cobra) package for command-line interface functionality.

## Features
- Fetch current weather information.
- Get weather forecasts for up to 16 days.
- Supports querying by city name or zip code.
- Allows specifying units (e.g., Celsius/Fahrenheit) and country code for international weather data.

## How to run

### Prerequisites
- Ensure you have [Golang](https://go.dev/doc/install) installed on your machine.

### Installation
1. Clone the repository 
```
git clone https://github.com/auxence-m/go-weather.git
```
2. Navigate to the Project Directory

```
cd go-weather
```

3. Install Dependencies

```
go mod tidy
```

4. Build the Application

```
go build
``` 

After building, you'll find the `go-weather` executable (`go-weather.exe` on Windows) in your project directory.

5. Run the Application

```
go-weather [command] [subcommand] --flag
```

## Available Commands

### 1. `current`: Get Current Weather
Each command and subcommand comes with a built-in `--help` flag to describe its usage.

#### Subcommands: `city` and `zip`
 - **City Example:**
```
go-weather current city london 
go-weather current city london --country ca --units I --detailed
```
For more details on the `current city` subcommand
``` 
go-weather current city --help

```

- **Zip Example:**
```
go-weather current zip 75001 --country fr --units S -detailed
go-weather current zip j4b --country ca --units I --detailed
go-weather current zip 94040
```
For more details on the `current zip` subcommand
``` 
go-weather current zip --help

```

##### Flags:

- `--country`: Specify the country code (e.g., `ca`, `fr`).
- `--units`: Specify the unit system (e.g., `I` for imperial, `S` for scientific, `M` for metric).
- `--detailed`: Provides additional detailed weather information (optional).


### 2. `forecast`: Get Weather Forecast

#### Subcommands: `city` and `zip`
- **City Example:**
```
go-weather forecast city london
go-weather forecast city madrid --count 5 --units S --detailed
```
For more details on the `forecast city` subcommand
``` 
go-weather forecast city --help

```

- **Zip Example:**
```
go-weather current zip 75001 --country fr --units S -detailed
go-weather forecast zip k1n --country ca --units I --detailed
go-weather forecast zip 94040
```
For more details on the `forecast zip` subcommand
``` 
go-weather forecast zip --help

```

##### Flags:

- `--country`: Specify the country code (e.g., `ca`, `fr`).
- `--units`: Specify the unit system (e.g., `I` for imperial, `S` for scientific, `M` for metric).
- `--count`: Specify the number of days for the forecast (default is 7).
- `--detailed` : Provides additional detailed weather information (optional).

### Additional Notes
- Country Codes: Use ISO 3166-1 alpha-2 codes (e.g., `ca` for Canada, `gb` for the United Kingdom, `es` for Spain).
- Units: `M` for metric (Celsius, meters, etc.), `I` for imperial (Fahrenheit, miles, etc.), `S` for scientific (Kelvin, meters, etc.).

### License
This project is licensed under the MIT License.