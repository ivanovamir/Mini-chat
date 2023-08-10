package configs

type LoggerConfig struct {
	Path  string `yaml:"path"`
	Level int    `yaml:"level"`
}
