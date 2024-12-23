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

	err := imagepb.RegisterImageServiceHandlerClient(ctx, mux, imagepb.NewImageServiceClient(gRPCConnection))
	if err != nil {
		return nil, err
	}

	//err := RegisterImag(ctx, mux, imageV1Client)
	//if err != nil {
	//	log.ErrorE(err)
	//	os.Exit(1)
	//}
	//
	//var opts []grpc.DialOption
	//err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	//if err != nil {
	//	log.Fatalf("Failed to register gRPC Gateway: %v", err)
	//}

	return mux, nil
}

func ConnectToGRPCServer(config *domain.Config) (*grpc.ClientConn, func() error, error) {
	connection, err := grpc.Dial(
		fmt.Sprintf("0.0.0.0:%s", config.GRPCPort),
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

func StartGRPCServer(configuration *domain.Config) (func(), error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configuration.GRPCPort))
	if err != nil {
		return nil, err
	}

	opts := grpc.ChainUnaryInterceptor(
	//s.auth.PopulateUser,
	//request.PopulateContext(),
	)
	grpcServer := grpc.NewServer(opts)

	// todo: attach controllers here

	err = grpcServer.Serve(listener)
	if err != nil {
		return nil, err
	}

	return grpcServer.GracefulStop, nil
}
