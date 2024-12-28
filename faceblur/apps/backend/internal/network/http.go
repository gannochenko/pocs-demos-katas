package network

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

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
		//runtime.WithIncomingHeaderMatcher(middleware.HeaderMatcher),
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
		//grpc.WithStreamInterceptor(grpcMiddleware.ChainStreamClient(
		//	grpcOpentracing.StreamClientInterceptor(grpcOpentracing.WithTracer(*s.tracer)),
		//)),
		//grpc.WithUnaryInterceptor(grpcMiddleware.ChainUnaryClient(
		//	grpcOpentracing.UnaryClientInterceptor(grpcOpentracing.WithTracer(*s.tracer)),
		//)),
	)
}
