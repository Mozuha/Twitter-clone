package route

import (
	"app/ent"
	"app/handlers"
	"app/middlewares"
	"app/services"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(entClient *ent.Client) *gin.Engine {
	router := gin.Default()
	service := services.New(entClient)

	router.POST("/query", middlewares.JWTAuth(entClient), handlers.GqlHandler(service))
	router.GET("/gqlplayground", handlers.PlaygroundHandler())

	return router
}
