package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"rest-server/internal/config"
)

func NewLogger(cfg config.LoggingConfig) (*slog.Logger, error) {
	level, err := parseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	output, err := parseOutput(cfg.Output)
	if err != nil {
		return nil, err
	}

	options := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler

	switch cfg.Format {
	case "text":
		handler = slog.NewTextHandler(output, options)
	case "json":
		handler = slog.NewJSONHandler(output, options)
	default:
		return nil, fmt.Errorf("invalid log format %q", cfg.Format)
	}

	return slog.New(handler), nil
}

func parseLevel(level string) (slog.Level, error) {
	switch level {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid log level %q", level)
	}
}

func parseOutput(output string) (io.Writer, error) {
	switch output {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return nil, fmt.Errorf("invalid log output %q", output)
	}
}
