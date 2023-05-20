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
	utils.LoadEnv()

	// expect cmd like: go run main.go -m test
	mode := flag.String("m", "dev", "mode flag to change db to use")
	flag.Parse()

	var (
		entClient *ent.Client
		err       error
	)
	if *mode == "test" {
		entClient, err = db.ConnectTestDB()
	} else {
		entClient, err = db.ConnectDB()
	}
	if err != nil {
		os.Exit(2)
	}
	defer entClient.Close()

	router := route.SetUpRouter(entClient)

	router.Run(":8080")
}
