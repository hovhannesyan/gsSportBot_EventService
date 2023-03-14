// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/pb/event.proto

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

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventClient interface {
	AddDiscipline(ctx context.Context, in *AddDisciplineRequest, opts ...grpc.CallOption) (*AddDisciplineResponse, error)
	DeleteDiscipline(ctx context.Context, in *DeleteDisciplineRequest, opts ...grpc.CallOption) (*DeleteDisciplineResponse, error)
	GetDisciplines(ctx context.Context, in *GetDisciplinesRequest, opts ...grpc.CallOption) (*GetDisciplinesResponse, error)
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) AddDiscipline(ctx context.Context, in *AddDisciplineRequest, opts ...grpc.CallOption) (*AddDisciplineResponse, error) {
	out := new(AddDisciplineResponse)
	err := c.cc.Invoke(ctx, "/event.Event/AddDiscipline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) DeleteDiscipline(ctx context.Context, in *DeleteDisciplineRequest, opts ...grpc.CallOption) (*DeleteDisciplineResponse, error) {
	out := new(DeleteDisciplineResponse)
	err := c.cc.Invoke(ctx, "/event.Event/DeleteDiscipline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) GetDisciplines(ctx context.Context, in *GetDisciplinesRequest, opts ...grpc.CallOption) (*GetDisciplinesResponse, error) {
	out := new(GetDisciplinesResponse)
	err := c.cc.Invoke(ctx, "/event.Event/GetDisciplines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, "/event.Event/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, "/event.Event/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/event.Event/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServer is the server API for Event service.
// All implementations must embed UnimplementedEventServer
// for forward compatibility
type EventServer interface {
	AddDiscipline(context.Context, *AddDisciplineRequest) (*AddDisciplineResponse, error)
	DeleteDiscipline(context.Context, *DeleteDisciplineRequest) (*DeleteDisciplineResponse, error)
	GetDisciplines(context.Context, *GetDisciplinesRequest) (*GetDisciplinesResponse, error)
	AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	mustEmbedUnimplementedEventServer()
}

// UnimplementedEventServer must be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) AddDiscipline(context.Context, *AddDisciplineRequest) (*AddDisciplineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDiscipline not implemented")
}
func (UnimplementedEventServer) DeleteDiscipline(context.Context, *DeleteDisciplineRequest) (*DeleteDisciplineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDiscipline not implemented")
}
func (UnimplementedEventServer) GetDisciplines(context.Context, *GetDisciplinesRequest) (*GetDisciplinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisciplines not implemented")
}
func (UnimplementedEventServer) AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedEventServer) DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {}

// UnsafeEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServer will
// result in compilation errors.
type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	s.RegisterService(&Event_ServiceDesc, srv)
}

func _Event_AddDiscipline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDisciplineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).AddDiscipline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Event/AddDiscipline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).AddDiscipline(ctx, req.(*AddDisciplineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_DeleteDiscipline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDisciplineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).DeleteDiscipline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Event/DeleteDiscipline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).DeleteDiscipline(ctx, req.(*DeleteDisciplineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_GetDisciplines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDisciplinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).GetDisciplines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Event/GetDisciplines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).GetDisciplines(ctx, req.(*GetDisciplinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Event/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Event/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Event/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Event_ServiceDesc is the grpc.ServiceDesc for Event service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDiscipline",
			Handler:    _Event_AddDiscipline_Handler,
		},
		{
			MethodName: "DeleteDiscipline",
			Handler:    _Event_DeleteDiscipline_Handler,
		},
		{
			MethodName: "GetDisciplines",
			Handler:    _Event_GetDisciplines_Handler,
		},
		{
			MethodName: "AddEvent",
			Handler:    _Event_AddEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Event_DeleteEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _Event_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/event.proto",
}
