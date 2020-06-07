package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

type DBClient struct {
	url string
}

func ComposeURL(port, user, password, db, host string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, db)
}

func NewDBClient(url string) (*DBClient, error) {
	return &DBClient{
		url: url,
	}, nil
}

func (db *DBClient) AddStates(states *[]State) error {
	stateBuff := make([]State, chunkSize)
	var counter = 0
	for i, state := range *states {
		if counter == chunkSize {
			counter = 0
			time.Sleep(dbDelay)
			if err := db.insertStates(&stateBuff); err != nil {
				return err
			}
		}
		stateBuff[counter] = state
		if (i + 1) == len(*states) {
			time.Sleep(dbDelay)
			if err := db.insertStates(&stateBuff); err != nil {
				return err
			}
		}
		counter++
	}

	return nil
}

func (db *DBClient) insertStates(states *[]State) error {
	conn, err := pgx.Connect(context.Background(), db.url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	var insertLines string

	for i, state := range *states {
		insertLines = insertLines + fmt.Sprintf("('%s', '%s', '%s')", state.ID, state.Name, state.CountryID)
		if i != (len(*states) - 1) {
			insertLines = insertLines + ", "
		}
	}
	query := fmt.Sprintf("INSERT INTO states (id, name, country_id) VALUES %s;", insertLines)

	if _, err := conn.Query(context.Background(), query); err != nil {
		return err
	}
	return nil
}

func (db *DBClient) AddCities(cities *[]City) error {
	cityBuff := make([]City, chunkSize)
	var counter = 0
	for i, city := range *cities {
		if counter == chunkSize {
			counter = 0
			time.Sleep(dbDelay)
			if err := db.insertCities(&cityBuff); err != nil {
				return err
			}
		}
		cityBuff[counter] = city
		if (i + 1) == len(*cities) {
			time.Sleep(dbDelay)
			if err := db.insertCities(&cityBuff); err != nil {
				return err
			}
		}
		counter++
	}

	return nil
}

func (db *DBClient) insertCities(cities *[]City) error {
	conn, err := pgx.Connect(context.Background(), db.url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	var insertLines string

	for i, city := range *cities {
		insertLines = insertLines + fmt.Sprintf("('%s', '%s', '%s')", city.ID, city.Name, city.StateID)
		if i != (len(*cities) - 1) {
			insertLines = insertLines + ", "
		}
	}
	query := fmt.Sprintf("INSERT INTO cities (id, name, state_id) VALUES %s;", insertLines)

	if _, err := conn.Query(context.Background(), query); err != nil {
		return err
	}
	return nil
}

func (db *DBClient) GetCountries(countryAlphaTwoCode string, countries *[]Country) error {
	conn, err := pgx.Connect(context.Background(), db.url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := fmt.Sprintf("SELECT * FROM countries WHERE alpha_two_code LIKE '%s';", countryAlphaTwoCode)

	result := make([]Country, 0)
	rows, _ := conn.Query(context.Background(), query)

	for rows.Next() {
		var id, name, alphaTwoCode, alphaThreeCode, phoneCodes string
		err := rows.Scan(&id, &name, &alphaTwoCode, &alphaThreeCode, &phoneCodes)
		if err != nil {
			return err
		}
		result = append(result, Country{id, name, alphaTwoCode, alphaThreeCode, phoneCodes})
	}
	*countries = result

	return rows.Err()
}
