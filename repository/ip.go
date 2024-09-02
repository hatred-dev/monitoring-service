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

func (i *ipRepository) GetIps(project *database.Project) []database.Ip {
	var ips []database.Ip
	opts := options.Find().SetProjection(bson.M{"_id": 0})
	filter := bson.M{"project_id": project.ID}
	cur, err := i.Find(i.Ctx, filter, opts)
	if err != nil {
		logger.Log.Warn(err)
	}
	err = cur.All(i.Ctx, &ips)
	if err != nil {
		logger.Log.Warn(err)
	}
	return ips
}

func (i *ipRepository) CreateIp(project database.Project, ip database.Ip) (primitive.ObjectID, error) {
	ip.ProjectID = project.ID
	res, err := i.InsertOne(i.Ctx, ip)
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
	err := i.FindOne(i.Ctx, filter, opts).Decode(&ipObj)
	if err != nil {
		return false
	}
	return ipObj.Active
}

func (i *ipRepository) UpdateIp(project database.Project, oldIp, newIp string) error {
	filter := bson.D{
		{Key: "project_id", Value: project.ID},
		{Key: "ip", Value: oldIp},
	}
	update := bson.M{"$set": bson.M{"ip": newIp}}
	res, err := i.UpdateOne(i.Ctx, filter, update)
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
		{Key: "project_id", Value: project.ID},
		{Key: "ip", Value: ip},
	}
	res, err := i.DeleteOne(i.Ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("ip not found")
	}
	return nil
}

func (i *ipRepository) SetIpState(ip *database.Ip, state bool) {
	filter := bson.M{
		"_id": ip.ID,
	}
	update := bson.M{"$set": bson.M{"active": state}}
	_, err := i.UpdateOne(i.Ctx, filter, update)
	if err != nil {
		logger.Log.Warn(err)
	}
}
