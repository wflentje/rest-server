package config

import (
	"fmt"
	"os"
	"time"
)

type ServerConfig struct {
	Address           string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func Default() Config {
	return Config{
		Server: ServerConfig{
			Address:           ":8080",
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
		Logging: LoggingConfig{
			Level:  "INFO",
			Format: "text",
			Output: "stdout",
		},
	}
}

func loadServer(cfg *Config) error {
	cfg.Server.Address = getEnv("SERVER_ADDRESS", cfg.Server.Address)

	var err error

	cfg.Server.ReadTimeout, err = getDuration("SERVER_READ_TIMEOUT", cfg.Server.ReadTimeout)
	if err != nil {
		return err
	}

	cfg.Server.ReadHeaderTimeout, err = getDuration("SERVER_READ_HEADER_TIMEOUT", cfg.Server.ReadHeaderTimeout)
	if err != nil {
		return err
	}

	cfg.Server.WriteTimeout, err = getDuration("SERVER_WRITE_TIMEOUT", cfg.Server.WriteTimeout)
	if err != nil {
		return err
	}

	cfg.Server.IdleTimeout, err = getDuration("SERVER_IDLE_TIMEOUT", cfg.Server.IdleTimeout)
	if err != nil {
		return err
	}

	return validateServer(cfg.Server)
}

func validateServer(cfg ServerConfig) error {
	if cfg.Address == "" {
		return fmt.Errorf("server address is required")
	}

	if cfg.ReadTimeout <= 0 {
		return fmt.Errorf("server read timeout must be greater than zero")
	}

	if cfg.ReadHeaderTimeout <= 0 {
		return fmt.Errorf("server read header timeout must be greater than zero")
	}

	if cfg.WriteTimeout <= 0 {
		return fmt.Errorf("server write timeout must be greater than zero")
	}

	if cfg.IdleTimeout <= 0 {
		return fmt.Errorf("server idle timeout must be greater than zero")
	}

	return nil
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getDuration(key string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("invalid duration for %s: %w", key, err)
	}

	return duration, nil
}
