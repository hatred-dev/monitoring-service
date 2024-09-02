package repository

import (
	"errors"
	"monitoring-service/models/api"
	"monitoring-service/models/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const Lookup = "$lookup"

func (p *projectRepository) GetProjects() []database.Project {
	var projects []database.Project
	pipeline := mongo.Pipeline{
		{{Key: Lookup, Value: bson.M{
			"from":         "ips",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "ips",
		}}},
		{{Key: Lookup, Value: bson.M{
			"from":         "services",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "services",
		}}},
	}
	cursor, err := p.Aggregate(p.Ctx, pipeline, nil)
	if err != nil {
		return nil
	}
	if err := cursor.All(p.Ctx, &projects); err != nil {
		return nil
	}
	return projects
}

func (p *projectRepository) ProjectExists(projectName string) (database.Project, error) {
	var project database.Project
	filter := bson.M{"project_name": projectName}
	err := p.FindOne(p.Ctx, filter).Decode(&project)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return database.Project{}, errors.New("project not found")
		}
		return database.Project{}, err
	}
	return project, nil
}

func (p *projectRepository) GetProjectByName(projectName string) (database.Project, error) {
	var project database.Project
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"project_name": projectName}}},
		{{Key: Lookup, Value: bson.M{
			"from":         "ips",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "ips",
		}}},
		{{Key: Lookup, Value: bson.M{
			"from":         "services",
			"localField":   "_id",
			"foreignField": "project_id",
			"as":           "services",
		}}},
	}
	cur, err := p.Aggregate(p.Ctx, pipeline)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return database.Project{}, errors.New("project not found")
		}
	}
	if cur.Next(p.Ctx) {
		err = cur.Decode(&project)
		if err != nil {
			return database.Project{}, err
		}
	}
	return project, nil
}

func (p *projectRepository) CreateProject(project database.Project) (primitive.ObjectID, error) {
	res, err := p.InsertOne(p.Ctx, project)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return primitive.NilObjectID, errors.New("project already exists")
		}
		return primitive.NilObjectID, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (p *projectRepository) UpdateProject(project database.Project, newProject api.Project) error {
	filter := bson.M{"_id": project.ID}
	update := bson.M{"$set": newProject}
	_, err := p.UpdateOne(p.Ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) DeleteProject(project database.Project) error {
	filter := bson.M{"_id": project.ID}
	_, err := p.DeleteOne(p.Ctx, filter)
	if err != nil {
		return err
	}
	filter = bson.M{
		"name": bson.M{
			"$ne": p.Collection.Name(),
		},
	}
	collections, err := DB.ListCollectionNames(p.Ctx, filter)
	if err != nil {
		return err
	}
	for _, collection := range collections {
		collection := DB.Collection(collection, nil)
		filter := bson.M{"project_id": project.ID}
		_, _ = collection.DeleteMany(p.Ctx, filter)
	}
	return nil
}
