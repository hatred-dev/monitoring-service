package handlers

import (
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetProjects(ctx *gin.Context) {
	projects := repository.ProjectRepository.GetProjects()
	ctx.JSON(http.StatusOK, projects)
}

func HandleGetProjectByName(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	ctx.JSON(http.StatusOK, project)
}

func HandleCreateProject(ctx *gin.Context) {
	var project database.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	projectCreated, err := repository.ProjectRepository.CreateProject(project)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id": projectCreated,
	})
}

func HandlePatchProject(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	var newProject api.Project
	if err := ctx.ShouldBindJSON(&newProject); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := repository.ProjectRepository.UpdateProject(project, newProject)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func HandleDeleteProject(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	err := repository.ProjectRepository.DeleteProject(project)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
