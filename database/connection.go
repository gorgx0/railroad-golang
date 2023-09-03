package database

import (
	"database/sql"
	"fmt"
	"railway/config"
)
import _ "github.com/lib/pq"

func GetDBConnection(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	dbString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Db)
	postgres, err := sql.Open("postgres", dbString)
	//postgres, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/railway?sslmode=disable")
	return postgres, err
}
