package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort string `envconfig:"DB_PORT" default:"3306"`
	DBUser string `envconfig:"DB_USER" default:"torasemi"`
	DBPass string `envconfig:"DB_PASS" default:"PR85UafO"`
	DBName string `envconfig:"DB_NAME" default:"torasemi-todo-api-local"`
}

func NewConfigFromEnv() Config {
	env := Config{}
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("action=envconfig.Process, status=error: %v", err)
	}
	return env
}
