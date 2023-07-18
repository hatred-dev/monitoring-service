package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB                *mongo.Database
	ProjectRepository *projectRepository
)

type projectRepository struct {
	*mongo.Collection
}

func InitRepository() {
	ProjectRepository = &projectRepository{
		DB.Collection("projects"),
	}
}

func setupIndexes() {
	idxOptions := options.Index().SetUnique(true)
	_, err := ProjectRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{"project_name", 1},
			{"ips.ip", 1},
			{"services.url", 1},
		},
		Options: idxOptions,
	})
	if err != nil {
		fmt.Println(err)
	}
}
