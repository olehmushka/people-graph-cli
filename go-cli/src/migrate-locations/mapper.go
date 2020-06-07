package main

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func removeEscape(text string) string {
	return regexp.MustCompile(`\'`).ReplaceAllString(text, "")
}

func getCountriesMapper(countries *[]getCountryResponse) *[]Country {
	mapped := make([]Country, len(*countries))

	for i, country := range *countries {
		mapped[i] = Country{
			ID:             uuid.New().String(),
			Name:           removeEscape(country.Name),
			AlphaTwoCode:   country.Alpha2Code,
			AlphaThreeCode: country.Alpha3Code,
			PhoneCodes:     strings.Join(country.CallingCodes[:], ","),
		}
	}

	return &mapped
}

func getLocationStatesMapper(countryID string, states *[]getLocationState) *[]State {
	mapped := make([]State, len(*states))

	for i, state := range *states {
		mapped[i] = State{
			ID:        uuid.New().String(),
			Name:      removeEscape(state.Name),
			CountryID: countryID,
		}
	}

	return &mapped
}

func getLocationCitiesMapper(stateID string, cities *[]getLocationCity) *[]City {
	mapped := make([]City, len(*cities))

	for i, city := range *cities {
		mapped[i] = City{
			ID:      uuid.New().String(),
			Name:    removeEscape(city.Name),
			StateID: stateID,
		}
	}

	return &mapped
}
