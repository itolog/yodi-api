package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/itolog/yodi-api/src/pkg/logging"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     string `env:"PORT" env-default:"8000"`
	Host     string `env:"HOST" env-default:"localhost"`
	PrefixV1 string `env:"PREFIX_V1" env-default:"/api/v1"`
}

const (
	DEV  = "development"
	PROD = "production"
)

func NewConfig() *Config {
	log := logging.GetLogger()
	appEnv := os.Getenv("APP_YP_ENV")

	if appEnv == PROD {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else if appEnv == DEV {
		err := godotenv.Load(".env.development")
		if err != nil {
			log.Fatal("Error loading .env.development file")
		}
	}

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Warn(err)
	}

	return &Config{
		Port:     cfg.Port,
		Host:     cfg.Host,
		PrefixV1: cfg.PrefixV1,
	}
}
