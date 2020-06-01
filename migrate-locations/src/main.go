package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	initConfig()
	// check env variables
	envVariables := []string{
		config.getCountriesHost,
		config.getLocationHost,
		config.getLocationAPIToken,
		config.getLocationUserEmail,
		config.getLocationAccessTokenPath,
		config.getLocationCountriesPath,
		config.getLocationStatesPath,
		config.getLocationCitiesPath,
	}

	for i, v := range envVariables {
		if v == "" {
			fmt.Printf("Error: %d %v shouldn't be empty\n", i, v)
			os.Exit(1)
		}
	}

	client := newGetLocationClient(&http.Client{})

	var countries []getLocationCountry
	if err := client.GetCountries(&countries); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	for i, country := range countries {
		if i > 10 {
			continue
		}
		fmt.Println(country.Name)
	}
}
