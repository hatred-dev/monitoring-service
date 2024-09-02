package api

type Project struct {
	ProjectName string `bson:"project_name" json:"project_name" validate:"required"`
}
