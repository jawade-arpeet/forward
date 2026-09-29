package handler

import "forward/internal/service"

type Handler struct {
	Health *HealthHandler
	Auth   *AuthHandler
}

func New(svc *service.Service) *Handler {
	return &Handler{
		Health: newHealthHandler(),
		Auth:   newAuthHandler(svc.Auth),
	}
}
