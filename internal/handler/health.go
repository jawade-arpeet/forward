package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthHandler struct{}

func newHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
