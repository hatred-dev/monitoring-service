package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"monitoring-service/src/api/handlers"
	"monitoring-service/src/configuration"
	srv "monitoring-service/src/services"
	"net/http"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group(configuration.AppConf.RootPrefix + "/api")
	apiGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "working",
		})
	})
	apiGroup.Use(func(context *gin.Context) {
		// executes only after request is successfully handled
		context.Next()
		if context.IsAborted() {
			return
		}

		if context.Request.Method != "GET" {
			fmt.Println("Reloading...")
			go srv.SupervisorObject.ReloadServices()
		}

	})
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
