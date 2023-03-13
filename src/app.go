package src

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"monitoring-service/database"
	"monitoring-service/src/api/router"
	"monitoring-service/src/configuration"
)

func init() {
	configuration.InitConfigurations()
}

func StartApplication() {
	db, err := sql.Open("postgres", configuration.DBConfig.DatabaseUrl())
	if err != nil {
		log.Fatal(err)
	}
	database.Conn = database.New(db)
	router.CreateRouter().Run(":8000")
}
