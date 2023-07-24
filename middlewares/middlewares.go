package middlewares

import (
	"monitoring-service/configuration"
	"monitoring-service/repository"
	"monitoring-service/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ReloadProjects(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		err := next(context)
		if context.Request().Method != http.MethodGet && err == nil {
			services.TimerService.TriggerReset()
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return nil
	}
}

func ProjectExists(next echo.HandlerFunc) echo.HandlerFunc {
	// if project exists, set it to current context to fetch later
	return func(context echo.Context) error {
		projectName := context.Param("project_name")
		if projectName == "" {
			return next(context)
		}
		project, err := repository.ProjectRepository.ProjectExists(projectName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		context.Set("project", project)
		return next(context)
	}
}

func VerifyRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := context.Request().Header.Get("api-key")
		if token != configuration.AppConf.ApiKey {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid api key")
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
