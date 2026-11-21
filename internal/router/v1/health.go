package v1

import (
	"forward/internal/handler"

	"github.com/labstack/echo/v5"
)

func mountHealthRouter(
	routerGrp *echo.Group,
	hdlr *handler.HealthHdlr,
) {
	healthGrp := routerGrp.Group("/health")

	healthGrp.GET("/check", hdlr.Check)
}
