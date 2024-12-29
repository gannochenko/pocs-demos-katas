package network

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"backend/interfaces"
	"backend/internal/domain"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer(controllers *Controllers, loggerService interfaces.LoggerService) *GRPCServer {
	opts := grpc.ChainUnaryInterceptor(
		MapError,
		PopulateUser,
		PopulateOperationID,
		GetLogRequest(loggerService),
	)
	grpcServer := grpc.NewServer(opts)

	for _, schemaItem := range APISchema {
		schemaItem.RegisterService(grpcServer, controllers)
	}

	return &GRPCServer{
		server: grpcServer,
	}
}

func (s *GRPCServer) Start(ctx context.Context, configuration *domain.Config) error {
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", configuration.GRPCPort))
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
