package util

import (
	"context"
	"fmt"
	"net/http"

	"api/pkg/logger"
	"api/pkg/syserr"
)

func withLogger(h AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := h(w, r)

		fields := make([]*logger.Field, 1)
		fields[0] = logger.F("query", fmt.Sprintf("%s %s", r.Method, r.RequestURI))

		ctx := r.Context()
		if err != nil {
			logByErrorCode(ctx, err, fields...)
		} else {
			logger.Info(ctx, "request handled", fields...)
		}

		return err
	}
}

var (
	errorToLogLevel = map[syserr.Code]func(ctx context.Context, message string, fields ...*logger.Field){
		syserr.InternalCode:       logger.Error,
		syserr.BadInputCode:       logger.Warning,
		syserr.NotFoundCode:       logger.Warning,
		syserr.NotImplementedCode: logger.Error,
	}
)

func logByErrorCode(ctx context.Context, err error, fields ...*logger.Field) {
	code := syserr.GetCode(err)
	fn := errorToLogLevel[code]
	if fn == nil {
		fn = logger.Error
	}

	fields = append(fields, convertErrorFieldsToLoggerFields(syserr.GetFields(err))...)
	fields = append(fields, logger.F("stack", syserr.GetStackFormatted(err)))

	fn(ctx, fmt.Sprintf("request handled with error: %s", err.Error()), fields...)
}

func convertErrorFieldsToLoggerFields(fields []*syserr.Field) []*logger.Field {
	result := make([]*logger.Field, len(fields))

	for index, field := range fields {
		result[index] = logger.F(field.Key, field.Value)
	}

	return result
}
