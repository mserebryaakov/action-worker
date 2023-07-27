package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	Debug           bool
	WorkerFrequency uint64
	RestRabbitUrl   string
}

func GetEnv() (env AppEnv, err error) {
	_ = godotenv.Load()

	debugEnv := lookupEnv("DEBUG", "false")
	debugResult := false
	if debugEnv == "true" {
		debugResult = true
	}

	workerFreqEnv := lookupEnv("WORKERFREQUENCY", "")
	workerFreqResult, err := strconv.ParseUint(workerFreqEnv, 10, 64)
	if err != nil || workerFreqResult < 1 {
		return AppEnv{}, errors.New("WORKERFREQUENCY is not a correct")
	}

	restRabbitUrl := lookupEnv("RABBITURL", "")
	if restRabbitUrl == "" {
		return AppEnv{}, errors.New("RABBITURL is not a correct")
	}

	env = AppEnv{
		Debug:           debugResult,
		WorkerFrequency: workerFreqResult,
		RestRabbitUrl:   restRabbitUrl,
	}

	return env, err
}

func lookupEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
