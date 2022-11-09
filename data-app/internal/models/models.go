package models

type ResponseRestApi struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}
