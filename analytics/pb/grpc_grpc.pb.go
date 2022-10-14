// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: grpc.proto

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

// RouteAnalyticsServiceClient is the client API for RouteAnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteAnalyticsServiceClient interface {
	SessionStartedEvent(ctx context.Context, in *SessionStartedEventRequest, opts ...grpc.CallOption) (*SessionStartedEventResponse, error)
	StageCompletedEvent(ctx context.Context, in *StageCompletedEventRequest, opts ...grpc.CallOption) (*StageCompletedEventResponse, error)
	ChapterCompletedEvent(ctx context.Context, in *ChapterCompletedEventRequest, opts ...grpc.CallOption) (*ChapterCompletedEventResponse, error)
	BasicStageCompletedEvent(ctx context.Context, in *BasicStageCompletedEventRequest, opts ...grpc.CallOption) (*BasicStageCompletedEventResponse, error)
	EventStageCompletedEvent(ctx context.Context, in *EventStageCompletedEventRequest, opts ...grpc.CallOption) (*EventStageCompletedEventResponse, error)
	BasicLevelNewScoreEvent(ctx context.Context, in *BasicLevelNewScoreEventRequest, opts ...grpc.CallOption) (*BasicLevelNewScoreEventResponse, error)
	EventLevelNewScoreEvent(ctx context.Context, in *EventLevelNewScoreEventRequest, opts ...grpc.CallOption) (*EventLevelNewScoreEventResponse, error)
	// FOR TESTING ONLY
	IsAlive(ctx context.Context, in *IsAliveRequest, opts ...grpc.CallOption) (*IsAliveResponse, error)
}

type routeAnalyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteAnalyticsServiceClient(cc grpc.ClientConnInterface) RouteAnalyticsServiceClient {
	return &routeAnalyticsServiceClient{cc}
}

func (c *routeAnalyticsServiceClient) SessionStartedEvent(ctx context.Context, in *SessionStartedEventRequest, opts ...grpc.CallOption) (*SessionStartedEventResponse, error) {
	out := new(SessionStartedEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/SessionStartedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) StageCompletedEvent(ctx context.Context, in *StageCompletedEventRequest, opts ...grpc.CallOption) (*StageCompletedEventResponse, error) {
	out := new(StageCompletedEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/StageCompletedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) ChapterCompletedEvent(ctx context.Context, in *ChapterCompletedEventRequest, opts ...grpc.CallOption) (*ChapterCompletedEventResponse, error) {
	out := new(ChapterCompletedEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/ChapterCompletedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) BasicStageCompletedEvent(ctx context.Context, in *BasicStageCompletedEventRequest, opts ...grpc.CallOption) (*BasicStageCompletedEventResponse, error) {
	out := new(BasicStageCompletedEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/BasicStageCompletedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) EventStageCompletedEvent(ctx context.Context, in *EventStageCompletedEventRequest, opts ...grpc.CallOption) (*EventStageCompletedEventResponse, error) {
	out := new(EventStageCompletedEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/EventStageCompletedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) BasicLevelNewScoreEvent(ctx context.Context, in *BasicLevelNewScoreEventRequest, opts ...grpc.CallOption) (*BasicLevelNewScoreEventResponse, error) {
	out := new(BasicLevelNewScoreEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/BasicLevelNewScoreEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) EventLevelNewScoreEvent(ctx context.Context, in *EventLevelNewScoreEventRequest, opts ...grpc.CallOption) (*EventLevelNewScoreEventResponse, error) {
	out := new(EventLevelNewScoreEventResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/EventLevelNewScoreEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeAnalyticsServiceClient) IsAlive(ctx context.Context, in *IsAliveRequest, opts ...grpc.CallOption) (*IsAliveResponse, error) {
	out := new(IsAliveResponse)
	err := c.cc.Invoke(ctx, "/RouteAnalyticsService/IsAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteAnalyticsServiceServer is the server API for RouteAnalyticsService service.
// All implementations must embed UnimplementedRouteAnalyticsServiceServer
// for forward compatibility
type RouteAnalyticsServiceServer interface {
	SessionStartedEvent(context.Context, *SessionStartedEventRequest) (*SessionStartedEventResponse, error)
	StageCompletedEvent(context.Context, *StageCompletedEventRequest) (*StageCompletedEventResponse, error)
	ChapterCompletedEvent(context.Context, *ChapterCompletedEventRequest) (*ChapterCompletedEventResponse, error)
	BasicStageCompletedEvent(context.Context, *BasicStageCompletedEventRequest) (*BasicStageCompletedEventResponse, error)
	EventStageCompletedEvent(context.Context, *EventStageCompletedEventRequest) (*EventStageCompletedEventResponse, error)
	BasicLevelNewScoreEvent(context.Context, *BasicLevelNewScoreEventRequest) (*BasicLevelNewScoreEventResponse, error)
	EventLevelNewScoreEvent(context.Context, *EventLevelNewScoreEventRequest) (*EventLevelNewScoreEventResponse, error)
	// FOR TESTING ONLY
	IsAlive(context.Context, *IsAliveRequest) (*IsAliveResponse, error)
	mustEmbedUnimplementedRouteAnalyticsServiceServer()
}

// UnimplementedRouteAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRouteAnalyticsServiceServer struct {
}

func (UnimplementedRouteAnalyticsServiceServer) SessionStartedEvent(context.Context, *SessionStartedEventRequest) (*SessionStartedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionStartedEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) StageCompletedEvent(context.Context, *StageCompletedEventRequest) (*StageCompletedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StageCompletedEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) ChapterCompletedEvent(context.Context, *ChapterCompletedEventRequest) (*ChapterCompletedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChapterCompletedEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) BasicStageCompletedEvent(context.Context, *BasicStageCompletedEventRequest) (*BasicStageCompletedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasicStageCompletedEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) EventStageCompletedEvent(context.Context, *EventStageCompletedEventRequest) (*EventStageCompletedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventStageCompletedEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) BasicLevelNewScoreEvent(context.Context, *BasicLevelNewScoreEventRequest) (*BasicLevelNewScoreEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasicLevelNewScoreEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) EventLevelNewScoreEvent(context.Context, *EventLevelNewScoreEventRequest) (*EventLevelNewScoreEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventLevelNewScoreEvent not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) IsAlive(context.Context, *IsAliveRequest) (*IsAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAlive not implemented")
}
func (UnimplementedRouteAnalyticsServiceServer) mustEmbedUnimplementedRouteAnalyticsServiceServer() {}

// UnsafeRouteAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteAnalyticsServiceServer will
// result in compilation errors.
type UnsafeRouteAnalyticsServiceServer interface {
	mustEmbedUnimplementedRouteAnalyticsServiceServer()
}

func RegisterRouteAnalyticsServiceServer(s grpc.ServiceRegistrar, srv RouteAnalyticsServiceServer) {
	s.RegisterService(&RouteAnalyticsService_ServiceDesc, srv)
}

func _RouteAnalyticsService_SessionStartedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionStartedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).SessionStartedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/SessionStartedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).SessionStartedEvent(ctx, req.(*SessionStartedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_StageCompletedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StageCompletedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).StageCompletedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/StageCompletedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).StageCompletedEvent(ctx, req.(*StageCompletedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_ChapterCompletedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChapterCompletedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).ChapterCompletedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/ChapterCompletedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).ChapterCompletedEvent(ctx, req.(*ChapterCompletedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_BasicStageCompletedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicStageCompletedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).BasicStageCompletedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/BasicStageCompletedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).BasicStageCompletedEvent(ctx, req.(*BasicStageCompletedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_EventStageCompletedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventStageCompletedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).EventStageCompletedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/EventStageCompletedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).EventStageCompletedEvent(ctx, req.(*EventStageCompletedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_BasicLevelNewScoreEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicLevelNewScoreEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).BasicLevelNewScoreEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/BasicLevelNewScoreEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).BasicLevelNewScoreEvent(ctx, req.(*BasicLevelNewScoreEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_EventLevelNewScoreEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventLevelNewScoreEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).EventLevelNewScoreEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/EventLevelNewScoreEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).EventLevelNewScoreEvent(ctx, req.(*EventLevelNewScoreEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteAnalyticsService_IsAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteAnalyticsServiceServer).IsAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteAnalyticsService/IsAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteAnalyticsServiceServer).IsAlive(ctx, req.(*IsAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteAnalyticsService_ServiceDesc is the grpc.ServiceDesc for RouteAnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteAnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RouteAnalyticsService",
	HandlerType: (*RouteAnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SessionStartedEvent",
			Handler:    _RouteAnalyticsService_SessionStartedEvent_Handler,
		},
		{
			MethodName: "StageCompletedEvent",
			Handler:    _RouteAnalyticsService_StageCompletedEvent_Handler,
		},
		{
			MethodName: "ChapterCompletedEvent",
			Handler:    _RouteAnalyticsService_ChapterCompletedEvent_Handler,
		},
		{
			MethodName: "BasicStageCompletedEvent",
			Handler:    _RouteAnalyticsService_BasicStageCompletedEvent_Handler,
		},
		{
			MethodName: "EventStageCompletedEvent",
			Handler:    _RouteAnalyticsService_EventStageCompletedEvent_Handler,
		},
		{
			MethodName: "BasicLevelNewScoreEvent",
			Handler:    _RouteAnalyticsService_BasicLevelNewScoreEvent_Handler,
		},
		{
			MethodName: "EventLevelNewScoreEvent",
			Handler:    _RouteAnalyticsService_EventLevelNewScoreEvent_Handler,
		},
		{
			MethodName: "IsAlive",
			Handler:    _RouteAnalyticsService_IsAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
