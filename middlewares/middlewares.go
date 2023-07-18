package middlewares

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"monitoring-service/services"
)

func ReloadProjects(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println(context)
		if context.Request().Method != "GET" {
			services.TimerService.ResetTimer()
		}
		return next(context)
	}
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if errs, ok := cv.Validator.Struct(i).(validator.ValidationErrors); ok {
		if errs != nil {
			for _, err := range errs {
				if fe, ok := err.(validator.FieldError); ok {
					ve := validator.ValidationErrors{fe}
					return ve
				}
			}
		}
	}
	return nil
}
