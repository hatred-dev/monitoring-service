package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
)

func HandleGetProjects(ctx echo.Context) error {
	projects := repository.ProjectRepository.GetProjects()
	err := ctx.JSON(http.StatusOK, projects)
	return err
}

func HandleGetProjectByName(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	err := ctx.JSON(http.StatusOK, project)
	return err
}

func HandleCreateProject(ctx echo.Context) error {
	var project database.Project
	if err := ctx.Bind(&project); err != nil {
		return err
	}
	err := ctx.Validate(&project)
	if err != nil {
		return err
	}
	projectCreated, err := repository.ProjectRepository.CreateProject(project)
	if err != nil {
		return err
	}
	err = ctx.JSON(http.StatusCreated, echo.Map{
		"id": projectCreated,
	})
	return err
}

func HandlePatchProject(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var newProject database.Project
	if err := ctx.Bind(&newProject); err != nil {
		return err
	}
	if err := ctx.Validate(&newProject); err != nil {
		return err
	}
	err := repository.ProjectRepository.UpdateProject(project, newProject)
	return err
}

func HandleDeleteProject(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	err := repository.ProjectRepository.DeleteProject(project)
	return err
}
