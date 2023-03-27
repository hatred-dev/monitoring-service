package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"monitoring-service/database"
	"net/http"
)

func Abort(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func CheckIfProjectExists(ctx *gin.Context, name string) bool {
	exists, _ := database.Conn.ProjectExists(ctx, name)
	if exists {
		Abort(ctx, errors.New("project already exists"))
	}
	return exists
}

func CheckIfServiceExists(ctx *gin.Context, projectName string, serviceName string) bool {
	exists, _ := database.Conn.ServiceExists(ctx, database.ServiceExistsParams{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	if !exists {
		Abort(ctx, errors.New("service already exists"))
	}
	return exists
}

func CheckIfIPExists(ctx *gin.Context, ip string) bool {
	exists, _ := database.Conn.IPExists(ctx, ip)
	if !exists {
		Abort(ctx, errors.New("ip already exists"))
	}
	return exists
}
