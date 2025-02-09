package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"backend/interfaces"
	"backend/internal/util/syserr"
	errorpb "backend/proto/common/error/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"backend/internal/domain"
)

type HTTPServer struct {
	server         *http.Server
	gRPCConnection *grpc.ClientConn

	configService   interfaces.ConfigService
	websocketServer interfaces.WebsocketServer
}

func NewHTTPServer(configService interfaces.ConfigService, websocketServer interfaces.WebsocketServer) *HTTPServer {
	return &HTTPServer{configService: configService, websocketServer: websocketServer}
}

func (s *HTTPServer) GetMux(ctx context.Context) (http.Handler, error) {
	mainRouter := mux.NewRouter()

	proxyMux, err := s.getGRPCProxyMux(ctx)
	if err != nil {
		return nil, syserr.Wrap(err, "could not create grpc proxy mux")
	}

	// todo: put some actual logic in here if needed
	mainRouter.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("1"))
	})
	mainRouter.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("1"))
	})
	mainRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		_ = withHTTPLogger(withHTTPErrorHandler(withHTTPContext(s.websocketServer.GetHandler())))(w, r)
	})
	mainRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyMux.ServeHTTP(w, r)
	})

	config, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not get config")
	}

	externalMux := withCorsMiddleware(mainRouter, config)

	return externalMux, nil
}

// getGRPCProxyMux creates a mux that serves all paths declared in the protobuf files
func (s *HTTPServer) getGRPCProxyMux(ctx context.Context) (*runtime.ServeMux, error) {
	proxyMux := runtime.NewServeMux(
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

	for _, schemaItem := range GRPCAPI {
		err := schemaItem.RegisterHTTPClient(ctx, proxyMux, s.gRPCConnection)
		if err != nil {
			return nil, err
		}
	}

	return proxyMux, nil
}

func (s *HTTPServer) Start(ctx context.Context, config *domain.Config) error {
	var err error
	s.gRPCConnection, err = s.connectToGRPCServer(config)
	if err != nil {
		return syserr.Wrap(err, "could not connect to gRPC server")
	}

	serverMux, err := s.GetMux(ctx)
	if err != nil {
		return err
	}

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Backend.HTTP.Port),
		Handler: serverMux,
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
		fmt.Sprintf("0.0.0.0:%d", config.Backend.GRPCPort),
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

	responseError := &errorpb.ErrorResponse{
		Error: lo.Ternary(grpcStatus.Code() == codes.Internal, "internal error occurred, contact support", grpcStatus.Message()),
	}

	json.NewEncoder(w).Encode(responseError)
}
