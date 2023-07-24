package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"monitoring-service/models/database"
)

func (p *projectRepository) GetProjects() []database.Project {
	var projects []database.Project
	pipeline := mongo.Pipeline{
		{{"$lookup", bson.M{
			"from":         "ips",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "ips",
		}}},
		{{"$lookup", bson.M{
			"from":         "services",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "services",
		}}},
	}
	cursor, err := p.Aggregate(context.Background(), pipeline, nil)
	if err != nil {
		return nil
	}
	if err != cursor.All(context.Background(), &projects) {
		return nil
	}
	return projects
}

func (p *projectRepository) ProjectExists(projectName string) (database.Project, error) {
	var project database.Project
	filter := bson.M{"project_name": projectName}
	err := p.FindOne(context.Background(), filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return database.Project{}, errors.New("project not found")
		}
		return database.Project{}, err
	}
	return project, nil
}

func (p *projectRepository) GetProjectByName(projectName string) (database.Project, error) {
	var project database.Project
	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"project_name": projectName}}},
		{{"$lookup", bson.M{
			"from":         "ips",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "ips",
		}}},
		{{"$lookup", bson.M{
			"from":         "services",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "services",
		}}},
	}
	cur, err := p.Aggregate(context.Background(), pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return database.Project{}, errors.New("project not found")
		}
	}
	if cur.Next(context.Background()) {
		err = cur.Decode(&project)
		if err != nil {
			return database.Project{}, err
		}
	}
	return project, nil
}

func (p *projectRepository) CreateProject(project database.Project) (primitive.ObjectID, error) {
	res, err := p.InsertOne(context.Background(), project)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return primitive.NilObjectID, errors.New("project already exists")
		}
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
