package main

import (
	"net/url"
	"regexp"

	"github.com/google/uuid"
)

func removeEscape(text string) string {
	return regexp.MustCompile(`\'`).ReplaceAllString(text, "")
}

func prepareUrlString(text string) string {
	return regexp.MustCompile(`\+`).ReplaceAllString(url.QueryEscape(text), " ")
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
