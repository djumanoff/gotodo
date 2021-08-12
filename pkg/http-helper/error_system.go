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

func (sys httpErrorSystem) BadRequest(code int, message string) Error {
	err := &HttpError{}
	err.System = sys.system
	err.Code = code
	err.StatusCode = http.StatusBadRequest
	err.statusCode = err.StatusCode
	err.Message = err.System + "/" +
		strconv.Itoa(err.StatusCode) + "/" +
		strconv.Itoa(code) + "::" +
		message

	return err
}

func (sys httpErrorSystem) InternalServerError(code int, message string) Error {
	err := &HttpError{}
	err.System = sys.system
	err.Code = code
	err.StatusCode = http.StatusInternalServerError
	err.statusCode = err.StatusCode
	err.Message = err.System + "/" +
		strconv.Itoa(err.StatusCode) + "/" +
		strconv.Itoa(code) + "::" +
		message

	return err
}

func (sys httpErrorSystem) NotFound(code int, message string) Error {
	err := &HttpError{}
	err.System = sys.system
	err.Code = code
	err.StatusCode = http.StatusNotFound
	err.statusCode = err.StatusCode
	err.Message = err.System + "/" +
		strconv.Itoa(err.StatusCode) + "/" +
		strconv.Itoa(code) + "::" +
		message

	return err
}

func (sys httpErrorSystem) Forbidden(code int, message string) Error {
	err := &HttpError{}
	err.System = sys.system
	err.Code = code
	err.StatusCode = http.StatusForbidden
	err.statusCode = err.StatusCode
	err.Message = err.System + "/" +
		strconv.Itoa(err.StatusCode) + "/" +
		strconv.Itoa(code) + "::" +
		message

	return err
}

func (sys httpErrorSystem) Unauthorized(code int, message string) Error {
	err := &HttpError{}
	err.System = sys.system
	err.Code = code
	err.StatusCode = http.StatusUnauthorized
	err.statusCode = err.StatusCode
	err.Message = err.System + "/" +
		strconv.Itoa(err.StatusCode) + "/" +
		strconv.Itoa(code) + "::" +
		message

	return err
}
