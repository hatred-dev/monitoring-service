package router

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/src/api/handlers"
	"monitoring-service/src/api/middlewares"
	"monitoring-service/src/configuration"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group(configuration.AppConf.RootPrefix + "/api")
	apiGroup.GET("/", handlers.HandleRoot)
	apiGroup.Use(middlewares.ReloadProjects)
	{
		projects := apiGroup.Group("/projects")
		{
			projects.POST("/", handlers.HandleCreateProject)
			projects.GET("/", handlers.HandleGetProjects)
			projects.GET("/:project_name", handlers.HandleGetProjectByName)
			projects.DELETE("/", handlers.HandleDeleteProject)
			projects.PATCH("/", handlers.HandlePatchProject)
		}
		services := apiGroup.Group("/services")
		{
			services.POST("/", handlers.HandleCreateService)
			services.GET("/", handlers.HandleGetServices)
			services.GET("/:project_name", handlers.HandleGetServicesByProjectName)
			services.DELETE("/", handlers.HandleDeleteService)
			services.PATCH("/", handlers.HandlePatchService)
		}
		ips := apiGroup.Group("/ips")
		{
			ips.POST("/", handlers.HandleCreateIP)
			ips.PATCH("/", handlers.HandleUpdateIP)
			ips.GET("/", handlers.HandleGetIPs)
			ips.GET("/:project_name", handlers.HandleGetIPsByProjectName)
			ips.DELETE("/", handlers.HandleDeleteIP)
		}
	}
	return r
}
