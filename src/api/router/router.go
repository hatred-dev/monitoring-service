package router

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/src/api/handlers"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/projects", handlers.HandleCreateProject)
		apiGroup.GET("/projects", handlers.HandleGetProjects)
	}
	return r
}
