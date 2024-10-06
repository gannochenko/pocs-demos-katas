package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"api/internal/types"
	pkgCtx "api/pkg/ctx"
	"api/pkg/logger"
	"api/pkg/syserr"
)

func WithLogger(next types.Handler) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		operationID := uuid.New().String()
		ctx := pkgCtx.WithOperationID(r.Context(), operationID)
		*r = *r.WithContext(ctx)

		w.Header().Set("X-Operation-ID", operationID)

		err := next(w, r)

		user := pkgCtx.GetUserEmail(r.Context())

		fields := make([]*logger.Field, 1)
		fields[0] = logger.F("query", fmt.Sprintf("%s %s", r.Method, r.RequestURI))
		if user != "" {
			fields = append(fields, logger.F("user", user))
		}

		ctx = r.Context()
		if err != nil {
			LogByErrorCode(ctx, errors.Wrap(err, "request handles with errors"), fields...)
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

func LogByErrorCode(ctx context.Context, err error, fields ...*logger.Field) {
	code := syserr.GetCode(err)
	fn := errorToLogLevel[code]
	if fn == nil {
		fn = logger.Error
	}

	fields = append(fields, convertErrorFieldsToLoggerFields(syserr.GetFields(err))...)
	fields = append(fields, logger.F("stack", syserr.GetStackFormatted(err)))

	fn(ctx, err.Error(), fields...)
}

func convertErrorFieldsToLoggerFields(fields []*syserr.Field) []*logger.Field {
	result := make([]*logger.Field, len(fields))

	for index, field := range fields {
		result[index] = logger.F(field.Key, field.Value)
	}

	return result
}
