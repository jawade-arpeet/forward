package handler

import (
	"forward/internal/errs"
	"forward/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func newAuthHandler(authSvc *service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *AuthHandler) SignUp(ctx *echo.Context) error {
	var req SignUpRequest
	if err := ctx.Bind(&req); err != nil {
		return errs.BadRequest(
			errs.CodeInvalidInput,
			"invalid payload",
			err,
		)
	}

	if err := validator.New().Struct(req); err != nil {
		return errs.BadRequest(
			errs.CodeInvalidInput,
			"invalid payload",
			err,
		)
	}

	if err := h.authSvc.SignUp(
		ctx.Request().Context(),
		req.Email,
		req.Password,
	); err != nil {
		return err
	}

	return ctx.JSON(
		http.StatusCreated,
		map[string]string{"message": "account signed up successfully"},
	)
}

type SignInPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *AuthHandler) SignIn(ctx *echo.Context) error {
	var req SignUpRequest
	if err := ctx.Bind(&req); err != nil {
		return errs.BadRequest(
			errs.CodeInvalidInput,
			"invalid payload",
			err,
		)
	}

	if err := validator.New().Struct(req); err != nil {
		return errs.BadRequest(
			errs.CodeInvalidInput,
			"invalid payload",
			err,
		)
	}

	token, err := h.authSvc.SignIn(
		ctx.Request().Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]string{
			"message":    "account signed in successfully",
			"account_id": token,
		},
	)
}
