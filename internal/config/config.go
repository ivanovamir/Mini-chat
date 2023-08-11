package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"mini_chat/pkg/logger"
	"mini_chat/pkg/postgres"
	"os"
)

type configOptions func(*Config)

type Config struct {
	HandlerConfig    ServerConfig              `yaml:"server"`
	AppVersion       AppConfig                 `yaml:"app"`
	LoggerConfig     logger.LoggerConfig       `yaml:"logger"`
	PostgresDBConfig postgres.PostgresDBConfig `yaml:"postgres"`
}

var globalConfigOptions []configOptions

func addGlobalConfigOption(opt configOptions) {
	globalConfigOptions = append(globalConfigOptions, opt)
}

func NewConfig() *Config {
	cfg := new(Config)

	if err := getEnv(); err != nil {
		return nil
	}

	validateServerMode()

	if err := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), cfg); err != nil {
		return nil
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil
	}

	for _, opt := range globalConfigOptions {
		opt(cfg)
	}

	return cfg
}

func getEnv() error {
	return godotenv.Load(".env")
}
