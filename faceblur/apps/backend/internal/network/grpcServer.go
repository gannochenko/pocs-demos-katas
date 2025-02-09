package network

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"backend/interfaces"
	"backend/internal/domain"
)

// https://github.com/grpc-ecosystem/grpc-gateway/

type GRPCAPIItem struct {
	RegisterHTTPClient  func(ctx context.Context, mux *runtime.ServeMux, gRPCConnection *grpc.ClientConn) error
	RegisterGRPCService func(grpcServer *grpc.Server, controllers *GRPCControllers)
}

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer(controllers *GRPCControllers, loggerService interfaces.LoggerService, authService interfaces.AuthService, userService interfaces.UserService) *GRPCServer {
	opts := grpc.ChainUnaryInterceptor(
		GRPCMapError,
		GRPCPopulateOperationID,
		GetGRPCUserPopulator(authService, userService),
		GRPCGetRequestLogger(loggerService),
	)
	grpcServer := grpc.NewServer(opts)

	for _, schemaItem := range GRPCAPI {
		schemaItem.RegisterGRPCService(grpcServer, controllers)
	}

	return &GRPCServer{
		server: grpcServer,
	}
}

func (s *GRPCServer) Start(ctx context.Context, configuration *domain.Config) error {
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", configuration.Backend.GRPCPort))
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
