package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"monitoring-service/database"
	"monitoring-service/src/api/models"
	"net/http"
)

func HandleGetProjects(ctx *gin.Context) {
	projects, err := database.Conn.GetProjects(ctx)
	if err != nil {
		errorSender(ctx, err)
		return
	}
	jsonProjects := make([]models.CreateProjectResponse, len(projects))
	for i, v := range projects {
		jsonProjects[i] = models.CreateProjectResponse{
			Id:          v.ID.String(),
			ProjectName: v.ProjectName,
			Active:      v.Active.Bool,
		}
	}
	ctx.JSON(http.StatusOK, jsonProjects)
}

func HandleCreateProject(ctx *gin.Context) {
	var project *models.CreateProjectRequest
	if err := ctx.ShouldBind(&project); err != nil {
		fmt.Println()
		errorSender(ctx, err)
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
		errorSender(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, models.CreateProjectResponse{
		Id:          projectCreated.ID.String(),
		ProjectName: project.ProjectName,
		Active:      project.Active,
	})
}

func errorSender(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
