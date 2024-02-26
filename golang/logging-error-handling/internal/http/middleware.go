package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	pkgContext "loggingerrorhandling/internal/context"
	"loggingerrorhandling/internal/logger"
	"loggingerrorhandling/internal/syserr"
)

const (
	OperationIDHeaderName = "X-Operation-Id"
)

func ResponseWriter(controllerFn func(w http.ResponseWriter, r *http.Request) ([]byte, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(pkgContext.WithOperationID(r.Context(), extractOperationID(r)))

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

	fields := make([]*logger.Field, 3)
	fields[0] = logger.NewFiled("method", r.Method)
	fields[1] = logger.NewFiled("url", fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery))
	fields[2] = logger.NewFiled("status", httpStatus)

	message := "request handled"

	// todo: we can also log the request body here, if needed

	if err != nil {
		message = fmt.Sprintf("%s with error: %s", message, err.Error())

		loggerFn = mapErrorToLoggerFunction(err)

		stack := make([]string, 0)

		var systemError *syserr.Error
		ok := errors.As(err, &systemError)
		if ok {
			for _, field := range systemError.GetFields() {
				fields = append(fields, logger.NewFiled(field.Key, field.Value))
			}

			stack = formatErrorStack(systemError.GetStack())
		} else {
			stack = getErrorStackFormatted(err)
		}

		fields = append(fields, logger.NewFiled("stack", stack))
	}

	loggerFn(r.Context(), message, fields...)
}

func mapErrorToHTTPStatus(err error) int {
	var systemError *syserr.Error
	ok := errors.As(err, &systemError)
	if ok {
		code := systemError.GetCode()

		switch code {
		case syserr.NotFoundCode:
			return http.StatusNotFound
		case syserr.BadInputCode:
			return http.StatusBadRequest
		}
	}

	return http.StatusInternalServerError
}

func mapErrorToLoggerFunction(err error) func(ctx context.Context, message string, fields ...*logger.Field) {
	var systemError *syserr.Error
	ok := errors.As(err, &systemError)
	if ok {
		code := systemError.GetCode()

		switch code {
		case syserr.InternalCode:
			return logger.Error
		case syserr.NotFoundCode:
			return logger.Warning
		case syserr.BadInputCode:
			return logger.Warning
		}
	}

	return logger.Error
}

func getErrorStackFormatted(e error) []string {
	return formatErrorStack(syserr.GetStack(e))
}

func formatErrorStack(stack []*syserr.ErrorStackItem) []string {
	result := make([]string, len(stack))

	for index, item := range stack {
		result[index] = fmt.Sprintf("%s:%s %s", item.File, item.Line, item.Function)
	}

	return result
}

func extractOperationID(r *http.Request) string {
	headers := r.Header
	for name, value := range headers {
		if name == OperationIDHeaderName {
			return value[0]
		}
	}

	return ""
}
