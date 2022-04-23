package config

import (
	"MafiaREST/utils"
	"os"
)

func setEnvConfigParam(name, defaultVal string, enableDefault bool) string {
	val := os.Getenv(name)

	if enableDefault && val == "" {
		err := os.Setenv(name, defaultVal)
		utils.PanicOnError("Failed to set env", err)
		val = os.Getenv(name)
	}

	return val
}
