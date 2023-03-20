package models

type CreateIPAddressReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	Ip          string `json:"ip" binding:"required"`
}

type CreateIPAddressResp struct {
	Id string `json:"id" binding:"required"`
	Ip string `json:"ip" binding:"required"`
}

type UpdateIPAddressReq struct {
	Ip    string `json:"ip" binding:"required"`
	NewIp string `json:"new_ip" binding:"required"`
}

type DeleteIPAddressReq struct {
	Ip string `json:"ip" binding:"required"`
}
