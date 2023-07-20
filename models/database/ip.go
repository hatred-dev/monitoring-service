package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ip struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Ip     string             `bson:"ip" json:"ip" validate:"required"`
	Active bool               `bson:"active" json:"active"`
}
