package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	Debug bool
}

func GetEnv() (env AppEnv, err error) {
	_ = godotenv.Load()

	debugEnv := lookupEnv("DEBUG", "false")
	debugResult := false
	if debugEnv == "true" {
		debugResult = true
	}

	env = AppEnv{
		Debug: debugResult,
	}

	return env, err
}

func lookupEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
