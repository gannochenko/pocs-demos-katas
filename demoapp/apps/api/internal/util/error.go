package util

import (
	"net/http"

	"api/internal/types"
	httpUtil "api/internal/util/http"
	"api/pkg/logger"
	"api/pkg/syserr"
)

var (
	errorToHTTPError = map[syserr.Code]int{
		syserr.InternalCode:       http.StatusInternalServerError,
		syserr.BadInputCode:       http.StatusBadRequest,
		syserr.NotFoundCode:       http.StatusBadRequest,
		syserr.NotImplementedCode: http.StatusNotImplemented,
	}
)

type ErrorResponse struct {
	Error   string   `json:"error"`
	Reasons []string `json:"reasons"`
}

func WithErrorHandler(next types.Handler) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := next(w, r)
		if err != nil {
			code := syserr.GetCode(err)
			httpCode := errorToHTTPError[code]
			if httpCode == 0 {
				httpCode = http.StatusInternalServerError
			}

			response := ErrorResponse{
				Error: err.Error(),
			}
			if code == syserr.InternalCode {
				response.Error = "internal error occurred"
			}
			if code == syserr.BadInputCode {
				fields := syserr.GetFields(err)
				for _, field := range fields {
					if field.Key == "reasons" {
						for _, reason := range field.Value.([]string) {
							response.Reasons = append(response.Reasons, reason)
						}
					}
				}
			}

			writeErr := httpUtil.EncodeJSONResponse(response, ToPtr(httpCode), w)
			if writeErr != nil {
				logger.Error(r.Context(), "could not write the error response")
			}
		}

		return err
	}
}
