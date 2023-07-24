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

func (i *ipRepository) CreateIp(project database.Project, ip database.Ip) (primitive.ObjectID, error) {
	ip.ProjectID = project.ID
	res, err := i.InsertOne(context.Background(), ip)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return primitive.NilObjectID, errors.New("ip already exists")
		}
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (i *ipRepository) GetIpState(ip *database.Ip) bool {
	var ipObj database.Ip
	opts := options.FindOne().SetProjection(bson.M{"active": 1})
	filter := bson.M{"_id": ip.ID}
	err := i.FindOne(context.Background(), filter, opts).Decode(&ipObj)
	if err != nil {
		return false
	}
	return ipObj.Active
}

func (i *ipRepository) UpdateIp(project database.Project, oldIp, newIp string) error {
	filter := bson.D{
		{"project_id", project.ID},
		{"ip", oldIp},
	}
	update := bson.M{"$set": bson.M{"ip": newIp}}
	res, err := i.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("ip not found")
	}
	return nil
}

func (i *ipRepository) DeleteIp(project database.Project, ip string) error {
	filter := bson.D{
		{"project_id", project.ID},
		{"ip", ip},
	}
	res, err := i.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("ip not found")
	}
	return nil
}

func (i *ipRepository) SetIpState(ip string, state bool) {
	filter := bson.M{"ips.ip": ip}
	update := bson.M{"$set": bson.M{"ips.$.active": state}}
	_, err := i.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}
