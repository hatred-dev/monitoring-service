package router

import (
	"monitoring-service/configuration"
	"monitoring-service/handlers"
	"monitoring-service/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.New()
	apiGroup := r.Group(configuration.GetApplicationConfig().RootPrefix + "/api")
	apiGroup.GET("/", handlers.HandleRoot)
	apiGroup.Use(middlewares.VerifyRequest(), middlewares.ProjectExists(), middlewares.ReloadProjects())
	{
		projects := apiGroup.Group("/projects")
		{
			// All projects operations
			projects.POST("", handlers.HandleCreateProject)
			projects.GET("", handlers.HandleGetProjects)
			// Specific project operations
			project := projects.Group("/:project_name")
			{
				project.GET("", handlers.HandleGetProjectByName)
				project.PATCH("", handlers.HandlePatchProject)
				project.DELETE("", handlers.HandleDeleteProject)
				// Ip operations
				project.GET("/ips", handlers.HandleGetIPs)
				project.POST("/ips", handlers.HandleCreateIP)
				project.PATCH("/ips", handlers.HandleUpdateIP)
				project.DELETE("/ips", handlers.HandleDeleteIP)
				// Service operations
				project.GET("/services", handlers.HandleGetServices)
				project.POST("/services", handlers.HandleCreateService)
				project.PATCH("/services", handlers.HandlePatchService)
				project.DELETE("/services", handlers.HandleDeleteService)
			}
		}
	}
	return r
}
