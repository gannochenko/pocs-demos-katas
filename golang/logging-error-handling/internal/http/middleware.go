package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"loggingerrorhandling/internal/logger"
	"loggingerrorhandling/internal/syserr"
)

func ResponseWriter(controllerFn func(w http.ResponseWriter, r *http.Request) ([]byte, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody, err := controllerFn(w, r)

		httpStatus := http.StatusOK

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			httpStatus = mapErrorToHTTPStatus(err)

			w.WriteHeader(httpStatus)
			responseBody, _ := json.Marshal(&ErrorResponse{
				Error: err.Error(),
			})
			_, _ = w.Write(responseBody)
		} else {
			w.WriteHeader(httpStatus)
			if len(responseBody) > 0 {
				_, _ = w.Write(responseBody)
			}
		}

		logRequest(r, err, httpStatus)
	})
}

func logRequest(r *http.Request, err error, httpStatus int) {
	loggerFn := logger.Info

	// todo: extract verb, path, query and request body and put into fields, along with the request ID and httpStatus
	// todo: read other fields from the context and also feed to the logging function

	if err != nil {
		loggerFn = logger.Error

		var systemError *syserr.Error
		ok := errors.As(err, &systemError)
		if ok {
			code := systemError.GetCode()

			switch code {
			case syserr.NotFound:
				loggerFn = logger.Warning
			case syserr.BadInput:
				loggerFn = logger.Warning
			}
		}
	}

	loggerFn(r.Context(), "request handled")
}

func mapErrorToHTTPStatus(err error) int {
	var systemError *syserr.Error
	ok := errors.As(err, &systemError)
	if ok {
		code := systemError.GetCode()

		switch code {
		case syserr.NotFound:
			return http.StatusNotFound
		case syserr.BadInput:
			return http.StatusBadRequest
		}
	}

	return http.StatusInternalServerError
}
