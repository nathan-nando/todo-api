package env

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func LoadEnv(env string) {
	cwd, _ := os.Getwd()

	var envFile string

	switch env {
	case "LOCAL":
		envFile = "local"
	case "STAGING":
		envFile = "staging"
	case "PRODUCTION":
		envFile = "product"
	default:
		envFile = ""
	}

	err := godotenv.Load(cwd + `/.env.` + envFile)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
			"cwd":   cwd,
		}).Fatal("Load .env error")
		os.Exit(-1)
	}
}
