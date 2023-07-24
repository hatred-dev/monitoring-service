package router

import (
	"monitoring-service/configuration"
	"monitoring-service/handlers"
	"monitoring-service/middlewares"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func CreateRouter() *echo.Echo {
	r := echo.New()
	r.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	apiGroup := r.Group(configuration.AppConf.RootPrefix + "/api")
	apiGroup.GET("/", handlers.HandleRoot)
	apiGroup.Use(middlewares.VerifyRequest, middlewares.ProjectExists, middlewares.ReloadProjects)
	{
		projects := apiGroup.Group("/projects")
		{
			// Project operations
			projects.POST("", handlers.HandleCreateProject)
			projects.GET("", handlers.HandleGetProjects)
			projects.GET("/:project_name", handlers.HandleGetProjectByName)
			projects.PATCH("/:project_name", handlers.HandlePatchProject)
			projects.DELETE("/:project_name", handlers.HandleDeleteProject)
			// Ip operations
			projects.POST("/:project_name/ips", handlers.HandleCreateIP)
			projects.PATCH("/:project_name/ips", handlers.HandleUpdateIP)
			projects.DELETE("/:project_name/ips", handlers.HandleDeleteIP)
			// Service operations
			projects.POST("/:project_name/services", handlers.HandleCreateService)
			projects.PATCH("/:project_name/services", handlers.HandlePatchService)
			projects.DELETE("/:project_name/services", handlers.HandleDeleteService)
		}
	}
	return r
}
