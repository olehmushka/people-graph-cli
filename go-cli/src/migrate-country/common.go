package main

import "time"

const (
	appJSON = "application/json"

	dbDelay   = 500 * time.Millisecond
	chunkSize = 50
)

type getLocationAuthResp struct {
	AuthToken string `json:"auth_token"`
}

type getLocationState struct {
	Name string `json:"state_name"`
}

type getLocationCity struct {
	Name string `json:"city_name"`
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
