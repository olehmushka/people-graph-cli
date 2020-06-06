package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	initConfig()

	getCountriesClient := newGetCountriesClient(&http.Client{})
	db, err := NewDBClient(ComposeURL(config.pgPort, config.pgUser, config.pgPassword, config.pgDB, config.pgHost))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	var countries2 []getCountryResponse
	if err := getCountriesClient.GetCountries(&countries2); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	if err := db.AddCountries(getCountriesMapper(&countries2)); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
