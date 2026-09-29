package v1

import (
	"forward/internal/handler"

	"github.com/labstack/echo/v5"
)

func mountAuthRouter(
	routerGrp *echo.Group,
	hdlr *handler.AuthHandler,
) {
	authGrp := routerGrp.Group("/auth")

	authGrp.POST("/sign-up", hdlr.SignUp)
	authGrp.POST("/sign-in", hdlr.SignIn)
}
