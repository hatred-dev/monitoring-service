package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "working",
	})
}
