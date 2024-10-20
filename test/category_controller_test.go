package test

import (
	"context"
	"database/sql"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/model/domain"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func SetupDB() *sql.DB{
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler{
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := SetupDB()
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": "laptop"}`)
	request := httptest.NewRequest("POST", "http://localhost:8080/api/category",requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T) {
	db := SetupDB()
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest("POST", "http://localhost:8080/api/category",requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	category := repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": "HAPEEE"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/category/" + strconv.Itoa(category.Id),requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	category := repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/category/" + strconv.Itoa(category.Id),requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestGetCategorySuccess(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	category := repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/category/" + strconv.Itoa(category.Id),nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestGetCategoryFailed(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/category/" + strconv.Itoa(4),nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	category := repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/category/" + strconv.Itoa(category.Id),nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/category/" + strconv.Itoa(6),nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestUnauthorized(t *testing.T) {
	db := SetupDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repoCategory := repository.NewCategoryRepository()
	category := repoCategory.Save(context.Background(), tx, domain.Category{
		Name: "hape",
	})
	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/category/" + strconv.Itoa(category.Id),nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}