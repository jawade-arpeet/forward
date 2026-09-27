package middleware

import (
	"uuid"

	"github.com/labstack/echo/v5"
)

type RequestMiddleware struct{}

func newRequestMiddleware() *RequestMiddleware {
	return &RequestMiddleware{}
}

func (m *RequestMiddleware) SetRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx *echo.Context) error {
		reqID := ctx.Request().Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.NewV7().String()
			ctx.Request().Header.Set("X-Request-ID", reqID)
		}

		ctx.Response().Header().Set("X-Request-ID", reqID)
		ctx.Set("request_id", reqID)

		return next(ctx)
	}
}
