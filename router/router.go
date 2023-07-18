package router

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"monitoring-service/configuration"
	"monitoring-service/handlers"
	"monitoring-service/middlewares"
)

func CreateRouter() *echo.Echo {
	r := CreateRouter()
	r.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	apiGroup := r.Group(configuration.AppConf.RootPrefix + "/api")
	apiGroup.GET("/", handlers.HandleRoot)
	apiGroup.Use(middlewares.ReloadProjects)
	{
		projects := apiGroup.Group("/projects")
		{
			projects.POST("/", handlers.HandleCreateProject)
			projects.GET("/", handlers.HandleGetProjects)
			projects.GET("/:project_name", handlers.HandleGetProjectByName)
			projects.PATCH("/:project_name", handlers.HandlePatchProject)
			projects.DELETE("/:project_name", handlers.HandleDeleteProject)
			projects.POST("/:project_name/ips", handlers.HandleCreateIP)
			projects.PATCH("/:project_name/ips", handlers.HandleUpdateIP)
			projects.DELETE("/:project_name/ips", handlers.HandleDeleteIP)
			projects.POST("/:project_name/services", handlers.HandleCreateService)
			projects.PATCH("/:project_name/services", handlers.HandlePatchService)
			projects.DELETE("/:project_name/services", handlers.HandleDeleteService)
		}
	}
	return r
}
