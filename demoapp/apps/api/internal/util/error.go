package util

import (
	"net/http"

	httpUtil "api/internal/util/http"
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

func withErrorHandler(h AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := h(w, r)
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

			_ = httpUtil.EncodeJSONResponse(response, ToPtr(httpCode), w)
		}

		return err
	}
}
