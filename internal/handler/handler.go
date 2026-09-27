package handler

import "forward/internal/service"

type Handler struct {
	Health *HealthHandler
}

func New(svc *service.Service) *Handler {
	return &Handler{
		Health: newHealthHandler(),
	}
}
