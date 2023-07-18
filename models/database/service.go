package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ServiceName string             `bson:"service_name" json:"service_name"`
	Url         string             `bson:"url" json:"url"`
	Active      bool               `bson:"active" json:"active"`
}
