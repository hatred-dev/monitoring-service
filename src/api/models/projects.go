package models

type CreateProjectReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	Active      bool   `json:"active"`
}

type CreateProjectResp struct {
	Id          string `json:"id"`
	ProjectName string `json:"project_name"`
	Active      bool   `json:"active"`
}

type PatchProjectReq struct {
	ProjectName    string `json:"project_name" binding:"required"`
	NewProjectName string `json:"new_project_name" binding:"required"`
	Active         bool   `json:"active"`
}

type DeleteProjectReq struct {
	ProjectName string `json:"project_name" binding:"required"`
}
