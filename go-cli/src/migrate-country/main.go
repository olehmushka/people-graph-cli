package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	initConfig()
	start := time.Now()
	countryName := config.countryName
	countryAlphaTwoCode := config.countryCode

	getLocationClient := newGetLocationClient(&http.Client{})
	db, err := NewDBClient(ComposeURL(config.pgPort, config.pgUser, config.pgPassword, config.pgDB, config.pgHost))
	if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		os.Exit(1)
	}

	var countries []Country
	if err := db.GetCountries(countryAlphaTwoCode, &countries); err != nil {
		fmt.Printf("Query country(country name: %s) Error: %v\n", countryName, err)
		os.Exit(1)
	}
	if len(countries) < 1 {
		fmt.Printf("Query country(country name: %s) Error: %v\n", countryName, "countries weren't found")
		os.Exit(1)
	}

	var basicStates []getLocationState
	if err := getLocationClient.GetStates(prepareUrlString(countryName), &basicStates); err != nil {
		fmt.Printf("GET states(country name: %s) Error: %v\n", countryAlphaTwoCode, err)
		os.Exit(1)
	}
	fmt.Printf("Query country(country name: %s) basicStates: %v\n", countryName, basicStates)
	states := getLocationStatesMapper(countries[0].ID, &basicStates)
	if err := db.AddStates(states); err != nil {
		fmt.Printf("Add states(country name: %s) Error: %v\n", countryName, err)
		os.Exit(1)
	}

	for _, state := range *states {
		var basicCities []getLocationCity
		if err := getLocationClient.GetCities(prepareUrlString(state.Name), &basicCities); err != nil {
			fmt.Printf("GET cities(country name: %s, state name: %s) Error: %v\n", countryName, state.Name, err)
			os.Exit(1)
		}
		cities := getLocationCitiesMapper(state.ID, &basicCities)
		if err := db.AddCities(cities); err != nil {
			fmt.Printf("Add cities(country name: %s, state name: %s) Error: %v\n", countryName, state.Name, err)
			os.Exit(1)
		}
	}
	end := time.Now()
	fmt.Printf("Finished after %v\n", end.Sub(start))
}
