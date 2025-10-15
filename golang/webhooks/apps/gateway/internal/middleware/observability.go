package middleware

import (
	"gateway/internal/interfaces"

	"github.com/labstack/echo/v4"

	"lib/logger"
)

func ObservabilityMiddleware( monitoringService interfaces.MonitoringService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				status := 500

				if httpErr, ok := err.(*echo.HTTPError); ok { // todo: use errors.As()
					status = httpErr.Code
				}

				if status >= 500 {
					monitoringService.AddInt64Counter(c.Request().Context(), "http_errors", "total", 1, "", "")
				}
			}

			monitoringService.AddInt64Counter(c.Request().Context(), "request_count", "total", 1, "", "")

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
