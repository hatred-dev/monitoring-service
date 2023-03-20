package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"monitoring-service/src/api/handlers"
	"monitoring-service/src/api/validators"
	srv "monitoring-service/src/services"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("project-exists", validators.ProjectExists)
		if err != nil {
			return nil
		}
	}
	apiGroup := r.Group("/api")
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
