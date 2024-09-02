package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ProjectName string             `bson:"project_name" json:"project_name" binding:"required"`
	Ips         []Ip               `bson:"ips,omitempty" json:"ips,omitempty"`
	Services    []Service          `bson:"services,omitempty" json:"services,omitempty"`
	Active      bool               `bson:"active" json:"active" binding:"required"`
}

func (p *Project) IpsEmpty() bool {
	return len(p.Ips) == 0
}

func (p *Project) ServicesEmpty() bool {
	return len(p.Services) == 0
}
