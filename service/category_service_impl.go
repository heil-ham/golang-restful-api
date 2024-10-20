package service

import (
	"context"
	"database/sql"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository 	repository.CategoryRepository
	DB 					*sql.DB
	validate 			*validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	errValidate := service.validate.Struct(request)
	helper.PanicIfError(errValidate)
	
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidate := service.validate.Struct(request)
	helper.PanicIfError(errValidate)
	
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	category, errFindById := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if errFindById != nil {
		panic(exception.NewNotFoundError(errFindById.Error())) 
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int32) {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	category, errFindById := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if errFindById != nil {
		panic(exception.NewNotFoundError(errFindById.Error())) 
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int32) web.CategoryResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	category, errFindById := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if errFindById != nil {
		panic(exception.NewNotFoundError(errFindById.Error())) 
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}