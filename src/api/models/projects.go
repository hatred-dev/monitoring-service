package models

type GetProjectByName struct {
	ProjectName string `json:"project_name" binding:"required"`
}

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
	ProjectName    string `json:"project_name" binding:"required,project-exists"`
	NewProjectName string `json:"new_project_name" binding:"required"`
	Active         bool   `json:"active"`
}

type DeleteProjectRequest struct {
	ProjectName string `json:"project_name" binding:"required,project-exists"`
}

type FullProjectInfo struct {
	ProjectName string
	Active      bool
	Ips         []struct {
		Ip     string
		Active bool
	}
	Services []struct {
		ServiceName string
		Url         string
		Active      bool
	}
}
