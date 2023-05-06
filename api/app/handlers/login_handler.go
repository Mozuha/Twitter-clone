package handlers

import (
	"api/controllers"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	loginService    services.LoginService       = services.NewLoginService()
	jwtService      services.JWTService         = services.NewJWTService()
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func LoginHandler(ctx *gin.Context) {
	token, err := loginController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, nil)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
