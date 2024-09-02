package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ip struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ProjectID primitive.ObjectID `bson:"project_id" json:"-"`
	Ip        string             `bson:"ip" json:"ip" binding:"required"`
	Active    bool               `bson:"active" json:"active"`
}
