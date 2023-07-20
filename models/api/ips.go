package api

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
