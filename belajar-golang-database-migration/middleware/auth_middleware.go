package middleware

import (
	"iqmalakur/belajar-golang-database-migration/helper"
	"iqmalakur/belajar-golang-database-migration/model/web"
	"net/http"
)

type AuthMiddleware struct {
	http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := request.Header.Get("X-API-KEY")

	if apiKey == "RAHASIA" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.NewWebResponse(http.StatusUnauthorized, nil)
		helper.WriteToResponseBody(writer, webResponse)
	}
}
