//go:build wireinject
// +build wireinject

package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
	app.NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitializedServer() *http.Server{
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}