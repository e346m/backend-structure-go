package http

import (
	usecase "github.com/e346m/bsg/internal/booking/usecases"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		usecase: *usecase,
	}
}
