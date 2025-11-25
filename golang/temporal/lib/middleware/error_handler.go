package middleware

import (
	"lib/logger"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(log *slog.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var (
			code = http.StatusInternalServerError
			msg  interface{}
		)

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message
		} else {
			msg = err.Error()
		}

		// Avoid leaking internal error messages to the client
		if code >= 500 || msg == err.Error() {
			msg = "Internal server error"
		}

		// Send JSON response
		if !c.Response().Committed {
			var err error

			if c.Request().Header.Get("Content-Type") == "application/json" {
				err = c.JSON(code, map[string]interface{}{
					"error": msg,
				})
			} else {
				err = c.String(code, "Internal server error")
			}

			if err != nil {
				logger.Error(c.Request().Context(), log, err.Error())
			}
		}
	}
}
