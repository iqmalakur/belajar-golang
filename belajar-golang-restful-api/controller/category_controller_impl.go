package controller

import (
	"iqmalakur/belajar-golang-restful-api/helper"
	"iqmalakur/belajar-golang-restful-api/model/web"
	"iqmalakur/belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	service service.CategoryService
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.service.Create(request.Context(), categoryCreateRequest)
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	idParam := params.ByName("id")
	id, err := strconv.Atoi(idParam)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.service.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idParam := params.ByName("id")
	id, err := strconv.Atoi(idParam)
	helper.PanicIfError(err)

	controller.service.Delete(request.Context(), id)
	webResponse := web.NewWebResponse(http.StatusOK, nil)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramId := params.ByName("id")
	id, err := strconv.Atoi(paramId)
	helper.PanicIfError(err)

	categoryResponse := controller.service.FindById(request.Context(), id)
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.service.FindAll(request.Context())
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponses)

	helper.WriteToResponseBody(writer, webResponse)
}
