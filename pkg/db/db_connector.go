package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDbConnection() (*sql.DB, error) {
	connectionString := "user=postgres password=postgres dbname=terminal-jeopardy sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
