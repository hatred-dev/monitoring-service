package middlewares

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/src/services"
)

func ReloadProjects(context *gin.Context) {
	// executes only after request is successfully handled
	context.Next()
	if context.IsAborted() {
		return
	}
	if context.Request.Method != "GET" {
		services.TimerService.ResetTimer()
	}
}
