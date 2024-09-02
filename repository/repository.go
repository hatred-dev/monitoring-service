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
	ServiceRepository *serviceRepository
	IpRepository      *ipRepository
)

type projectRepository struct {
	*mongo.Collection
	Ctx context.Context
}

type serviceRepository struct {
	*mongo.Collection
	Ctx context.Context
}

type ipRepository struct {
	*mongo.Collection
	Ctx context.Context
}

func setupIndexes() {
	idxOptions := options.Index().SetUnique(true)
	_, err := ProjectRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "project_name", Value: 1},
		},
		Options: idxOptions,
	})
	if err != nil {
		fmt.Println(err)
	}
	_, err = ServiceRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "project_id", Value: 1},
			{Key: "url", Value: 1},
		},
		Options: idxOptions,
	})
	if err != nil {
		fmt.Println(err)
	}
	_, err = IpRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "project_id", Value: 1},
			{Key: "ip", Value: 1},
		},
		Options: idxOptions,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func InitRepository() {
	ProjectRepository = &projectRepository{
		DB.Collection("projects"),
		context.Background(),
	}
	ServiceRepository = &serviceRepository{
		DB.Collection("services"),
		context.Background(),
	}
	IpRepository = &ipRepository{
		DB.Collection("ips"),
		context.Background(),
	}
	setupIndexes()
}
