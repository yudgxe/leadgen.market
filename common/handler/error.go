package handler

import "net/http"

type HttpCodeError struct {
	Code    int
	Message string
}

func (e HttpCodeError) Error() string {
	return e.Message
}

func NewHttpError(code int, msg string) HttpCodeError {
	return HttpCodeError{
		Code:    code,
		Message: msg,
	}
}

func NewHttpErrorBadRequest(msg string) HttpCodeError {
	return NewHttpError(http.StatusBadRequest, msg)
}
