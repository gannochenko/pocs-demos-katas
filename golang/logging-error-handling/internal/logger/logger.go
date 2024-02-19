package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func Warning(ctx context.Context, message string) {
	logger.Warn(message)
}

func Error(ctx context.Context, message string) {
	logger.Error(message)
}

func Info(ctx context.Context, message string) {
	logger.Info(message)
}
