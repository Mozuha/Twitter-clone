package main

import (
	"api/controllers"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userService    services.UserService       = services.New()
	userController controllers.UserController = controllers.New(userService)
)

func GetUsersHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userController.GetUsers())
}

func GetUserByIdHandler(ctx *gin.Context) {
	user, err := userController.GetUserById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": err.Error()})
	}
	ctx.JSON(http.StatusOK, user)
}

func CreateUserHandler(ctx *gin.Context) {
	user, err := userController.CreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, user)
}

func main() {
	router := gin.Default()

	router.GET("/users", GetUsersHandler)

	router.GET("/users/:id", GetUserByIdHandler)

	router.POST("/user", CreateUserHandler)

	router.Run("localhost:8080")
}
