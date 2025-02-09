package network

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	imagepbV1 "backend/proto/image/v1"
)

type GRPCControllers struct {
	ImageServiceV1 imagepbV1.ImageServiceServer
}

var (
	GRPCAPI = []GRPCAPIItem{
		{
			RegisterHTTPClient: func(ctx context.Context, mux *runtime.ServeMux, gRPCConnection *grpc.ClientConn) error {
				return imagepbV1.RegisterImageServiceHandlerClient(ctx, mux, imagepbV1.NewImageServiceClient(gRPCConnection))
			},
			RegisterGRPCService: func(grpcServer *grpc.Server, controllers *GRPCControllers) {
				imagepbV1.RegisterImageServiceServer(grpcServer, controllers.ImageServiceV1)
			},
		},
	}
)
