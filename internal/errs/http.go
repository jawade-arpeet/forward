package errs

import (
	"fmt"
	"net/http"
)

type Code string

const (
	CodeInternal     Code = "INTERNAL_SERVER_ERROR"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeNotFound     Code = "NOT_FOUND"
	CodeConflict     Code = "CONFLICT"
	CodeInvalidInput Code = "INVALID_INPUT"
)

type HTTPError struct {
	Code    Code   `json:"code"`
	Status  int    `json:"-"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func new(code Code, status int, msg string, err error) *HTTPError {
	return &HTTPError{Code: code, Status: status, Message: msg, Err: err}
}

func (e *HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}

	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func BadRequest(code Code, msg string, err error) *HTTPError {
	return &HTTPError{
		Code:    code,
		Status:  http.StatusBadRequest,
		Message: msg,
		Err:     err,
	}
}

func Unauthorized(code Code, msg string, err error) *HTTPError {
	return &HTTPError{
		Code:    code,
		Status:  http.StatusUnauthorized,
		Message: msg,
		Err:     err,
	}
}

func NotFound(code Code, msg string, err error) *HTTPError {
	return &HTTPError{
		Code:    code,
		Status:  http.StatusNotFound,
		Message: msg,
		Err:     err,
	}
}

func Conflict(code Code, msg string, err error) *HTTPError {
	return &HTTPError{
		Code:    code,
		Status:  http.StatusConflict,
		Message: msg,
		Err:     err,
	}
}

func Internal(err error) *HTTPError {
	return new(
		CodeInternal,
		http.StatusInternalServerError,
		"something went wrong",
		err,
	)
}
