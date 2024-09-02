package repository

import (
	"errors"
	"monitoring-service/logger"
	"monitoring-service/models/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *serviceRepository) GetServices(project *database.Project) []database.Service {
	var services []database.Service
	opts := options.Find().SetProjection(bson.M{"_id": 0})
	filter := bson.M{"project_id": project.ID}
	cur, err := s.Find(s.Ctx, filter, opts)
	if err != nil {
		logger.Log.Warn(err)
	}
	err = cur.All(s.Ctx, &services)
	if err != nil {
		logger.Log.Warn(err)
	}
	return services
}

func (s *serviceRepository) CreateService(project database.Project, service database.Service) (primitive.ObjectID, error) {
	service.ProjectID = project.ID
	res, err := s.InsertOne(s.Ctx, service)
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
		{Key: "project_id", Value: project.ID},
		{Key: "service_name", Value: serviceName},
	}
	update := bson.M{"$set": service}
	res, err := s.UpdateOne(s.Ctx, filter, update)
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
		{Key: "project_id", Value: project.ID},
		{Key: "service_name", Value: serviceName},
	}
	_, err := s.DeleteOne(s.Ctx, filter)
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
	err := s.FindOne(s.Ctx, filter).Decode(&serviceObj)
	if err != nil {
		logger.Log.Warn(err)
	}
	return serviceObj.Active
}

func (s *serviceRepository) SetServiceState(service *database.Service, state bool) {
	filter := bson.M{
		"_id": service.ID,
	}
	update := bson.M{"$set": bson.M{"active": state}}
	_, err := s.UpdateOne(s.Ctx, filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}
