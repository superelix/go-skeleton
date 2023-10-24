package common

import (
	"github.com/joho/godotenv"
)

var EnvMap map[string]string

func init() {
	err := godotenv.Load(ProjectEnvPath() + ".env")
	if err != nil {
		GetLogger().Fatalf("Failed to load the environment variables, %s", err)
		panic(err)
	}

	EnvMap, err = godotenv.Read()
	if err != nil {
		GetLogger().Fatalf("Failed to read the environment variables, %s", err)
		panic(err)
	}
}
