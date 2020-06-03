package main

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
