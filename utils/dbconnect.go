package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "user=appuser dbname=crypto sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckErr("sql.Open", err)
	return db
}
