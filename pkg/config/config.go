package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBConfig
}

type DBConfig struct {
	DBHost string `envconfig:"DB_HOST" default:"db"`
	DBPort string `envconfig:"DB_PORT" default:"3306"`
	DBUser string `envconfig:"DB_USERNAME" default:"torasemi"`
	DBPass string `envconfig:"DB_PASSWORD" default:"PR85UafO"`
	DBName string `envconfig:"DB_DATABASE" default:"torasemi-todo-api-local"`
}

func NewConfigFromEnv() Config {
	env := Config{}
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("action=envconfig.Process, status=error: %v", err)
	}
	return env
}
