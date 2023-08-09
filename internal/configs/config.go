package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type configOptions func(*Config)

type Config struct {
	HandlerConfig serverConfig `yaml:"server"`
	LoggerConfig  logConfig    `yaml:"logger"`
}

var globalConfigOptions []configOptions

func addGlobalConfigOption(opt configOptions) {
	globalConfigOptions = append(globalConfigOptions, opt)
}

func NewConfig(configPath string) *Config {
	cfg := new(Config)

	ValidateServerMode()

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil
	}

	for _, opt := range globalConfigOptions {
		opt(cfg)
	}

	return cfg
}
