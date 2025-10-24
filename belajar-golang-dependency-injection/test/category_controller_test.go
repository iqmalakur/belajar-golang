package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"iqmalakur/belajar-golang-dependency-injection/app"
	"iqmalakur/belajar-golang-dependency-injection/controller"
	"iqmalakur/belajar-golang-dependency-injection/helper"
	"iqmalakur/belajar-golang-dependency-injection/middleware"
	"iqmalakur/belajar-golang-dependency-injection/model/domain"
	"iqmalakur/belajar-golang-dependency-injection/repository"
	"iqmalakur/belajar-golang-dependency-injection/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_dependency_injection_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
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

func createCategory(db *sql.DB) domain.Category {
	tx, _ := db.Begin()
	defer tx.Commit()

	repository := repository.NewCategoryRepository()
	return repository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
}

func decodeResponse(body io.Reader) map[string]any {
	response := map[string]any{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	helper.PanicIfError(err)

	response["code"] = int(response["code"].(float64))

	return response
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	requestBody := strings.NewReader(`{"name": "Gadget"}`)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusOK, response["code"])
	assert.Equal(t, http.StatusText(http.StatusOK), response["status"])

	data := response["data"].(map[string]any)

	assert.Equal(t, "Gadget", data["name"])
	assert.NotZero(t, data["id"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	requestBody := strings.NewReader(`{}`)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusBadRequest, response["code"])
	assert.Equal(t, http.StatusText(http.StatusBadRequest), response["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	requestBody := strings.NewReader(`{"name": "Fashion"}`)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusOK, response["code"])
	assert.Equal(t, http.StatusText(http.StatusOK), response["status"])

	data := response["data"].(map[string]any)

	assert.Equal(t, "Fashion", data["name"])
	assert.Equal(t, category.Id, int(data["id"].(float64)))
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	requestBody := strings.NewReader(`{}`)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusBadRequest, response["code"])
	assert.Equal(t, http.StatusText(http.StatusBadRequest), response["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusOK, response["code"])
	assert.Equal(t, http.StatusText(http.StatusOK), response["status"])

	data := response["data"].(map[string]any)

	assert.Equal(t, category.Name, data["name"])
	assert.Equal(t, category.Id, int(data["id"].(float64)))
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/categories/0", nil)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusNotFound, response["code"])
	assert.Equal(t, http.StatusText(http.StatusNotFound), response["status"])
	assert.Equal(t, "category is not found", response["data"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	category := createCategory(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusOK, response["code"])
	assert.Equal(t, http.StatusText(http.StatusOK), response["status"])
	assert.Equal(t, nil, response["data"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/api/categories/0", nil)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusNotFound, response["code"])
	assert.Equal(t, http.StatusText(http.StatusNotFound), response["status"])
	assert.Equal(t, "category is not found", response["data"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()

	category1 := repository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	category2 := repository.Save(context.Background(), tx, domain.Category{
		Name: "Food",
	})

	tx.Commit()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	request.Header.Set("X-API-KEY", "RAHASIA")

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusOK, response["code"])
	assert.Equal(t, http.StatusText(http.StatusOK), response["status"])

	categoryResponse1 := response["data"].([]any)[0].(map[string]any)
	categoryResponse2 := response["data"].([]any)[1].(map[string]any)

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	response := decodeResponse(result.Body)

	assert.Equal(t, http.StatusUnauthorized, response["code"])
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), response["status"])
	assert.Equal(t, nil, response["data"])
}
