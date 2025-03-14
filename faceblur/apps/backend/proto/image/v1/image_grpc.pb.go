// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: image/v1/image.proto

package imagepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ImageService_GetUploadURL_FullMethodName = "/faceblur.image.v1.ImageService/GetUploadURL"
	ImageService_SubmitImage_FullMethodName  = "/faceblur.image.v1.ImageService/SubmitImage"
	ImageService_ListImages_FullMethodName   = "/faceblur.image.v1.ImageService/ListImages"
)

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	// GetUploadURL returns a new signed URL for image upload
	GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error)
	// SubmitImage creates a new image and puts it to the processing queue
	SubmitImage(ctx context.Context, in *SubmitImageRequest, opts ...grpc.CallOption) (*SubmitImageResponse, error)
	// ListImages returns a list of user images, paginated and sorted by creation date
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUploadURLResponse)
	err := c.cc.Invoke(ctx, ImageService_GetUploadURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) SubmitImage(ctx context.Context, in *SubmitImageRequest, opts ...grpc.CallOption) (*SubmitImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitImageResponse)
	err := c.cc.Invoke(ctx, ImageService_SubmitImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, ImageService_ListImages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility.
type ImageServiceServer interface {
	// GetUploadURL returns a new signed URL for image upload
	GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error)
	// SubmitImage creates a new image and puts it to the processing queue
	SubmitImage(context.Context, *SubmitImageRequest) (*SubmitImageResponse, error)
	// ListImages returns a list of user images, paginated and sorted by creation date
	ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageServiceServer struct{}

func (UnimplementedImageServiceServer) GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadURL not implemented")
}
func (UnimplementedImageServiceServer) SubmitImage(context.Context, *SubmitImageRequest) (*SubmitImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitImage not implemented")
}
func (UnimplementedImageServiceServer) ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}
func (UnimplementedImageServiceServer) testEmbeddedByValue()                      {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	// If the following call pancis, it indicates UnimplementedImageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_GetUploadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetUploadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_GetUploadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetUploadURL(ctx, req.(*GetUploadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_SubmitImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).SubmitImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_SubmitImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).SubmitImage(ctx, req.(*SubmitImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_ListImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "faceblur.image.v1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUploadURL",
			Handler:    _ImageService_GetUploadURL_Handler,
		},
		{
			MethodName: "SubmitImage",
			Handler:    _ImageService_SubmitImage_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _ImageService_ListImages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image/v1/image.proto",
}
