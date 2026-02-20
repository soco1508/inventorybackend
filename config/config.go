package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerHost string `envconfig:"SERVER_HOST"`
	ServerPort string `envconfig:"SERVER_PORT"`
}

type Database struct {
	Host     string `envocnfig:"DATABASE_HOST" required:"true"`
	Port     string `envocnfig:"DATABASE_PORT" required:"true"`
	Username string `envocnfig:"DATABASE_USERNAME" required:"true"`
	Password string `envocnfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envocnfig:"DATABASE_NAME" required:"true"`
}

func NewParsedConfig() (Config, error) {
	config := Config{}
	if err := godotenv.Load(".env"); err != nil {
		return config, fmt.Errorf("cannot load env, err:\n %+v", err)
	}
	if err := envconfig.Process("", &config); err != nil {
		return config, fmt.Errorf("cannot process env, err:\n %+v", err)
	}
	return config, nil
}
