package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// detect running environment (docker or local) for choosing appropriate db host name
var runningEnv string

const projectRoot = "twitter-clone"

func LoadEnv() string {
	// No need to load env values if this is running inside docker container as they're already set
	if tmp := os.Getenv("API_PORT"); tmp != "" {
		runningEnv = "docker"
		return runningEnv
	}

	runningEnv = "local"

	re := regexp.MustCompile(`^(.*` + projectRoot + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}

	return runningEnv
}
