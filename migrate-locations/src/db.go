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

func (db *DBClient) AddCountries(countries *[]Country) error {
	const chunkSize = 25
	countryBuff := make([]Country, chunkSize)
	var counter = 0
	for _, country := range *countries {
		if counter == chunkSize {
			counter = 0
			time.Sleep(1 * time.Second)
			if err := db.insertCountries(&countryBuff); err != nil {
				return err
			}
		}
		countryBuff[counter] = country
		counter++
	}

	return nil
}

func (db *DBClient) insertCountries(countries *[]Country) error {
	conn, err := pgx.Connect(context.Background(), db.url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	var insertLines string

	for i, country := range *countries {
		insertLines = insertLines + fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", country.ID, country.Name, country.AlphaTwoCode, country.AlphaThreeCode, country.PhoneCodes)
		if i != (len(*countries) - 1) {
			insertLines = insertLines + ", "
		}
	}
	query := fmt.Sprintf("INSERT INTO countries (id, name, alpha_two_code, alpha_three_code, phone_codes) VALUES %s;", insertLines)

	if _, err := conn.Query(context.Background(), query); err != nil {
		return err
	}
	return nil
}
