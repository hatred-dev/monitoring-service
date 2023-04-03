package src

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"monitoring-service/database"
	"monitoring-service/src/api/router"
	"monitoring-service/src/configuration"
	"monitoring-service/src/logger"
	"monitoring-service/src/services"
)

func init() {
	configuration.InitConfigurations()
	logger.InitLogger()
	go services.TimerService.StartTimer()
}

func StartApplication() {
	db, err := sql.Open("postgres", configuration.DBConfig.DatabaseUrl())
	if err != nil {
		log.Fatal(err)
	}
	database.Conn = database.New(db)
	services.RunMigrations(db)
	services.StartServices()
	err = router.CreateRouter().Run(":8000")
	if err != nil {
		fmt.Println(err)
	}
}
