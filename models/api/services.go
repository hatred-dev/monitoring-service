package api

type DeleteServiceReq struct {
	ServiceName string `json:"service_name" binding:"required"`
}
