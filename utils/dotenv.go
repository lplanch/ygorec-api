package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {
	env := make(chan string, 1)

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") != "production" {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	return <-env
}
