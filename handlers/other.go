package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleRoot(ctx echo.Context) error {
	err := ctx.JSON(http.StatusOK, echo.Map{
		"status": "working",
	})
	if err != nil {
		return err
	}
	return nil
}
