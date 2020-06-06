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
