package handlers

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/database"
	"monitoring-service/src/api/helpers"
	"monitoring-service/src/api/models"
	"net/http"
)

func HandleCreateService(ctx *gin.Context) {
	var service *models.CreateServiceRequest
	if err := ctx.ShouldBind(&service); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	if exists := helpers.CheckIfProjectExists(ctx, service.ProjectName); !exists {
		return
	}
	createdService, err := database.Conn.CreateService(ctx, database.CreateServiceParams{
		ProjectName: service.ProjectName,
		ServiceName: service.ServiceName,
		Url:         service.Url,
	})
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, models.CreateServiceResponse{
		Id:          createdService.ID.String(),
		ServiceName: createdService.ServiceName,
		Url:         createdService.Url,
	})
}

func HandleGetServices(ctx *gin.Context) {
	services, err := database.Conn.GetServices(ctx)
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, services)
}

func HandleGetServicesByProjectName(ctx *gin.Context) {
	projectName := ctx.Param("project_name")
	if exists := helpers.CheckIfProjectExists(ctx, projectName); !exists {
		return
	}
	services, err := database.Conn.GetServicesByProjectName(ctx, projectName)
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	servicesResponse := make([]*models.GetServicesByProjectNameResponse, len(services))
	for i, v := range services {
		servicesResponse[i] = &models.GetServicesByProjectNameResponse{
			Id:          v.ID.String(),
			ServiceName: v.ServiceName,
			Url:         v.Url,
		}
	}
	ctx.JSON(http.StatusOK, servicesResponse)
}

func HandleDeleteService(ctx *gin.Context) {
	var service *models.DeleteServiceRequest
	if err := ctx.ShouldBind(&service); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	if exists := helpers.CheckIfServiceExists(
		ctx, service.ProjectName, service.ServiceName); !exists {
		return
	}
	err := database.Conn.DeleteService(ctx, database.DeleteServiceParams{
		ProjectName: service.ProjectName,
		ServiceName: service.ServiceName,
	})
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
}
