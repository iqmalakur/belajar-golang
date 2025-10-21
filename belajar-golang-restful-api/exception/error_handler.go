package exception

import (
	"iqmalakur/belajar-golang-restful-api/helper"
	"iqmalakur/belajar-golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	switch err.(type) {
	case NotFoundError:
		notFoundError(writer, err)
	case validator.ValidationErrors:
		validationiErrors(writer, err)
	default:
		internalServerError(writer, err)
	}
}

func internalServerError(writer http.ResponseWriter, err any) {
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.NewWebResponse(http.StatusInternalServerError, err)
	helper.WriteToResponseBody(writer, webResponse)
}

func notFoundError(writer http.ResponseWriter, err any) {
	exception := err.(NotFoundError)

	writer.WriteHeader(http.StatusNotFound)
	webResponse := web.NewWebResponse(http.StatusNotFound, exception.Error())
	helper.WriteToResponseBody(writer, webResponse)
}

func validationiErrors(writer http.ResponseWriter, err any) {
	exception := err.(validator.ValidationErrors)

	writer.WriteHeader(http.StatusBadRequest)
	webResponse := web.NewWebResponse(http.StatusBadRequest, exception.Error())
	helper.WriteToResponseBody(writer, webResponse)
}
