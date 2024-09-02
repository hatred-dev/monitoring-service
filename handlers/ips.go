package handlers

import (
    "monitoring-service/models/api"
    "monitoring-service/models/database"
    "monitoring-service/repository"
    "net/http"

    "github.com/gin-gonic/gin"
)

func HandleGetIPs(ctx *gin.Context) {
    project := ctx.MustGet("project").(database.Project)
    ips := repository.IpRepository.GetIps(&project)
    ctx.JSON(http.StatusOK, ips)
}

func HandleCreateIP(ctx *gin.Context) {
    project := ctx.MustGet("project").(database.Project)
    var ip *api.CreateIPAddressReq
    if err := ctx.ShouldBindJSON(&ip); err != nil {
        _ = ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }
    id, err := repository.IpRepository.CreateIp(project, database.Ip{
        Ip: ip.Ip,
    })
    if err != nil {
        _ = ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{
        "id": id,
    })
}

func HandleUpdateIP(ctx *gin.Context) {
    project := ctx.MustGet("project").(database.Project)
    var ip *api.UpdateIPAddressReq
    if err := ctx.ShouldBindJSON(&ip); err != nil {
        _ = ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }
    err := repository.IpRepository.UpdateIp(project, ip.Ip, ip.NewIp)
    if err != nil {
        _ = ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

}

func HandleDeleteIP(ctx *gin.Context) {
    project := ctx.MustGet("project").(database.Project)
    var ip *api.DeleteIPAddressReq
    if err := ctx.ShouldBindJSON(&ip); err != nil {
        _ = ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }
    err := repository.IpRepository.DeleteIp(project, ip.Ip)
    if err != nil {
        _ = ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }
}
