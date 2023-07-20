package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ServiceName string             `bson:"service_name" json:"service_name" validate:"required"`
	Url         string             `bson:"url" json:"url" validate:"required"`
	Active      bool               `bson:"active" json:"active,default:true"`
}
