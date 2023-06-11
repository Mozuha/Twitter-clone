package main

import (
	"app/db"
	"app/ent"
	"app/route"
	"app/utils"
	"flag"
	"os"
)

func main() {
	runningEnv, err := utils.LoadEnv()
	if err != nil {
		os.Exit(2)
	}

	// expect cmd like: go run main.go -m test
	mode := flag.String("m", "dev", "mode flag to change db to use")
	flag.Parse()

	var entClient *ent.Client
	if *mode == "test" {
		entClient, err = db.ConnectTestDB(runningEnv)
	} else {
		entClient, err = db.ConnectDB(runningEnv)
	}
	if err != nil {
		os.Exit(2)
	}
	defer entClient.Close()

	redisStore, err := db.SetUpRedisStore(runningEnv)
	if err != nil {
		os.Exit(2)
	}

	router := route.SetUpRouter(entClient, redisStore)

	router.Run(":8080")
}
