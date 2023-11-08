package entity

type HTTPResponse struct {
	Message    string           `json:"message"`
	IsSuccess  bool             `json:"is_success"`
	Data       interface{}      `json:"data"`
}