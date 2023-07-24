package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"monitoring-service/logger"
	"monitoring-service/models/database"
)

func (s *serviceRepository) CreateService(project database.Project, service database.Service) (primitive.ObjectID, error) {
	service.ProjectID = project.ID
	res, err := s.InsertOne(context.Background(), service)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return primitive.NilObjectID, errors.New("service already exists")
		}
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (s *serviceRepository) UpdateService(project database.Project, serviceName string, service database.Service) error {
	filter := bson.D{
		{"project_id", project.ID},
		{"service_name", serviceName},
	}
	update := bson.M{"$set": service}
	res, err := s.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("service not found")
	}
	return nil
}

func (s *serviceRepository) DeleteService(project database.Project, serviceName string) error {
	filter := bson.D{
		{"project_name", project.ProjectName},
		{"services.service_name", serviceName},
	}
	_, err := s.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceRepository) GetServiceState(project, name string) bool {
	var serviceObj database.Service
	filter := bson.D{
		{"project_name", project},
		{"services.service_name", name},
	}
	err := s.FindOne(context.Background(), filter).Decode(&serviceObj)
	if err != nil {
		return false
	}
	return serviceObj.Active
}

func (s *serviceRepository) SetServiceState(project, name string, state bool) {
	filter := bson.D{
		{"project_name", project},
		{"services.service_name", name},
	}
	update := bson.M{"$set": bson.M{"services.$.active": state}}
	_, err := s.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}
