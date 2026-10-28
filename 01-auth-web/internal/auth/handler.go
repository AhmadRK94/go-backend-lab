package auth

import "github.com/go-playground/validator/v10"

type Handler struct {
	service  *Service
	validate *validator.Validate
}

func NewHandler(service *Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		validate: validator,
	}
}
