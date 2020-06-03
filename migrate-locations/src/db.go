package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type DBClient struct {
	conn *pgx.Conn
}

func ComposeURL(port, user, password, db, host string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, db)
}

func NewDBClient(url string) (*DBClient, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &DBClient{
		conn: conn,
	}, nil
}

func (db *DBClient) Close() {
	db.conn.Close(context.Background())

}
