package models

type DeleteServiceRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
}

type CreateServiceRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	Url         string `json:"url" binding:"required"`
}

type CreateServiceResponse struct {
	Id          string `json:"id"`
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}

type GetServicesResponse struct {
	Id          string `json:"id" binding:"required"`
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name" binding:"required"`
	Url         string `json:"url" binding:"required"`
}

type GetServicesByProjectNameResponse struct {
	Id          string `json:"id"`
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}
