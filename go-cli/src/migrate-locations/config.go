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

	pgPort     string
	pgUser     string
	pgPassword string
	pgDB       string
	pgHost     string
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

	config.pgPort = os.Getenv("POSTGRES_PORT")
	config.pgUser = os.Getenv("POSTGRES_USER")
	config.pgPassword = os.Getenv("POSTGRES_PASSWORD")
	config.pgDB = os.Getenv("POSTGRES_DB")
	config.pgHost = os.Getenv("POSTGRES_HOST")

	envVariables := []string{
		config.getCountriesHost,
		config.getLocationHost,
		config.getLocationAPIToken,
		config.getLocationUserEmail,
		config.getLocationAccessTokenPath,
		config.getLocationCountriesPath,
		config.getLocationStatesPath,
		config.getLocationCitiesPath,
		config.pgPort,
		config.pgUser,
		config.pgPassword,
		config.pgDB,
		config.pgHost,
	}

	for i, v := range envVariables {
		if v == "" {
			fmt.Printf("Error: %d %v shouldn't be empty\n", i, v)
			os.Exit(1)
		}
	}
}
