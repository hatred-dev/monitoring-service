package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
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
	if err := ctx.Bind(&ip); err != nil {
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
	if err := ctx.Bind(&ip); err != nil {
		return err
	}
	err := repository.IpRepository.UpdateIp(project, ip.Ip, ip.NewIp)
	return err
}

func HandleDeleteIP(ctx echo.Context) error {
	project := ctx.Get("project").(database.Project)
	var ip *api.DeleteIPAddressReq
	if err := ctx.Bind(&ip); err != nil {
		return err
	}
	if err := ctx.Validate(&ip); err != nil {
		return err
	}
	err := repository.IpRepository.DeleteIp(project, ip.Ip)
	return err
}
