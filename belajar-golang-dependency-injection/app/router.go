package app

import (
	"iqmalakur/belajar-golang-dependency-injection/controller"
	"iqmalakur/belajar-golang-dependency-injection/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
