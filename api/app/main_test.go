package main

import (
	"api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetUsersHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users", GetUsersHandler)
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}

func TestGetUserByIdHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users/:id", GetUserByIdHandler)
	req, _ := http.NewRequest("GET", "/users/3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
}

func TestGetUserByIdHandlerNotFound(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users/:id", GetUserByIdHandler)
	req, _ := http.NewRequest("GET", "/users/100", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateUserHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.POST("/user", CreateUserHandler)
	user := models.User{
		DisplayName: "test4",
		Username:    "test 4",
		Email:       "test4@ymail.ne.jp",
		Password:    "12345",
	}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateUserHandlerBadRequest(t *testing.T) {
	r := SetUpRouter()
	r.POST("/user", CreateUserHandler)

	// lacking Username
	user := models.User{
		DisplayName: "test4",
		Email:       "test4@ymail.ne.jp",
		Password:    "12345",
	}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
