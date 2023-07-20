package api

type DeleteServiceReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
}
