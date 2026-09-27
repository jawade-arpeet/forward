package router

import (
	"forward/internal/handler"
	v1 "forward/internal/router/v1"

	"github.com/labstack/echo/v5"
)

func New(hdlr *handler.Handler) *echo.Echo {
	e := echo.New()

	apiGrp := e.Group("/api")

	v1.MountV1Router(apiGrp, hdlr)

	return e
}
