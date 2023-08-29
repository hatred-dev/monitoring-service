package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"monitoring-service/utils"
	"net/http"
)

func HandleGetIPs(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	ips := repository.IpRepository.GetIps(&project)
	err := ctx.JSON(http.StatusOK, ips)
	return err
}

func HandleCreateIP(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var ip *api.CreateIPAddressReq
	if err := utils.BindAndValidate(ctx, &ip); err != nil {
		return err
	}
	id, err := repository.IpRepository.CreateIp(project, database.Ip{
		Ip: ip.Ip,
	})
	if err != nil {
		return err
	}
	err = ctx.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
	return err
}

func HandleUpdateIP(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var ip *api.UpdateIPAddressReq
	if err := utils.BindAndValidate(ctx, &ip); err != nil {
		return err
	}
	err := repository.IpRepository.UpdateIp(project, ip.Ip, ip.NewIp)
	return err
}

func HandleDeleteIP(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var ip *api.DeleteIPAddressReq
	if err := utils.BindAndValidate(ctx, &ip); err != nil {
		return err
	}
	err := repository.IpRepository.DeleteIp(project, ip.Ip)
	return err
}
