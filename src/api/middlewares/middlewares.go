package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	srv "monitoring-service/src/services"
)

func ReloadProjects(context *gin.Context) {
	// executes only after request is successfully handled
	context.Next()
	if context.IsAborted() {
		return
	}

	if context.Request.Method != "GET" {
		fmt.Println("Reloading...")
		go srv.SupervisorObject.ReloadServices()
	}
}
