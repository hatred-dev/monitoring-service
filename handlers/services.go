package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
)

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
	return nil
}

func HandlePatchService(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var service *api.UpdateServiceReq
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	err := repository.ServiceRepository.UpdateService(project, service.ServiceName, service.Settings)
	if err != nil {
		return err
	}
	return nil
}

func HandleDeleteService(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var service *api.DeleteServiceReq
	if err := ctx.Bind(&service); err != nil {
		return err
	}
	err := repository.ServiceRepository.DeleteService(project, service.ServiceName)
	if err != nil {
		return err
	}
	return nil
}
