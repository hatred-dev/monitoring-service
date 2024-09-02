package main

import (
	"context"
	"fmt"
	"monitoring-service/configuration"
	"monitoring-service/logger"
	"monitoring-service/repository"
	"monitoring-service/router"
	"monitoring-service/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	configuration.InitConfigurations()
	logger.InitLogger()
	go services.TimerService.Watch()
}

func main() {
	mongoClient, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(configuration.GetMongoConfig().DatabaseUrl()),
	)
	if err != nil {
		logger.Log.Fatal(err)
	}
	repository.DB = mongoClient.Database("monitoring")
	repository.InitRepository()
	services.StartServices()
	err = router.CreateRouter().Run(":8000")
	if err != nil {
		fmt.Println(err)
	}
}
