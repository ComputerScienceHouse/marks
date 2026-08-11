package models

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type UnauthorizedResponse struct {
	Message string `json:"message"`
}
