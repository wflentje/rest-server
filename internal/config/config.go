package config

type Config struct {
	Server  ServerConfig
	Logging LoggingConfig
}

func Load() (Config, error) {
	cfg := Default()

	if err := loadServer(&cfg); err != nil {
		return Config{}, err
	}

	if err := loadLogging(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
