package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	categoryResponse := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}