package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"monitoring-service/configuration"
	"monitoring-service/logger"
	"monitoring-service/repository"
	"monitoring-service/router"
	"monitoring-service/services"
)

func init() {
	configuration.InitConfigurations()
	logger.InitLogger()
	go services.TimerService.StartTimer()
}

func main() {
	mongoClient, err := mongo.Connect(nil, options.Client().ApplyURI(configuration.MDBConfig.DatabaseUrl()))
	if err != nil {
		logger.Log.Fatal(err)
	}
	repository.DB = mongoClient.Database("monitoring")
	repository.InitRepository()
	services.StartServices()
	err = router.CreateRouter().Start(":8000")
	if err != nil {
		fmt.Println(err)
	}
}
