package v1

import (
	"forward/internal/handler"

	"github.com/labstack/echo/v5"
)

func MountV1Router(
	routerGrp *echo.Group,
	hdlr *handler.Handler,
) {
	v1Grp := routerGrp.Group("/v1")

	mountHealthRouter(v1Grp, hdlr.Health)
}
