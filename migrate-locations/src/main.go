package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	initConfig()

	client := newGetLocationClient(&http.Client{})
	db, err := NewDBClient(ComposeURL(config.pgPort, config.pgUser, config.pgPassword, config.pgDB, config.pgHost))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

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
