package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type getLocationClient struct {
	baseURL string
	client  httpClient
	bearer  string
}

func newGetLocationClient(client httpClient) *getLocationClient {
	return &getLocationClient{
		baseURL: config.getLocationHost,
		client:  client,
	}
}

func (gl *getLocationClient) GetStates(countryName string, states *[]getLocationState) error {
	fmt.Printf("Query country === countryName: %v\n", countryName)
	resp, err := gl.request(http.MethodGet, config.getLocationStatesPath+"/"+countryName, nil, appJSON)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(string(b))
	}

	return json.Unmarshal(b, states)
}

func (gl *getLocationClient) GetCities(stateName string, states *[]getLocationCity) error {
	resp, err := gl.request(http.MethodGet, config.getLocationCitiesPath+"/"+stateName, nil, appJSON)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(string(b))
	}

	return json.Unmarshal(b, states)
}

func (gl *getLocationClient) getToken() (string, error) {
	var body io.Reader
	url := fmt.Sprintf("%s%s", gl.baseURL, config.getLocationAccessTokenPath)
	req, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", appJSON)
	req.Header.Set("api-token", config.getLocationAPIToken)
	req.Header.Set("user-email", config.getLocationUserEmail)

	resp, err := gl.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(br))
	}

	var authToken getLocationAuthResp
	if err := json.Unmarshal(br, &authToken); err != nil {
		return "", err
	}

	return authToken.AuthToken, nil
}

func (gl *getLocationClient) request(method, path string, body io.Reader, contentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", gl.baseURL, path)

	return gl.requestByURL(method, url, body, contentType)
}

func (gl *getLocationClient) requestByURL(method, url string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if gl.bearer == "" {
		gl.bearer, err = gl.getToken()
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gl.bearer))
	req.Header.Set("Content-Type", contentType)

	return gl.client.Do(req)
}

type getCountriesClient struct {
	baseURL string
	client  httpClient
}

func newGetCountriesClient(client httpClient) *getCountriesClient {
	return &getCountriesClient{
		baseURL: config.getCountriesHost,
		client:  client,
	}
}

func (gc *getCountriesClient) request(method, path string, body io.Reader, contentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", gc.baseURL, path)

	return gc.requestByURL(method, url, body, contentType)
}

func (gc *getCountriesClient) requestByURL(method, url string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return gc.client.Do(req)
}
