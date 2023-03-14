package models

type AddIpAddressRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
	Ip          string `json:"ip" binding:"required"`
}
