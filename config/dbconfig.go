package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var database *sql.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "postgres"
	password := "admin"
	dbName := "ratings"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	database, e = sql.Open("postgres", dsn)

	if e != nil {
		panic(e)
	}

	// Ping the database to verify the connection
	if e = database.Ping(); e != nil {
		panic(e)
	}
}

func DB() *sql.DB {
	return database
}
