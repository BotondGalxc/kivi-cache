package cmd

import (
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

func Logger(level slog.Leveler) *slog.Logger {
	if logger == nil {
		opts := &slog.HandlerOptions{
			Level: level,
		}
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
		slog.SetDefault(logger)
		logger.Debug("Set Default Log Level", "level", opts.Level)
	}

	return logger
}
