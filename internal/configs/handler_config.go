package configs

type serverMode int

const (
	release serverMode = iota + 1
	debug
)

type serverConfig struct {
	ListenAddr string     `yaml:"address"`
	ListenPort string     `yaml:"port"`
	ListenType string     `yaml:"type"`
	Mode       serverMode `yaml:"mode"`
}

func ValidateServerMode() {
	addGlobalConfigOption(func(cfg *Config) {
		switch cfg.HandlerConfig.Mode {
		case release, debug:
		default:
			cfg.HandlerConfig.Mode = release
		}
	})
}
