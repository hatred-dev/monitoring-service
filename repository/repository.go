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
}

type serviceRepository struct {
	*mongo.Collection
}

type ipRepository struct {
	*mongo.Collection
}

func setupIndexes() {
	idxOptions := options.Index().SetUnique(true)
	_, err := ProjectRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{"project_name", 1},
		},
		Options: idxOptions,
	})
	_, err = ServiceRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{"project_id", 1},
			{"url", 1},
		},
		Options: idxOptions,
	})
	_, err = IpRepository.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{"project_id", 1},
			{"ip", 1},
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
	}
	ServiceRepository = &serviceRepository{
		DB.Collection("services"),
	}
	IpRepository = &ipRepository{
		DB.Collection("ips"),
	}
	setupIndexes()
}
