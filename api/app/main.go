package main

import (
	"api/db"
	"api/route"
	"api/utils"
	"os"
)

func main() {
	runningEnv := utils.LoadEnv()

	entClient, err := db.NewEntClient(runningEnv)
	if err != nil {
		os.Exit(2)
	}
	defer entClient.Close()

	router := route.SetUpRouter()

	router.Run(":8080")
}
