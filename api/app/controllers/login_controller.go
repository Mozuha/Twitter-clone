package controllers

// import (
// 	"app/dto"
// 	"app/services"
// 	"errors"

// 	"github.com/gin-gonic/gin"
// )

// type LoginController interface {
// 	Login(ctx *gin.Context) (string, error)
// }

// type loginController struct {
// 	loginService services.LoginService
// 	jwtService   services.JWTService
// }

// func NewLoginController(loginService services.LoginService, jwtService services.JWTService) LoginController {
// 	return &loginController{
// 		loginService: loginService,
// 		jwtService:   jwtService,
// 	}
// }

// func (controller *loginController) Login(ctx *gin.Context) (string, error) {
// 	var credentials dto.EmailLoginRequest
// 	err := ctx.ShouldBindJSON(&credentials)
// 	if err != nil {
// 		return "", err
// 	}

// 	// TODO: hash password at here using bcrypt

// 	if isAuthenticated := controller.loginService.Login(credentials.Email, credentials.Password); isAuthenticated {
// 		return controller.jwtService.GenerateToken(credentials.Email)
// 	} else {
// 		return "", errors.New("Authentication failed")
// 	}
// }
