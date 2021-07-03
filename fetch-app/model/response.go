package model

type SuccessResponse struct {
	Status int         `json:"status_code"`
	Data   interface{} `json:"data"`
}

type BadResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
}