package database

import (
	"database/sql"
	"log"
)
import _ "github.com/lib/pq"

func GetDBConnection() *sql.DB {
	postgres, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/railway?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	return postgres
}
