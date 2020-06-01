package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type configuration struct {
	getCountriesHost string
	getLocationHost  string

	getLocationAPIToken  string
	getLocationUserEmail string

	getLocationAccessTokenPath string
	getLocationCountriesPath   string
	getLocationStatesPath      string
	getLocationCitiesPath      string
}

var config *configuration

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	config = &configuration{}
	config.getCountriesHost = os.Getenv("GET_COUNTRIES_HOST")
	config.getLocationHost = os.Getenv("GET_LOCATION_HOST")

	config.getLocationAPIToken = os.Getenv("GET_LOCATION_API_TOKEN")
	config.getLocationUserEmail = os.Getenv("GET_LOCATION_USER_EMAIL")

	config.getLocationAccessTokenPath = os.Getenv("GET_LOCATION_GET_ACCESS_TOKEN_PATH")
	config.getLocationCountriesPath = os.Getenv("GET_LOCATION_GET_COUNTRIES_PATH")
	config.getLocationStatesPath = os.Getenv("GET_LOCATION_GET_STATES_PATH")
	config.getLocationCitiesPath = os.Getenv("GET_LOCATION_GET_CITIES_PATH")
}

const (
	appJSON = "application/json"
)

type getLocationAuthResp struct {
	AuthToken string `json:"auth_token"`
}

type getLocationCountry struct {
	Name       string `json:"country_name"`
	Alpha2Code string `json:"country_short_name"`
	PhoneCode  int    `json:"country_phone_code"`
}
