package handlers

import (
	"github.com/labstack/echo/v4"
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
)

func HandleCreateIP(ctx echo.Context) error {
	projectName := ctx.Param("project_name")
	var ip *api.CreateIPAddressReq
	if err := ctx.Bind(&ip); err != nil {
		return err
	}
	err := repository.ProjectRepository.CreateIp(projectName, database.Ip{
		Ip: ip.Ip,
	})
	if err != nil {
		return err
	}
	err = ctx.JSON(http.StatusCreated, echo.Map{})
	return nil
}

func HandleUpdateIP(ctx echo.Context) error {
	var ip *api.UpdateIPAddressReq
	if err := ctx.Bind(&ip); err != nil {
		return err
	}
	err := repository.ProjectRepository.UpdateIp(ip.Ip, ip.NewIp)
	if err != nil {
		return err
	}
	return nil
}

func HandleDeleteIP(ctx echo.Context) error {
	var ip *api.DeleteIPAddressReq
	if err := ctx.Bind(&ip); err != nil {
		return err
	}
	err := repository.ProjectRepository.DeleteIp(ip.Ip)
	if err != nil {
		return err
	}
	return nil
}
