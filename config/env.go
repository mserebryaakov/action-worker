package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	Debug               bool
	WorkerFrequency     uint64
	RestRabbitUrl       string
	RestRabbitRouteCode string
	RestRabbitRoutePass string
	ElmaUrl             string
	ElmaToken           string
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

	restRabbitRouteCode := lookupEnv("RABBITROUTECODE", "")
	if restRabbitRouteCode == "" {
		return AppEnv{}, errors.New("RABBITROUTECODE is not a correct")
	}

	restRabbitRoutePass := lookupEnv("RABBITROUTEPASS", "")
	if restRabbitRoutePass == "" {
		return AppEnv{}, errors.New("RABBITROUTEPASS is not a correct")
	}

	elmaUrl := lookupEnv("ELMAURL", "")
	if elmaUrl == "" {
		return AppEnv{}, errors.New("ELMAURL is not a correct")
	}

	env = AppEnv{
		Debug:               debugResult,
		WorkerFrequency:     workerFreqResult,
		RestRabbitUrl:       restRabbitUrl,
		RestRabbitRouteCode: restRabbitRouteCode,
		RestRabbitRoutePass: restRabbitRoutePass,
		ElmaUrl:             elmaUrl,
		ElmaToken:           lookupEnv("ELMATOKEN", ""),
	}

	return env, err
}

func lookupEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
