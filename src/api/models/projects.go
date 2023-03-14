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

type PatchProjectRequest struct {
	ProjectName    string `json:"project_name" binding:"required"`
	NewProjectName string `json:"new_project_name" binding:"required"`
	Active         bool   `json:"active"`
}

type DeleteProjectRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
}
