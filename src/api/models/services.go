package models

type DeleteServiceReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
}

type CreateServiceReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	Url         string `json:"url" binding:"required"`
}

type CreateServiceResp struct {
	Id          string `json:"id"`
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}

type GetServicesResp struct {
	Id          string `json:"id" binding:"required"`
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name" binding:"required"`
	Url         string `json:"url" binding:"required"`
}

type GetServicesByProjectNameResp struct {
	Id          string `json:"id"`
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}
