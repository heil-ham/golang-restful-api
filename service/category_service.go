package service

import (
	"context"
	"golang-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int32) 
	FindById(ctx context.Context, categoryId int32) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}