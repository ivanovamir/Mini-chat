package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type configOptions func(*Config)

type Config struct {
	HandlerConfig ServerConfig `yaml:"server"`
	LoggerConfig  LoggerConfig `yaml:"logger"`
	AppVersion    AppConfig    `yaml:"app"`
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

	for _, opt := range globalConfigOptions {
		opt(cfg)
	}

	return cfg
}

func getEnv() error {
	return godotenv.Load(".env")
}
