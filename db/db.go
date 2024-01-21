package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connection := fmt.Sprintf("user=postgres dbname=postgres password=%s host=localhost port=5433 sslmode=disable", "postgres")

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
