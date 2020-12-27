package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Connection struct {
	Conn *sql.DB
}

func New() (*Connection, error) {
	db, err := sql.Open("postgres", "host=localhost password=cumcum123 user=userpg port=5432 dbname=postgres sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &Connection{Conn: db}, nil

}