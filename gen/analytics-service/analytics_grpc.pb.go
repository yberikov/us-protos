// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: analytics-service/analytics.proto

package analytics_microservice

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
	AnalyticsService_LogURLAccess_FullMethodName         = "/analytics.AnalyticsService/LogURLAccess"
	AnalyticsService_GetURLStats_FullMethodName          = "/analytics.AnalyticsService/GetURLStats"
	AnalyticsService_GetURLStatsForPeriod_FullMethodName = "/analytics.AnalyticsService/GetURLStatsForPeriod"
)

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	// LogURLAccess logs the access of a shortened URL.
	LogURLAccess(ctx context.Context, in *LogURLAccessRequest, opts ...grpc.CallOption) (*LogURLAccessResponse, error)
	// GetURLStats retrieves statistics for a specific URL.
	GetURLStats(ctx context.Context, in *GetURLStatsRequest, opts ...grpc.CallOption) (*GetURLStatsResponse, error)
	// GetURLStatsForPeriod retrieves statistics for a specific URL within a time period.
	GetURLStatsForPeriod(ctx context.Context, in *GetURLStatsForPeriodRequest, opts ...grpc.CallOption) (*GetURLStatsForPeriodResponse, error)
}

type analyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsServiceClient(cc grpc.ClientConnInterface) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) LogURLAccess(ctx context.Context, in *LogURLAccessRequest, opts ...grpc.CallOption) (*LogURLAccessResponse, error) {
	out := new(LogURLAccessResponse)
	err := c.cc.Invoke(ctx, AnalyticsService_LogURLAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsServiceClient) GetURLStats(ctx context.Context, in *GetURLStatsRequest, opts ...grpc.CallOption) (*GetURLStatsResponse, error) {
	out := new(GetURLStatsResponse)
	err := c.cc.Invoke(ctx, AnalyticsService_GetURLStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsServiceClient) GetURLStatsForPeriod(ctx context.Context, in *GetURLStatsForPeriodRequest, opts ...grpc.CallOption) (*GetURLStatsForPeriodResponse, error) {
	out := new(GetURLStatsForPeriodResponse)
	err := c.cc.Invoke(ctx, AnalyticsService_GetURLStatsForPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
// All implementations must embed UnimplementedAnalyticsServiceServer
// for forward compatibility
type AnalyticsServiceServer interface {
	// LogURLAccess logs the access of a shortened URL.
	LogURLAccess(context.Context, *LogURLAccessRequest) (*LogURLAccessResponse, error)
	// GetURLStats retrieves statistics for a specific URL.
	GetURLStats(context.Context, *GetURLStatsRequest) (*GetURLStatsResponse, error)
	// GetURLStatsForPeriod retrieves statistics for a specific URL within a time period.
	GetURLStatsForPeriod(context.Context, *GetURLStatsForPeriodRequest) (*GetURLStatsForPeriodResponse, error)
	mustEmbedUnimplementedAnalyticsServiceServer()
}

// UnimplementedAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServiceServer struct {
}

func (UnimplementedAnalyticsServiceServer) LogURLAccess(context.Context, *LogURLAccessRequest) (*LogURLAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogURLAccess not implemented")
}
func (UnimplementedAnalyticsServiceServer) GetURLStats(context.Context, *GetURLStatsRequest) (*GetURLStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLStats not implemented")
}
func (UnimplementedAnalyticsServiceServer) GetURLStatsForPeriod(context.Context, *GetURLStatsForPeriodRequest) (*GetURLStatsForPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLStatsForPeriod not implemented")
}
func (UnimplementedAnalyticsServiceServer) mustEmbedUnimplementedAnalyticsServiceServer() {}

// UnsafeAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServiceServer will
// result in compilation errors.
type UnsafeAnalyticsServiceServer interface {
	mustEmbedUnimplementedAnalyticsServiceServer()
}

func RegisterAnalyticsServiceServer(s grpc.ServiceRegistrar, srv AnalyticsServiceServer) {
	s.RegisterService(&AnalyticsService_ServiceDesc, srv)
}

func _AnalyticsService_LogURLAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogURLAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).LogURLAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_LogURLAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).LogURLAccess(ctx, req.(*LogURLAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsService_GetURLStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).GetURLStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_GetURLStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).GetURLStats(ctx, req.(*GetURLStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsService_GetURLStatsForPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLStatsForPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).GetURLStatsForPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_GetURLStatsForPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).GetURLStatsForPeriod(ctx, req.(*GetURLStatsForPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalyticsService_ServiceDesc is the grpc.ServiceDesc for AnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogURLAccess",
			Handler:    _AnalyticsService_LogURLAccess_Handler,
		},
		{
			MethodName: "GetURLStats",
			Handler:    _AnalyticsService_GetURLStats_Handler,
		},
		{
			MethodName: "GetURLStatsForPeriod",
			Handler:    _AnalyticsService_GetURLStatsForPeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analytics-service/analytics.proto",
}
