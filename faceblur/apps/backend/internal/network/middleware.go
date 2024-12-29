package network

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"backend/interfaces"
	"backend/internal/domain"
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"
	"backend/internal/util/types"
)

func PopulateUser(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tokens := md["authorization"]

	if len(tokens) > 0 {
		// todo: parse and verify token here

		ctx = ctxUtil.WithUser(ctx, domain.Userr{
			Email: "foo@bar.baz",
		})
	}

	return handler(ctx, req)
}

func PopulateOperationID(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	operationID := uuid.New().String()
	ctx = ctxUtil.WithOperationID(ctx, operationID)

	_ = grpc.SetHeader(ctx, metadata.Pairs(
		"X-Operation-Id", operationID,
	))

	return handler(ctxUtil.WithOperationID(ctx, uuid.New().String()), req)
}

func GetLogRequest(loggerService interfaces.LoggerService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)

		md, _ := metadata.FromIncomingContext(ctx)
		// httpContentType := md.Get("grpcgateway-content-type")
		httpMethod := types.Unwrap(md.Get("http-method"), "")
		httpRequestPath := types.Unwrap(md.Get("http-request-path"), "")

		b := bytes.NewBuffer([]byte{})
		_ = json.NewEncoder(b).Encode(req)

		var fields []*logger.Field

		fields = append(
			fields,
			logger.F("http_request", fmt.Sprintf("%s %s", httpMethod, httpRequestPath)),
			logger.F("grpc_request", info.FullMethod),
			logger.F("request_data", strings.Trim(b.String(), " \r\n")),
		)

		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "request handled with error"), fields...)
		} else {
			loggerService.Info(ctx, "request handled", fields...)
		}

		return resp, err
	}
}

var (
	syserrCodeToGrpc = map[syserr.Code]codes.Code{
		syserr.BadInputCode:       codes.InvalidArgument,
		syserr.UnauthorisedCode:   codes.PermissionDenied,
		syserr.NotFoundCode:       codes.NotFound,
		syserr.NotImplementedCode: codes.Unimplemented,
		syserr.InternalCode:       codes.Internal,
	}
)

func MapError(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	resp, err := handler(ctx, req)

	if err != nil {
		switch e := err.(type) {
		case *syserr.Error:
			code := syserr.GetCode(e)
			grpcCode, ok := syserrCodeToGrpc[code]
			if ok {
				return nil, status.Error(grpcCode, err.Error())
			}
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return resp, nil
}
