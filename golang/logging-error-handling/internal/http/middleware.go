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
		r = r.WithContext(pkgContext.WithOperationID(r.Context(), extractProvidedOperationID(r)))

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set(OperationIDHeaderName, pkgContext.GetOperationID(r.Context()))

		defer func() {
			if rec := recover(); rec != nil {
				httpStatus := http.StatusInternalServerError
				err := syserr.Internal(fmt.Sprintf("PANIC: %v", rec))

				writeError(httpStatus, extractPublicMessage(syserr.InternalCode, err), w)
				logRequest(r, err, httpStatus)
			}
		}()

		responseBody, err := controllerFn(w, r)

		httpStatus := http.StatusOK

		if err != nil {
			errorCode := extractErrorCode(err)
			httpStatus = mapErrorCodeToHTTPStatus(errorCode, err)
			writeError(httpStatus, extractPublicMessage(errorCode, err), w)
		} else {
			w.WriteHeader(httpStatus)
			if len(responseBody) > 0 {
				_, _ = w.Write(responseBody)
			}
		}

		logRequest(r, err, httpStatus)
	})
}

func writeError(httpStatus int, publicErrorMessage string, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)

	responseBody, _ := json.Marshal(&ErrorResponse{
		Error: publicErrorMessage,
	})
	_, _ = w.Write(responseBody)
}

func logRequest(r *http.Request, err error, httpStatus int) {
	loggerFn := logger.Info

	fields := make([]*logger.Field, 3)
	fields[0] = logger.NewFiled("method", r.Method)
	fields[1] = logger.NewFiled("url", fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery))
	fields[2] = logger.NewFiled("status", httpStatus)

	loggedMessage := "request handled"

	// todo: we can also log the request body here, if needed

	if err != nil {
		errorCode := extractErrorCode(err)

		loggedMessage = fmt.Sprintf("%s with error: %s", loggedMessage, err.Error())
		loggerFn = mapErrorCodeToLoggerFunction(errorCode)

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

	loggerFn(r.Context(), loggedMessage, fields...)
}

func mapErrorCodeToHTTPStatus(errorCode syserr.Code, err error) int {
	switch errorCode {
	case syserr.NotFoundCode:
		return http.StatusNotFound
	case syserr.BadInputCode:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func extractPublicMessage(errorCode syserr.Code, err error) string {
	switch errorCode {
	case syserr.InternalCode:
		return "internal error occurred"
	default:
		return err.Error()
	}
}

func mapErrorCodeToLoggerFunction(errorCode syserr.Code) func(ctx context.Context, message string, fields ...*logger.Field) {
	switch errorCode {
	case syserr.InternalCode:
		return logger.Error
	case syserr.NotFoundCode:
		return logger.Warning
	case syserr.BadInputCode:
		return logger.Warning
	default:
		return logger.Error
	}
}

func extractErrorCode(err error) syserr.Code {
	var systemError *syserr.Error
	ok := errors.As(err, &systemError)
	if ok {
		return systemError.GetCode()
	}

	return syserr.InternalCode
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

func extractProvidedOperationID(r *http.Request) string {
	headers := r.Header
	for name, value := range headers {
		if name == OperationIDHeaderName {
			return value[0]
		}
	}

	return ""
}
