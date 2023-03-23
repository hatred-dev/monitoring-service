package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRoot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "working",
	})
}
