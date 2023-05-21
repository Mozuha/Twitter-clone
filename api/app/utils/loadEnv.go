package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// running in local: twitter-clone/.env
// running in go container: foo/app/.env
const projectRoot = "twitter-clone/"
const goAppRoot = "app/"

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
		re = regexp.MustCompile(`^(.*` + goAppRoot + `)`)
		rootPath = re.Find([]byte(cwd))
		runningEnvironment = "docker"
	}

	if err := godotenv.Load(string(rootPath) + `.env`); err != nil {
		return "", fmt.Errorf("Failed to load env file: %v", err)
	}

	return runningEnvironment, nil
}
