package network

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	imagepb "backend/proto/image/v1"

	"backend/internal/domain"
)

// https://github.com/grpc-ecosystem/grpc-gateway/

type APISchemaItem struct {
	RegisterClient  func(ctx context.Context, mux *runtime.ServeMux, gRPCConnection *grpc.ClientConn) error
	RegisterService func(grpcServer *grpc.Server, controllers *Controllers)
}

type Controllers struct {
	ImageServiceV1 imagepb.ImageServiceServer
}

var (
	APISchema = []APISchemaItem{
		{
			RegisterClient: func(ctx context.Context, mux *runtime.ServeMux, gRPCConnection *grpc.ClientConn) error {
				return imagepb.RegisterImageServiceHandlerClient(ctx, mux, imagepb.NewImageServiceClient(gRPCConnection))
			},
			RegisterService: func(grpcServer *grpc.Server, controllers *Controllers) {
				imagepb.RegisterImageServiceServer(grpcServer, controllers.ImageServiceV1)
			},
		},
	}
)

func GetMux(ctx context.Context, gRPCConnection *grpc.ClientConn) (http.Handler, error) {
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
		err := schemaItem.RegisterClient(ctx, mux, gRPCConnection)
		if err != nil {
			return nil, err
		}
	}

	return mux, nil
}

func StartHTTPServer(ctx context.Context, config *domain.Config, mux http.Handler) (func() error, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTPPort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	err := server.ListenAndServe()
	if err != nil {
		return nil, err
	}

	return func() error {
		return server.Shutdown(ctx)
	}, nil
}

func ConnectToGRPCServer(config *domain.Config) (*grpc.ClientConn, func() error, error) {
	connection, err := grpc.Dial(
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
	if err != nil {
		return nil, nil, err
	}

	return connection, connection.Close, nil
}

func StartGRPCServer(ctx context.Context, configuration *domain.Config, controllers *Controllers) (func(), error) {
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", configuration.GRPCPort))
	//listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configuration.GRPCPort))
	if err != nil {
		return nil, err
	}

	opts := grpc.ChainUnaryInterceptor(
	//s.auth.PopulateUser,
	//request.PopulateContext(),
	)
	grpcServer := grpc.NewServer(opts)

	for _, schemaItem := range APISchema {
		schemaItem.RegisterService(grpcServer, controllers)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		return nil, err
	}

	return grpcServer.GracefulStop, nil
}
