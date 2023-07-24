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
	if err != nil {
		return err
	}
	return nil
}

func HandleGetProjectByName(ctx echo.Context) error {
	projectName := ctx.Param("project_name")
	project, err := repository.ProjectRepository.GetProjectByName(projectName)
	if err != nil {
		return err
	}
	err = ctx.JSON(http.StatusOK, project)
	if err != nil {
		return err
	}
	return nil
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
	if err != nil {
		return err
	}
	return nil
}

func HandlePatchProject(ctx echo.Context) error {
	projectName := ctx.Param("project_name")
	var newProject database.Project
	if err := ctx.Bind(&newProject); err != nil {
		return err
	}
	if err := ctx.Validate(&newProject); err != nil {
		return err
	}
	err := repository.ProjectRepository.UpdateProject(projectName, newProject)
	if err != nil {
		return err
	}
	return nil
}

func HandleDeleteProject(ctx echo.Context) error {
	projectName := ctx.Param("project_name")
	err := repository.ProjectRepository.DeleteProject(projectName)
	if err != nil {
		return err
	}
	return nil
}
