package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
)

func HandleCreateService(ctx echo.Context) error {
	projectName := ctx.Param("project_name")
	var service database.Service
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	err := repository.ProjectRepository.CreateService(projectName, service)
	if err != nil {
		return err
	}
	return nil
}

func HandlePatchService(ctx echo.Context) error {
	return nil
}

func HandleDeleteService(ctx echo.Context) error {
	projectName := ctx.Param("project_name")
	var service *api.DeleteServiceReq
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	err := repository.ProjectRepository.DeleteService(projectName, service.ServiceName)
	if err != nil {
		return err
	}
	return nil
}
