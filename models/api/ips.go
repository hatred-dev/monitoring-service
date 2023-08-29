package api

import "monitoring-service/models/database"

type CreateIPAddressReq struct {
	Ip string `json:"ip" validate:"required"`
}

type UpdateIPAddressReq struct {
	Ip    string `json:"ip" binding:"required"`
	NewIp string `json:"new_ip" binding:"required"`
}

type DeleteIPAddressReq struct {
	Ip string `json:"ip" binding:"required"`
}

type UpdateServiceReq struct {
	ServiceName string           `json:"service_name" binding:"required"`
	Service     database.Service `json:"settings" binding:"required"`
}
