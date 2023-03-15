package handlers

import (
	"github.com/gin-gonic/gin"
	"monitoring-service/database"
	"monitoring-service/src/api/helpers"
	"monitoring-service/src/api/models"
	"net/http"
)

func HandleCreateIP(ctx *gin.Context) {
	var ip *models.CreateIPAddressRequest
	if err := ctx.ShouldBind(&ip); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	projectExists := helpers.CheckIfProjectExists(ctx, ip.ProjectName)
	if !projectExists {
		return
	}
	newIp, err := database.Conn.CreateIP(ctx, database.CreateIPParams{
		ProjectName: ip.ProjectName,
		Ip:          ip.Ip,
	})
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, models.CreateIPAddressResponse{
		Id: newIp.ID.String(),
		Ip: newIp.Ip,
	})
}

func HandleUpdateIP(ctx *gin.Context) {
	var ip *models.UpdateIPAddressRequest
	if err := ctx.ShouldBind(&ip); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	exists := helpers.CheckIfIPExists(ctx, ip.Ip)
	if !exists {
		return
	}
	err := database.Conn.UpdateIP(ctx, database.UpdateIPParams{
		Ip:   ip.Ip,
		Ip_2: ip.NewIp,
	})
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
}

func HandleDeleteIP(ctx *gin.Context) {
	var ip *models.DeleteIPAddressRequest
	if err := ctx.ShouldBind(&ip); err != nil {
		helpers.Abort(ctx, err)
		return
	}
	exists := helpers.CheckIfIPExists(ctx, ip.Ip)
	if !exists {
		return
	}
	err := database.Conn.DeleteIP(ctx, ip.Ip)
	if err != nil {
		helpers.Abort(ctx, err)
		return
	}
}

func HandleGetIPs(ctx *gin.Context) {
	ips, _ := database.Conn.GetAllIPs(ctx)
	ctx.JSON(http.StatusOK, ips)
}

func HandleGetIPsByProjectName(ctx *gin.Context) {
	projectName := ctx.Param("project_name")
	projectExists := helpers.CheckIfProjectExists(ctx, projectName)
	if !projectExists {
		return
	}
	ips, err := database.Conn.GetIPsByProjectName(ctx, projectName)
	if err != nil {
		helpers.Abort(ctx, err)
	}
	ctx.JSON(http.StatusOK, ips)
}
