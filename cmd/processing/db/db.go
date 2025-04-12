package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := "postgres://user:password@processing-postgres:5432/database?sslmode=disable"
	//connStr := "postgres://user:password@localhost:9002/database?sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot ping database: ", err)
	}

	queryBytes, err := ioutil.ReadFile("./configs/data.sql")
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	getBovineCountQuery := string(queryBytes)

	_, err = DB.Exec(getBovineCountQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to database.")
}
