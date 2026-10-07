package model

type APIResponse struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
	PopUp   string `json:"popUp"`
}
