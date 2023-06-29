package route

import (
	"app/ent"
	"app/handlers"
	"app/middlewares"
	"app/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(entClient *ent.Client, redisStore redis.Store) *gin.Engine {
	router := gin.Default()
	service := services.New(entClient)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	// set session instance to gin context
	router.Use(sessions.Sessions("mysession", redisStore))

	router.POST("/query", middlewares.AuthMiddleware(), handlers.GqlHandler(service))
	router.GET("/gqlplayground", handlers.PlaygroundHandler())

	return router
}
