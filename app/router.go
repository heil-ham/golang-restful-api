package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/category", categoryController.Create)
	router.PUT("/api/category/:categoryId", categoryController.Update)
	router.DELETE("/api/category/:categoryId", categoryController.Delete)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.GET("/api/category", categoryController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}