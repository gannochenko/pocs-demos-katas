package http

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	imagepb "api/proto/image/v1"
)

// https://github.com/grpc-ecosystem/grpc-gateway/

func connectGrpcServer() *grpc.ClientConn {
	conn, err := grpc.Dial(
		fmt.Sprintf("0.0.0.0:%s", "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithStreamInterceptor(grpcMiddleware.ChainStreamClient(
		//	grpcOpentracing.StreamClientInterceptor(grpcOpentracing.WithTracer(*s.tracer)),
		//)),
		//grpc.WithUnaryInterceptor(grpcMiddleware.ChainUnaryClient(
		//	grpcOpentracing.UnaryClientInterceptor(grpcOpentracing.WithTracer(*s.tracer)),
		//)),
	)

	if err != nil {
		os.Exit(1)
	}

	return conn
}

func GetMux(ctx context.Context) (http.Handler, error) {
	conn := connectGrpcServer()

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

	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := gw.RegisterImageServiceServer(ctx, mux, *grpcServerEndpoint, opts)
	//if err != nil {
	//	return err
	//}

	err := imagepb.RegisterImageServiceHandlerClient(ctx, mux, imagepb.NewImageServiceClient(conn))
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
