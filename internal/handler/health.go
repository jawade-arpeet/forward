package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthHdlr struct{}

func newHealthHdlr() *HealthHdlr {
	return &HealthHdlr{}
}

func (h *HealthHdlr) Check(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
