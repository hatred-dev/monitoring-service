package models

type CreateIPAddressRequest struct {
	ProjectName string `json:"project_name" binding:"required"`
	Ip          string `json:"ip" binding:"required"`
}

type CreateIPAddressResponse struct {
	Id string `json:"id" binding:"required"`
	Ip string `json:"ip" binding:"required"`
}

type UpdateIPAddressRequest struct {
	Ip    string `json:"ip" binding:"required"`
	NewIp string `json:"new_ip" binding:"required"`
}
