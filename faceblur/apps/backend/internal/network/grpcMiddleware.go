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
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"
	"backend/internal/util/types"
)

func GetGRPCUserPopulator(authService interfaces.AuthService, userService interfaces.UserService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		token, err := authService.ExtractToken(ctx)
		if err != nil {
			return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not validate token")
		}

		sup, _, err := authService.ValidateToken(ctx, token)
		if err != nil {
			return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not validate token")
		}

		//sup := "auth0:19482" // debug user

		user, err := userService.GetUserBySUP(ctx, nil, sup)
		if err != nil {
			return nil, syserr.Wrap(err, "error getting user", syserr.F("sup", sup))
		} else {
			ctx = ctxUtil.WithUser(ctx, *user)
		}

		return handler(ctx, req)
	}
}

func GRPCPopulateOperationID(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	operationID := uuid.New().String()
	ctx = ctxUtil.WithOperationID(ctx, operationID)

	_ = grpc.SetHeader(ctx, metadata.Pairs(
		"X-Operation-Id", operationID,
	))

	return handler(ctxUtil.WithOperationID(ctx, uuid.New().String()), req)
}

func GRPCGetRequestLogger(loggerService interfaces.LoggerService) grpc.UnaryServerInterceptor {
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

func GRPCMapError(
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
