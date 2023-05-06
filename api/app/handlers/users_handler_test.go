package handlers

import (
	"api/middlewares"
	"api/models"
	"api/services"
	"api/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UsersHandlerTestSuite struct {
	suite.Suite
	r           *gin.Engine
	authedToken string
}

func (s *UsersHandlerTestSuite) SetupTest() {
	utils.LoadEnv()
	gin.SetMode(gin.TestMode)

	s.r = gin.Default()
	apiRoutes := s.r.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/users", GetUsersHandler)

		apiRoutes.GET("/users/:id", GetUserByIdHandler)

		apiRoutes.POST("/user", CreateUserHandler)
	}

	s.authedToken, _ = services.NewJWTService().GenerateToken("test1@gmail.com")
}

func TestUsersHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UsersHandlerTestSuite))
}

func (s *UsersHandlerTestSuite) TestGetUsersHandlerSuccess() {
	req, _ := http.NewRequest("GET", "/api/users", nil)
	req.Header.Add("Authorization", "Bearer "+s.authedToken)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	s.Equal(http.StatusOK, w.Code)
	s.NotEmpty(users)
}

func (s *UsersHandlerTestSuite) TestGetUserByIdHandlerSuccess() {
	req, _ := http.NewRequest("GET", "/api/users/3", nil)
	req.Header.Add("Authorization", "Bearer "+s.authedToken)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	s.Equal(http.StatusOK, w.Code)
	s.NotEmpty(user)
}

func (s *UsersHandlerTestSuite) TestGetUserByIdHandlerNotFound() {
	req, _ := http.NewRequest("GET", "/api/users/100", nil)
	req.Header.Add("Authorization", "Bearer "+s.authedToken)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	s.Equal(http.StatusNotFound, w.Code)
}

func (s *UsersHandlerTestSuite) TestCreateUserHandlerSuccess() {
	user := models.User{
		DisplayName: "test4",
		Username:    "test 4",
		Email:       "test4@ymail.ne.jp",
		Password:    "12345",
	}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+s.authedToken)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	s.Equal(http.StatusCreated, w.Code)
}

func (s *UsersHandlerTestSuite) TestCreateUserHandlerBadRequest() {
	// lacking Username
	user := models.User{
		DisplayName: "test4",
		Email:       "test4@ymail.ne.jp",
		Password:    "12345",
	}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+s.authedToken)
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}
