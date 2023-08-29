package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
)

func HandleGetServices(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	services := repository.ServiceRepository.GetServices(&project)
	err := ctx.JSON(http.StatusOK, services)
	return err
}

func HandleCreateService(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var service database.Service
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	if err := ctx.Validate(&service); err != nil {
		return err
	}
	id, err := repository.ServiceRepository.CreateService(project, service)
	if err != nil {
		return err
	}
	err = ctx.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
	return err
}

func HandlePatchService(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var service *api.UpdateServiceReq
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	err := repository.ServiceRepository.UpdateService(project, service.ServiceName, service.Service)
	return err
}

func HandleDeleteService(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var service *api.DeleteServiceReq
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	err := repository.ServiceRepository.DeleteService(project, service.ServiceName)
	return err
}
