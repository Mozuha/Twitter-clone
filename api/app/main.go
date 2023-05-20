package main

import (
	"api/db"
	"api/ent"
	"api/route"
	"api/utils"
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

	router := route.SetUpRouter(entClient)

	router.Run(":8080")
}
