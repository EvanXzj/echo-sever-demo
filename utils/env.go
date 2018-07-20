package utils

import (
	"os"
)

// GetEnv : get an environment variable or its default value
func GetEnv(k string) string {
	env := os.Getenv(k)

	// default fallbacks
	if len(env) == 0 {
		switch k {
		case "APP_PORT":
			return "7447"
		case "JWT_SECRET":
			return "verygoodsecret"
		default:
			panic("unrecognized env variable")
		}
	}

	return env
}
