//cmd/weather.go 
package cmd

import (
	"fmt"
	"weather/config"
	"github.com/go-resty/resty/v2"
    "github.com/spf13/cobra"
	"os"
)

var location string
var format string

var rootCmd = &cobra.Command{
	Use: "weather",
	Short:"A CLI tool for fetching weather information",
	Long: "Weather is a command line tool that allows you to fetch weather information from various sources.",
	Run: func (cmd *cobra.Command, args []string)  {
		apiKey := config.GetAPIKey()
		if apiKey == "" {
			fmt.Println("Error: API key not found.")
			return	
		}

		if location == "" {
			fmt.Println("Please provide a location: ")
			return
		}

		client := resty.New()
		resp, err := client.R().
			SetQueryParams(map[string]string{"access_key": apiKey, "query", location}).
			Get("http://api.weatherstack.com/current")

		if err != nil {
			fmt.Println("Error in the application", err)
			return
		}

		fmt.Println("Response: ", resp.String())
	},
}


// fmt.Println("Your WeatherCli application welcomes you, here you can check the weather.")		

func Execute() {
	rootCmd.Flags().StringVarP(&location, "location", "l", "", "location to obtain the weather")
	// rootCmd.Flags().StringVarP(&format, "format", "f", "json", "Output format (json, cvs, plain)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)


	}
}