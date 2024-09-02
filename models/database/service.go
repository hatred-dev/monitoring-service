package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ProjectID   primitive.ObjectID `bson:"project_id,omitempty" json:"-"`
	ServiceName string             `bson:"service_name" json:"service_name" binding:"required"`
	Url         string             `bson:"url" json:"url" binding:"required"`
	Active      bool               `bson:"active" json:"active"`
}
