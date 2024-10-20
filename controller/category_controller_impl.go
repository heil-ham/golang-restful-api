package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.categoryService.Create(request.Context(),categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId, errCategoryId := strconv.Atoi(params.ByName("categoryId")) 
	helper.PanicIfError(errCategoryId)
	categoryUpdateRequest.Id = int32(categoryId)

	categoryResponse := controller.categoryService.Update(request.Context(),categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, errCategoryId := strconv.Atoi(params.ByName("categoryId")) 
	helper.PanicIfError(errCategoryId)

	controller.categoryService.Delete(request.Context(),int32(categoryId))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, errCategoryId := strconv.Atoi(params.ByName("categoryId")) 
	helper.PanicIfError(errCategoryId)

	categoryResponse := controller.categoryService.FindById(request.Context(),int32(categoryId))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.categoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
