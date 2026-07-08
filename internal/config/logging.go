package config

import (
	"fmt"
)

type LoggingConfig struct {
	Level  string
	Format string
	Output string
}

func loadLogging(cfg *Config) error {
	cfg.Logging.Level = getEnv("LOG_LEVEL", cfg.Logging.Level)
	cfg.Logging.Format = getEnv("LOG_FORMAT", cfg.Logging.Format)
	cfg.Logging.Output = getEnv("LOG_OUTPUT", cfg.Logging.Output)

	return validateLogging(cfg.Logging)
}

func validateLogging(cfg LoggingConfig) error {
	switch cfg.Level {
	case "DEBUG", "INFO", "WARN", "ERROR":
	default:
		return fmt.Errorf("invalid log level %q", cfg.Level)
	}

	switch cfg.Format {
	case "text", "json":
	default:
		return fmt.Errorf("invalid log format %q", cfg.Format)
	}

	switch cfg.Output {
	case "stdout", "stderr":
	default:
		return fmt.Errorf("invalid log output %q", cfg.Output)
	}

	return nil
}
