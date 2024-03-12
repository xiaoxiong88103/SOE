// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: config.proto

package config

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

const (
	SystemMetrics_GethardwareInfo_FullMethodName = "/monitor.SystemMetrics/GethardwareInfo"
	SystemMetrics_GetSystemInfo_FullMethodName   = "/monitor.SystemMetrics/GetSystemInfo"
)

// SystemMetricsClient is the client API for SystemMetrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemMetricsClient interface {
	GethardwareInfo(ctx context.Context, in *HardwareInfo, opts ...grpc.CallOption) (*Response, error)
	GetSystemInfo(ctx context.Context, in *SystemInfo, opts ...grpc.CallOption) (*Response, error)
}

type systemMetricsClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemMetricsClient(cc grpc.ClientConnInterface) SystemMetricsClient {
	return &systemMetricsClient{cc}
}

func (c *systemMetricsClient) GethardwareInfo(ctx context.Context, in *HardwareInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, SystemMetrics_GethardwareInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMetricsClient) GetSystemInfo(ctx context.Context, in *SystemInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, SystemMetrics_GetSystemInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemMetricsServer is the server API for SystemMetrics service.
// All implementations must embed UnimplementedSystemMetricsServer
// for forward compatibility
type SystemMetricsServer interface {
	GethardwareInfo(context.Context, *HardwareInfo) (*Response, error)
	GetSystemInfo(context.Context, *SystemInfo) (*Response, error)
	mustEmbedUnimplementedSystemMetricsServer()
}

// UnimplementedSystemMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedSystemMetricsServer struct {
}

func (UnimplementedSystemMetricsServer) GethardwareInfo(context.Context, *HardwareInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GethardwareInfo not implemented")
}
func (UnimplementedSystemMetricsServer) GetSystemInfo(context.Context, *SystemInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedSystemMetricsServer) mustEmbedUnimplementedSystemMetricsServer() {}

// UnsafeSystemMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemMetricsServer will
// result in compilation errors.
type UnsafeSystemMetricsServer interface {
	mustEmbedUnimplementedSystemMetricsServer()
}

func RegisterSystemMetricsServer(s grpc.ServiceRegistrar, srv SystemMetricsServer) {
	s.RegisterService(&SystemMetrics_ServiceDesc, srv)
}

func _SystemMetrics_GethardwareInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HardwareInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMetricsServer).GethardwareInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMetrics_GethardwareInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMetricsServer).GethardwareInfo(ctx, req.(*HardwareInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMetrics_GetSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMetricsServer).GetSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMetrics_GetSystemInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMetricsServer).GetSystemInfo(ctx, req.(*SystemInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemMetrics_ServiceDesc is the grpc.ServiceDesc for SystemMetrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemMetrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.SystemMetrics",
	HandlerType: (*SystemMetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GethardwareInfo",
			Handler:    _SystemMetrics_GethardwareInfo_Handler,
		},
		{
			MethodName: "GetSystemInfo",
			Handler:    _SystemMetrics_GetSystemInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}