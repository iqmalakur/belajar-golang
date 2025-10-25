//go:build wireinject
// +build wireinject

package main

import (
	"iqmalakur/belajar-golang-dependency-injection/app"
	"iqmalakur/belajar-golang-dependency-injection/controller"
	"iqmalakur/belajar-golang-dependency-injection/middleware"
	"iqmalakur/belajar-golang-dependency-injection/repository"
	"iqmalakur/belajar-golang-dependency-injection/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer(options ...validator.Option) *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
