package handlers

import (
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetServices(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	services := repository.ServiceRepository.GetServices(&project)
	ctx.JSON(http.StatusOK, services)
}

func HandleCreateService(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	var service database.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := repository.ServiceRepository.CreateService(project, service)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func HandlePatchService(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	var service *api.UpdateServiceReq
	if err := ctx.ShouldBindJSON(&service); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := repository.ServiceRepository.UpdateService(project, service.ServiceName, service.Service)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func HandleDeleteService(ctx *gin.Context) {
	project := ctx.MustGet("project").(database.Project)
	var service *api.DeleteServiceReq
	if err := ctx.ShouldBindJSON(&service); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := repository.ServiceRepository.DeleteService(project, service.ServiceName)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
