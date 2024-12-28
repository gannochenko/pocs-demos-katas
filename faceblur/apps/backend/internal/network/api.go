package network

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	imagepb "backend/proto/image/v1"
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
