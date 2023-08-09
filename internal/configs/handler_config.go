package configs

type serverMode int

const (
	release serverMode = iota + 1
	debug
)

type ServerConfig struct {
	ListenAddr string     `yaml:"address"`
	ListenPort string     `yaml:"port"`
	ListenType string     `yaml:"type"`
	Mode       serverMode `yaml:"mode"`
}

func validateServerMode() {
	addGlobalConfigOption(func(cfg *Config) {
		switch cfg.HandlerConfig.Mode {
		case release, debug:
		default:
			cfg.HandlerConfig.Mode = release
		}
	})
}
