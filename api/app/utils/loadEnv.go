package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectRoot = "twitter-clone"

func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectRoot + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}
}
