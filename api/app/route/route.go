package route

import (
	"app/ent"
	"app/handlers"
	"app/middlewares"
	"app/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(entClient *ent.Client, redisStore redis.Store) *gin.Engine {
	router := gin.Default()
	service := services.New(entClient)

	// set session instance to gin context
	router.Use(sessions.Sessions("mysession", redisStore))

	router.POST("/query", middlewares.AuthMiddleware(), handlers.GqlHandler(service))
	router.GET("/gqlplayground", handlers.PlaygroundHandler())

	return router
}
