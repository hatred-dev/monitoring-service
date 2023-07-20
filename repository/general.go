package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"monitoring-service/logger"
	"monitoring-service/models/database"
)

func (p *projectRepository) GetProjects() []database.Project {
	var projects []database.Project
	cursor, err := p.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		return nil
	}
	if err != cursor.All(context.Background(), &projects) {
		return nil
	}
	return projects
}

func (p *projectRepository) GetProjectByName(projectName string) database.Project {
	var project database.Project
	filter := bson.M{"project_name": projectName}
	err := p.FindOne(context.Background(), filter).Decode(&project)
	if err != nil {
		return database.Project{}
	}
	return project
}

func (p *projectRepository) CreateProject(project database.Project) (primitive.ObjectID, error) {
	res, err := p.InsertOne(context.Background(), project)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (p *projectRepository) UpdateProject(projectName string, project database.Project) error {
	filter := bson.M{"project_name": projectName}
	update := bson.M{"$set": project}
	_, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) DeleteProject(projectName string) error {
	filter := bson.M{"project_name": projectName}
	_, err := p.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) GetIpState(ip string) bool {
	var ipObj database.Ip
	filter := bson.M{"ips.ip": ip}
	err := p.FindOne(context.Background(), filter).Decode(&ipObj)
	if err != nil {
		return false
	}
	return ipObj.Active
}

func (p *projectRepository) CreateIp(projectName string, ip database.Ip) error {
	filter := bson.M{"project_name": projectName}
	update := bson.M{"$addToSet": bson.M{"ips": ip}}
	res, err := p.UpdateOne(context.Background(), filter, update)
	if res.MatchedCount == 0 {
		return errors.New("project not found")
	}
	if res.ModifiedCount == 0 {
		return errors.New("ip already exists")
	}
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) UpdateIp(oldIp, newIp string) error {
	filter := bson.M{"ips.ip": oldIp}
	update := bson.M{"$set": bson.M{"ips.$": newIp}}
	_, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) DeleteIp(ip string) error {
	filter := bson.M{"ips.ip": ip}
	update := bson.M{"$pull": bson.M{"ips": bson.M{"ip": ip}}}
	_, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) SetIpState(ip string, state bool) {
	filter := bson.M{"ips.ip": ip}
	update := bson.M{"$set": bson.M{"ips.$.active": state}}
	_, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}

func (p *projectRepository) CreateService(projectName string, service database.Service) (primitive.ObjectID, error) {
	filter := bson.M{"project_name": projectName}
	update := bson.M{"$push": bson.M{"services": service}}
	res, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.UpsertedID.(primitive.ObjectID), nil
}

func (p *projectRepository) UpdateService(projectName, serviceName string, service database.Service) error {
	filter := bson.D{
		{"project_name", projectName},
		{"services.service_name", serviceName},
	}
	update := bson.M{"$set": bson.M{"services.$": service}}
	_, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) DeleteService(projectName, serviceName string) error {
	filter := bson.D{
		{"project_name", projectName},
		{"services.service_name", serviceName},
	}
	_, err := p.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) GetServiceState(project, name string) bool {
	var serviceObj database.Service
	filter := bson.D{
		{"project_name", project},
		{"services.service_name", name},
	}
	err := p.FindOne(context.Background(), filter).Decode(&serviceObj)
	if err != nil {
		return false
	}
	return serviceObj.Active
}

func (p *projectRepository) SetServiceState(project, name string, state bool) {
	filter := bson.D{
		{"project_name", project},
		{"services.service_name", name},
	}
	update := bson.M{"$set": bson.M{"services.$.active": state}}
	_, err := p.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}
