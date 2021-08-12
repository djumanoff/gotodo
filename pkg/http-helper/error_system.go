package http_helper

import (
	"net/http"
	"strconv"
)

type ErrorSystem interface {
	BadRequest(code int, message string) Error

	InternalServerError(code int, message string) Error

	NotFound(code int, message string) Error

	Forbidden(code int, message string) Error

	Unauthorized(code int, message string) Error
}

type httpErrorSystem struct {
	system string
}

func NewErrorSystem(system string) ErrorSystem {
	return &httpErrorSystem{system}
}

func (sys httpErrorSystem) NewError(code int, status int, message string) Error {
	err := &HttpError{}
	err.Code = code
	err.Status = status
	err.statusCode = err.Status
	err.Message = err.System + "/" +
		strconv.Itoa(err.Status) + "/" +
		strconv.Itoa(err.Code) + "::" +
		message

	return err
}

func (sys httpErrorSystem) BadRequest(code int, message string) Error {
	return sys.NewError(code, http.StatusBadRequest, message)
}

func (sys httpErrorSystem) InternalServerError(code int, message string) Error {
	return sys.NewError(code, http.StatusInternalServerError, message)
}

func (sys httpErrorSystem) NotFound(code int, message string) Error {
	return sys.NewError(code, http.StatusNotFound, message)
}

func (sys httpErrorSystem) Forbidden(code int, message string) Error {
	return sys.NewError(code, http.StatusForbidden, message)
}

func (sys httpErrorSystem) Unauthorized(code int, message string) Error {
	return sys.NewError(code, http.StatusUnauthorized, message)
}
