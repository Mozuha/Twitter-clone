package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectRoot = "twitter-clone/"

var runningEnvironment string

func LoadEnv() (string, error) {
	runningEnvironment = "local"
	re := regexp.MustCompile(`^(.*` + projectRoot + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	// projectRoot match was not found; it's in go container either with env values set or with env file
	if rootPath == nil {
		// No need to load env file if this is running inside some environment where env values are already set
		if tmp := os.Getenv("API_CONTAINER_NAME"); tmp != "" {
			runningEnvironment = "docker"
			return runningEnvironment, nil
		}

		// Otherwise env file might be in this directory
		rootPath = make([]byte, 0)
		runningEnvironment = "docker"
	}

	if err := godotenv.Load(string(rootPath) + `.env`); err != nil {
		return "", fmt.Errorf("Failed to load env file: %v", err)
	}

	return runningEnvironment, nil
}
