// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GripperServiceClient is the client API for GripperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GripperServiceClient interface {
	// Open opens a gripper of the underlying robot.
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
	// Grab requests a gripper of the underlying robot to grab.
	Grab(ctx context.Context, in *GrabRequest, opts ...grpc.CallOption) (*GrabResponse, error)
	// Stop stops a robot's gripper
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type gripperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGripperServiceClient(cc grpc.ClientConnInterface) GripperServiceClient {
	return &gripperServiceClient{cc}
}

func (c *gripperServiceClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.gripper.v1.GripperService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) Grab(ctx context.Context, in *GrabRequest, opts ...grpc.CallOption) (*GrabResponse, error) {
	out := new(GrabResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.gripper.v1.GripperService/Grab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.gripper.v1.GripperService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripperServiceServer is the server API for GripperService service.
// All implementations must embed UnimplementedGripperServiceServer
// for forward compatibility
type GripperServiceServer interface {
	// Open opens a gripper of the underlying robot.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
	// Grab requests a gripper of the underlying robot to grab.
	Grab(context.Context, *GrabRequest) (*GrabResponse, error)
	// Stop stops a robot's gripper
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	mustEmbedUnimplementedGripperServiceServer()
}

// UnimplementedGripperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGripperServiceServer struct {
}

func (UnimplementedGripperServiceServer) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedGripperServiceServer) Grab(context.Context, *GrabRequest) (*GrabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grab not implemented")
}
func (UnimplementedGripperServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedGripperServiceServer) mustEmbedUnimplementedGripperServiceServer() {}

// UnsafeGripperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GripperServiceServer will
// result in compilation errors.
type UnsafeGripperServiceServer interface {
	mustEmbedUnimplementedGripperServiceServer()
}

func RegisterGripperServiceServer(s grpc.ServiceRegistrar, srv GripperServiceServer) {
	s.RegisterService(&GripperService_ServiceDesc, srv)
}

func _GripperService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.gripper.v1.GripperService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_Grab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).Grab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.gripper.v1.GripperService/Grab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).Grab(ctx, req.(*GrabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.gripper.v1.GripperService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GripperService_ServiceDesc is the grpc.ServiceDesc for GripperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GripperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.component.gripper.v1.GripperService",
	HandlerType: (*GripperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _GripperService_Open_Handler,
		},
		{
			MethodName: "Grab",
			Handler:    _GripperService_Grab_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _GripperService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/component/gripper/v1/gripper.proto",
}
