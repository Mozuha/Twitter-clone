package main

import (
	"api/controller"
	"api/services"

	"github.com/gin-gonic/gin"
)

var (
	userService    services.UserService      = services.New()
	userController controller.UserController = controller.New(userService)
)

func main() {
	router := gin.Default()

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.GetUsers())
	})

	router.Run("localhost:8080")
}
