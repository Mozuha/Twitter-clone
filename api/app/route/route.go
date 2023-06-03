package route

import (
	"app/ent"
	"app/handlers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(entClient *ent.Client) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.AuthorizeJWT(entClient))
	router.POST("/query", handlers.GqlHandler(entClient))
	router.GET("/gqlplayground", handlers.PlaygroundHandler())

	return router
}
