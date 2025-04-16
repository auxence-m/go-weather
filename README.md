## This a whether cli app  made in golang

The app name is go-weather. go-weather is a weather cli app in golang using the [spf13 cobra package](https://github.com/spf13/cobra)
The app will print in the console the current weather data for a specific city of for a specific zip code

## How to run

1. Clone the repository `git clone https://github.com/Auxence-M/go-weather.git`
2. Change your directory to the project directory `cd go-weather`
3. Install dependencies `go mod tidy`
4. Modify the _.config.yaml_ file and provide your own open weather api key.
5. Run `go build` to build the app, a _go-weather.exe_ file should appear in the directory
6. You can now start using the app by using the following `./go-weather [command] --flag`

## List of commands with examples

### city command

```
go-weather city london 
go-weather city london --country ca -units I --detailed
go-weather city montreal -c ca -u S -d
```

### zip command

```
go-weather zip h1a -c ca 
go-weather zip j4b --country ca --units I --detailed
go-weather zip 75001 -c fr -u S -d
go-weather zip 94040
```