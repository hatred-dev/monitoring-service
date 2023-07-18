package api

type DeleteProjectReq struct {
	ProjectName string `json:"project_name" binding:"required"`
}
