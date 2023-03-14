package helpers

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/database"
	"net/http"
)

func SendError(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

func Abort(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func CheckIfProjectExists(ctx *gin.Context, name string) bool {
	exists, _ := database.Conn.ProjectExists(ctx, name)
	if !exists {
		SendError(ctx, "project does not exist")
	}
	return exists
}

func CheckIfServiceExists(ctx *gin.Context, projectName string, serviceName string) bool {
	exists, _ := database.Conn.ServiceExists(ctx, database.ServiceExistsParams{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	if !exists {
		SendError(ctx, "service does not exist")
	}
	return exists
}
