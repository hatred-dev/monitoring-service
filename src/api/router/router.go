package router

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/src/api/handlers"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
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
