package config

import (
	"errors"
	"fmt"
	"os"
)

const envName = "ENVIRONMENT_TYPE"

func GetEnvType() (string, error) {
	variable, exists := os.LookupEnv(envName)
	var confFile string = "./config/dev.yml"

	if !exists {
		return "./config/dev.yml", errors.New("ENVIRONMENT_TYPE not defined")
	}

	switch variable {
	case "PROD":
		confFile = "./config/prod.yml"
	case "TEST":
		confFile = "./config/test.yml"
	case "DEV":
		confFile = "./config/dev.yml"
	case "AUTOTEST":
		confFile = "../../config/dev.yml"
	}

	if _, err := os.Stat(confFile); errors.Is(err, os.ErrNotExist) {
		return "", errors.New(fmt.Sprintf("%s does not exist ; env type: %s", confFile, variable))
	}

	return confFile, nil
}
