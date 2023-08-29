package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"monitoring-service/logger"
	"monitoring-service/models/database"
)

func (s *serviceRepository) GetServices(project *database.Project) []database.Service {
	var services []database.Service
	opts := options.Find().SetProjection(bson.M{"_id": 0, "service_name": 1, "active": 1})
	filter := bson.M{"project_id": project.ID}
	cur, err := s.Find(context.Background(), filter, opts)
	if err != nil {
		logger.Log.Warn(err)
	}
	err = cur.All(context.Background(), &services)
	if err != nil {
		logger.Log.Warn(err)
	}
	return services
}

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
		{"project_id", project.ID},
		{"service_name", serviceName},
	}
	_, err := s.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceRepository) GetServiceState(service *database.Service) bool {
	var serviceObj database.Service
	filter := bson.M{
		"_id": service.ID,
	}
	err := s.FindOne(context.Background(), filter).Decode(&serviceObj)
	if err != nil {
		return false
	}
	return serviceObj.Active
}

func (s *serviceRepository) SetServiceState(service *database.Service, state bool) {
	filter := bson.M{
		"_id": service.ID,
	}
	update := bson.M{"$set": bson.M{"active": state}}
	_, err := s.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}
