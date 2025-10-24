package web

import "net/http"

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func NewWebResponse(code int, data any) WebResponse {
	return WebResponse{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}
}
