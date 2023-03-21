package src

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"monitoring-service/database"
	"monitoring-service/src/api/router"
	"monitoring-service/src/configuration"
	"monitoring-service/src/services"
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
	services.StartServices()
	err = router.CreateRouter().Run(":8000")
	if err != nil {
		fmt.Println(err)
	}
}
