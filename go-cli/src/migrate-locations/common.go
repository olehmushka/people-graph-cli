package main

import "time"

const (
	appJSON = "application/json"

	dbDelay   = 1 * time.Second
	chunkSize = 10
)

type getLocationAuthResp struct {
	AuthToken string `json:"auth_token"`
}

type getLocationCountry struct {
	Name       string `json:"country_name"`
	Alpha2Code string `json:"country_short_name"`
	PhoneCode  int    `json:"country_phone_code"`
}

type getLocationState struct {
	Name string `json:"state_name"`
}

type getLocationCity struct {
	Name string `json:"city_name"`
}

type getCountryCurrencyResponse struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type getCountryLanguageResponse struct {
	ISO6391    string `json:"iso639_1"`
	ISO6392    string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type getCountryRegionalBlocResponse struct {
	Acronym       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}

type getCountryResponse struct {
	Name           string                            `json:"name"`
	TopLevelDomain []string                          `json:"topLevelDomain"`
	Alpha2Code     string                            `json:"alpha2Code"`
	Alpha3Code     string                            `json:"alpha3Code"`
	CallingCodes   []string                          `json:"callingCodes"`
	Capital        string                            `json:"capital"`
	AltSpellings   []string                          `json:"altSpellings"`
	Region         string                            `json:"region"`
	Subregion      string                            `json:"subregion"`
	Population     int64                             `json:"population"`
	LatLng         []float64                         `json:"latlng"`
	Demonym        string                            `json:"demonym"`
	Area           float64                           `json:"area"`
	Gini           float64                           `json:"gini"`
	Timezones      []string                          `json:"timezones"`
	Borders        []string                          `json:"borders"`
	NativeName     string                            `json:"nativeName"`
	NumericCode    string                            `json:"numericCode"`
	Currencies     *[]getCountryCurrencyResponse     `json:"currencies"`
	Languages      *[]getCountryLanguageResponse     `json:"languages"`
	Translations   map[string]string                 `json:"translations"`
	Flag           string                            `json:"flag"`
	RegionalBlocs  *[]getCountryRegionalBlocResponse `json:"regionalBlocs"`
	Cioc           string                            `json:"cioc"`
}

type Country struct {
	ID             string
	Name           string
	AlphaTwoCode   string
	AlphaThreeCode string
	PhoneCodes     string
}

type State struct {
	ID        string
	Name      string
	CountryID string
}

type City struct {
	ID      string
	Name    string
	StateID string
}
