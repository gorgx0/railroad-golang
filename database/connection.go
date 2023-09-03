package database

import (
	"database/sql"
)
import _ "github.com/lib/pq"

func GetDBConnection() (*sql.DB, error) {
	postgres, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/railway?sslmode=disable")
	return postgres, err
}
