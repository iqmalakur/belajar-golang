package controller

import (
	"iqmalakur/belajar-golang-dependency-injection/helper"
	"iqmalakur/belajar-golang-dependency-injection/model/web"
	"iqmalakur/belajar-golang-dependency-injection/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{Service: service}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.Service.Create(request.Context(), categoryCreateRequest)
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

	categoryResponse := controller.Service.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idParam := params.ByName("id")
	id, err := strconv.Atoi(idParam)
	helper.PanicIfError(err)

	controller.Service.Delete(request.Context(), id)
	webResponse := web.NewWebResponse(http.StatusOK, nil)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramId := params.ByName("id")
	id, err := strconv.Atoi(paramId)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.FindById(request.Context(), id)
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.Service.FindAll(request.Context())
	webResponse := web.NewWebResponse(http.StatusOK, categoryResponses)

	helper.WriteToResponseBody(writer, webResponse)
}
