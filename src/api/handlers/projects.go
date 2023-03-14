package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"monitoring-service/database"
	"monitoring-service/src/api/helpers"
	"monitoring-service/src/api/models"
	"net/http"
)

func HandleGetProjects(ctx *gin.Context) {
	projects, err := database.Conn.GetProjects(ctx)
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

func HandleGetProjectByName(ctx *gin.Context) {
	projectName := ctx.Param("project_name")
	project, err := database.Conn.GetProjectByName(ctx, projectName)
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, project)
}

func HandleCreateProject(ctx *gin.Context) {
	var project *models.CreateProjectRequest
	if err := ctx.ShouldBind(&project); err != nil {
		helpers.Abort(ctx, err)
		return
	}

	projectCreated, err := database.Conn.CreateProject(ctx, database.CreateProjectParams{
		ProjectName: project.ProjectName,
		Active: sql.NullBool{
			Bool:  project.Active,
			Valid: true,
		},
	})
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, models.CreateProjectResponse{
		Id:          projectCreated.ID.String(),
		ProjectName: project.ProjectName,
		Active:      project.Active,
	})
}

func HandlePatchProject(ctx *gin.Context) {
	var newProject *models.PatchProjectRequest
	if err := ctx.ShouldBind(&newProject); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	if exists := helpers.CheckIfProjectExists(ctx, newProject.ProjectName); !exists {
		return
	}
	err := database.Conn.UpdateProject(ctx, database.UpdateProjectParams{
		ProjectName:   newProject.ProjectName,
		ProjectName_2: newProject.NewProjectName,
		Active: sql.NullBool{
			Bool:  newProject.Active,
			Valid: true,
		},
	})
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
}

func HandleDeleteProject(ctx *gin.Context) {
	var project *models.DeleteProjectRequest

	if err := ctx.ShouldBind(&project); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	if exists := helpers.CheckIfProjectExists(ctx, project.ProjectName); !exists {
		return
	}
	if err := database.Conn.DeleteProject(ctx, project.ProjectName); err != nil {
		helpers.Abort(ctx, err)
	}
}
