package route

import (
	"api/handlers"
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// Authentication + Token creation
	router.POST("/login", handlers.LoginHandler)

	// JWT auth middleware applies to "/api" only
	apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/users", handlers.GetUsersHandler)

		apiRoutes.GET("/users/:id", handlers.GetUserByIdHandler)

		apiRoutes.POST("/user", handlers.CreateUserHandler)
	}

	return router
}
