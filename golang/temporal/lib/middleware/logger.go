package middleware

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"lib/ctx"
	"lib/logger"
)

func LoggerMiddleware(log *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			
			operationID := uuid.New().String()

			ctx := ctx.WithOperationID(c.Request().Context(), operationID)
			c.SetRequest(c.Request().WithContext(ctx))

			c.Response().Header().Set("X-Operation-ID", operationID)

			err := next(c)
			
			duration := time.Since(start).Milliseconds()

			if err != nil {
				loggerFn := logger.Error
				message := err.Error()
				status := 500

				if httpErr, ok := err.(*echo.HTTPError); ok { // todo: use errors.As()
					loggerFn = getLoggerByHTTPCode(httpErr.Code)
				}

				loggerFn(ctx, log, fmt.Sprintf("HTTP request failed: %s", message),
					logger.F("method", c.Request().Method),
					logger.F("path", c.Request().URL.Path),
					logger.F("duration_ms", duration),
					logger.F("status", status),
				)
			} else {
				// Get response status
				status := c.Response().Status
				if status == 0 {
					status = 200 // Default status if not set
				}
				
				// Log the request
				logger.Info(ctx, log, "HTTP request processed",
					logger.F("method", c.Request().Method),
					logger.F("path", c.Request().URL.Path),
					logger.F("status", status),
					logger.F("duration_ms", duration),
				)
			}
			
			return err
		}
	}
}

func getLoggerByHTTPCode(status int) logger.LoggerFn {
	switch {
	case status >= 500:
		return logger.Error
	case status >= 400:
		return logger.Warning
	case status >= 300:
		return logger.Info
	default:
		return logger.Info
	}
}
