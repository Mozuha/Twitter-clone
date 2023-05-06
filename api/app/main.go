package main

import (
	"api/route"
	"api/utils"
)

func main() {
	utils.LoadEnv()

	router := route.SetUpRouter()

	router.Run(":8080")
}
