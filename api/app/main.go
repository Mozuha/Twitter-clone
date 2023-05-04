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

func main() {
	router := gin.Default()

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.GetUsers())
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		user, err := userController.GetUserById(ctx)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": err.Error()})
		}
		ctx.JSON(http.StatusOK, user)
	})

	router.POST("/user", func(ctx *gin.Context) {
		user, err := userController.CreateUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusCreated, user)
	})

	router.Run("localhost:8080")
}
