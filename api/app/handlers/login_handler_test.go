package handlers

import (
	"app/dto"
	"app/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type LoginHandlerTestSuite struct {
	suite.Suite
	r *gin.Engine
}

func (s *LoginHandlerTestSuite) SetupTest() {
	if _, err := utils.LoadEnv(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	gin.SetMode(gin.TestMode)

	s.r = gin.Default()
	s.r.POST("/login", LoginHandler)
}

func TestLoginHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(LoginHandlerTestSuite))
}

func (s *LoginHandlerTestSuite) TestLoginHandlerSuccess() {
	credential := dto.EmailLoginRequest{
		Email:    "test1@gmail.com",
		Password: "pass",
	}
	jsonBody, _ := json.Marshal(credential)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
}

func (s *LoginHandlerTestSuite) TestLoginHandlerUnauthorized() {
	credential := dto.EmailLoginRequest{
		Email:    "unauth@gmail.com",
		Password: "letmein",
	}
	jsonBody, _ := json.Marshal(credential)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)

	s.Equal(http.StatusUnauthorized, w.Code)
}
