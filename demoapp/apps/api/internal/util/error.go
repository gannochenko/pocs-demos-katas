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
		syserr.NotImplementedCode: http.StatusInternalServerError,
	}
)

type ErrorResponse struct {
	Message string `json:"message"`
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
				Message: err.Error(),
			}
			if code == syserr.InternalCode {
				response.Message = "internal error occurred"
			}

			writeErr := httpUtil.EncodeJSONResponse(response, ToPtr(httpCode), w)
			if writeErr != nil {
				logger.Error(r.Context(), "could not write the error response")
			}
		}

		return err
	}
}
