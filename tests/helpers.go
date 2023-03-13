package tests

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	testDB = "testdb"

	dbConnection *sql.DB
)

func createConnection(connInfo string) {
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}
	dbConnection = db
}

func dropConnection() {
	if dbConnection != nil {
		err := dbConnection.Close()
		if err != nil {
			panic(err)
		}
		dbConnection = nil
	}
}

func createDB() {
	createConnection("user=postgres password=postgres host=0.0.0.0 sslmode=disable")
	_, err := dbConnection.Exec("create database " + testDB)
	if err != nil {
		fmt.Println(err) //only signals that database already exists
	}
	dropConnection()
	connInfo := fmt.Sprintf("postgres://postgres:postgres@0.0.0.0:5432/%s?sslmode=disable", testDB)
	createConnection(connInfo)
	runMigrations()
}

func dropDB() {
	dropConnection()
	connInfo := fmt.Sprintf("postgres://postgres:postgres@0.0.0.0:5432/%s?sslmode=disable", "postgres")
	createConnection(connInfo)
	_, err := dbConnection.Exec(fmt.Sprintf("drop database %s with (force)", testDB))
	if err != nil {
		panic(err)
	}
}

func runMigrations() {
	driver, err := postgres.WithInstance(dbConnection, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://../migrations", testDB, driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		fmt.Println(err)
	}
}
