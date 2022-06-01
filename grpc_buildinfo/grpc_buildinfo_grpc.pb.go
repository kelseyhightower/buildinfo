// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: grpc_buildinfo/grpc_buildinfo.proto

package grpc_buildinfo

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

// BuildInfoServiceClient is the client API for BuildInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildInfoServiceClient interface {
	GetBuildInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*BuildInfo, error)
}

type buildInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildInfoServiceClient(cc grpc.ClientConnInterface) BuildInfoServiceClient {
	return &buildInfoServiceClient{cc}
}

func (c *buildInfoServiceClient) GetBuildInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*BuildInfo, error) {
	out := new(BuildInfo)
	err := c.cc.Invoke(ctx, "/buildinfo.BuildInfoService/GetBuildInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildInfoServiceServer is the server API for BuildInfoService service.
// All implementations must embed UnimplementedBuildInfoServiceServer
// for forward compatibility
type BuildInfoServiceServer interface {
	GetBuildInfo(context.Context, *EmptyRequest) (*BuildInfo, error)
	mustEmbedUnimplementedBuildInfoServiceServer()
}

// UnimplementedBuildInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBuildInfoServiceServer struct {
}

func (UnimplementedBuildInfoServiceServer) GetBuildInfo(context.Context, *EmptyRequest) (*BuildInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildInfo not implemented")
}
func (UnimplementedBuildInfoServiceServer) mustEmbedUnimplementedBuildInfoServiceServer() {}

// UnsafeBuildInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildInfoServiceServer will
// result in compilation errors.
type UnsafeBuildInfoServiceServer interface {
	mustEmbedUnimplementedBuildInfoServiceServer()
}

func RegisterBuildInfoServiceServer(s grpc.ServiceRegistrar, srv BuildInfoServiceServer) {
	s.RegisterService(&BuildInfoService_ServiceDesc, srv)
}

func _BuildInfoService_GetBuildInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildInfoServiceServer).GetBuildInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildinfo.BuildInfoService/GetBuildInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildInfoServiceServer).GetBuildInfo(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuildInfoService_ServiceDesc is the grpc.ServiceDesc for BuildInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuildInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildinfo.BuildInfoService",
	HandlerType: (*BuildInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuildInfo",
			Handler:    _BuildInfoService_GetBuildInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_buildinfo/grpc_buildinfo.proto",
}