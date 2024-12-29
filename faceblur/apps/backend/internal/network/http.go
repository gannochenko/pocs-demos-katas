package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"backend/internal/util/syserr"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"backend/internal/domain"
)

type HTTPServer struct {
	server         *http.Server
	gRPCConnection *grpc.ClientConn
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) GetMux(ctx context.Context) (http.Handler, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("*", &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			return runtime.DefaultHeaderMatcher(s)
		}),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs(
				"http-request-path", req.URL.Path,
				"http-method", req.Method,
				"http-query-params", req.URL.RawQuery,
			)
		}),
		runtime.WithOutgoingHeaderMatcher(func(s string) (string, bool) {
			if s == "x-operation-id" {
				return "X-Operation-Id", true
			}
			return runtime.DefaultHeaderMatcher(s)
		}),
		runtime.WithErrorHandler(customErrorHandler),
	)

	for _, schemaItem := range APISchema {
		err := schemaItem.RegisterClient(ctx, mux, s.gRPCConnection)
		if err != nil {
			return nil, err
		}
	}

	return mux, nil
}

func (s *HTTPServer) Start(ctx context.Context, config *domain.Config) error {
	var err error
	s.gRPCConnection, err = s.connectToGRPCServer(config)
	if err != nil {
		return syserr.Wrap(err, "could not connect to gRPC server")
	}

	mux, err := s.GetMux(ctx)
	if err != nil {
		return err
	}

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTPPort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	var errors []string

	closeErr := s.gRPCConnection.Close()
	if closeErr != nil {
		errors = append(errors, syserr.Wrap(closeErr, "could not close grpc connection").Error())
	}

	shutdownErr := s.server.Shutdown(ctx)
	if shutdownErr != nil {
		errors = append(errors, shutdownErr.Error())
	}

	if len(errors) > 0 {
		return syserr.NewInternal(fmt.Sprintf("could not shut down HTTP server: %s", strings.Join(errors, ", ")))
	}

	return nil
}

func (s *HTTPServer) connectToGRPCServer(config *domain.Config) (*grpc.ClientConn, error) {
	return grpc.Dial(
		fmt.Sprintf("0.0.0.0:%d", config.GRPCPort),
		// the connection takes place in the VPC tier, no security is needed
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func customErrorHandler(ctx context.Context, mux *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	grpcStatus, _ := status.FromError(err)

	httpStatus := http.StatusInternalServerError

	switch grpcStatus.Code() {
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	case codes.PermissionDenied:
		httpStatus = http.StatusForbidden
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
	case codes.FailedPrecondition:
		httpStatus = http.StatusPreconditionFailed
	case codes.Unavailable:
		httpStatus = http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		httpStatus = http.StatusGatewayTimeout
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	// todo: use protobuf message here
	json.NewEncoder(w).Encode(map[string]string{
		"error": grpcStatus.Message(),
	})
}
