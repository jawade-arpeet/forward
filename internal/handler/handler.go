package handler

import "forward/internal/service"

type Handler struct {
	Health *HealthHdlr
}

func New(svc *service.Service) *Handler {
	return &Handler{
		Health: newHealthHdlr(),
	}
}
