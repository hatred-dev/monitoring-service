package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

func abort(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
