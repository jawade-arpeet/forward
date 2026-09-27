package router

import (
	"forward/internal/handler"
	"forward/internal/middleware"
	v1 "forward/internal/router/v1"

	"github.com/labstack/echo/v5"
)

func New(mw *middleware.Middleware, hdlr *handler.Handler) *echo.Echo {
	e := echo.New()

	e.Use(mw.Request.SetRequestID)

	apiGrp := e.Group("/api")

	v1.MountV1Router(apiGrp, hdlr)

	return e
}
