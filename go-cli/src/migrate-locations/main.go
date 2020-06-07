package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	initConfig()
	start := time.Now()

	getCountriesClient := newGetCountriesClient(&http.Client{})
	db, err := NewDBClient(ComposeURL(config.pgPort, config.pgUser, config.pgPassword, config.pgDB, config.pgHost))
	if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		os.Exit(1)
	}

	// 1 Step: Insert Countries
	var advancedCountries []getCountryResponse
	if err := getCountriesClient.GetCountries(&advancedCountries); err != nil {
		fmt.Printf("GET countries Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Countries count = %d\n", len(advancedCountries))
	countries := getCountriesMapper(&advancedCountries)
	if err := db.AddCountries(countries); err != nil {
		fmt.Printf("Add countries Error: %v\n", err)
		os.Exit(1)
	}

	// 2 Step: Insert States
	getLocationClient := newGetLocationClient(&http.Client{})
	var basicCountries []getLocationCountry
	if err := getLocationClient.GetCountries(&basicCountries); err != nil {
		fmt.Printf("GET basic countries Error: %v\n", err)
		os.Exit(1)
	}
	var country *getLocationCountry
	for _, ac := range *countries {
		for _, bc := range basicCountries {
			if strings.Compare(bc.Alpha2Code, ac.AlphaTwoCode) == 0 {
				country = &bc
				break
			}
		}
		if country == nil {
			continue
		}
		startCountry := time.Now()

		var basicStates []getLocationState
		if err := getLocationClient.GetStates(url.QueryEscape(country.Name), &basicStates); err != nil {
			fmt.Printf("GET states(country name: %s) Error: %v\n", country.Name, err)
			os.Exit(1)
		}
		states := getLocationStatesMapper(ac.ID, &basicStates)
		if err := db.AddStates(states); err != nil {
			fmt.Printf("Add states(country name: %s) Error: %v\n", country.Name, err)
			os.Exit(1)
		}

		// 3 Step: Insert Cities
		for _, state := range *states {
			var basicCities []getLocationCity
			if err := getLocationClient.GetCities(url.QueryEscape(state.Name), &basicCities); err != nil {
				fmt.Printf("GET cities(country name: %s, state name: %s) Error: %v\n", country.Name, state.Name, err)
				os.Exit(1)
			}
			cities := getLocationCitiesMapper(state.ID, &basicCities)
			if err := db.AddCities(cities); err != nil {
				fmt.Printf("Add cities(country name: %s, state name: %s) Error: %v\n", country.Name, state.Name, err)
				os.Exit(1)
			}
		}
		endCountry := time.Now()
		fmt.Printf("Country(%s) is finished after %v\n", ac.Name, endCountry.Sub(startCountry))
		country = nil
	}
	end := time.Now()
	fmt.Printf("Finished after %v\n", end.Sub(start))
}
