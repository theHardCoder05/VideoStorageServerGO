// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: service_content.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VideoStorageServiceClient is the client API for VideoStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStorageServiceClient interface {
	CreateContent(ctx context.Context, in *CreateContentRequest, opts ...grpc.CallOption) (*CreateContentResponse, error)
}

type videoStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStorageServiceClient(cc grpc.ClientConnInterface) VideoStorageServiceClient {
	return &videoStorageServiceClient{cc}
}

func (c *videoStorageServiceClient) CreateContent(ctx context.Context, in *CreateContentRequest, opts ...grpc.CallOption) (*CreateContentResponse, error) {
	out := new(CreateContentResponse)
	err := c.cc.Invoke(ctx, "/db.VideoStorageService/CreateContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoStorageServiceServer is the server API for VideoStorageService service.
// All implementations must embed UnimplementedVideoStorageServiceServer
// for forward compatibility
type VideoStorageServiceServer interface {
	CreateContent(context.Context, *CreateContentRequest) (*CreateContentResponse, error)
	mustEmbedUnimplementedVideoStorageServiceServer()
}

// UnimplementedVideoStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoStorageServiceServer struct {
}

func (UnimplementedVideoStorageServiceServer) CreateContent(context.Context, *CreateContentRequest) (*CreateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContent not implemented")
}
func (UnimplementedVideoStorageServiceServer) mustEmbedUnimplementedVideoStorageServiceServer() {}

// UnsafeVideoStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStorageServiceServer will
// result in compilation errors.
type UnsafeVideoStorageServiceServer interface {
	mustEmbedUnimplementedVideoStorageServiceServer()
}

func RegisterVideoStorageServiceServer(s grpc.ServiceRegistrar, srv VideoStorageServiceServer) {
	s.RegisterService(&VideoStorageService_ServiceDesc, srv)
}

func _VideoStorageService_CreateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoStorageServiceServer).CreateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.VideoStorageService/CreateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoStorageServiceServer).CreateContent(ctx, req.(*CreateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoStorageService_ServiceDesc is the grpc.ServiceDesc for VideoStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.VideoStorageService",
	HandlerType: (*VideoStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContent",
			Handler:    _VideoStorageService_CreateContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_content.proto",
}
