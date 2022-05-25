package zhttp

import (
	"fmt"
	"net/http"
)

type Error interface {
	HTTPError() (int, string)
}

type httpError struct {
	error
	statusCode int
}

func (e *httpError) HTTPError() (int, string) {
	return e.statusCode, e.Error()
}

func errorWithStatus(statusCode int, err error) *httpError {
	return &httpError{
		error:      err,
		statusCode: statusCode,
	}
}

// BadRequest return 400 with error
func BadRequest(err error) *httpError {
	return errorWithStatus(http.StatusBadRequest, err)
}

// BadRequestf return 400 with format error
func BadRequestf(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusBadRequest, fmt.Errorf(format, args...))
}

// Unauthorized return 401 with error
func Unauthorized(err error) *httpError {
	return errorWithStatus(http.StatusUnauthorized, err)
}

// Unauthorizedf return 401 with format error
func Unauthorizedf(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusUnauthorized, fmt.Errorf(format, args...))
}

// PaymentRequired return 402 with error
func PaymentRequired(err error) *httpError {
	return errorWithStatus(http.StatusPaymentRequired, err)
}

// PaymentRequired return 402 with format error
func PaymentRequiredf(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusUnauthorized, fmt.Errorf(format, args...))
}

// Forbidden return 403 with error
func Forbidden(err error) *httpError {
	return errorWithStatus(http.StatusForbidden, err)
}

// Forbiddenf return 403 with format error
func Forbiddenf(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusForbidden, fmt.Errorf(format, args...))
}

// NotFound return 404 with error
func NotFound(err error) *httpError {
	return errorWithStatus(http.StatusNotFound, err)
}

// NotFoundf return 404 with format error
func NotFoundf(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusNotFound, fmt.Errorf(format, args...))
}

// NotFound return 406 with error
func NotAcceptable(err error) *httpError {
	return errorWithStatus(http.StatusNotAcceptable, err)
}

// NotAcceptablef return 406 with format error
func NotAcceptablef(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusNotFound, fmt.Errorf(format, args...))
}

// InternalServerError return 500 with error
func InternalServerError(err error) *httpError {
	return errorWithStatus(http.StatusInternalServerError, err)
}

// NotAcceptablef return 500 with format error
func InternalServerErrorf(format string, args ...any) *httpError {
	return errorWithStatus(http.StatusInternalServerError, fmt.Errorf(format, args...))
}