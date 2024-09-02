package middlewares

import (
    "errors"
    "monitoring-service/configuration"
    "monitoring-service/repository"
    "monitoring-service/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

func ReloadProjects() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Next()
        if ctx.Request.Method != http.MethodGet && !ctx.IsAborted() {
            services.TimerService.TriggerReset()
        }
    }
}

func ProjectExists() gin.HandlerFunc {
    // if project exists, set it to current context to fetch later
    return func(ctx *gin.Context) {
        projectName := ctx.Param("project_name")
        if projectName == "" {
            ctx.Next()
            return
        }
        project, err := repository.ProjectRepository.ProjectExists(projectName)
        if err != nil {
            _ = ctx.AbortWithError(http.StatusBadRequest, err)
            return
        }
        ctx.Set("project", project)
    }
}

func VerifyRequest() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.Request.Header.Get("api-key")
        if token != configuration.GetApplicationConfig().ApiKey {
            _ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid api key"))
            return
        }
    }
}
