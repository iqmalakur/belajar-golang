package web

import "net/http"

type WebResponse struct {
	Code   int
	Status string
	Data   any
}

func NewWebResponse(code int, data any) WebResponse {
	return WebResponse{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}
}
