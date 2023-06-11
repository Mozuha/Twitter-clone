package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions/redis"
)

func SetUpRedisStore(runningEnv string) (redis.Store, error) {
	var (
		host string
		port string
	)

	if runningEnv == "docker" {
		host = os.Getenv("REDIS_HOST")
	} else {
		host = "localhost"
	}

	port = os.Getenv("REDIS_PORT")

	redisInfo := fmt.Sprintf("%s:%s", host, port)

	log.Println("opening connection to redis store...")
	store, err := redis.NewStore(10, "tcp", redisInfo, "", []byte(os.Getenv("REDIS_KEY")))
	if err != nil {
		log.Fatalf("failed opening connection to redis: %v", err)
		return nil, err
	}

	log.Println("connected to redis")

	// Setting options somehow prohibit the session from being registered on redis; comment it out for now
	// Idially, set MaxAge to be that of the refresh token
	// whether setting httponly to true depends on frontend
	// whether setting maxage depends on frontend (use cookie or not)
	// store.Options(sessions.Options{Path: "/", Secure: true, SameSite: http.SameSiteLaxMode})

	return store, nil
}
