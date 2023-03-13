package models

type CreateProjectRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
	Active      bool   `json:"active"`
}

type CreateProjectResponse struct {
	Id          string `json:"id"`
	ProjectName string `json:"projectName"`
	Active      bool   `json:"active"`
}

type CreateServiceRequest struct {
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}

type AddIpAddressRequest struct {
	ProjectName string `json:"project_name"`
	Ip          string `json:"ip"`
}
