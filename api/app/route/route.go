package route

import (
	"app/ent"
	"app/handlers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(entClient *ent.Client) *gin.Engine {
	router := gin.Default()
	userHandlers := handlers.NewUsersHandler(entClient)

	// Authentication + Token creation
	router.POST("/login", handlers.LoginHandler)

	// JWT auth middleware applies to "/api" only
	apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/users", userHandlers.GetUsersHandler)

		apiRoutes.GET("/users/:id", userHandlers.GetUserByIdHandler)

		apiRoutes.POST("/user", userHandlers.CreateUserHandler)
	}

	return router
}
