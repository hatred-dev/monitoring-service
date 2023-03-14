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
	exists, _ := database.Conn.IPExists(ctx, ip.Ip)
	if !exists {
		helpers.SendError(ctx, "ip does not exist")
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
