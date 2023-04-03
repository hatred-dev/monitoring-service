package services

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"monitoring-service/src/configuration"
	"monitoring-service/src/logger"
)

func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", configuration.DBConfig.DatabaseUrl(), driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		logger.Log.Error(err)
	}
}
