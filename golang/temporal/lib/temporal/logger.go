package temporal

import (
	"context"
	"fmt"
	"log/slog"

	"lib/logger"
)

// TemporalLogger wraps lib/logger to implement Temporal's log.Logger interface
type TemporalLogger struct {
	ctx context.Context
	log *slog.Logger
}

// NewTemporalLogger creates a new logger wrapper for Temporal
func NewTemporalLogger(ctx context.Context, log *slog.Logger) *TemporalLogger {
	return &TemporalLogger{
		ctx: ctx,
		log: log,
	}
}

// Debug logs a debug message
func (l *TemporalLogger) Debug(msg string, keyvals ...interface{}) {
	fields := convertKeyValsToFields(keyvals)
	l.log.Debug(msg, fields...)
}

// Info logs an info message
func (l *TemporalLogger) Info(msg string, keyvals ...interface{}) {
	logger.Info(l.ctx, l.log, msg, convertToLoggerFields(keyvals)...)
}

// Warn logs a warning message
func (l *TemporalLogger) Warn(msg string, keyvals ...interface{}) {
	logger.Warning(l.ctx, l.log, msg, convertToLoggerFields(keyvals)...)
}

// Error logs an error message
func (l *TemporalLogger) Error(msg string, keyvals ...interface{}) {
	logger.Error(l.ctx, l.log, msg, convertToLoggerFields(keyvals)...)
}

// convertKeyValsToFields converts key-value pairs to slog attributes
func convertKeyValsToFields(keyvals []interface{}) []interface{} {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "MISSING_VALUE")
	}
	return keyvals
}

// convertToLoggerFields converts key-value pairs to lib/logger Fields
func convertToLoggerFields(keyvals []interface{}) []*logger.Field {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "MISSING_VALUE")
	}

	fields := make([]*logger.Field, 0, len(keyvals)/2)
	for i := 0; i < len(keyvals); i += 2 {
		key := fmt.Sprintf("%v", keyvals[i])
		value := keyvals[i+1]
		fields = append(fields, logger.F(key, value))
	}

	return fields
}
