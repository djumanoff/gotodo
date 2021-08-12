package http_helper

import (
	"net/http"
)

type Response interface {
	StatusCode() int

	Response() interface{}

	Headers() []http.Header

	AddHeader(header http.Header) Response
}

type response struct {
	data interface{}

	statusCode int

	headers []http.Header
}

func (r response) StatusCode() int {
	return r.statusCode
}

func (r response) Response() interface{} {
	return r.data
}

func (r response) Headers() []http.Header {
	return r.headers
}

func (r *response) AddHeader(header http.Header) Response {
	r.headers = append(r.headers, header)
	return r
}

func OK(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusOK, headers: nil}
}

func Created(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusCreated, headers: nil}
}
