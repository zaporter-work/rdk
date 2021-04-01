// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RobotServiceClient is the client API for RobotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotServiceClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	StatusStream(ctx context.Context, in *StatusStreamRequest, opts ...grpc.CallOption) (RobotService_StatusStreamClient, error)
	DoAction(ctx context.Context, in *DoActionRequest, opts ...grpc.CallOption) (*DoActionResponse, error)
	ControlBase(ctx context.Context, in *ControlBaseRequest, opts ...grpc.CallOption) (*ControlBaseResponse, error)
	ArmCurrentPosition(ctx context.Context, in *ArmCurrentPositionRequest, opts ...grpc.CallOption) (*ArmCurrentPositionResponse, error)
	ArmCurrentJointPositions(ctx context.Context, in *ArmCurrentJointPositionsRequest, opts ...grpc.CallOption) (*ArmCurrentJointPositionsResponse, error)
	MoveArmToPosition(ctx context.Context, in *MoveArmToPositionRequest, opts ...grpc.CallOption) (*MoveArmToPositionResponse, error)
	MoveArmToJointPositions(ctx context.Context, in *MoveArmToJointPositionsRequest, opts ...grpc.CallOption) (*MoveArmToJointPositionsResponse, error)
	ControlGripper(ctx context.Context, in *ControlGripperRequest, opts ...grpc.CallOption) (*ControlGripperResponse, error)
	BoardStatus(ctx context.Context, in *BoardStatusRequest, opts ...grpc.CallOption) (*BoardStatusResponse, error)
	ControlBoardMotor(ctx context.Context, in *ControlBoardMotorRequest, opts ...grpc.CallOption) (*ControlBoardMotorResponse, error)
	ControlBoardServo(ctx context.Context, in *ControlBoardServoRequest, opts ...grpc.CallOption) (*ControlBoardServoResponse, error)
	CameraFrame(ctx context.Context, in *CameraFrameRequest, opts ...grpc.CallOption) (*CameraFrameResponse, error)
	RenderCameraFrame(ctx context.Context, in *CameraFrameRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type robotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotServiceClient(cc grpc.ClientConnInterface) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) StatusStream(ctx context.Context, in *StatusStreamRequest, opts ...grpc.CallOption) (RobotService_StatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RobotService_ServiceDesc.Streams[0], "/proto.api.v1.RobotService/StatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotServiceStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RobotService_StatusStreamClient interface {
	Recv() (*StatusStreamResponse, error)
	grpc.ClientStream
}

type robotServiceStatusStreamClient struct {
	grpc.ClientStream
}

func (x *robotServiceStatusStreamClient) Recv() (*StatusStreamResponse, error) {
	m := new(StatusStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *robotServiceClient) DoAction(ctx context.Context, in *DoActionRequest, opts ...grpc.CallOption) (*DoActionResponse, error) {
	out := new(DoActionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/DoAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ControlBase(ctx context.Context, in *ControlBaseRequest, opts ...grpc.CallOption) (*ControlBaseResponse, error) {
	out := new(ControlBaseResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/ControlBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ArmCurrentPosition(ctx context.Context, in *ArmCurrentPositionRequest, opts ...grpc.CallOption) (*ArmCurrentPositionResponse, error) {
	out := new(ArmCurrentPositionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/ArmCurrentPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ArmCurrentJointPositions(ctx context.Context, in *ArmCurrentJointPositionsRequest, opts ...grpc.CallOption) (*ArmCurrentJointPositionsResponse, error) {
	out := new(ArmCurrentJointPositionsResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/ArmCurrentJointPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) MoveArmToPosition(ctx context.Context, in *MoveArmToPositionRequest, opts ...grpc.CallOption) (*MoveArmToPositionResponse, error) {
	out := new(MoveArmToPositionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/MoveArmToPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) MoveArmToJointPositions(ctx context.Context, in *MoveArmToJointPositionsRequest, opts ...grpc.CallOption) (*MoveArmToJointPositionsResponse, error) {
	out := new(MoveArmToJointPositionsResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/MoveArmToJointPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ControlGripper(ctx context.Context, in *ControlGripperRequest, opts ...grpc.CallOption) (*ControlGripperResponse, error) {
	out := new(ControlGripperResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/ControlGripper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) BoardStatus(ctx context.Context, in *BoardStatusRequest, opts ...grpc.CallOption) (*BoardStatusResponse, error) {
	out := new(BoardStatusResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/BoardStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ControlBoardMotor(ctx context.Context, in *ControlBoardMotorRequest, opts ...grpc.CallOption) (*ControlBoardMotorResponse, error) {
	out := new(ControlBoardMotorResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/ControlBoardMotor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ControlBoardServo(ctx context.Context, in *ControlBoardServoRequest, opts ...grpc.CallOption) (*ControlBoardServoResponse, error) {
	out := new(ControlBoardServoResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/ControlBoardServo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) CameraFrame(ctx context.Context, in *CameraFrameRequest, opts ...grpc.CallOption) (*CameraFrameResponse, error) {
	out := new(CameraFrameResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/CameraFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) RenderCameraFrame(ctx context.Context, in *CameraFrameRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/proto.api.v1.RobotService/RenderCameraFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServiceServer is the server API for RobotService service.
// All implementations must embed UnimplementedRobotServiceServer
// for forward compatibility
type RobotServiceServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	StatusStream(*StatusStreamRequest, RobotService_StatusStreamServer) error
	DoAction(context.Context, *DoActionRequest) (*DoActionResponse, error)
	ControlBase(context.Context, *ControlBaseRequest) (*ControlBaseResponse, error)
	ArmCurrentPosition(context.Context, *ArmCurrentPositionRequest) (*ArmCurrentPositionResponse, error)
	ArmCurrentJointPositions(context.Context, *ArmCurrentJointPositionsRequest) (*ArmCurrentJointPositionsResponse, error)
	MoveArmToPosition(context.Context, *MoveArmToPositionRequest) (*MoveArmToPositionResponse, error)
	MoveArmToJointPositions(context.Context, *MoveArmToJointPositionsRequest) (*MoveArmToJointPositionsResponse, error)
	ControlGripper(context.Context, *ControlGripperRequest) (*ControlGripperResponse, error)
	BoardStatus(context.Context, *BoardStatusRequest) (*BoardStatusResponse, error)
	ControlBoardMotor(context.Context, *ControlBoardMotorRequest) (*ControlBoardMotorResponse, error)
	ControlBoardServo(context.Context, *ControlBoardServoRequest) (*ControlBoardServoResponse, error)
	CameraFrame(context.Context, *CameraFrameRequest) (*CameraFrameResponse, error)
	RenderCameraFrame(context.Context, *CameraFrameRequest) (*httpbody.HttpBody, error)
	mustEmbedUnimplementedRobotServiceServer()
}

// UnimplementedRobotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotServiceServer struct {
}

func (UnimplementedRobotServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedRobotServiceServer) StatusStream(*StatusStreamRequest, RobotService_StatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StatusStream not implemented")
}
func (UnimplementedRobotServiceServer) DoAction(context.Context, *DoActionRequest) (*DoActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoAction not implemented")
}
func (UnimplementedRobotServiceServer) ControlBase(context.Context, *ControlBaseRequest) (*ControlBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlBase not implemented")
}
func (UnimplementedRobotServiceServer) ArmCurrentPosition(context.Context, *ArmCurrentPositionRequest) (*ArmCurrentPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArmCurrentPosition not implemented")
}
func (UnimplementedRobotServiceServer) ArmCurrentJointPositions(context.Context, *ArmCurrentJointPositionsRequest) (*ArmCurrentJointPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArmCurrentJointPositions not implemented")
}
func (UnimplementedRobotServiceServer) MoveArmToPosition(context.Context, *MoveArmToPositionRequest) (*MoveArmToPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveArmToPosition not implemented")
}
func (UnimplementedRobotServiceServer) MoveArmToJointPositions(context.Context, *MoveArmToJointPositionsRequest) (*MoveArmToJointPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveArmToJointPositions not implemented")
}
func (UnimplementedRobotServiceServer) ControlGripper(context.Context, *ControlGripperRequest) (*ControlGripperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlGripper not implemented")
}
func (UnimplementedRobotServiceServer) BoardStatus(context.Context, *BoardStatusRequest) (*BoardStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BoardStatus not implemented")
}
func (UnimplementedRobotServiceServer) ControlBoardMotor(context.Context, *ControlBoardMotorRequest) (*ControlBoardMotorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlBoardMotor not implemented")
}
func (UnimplementedRobotServiceServer) ControlBoardServo(context.Context, *ControlBoardServoRequest) (*ControlBoardServoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControlBoardServo not implemented")
}
func (UnimplementedRobotServiceServer) CameraFrame(context.Context, *CameraFrameRequest) (*CameraFrameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CameraFrame not implemented")
}
func (UnimplementedRobotServiceServer) RenderCameraFrame(context.Context, *CameraFrameRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderCameraFrame not implemented")
}
func (UnimplementedRobotServiceServer) mustEmbedUnimplementedRobotServiceServer() {}

// UnsafeRobotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotServiceServer will
// result in compilation errors.
type UnsafeRobotServiceServer interface {
	mustEmbedUnimplementedRobotServiceServer()
}

func RegisterRobotServiceServer(s grpc.ServiceRegistrar, srv RobotServiceServer) {
	s.RegisterService(&RobotService_ServiceDesc, srv)
}

func _RobotService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_StatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatusStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RobotServiceServer).StatusStream(m, &robotServiceStatusStreamServer{stream})
}

type RobotService_StatusStreamServer interface {
	Send(*StatusStreamResponse) error
	grpc.ServerStream
}

type robotServiceStatusStreamServer struct {
	grpc.ServerStream
}

func (x *robotServiceStatusStreamServer) Send(m *StatusStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RobotService_DoAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).DoAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/DoAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).DoAction(ctx, req.(*DoActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ControlBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ControlBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/ControlBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ControlBase(ctx, req.(*ControlBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ArmCurrentPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArmCurrentPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ArmCurrentPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/ArmCurrentPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ArmCurrentPosition(ctx, req.(*ArmCurrentPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ArmCurrentJointPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArmCurrentJointPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ArmCurrentJointPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/ArmCurrentJointPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ArmCurrentJointPositions(ctx, req.(*ArmCurrentJointPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_MoveArmToPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveArmToPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).MoveArmToPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/MoveArmToPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).MoveArmToPosition(ctx, req.(*MoveArmToPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_MoveArmToJointPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveArmToJointPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).MoveArmToJointPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/MoveArmToJointPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).MoveArmToJointPositions(ctx, req.(*MoveArmToJointPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ControlGripper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlGripperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ControlGripper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/ControlGripper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ControlGripper(ctx, req.(*ControlGripperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_BoardStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).BoardStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/BoardStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).BoardStatus(ctx, req.(*BoardStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ControlBoardMotor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlBoardMotorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ControlBoardMotor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/ControlBoardMotor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ControlBoardMotor(ctx, req.(*ControlBoardMotorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ControlBoardServo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlBoardServoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ControlBoardServo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/ControlBoardServo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ControlBoardServo(ctx, req.(*ControlBoardServoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_CameraFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CameraFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).CameraFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/CameraFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).CameraFrame(ctx, req.(*CameraFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_RenderCameraFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CameraFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).RenderCameraFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.RobotService/RenderCameraFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).RenderCameraFrame(ctx, req.(*CameraFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotService_ServiceDesc is the grpc.ServiceDesc for RobotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.v1.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _RobotService_Status_Handler,
		},
		{
			MethodName: "DoAction",
			Handler:    _RobotService_DoAction_Handler,
		},
		{
			MethodName: "ControlBase",
			Handler:    _RobotService_ControlBase_Handler,
		},
		{
			MethodName: "ArmCurrentPosition",
			Handler:    _RobotService_ArmCurrentPosition_Handler,
		},
		{
			MethodName: "ArmCurrentJointPositions",
			Handler:    _RobotService_ArmCurrentJointPositions_Handler,
		},
		{
			MethodName: "MoveArmToPosition",
			Handler:    _RobotService_MoveArmToPosition_Handler,
		},
		{
			MethodName: "MoveArmToJointPositions",
			Handler:    _RobotService_MoveArmToJointPositions_Handler,
		},
		{
			MethodName: "ControlGripper",
			Handler:    _RobotService_ControlGripper_Handler,
		},
		{
			MethodName: "BoardStatus",
			Handler:    _RobotService_BoardStatus_Handler,
		},
		{
			MethodName: "ControlBoardMotor",
			Handler:    _RobotService_ControlBoardMotor_Handler,
		},
		{
			MethodName: "ControlBoardServo",
			Handler:    _RobotService_ControlBoardServo_Handler,
		},
		{
			MethodName: "CameraFrame",
			Handler:    _RobotService_CameraFrame_Handler,
		},
		{
			MethodName: "RenderCameraFrame",
			Handler:    _RobotService_RenderCameraFrame_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StatusStream",
			Handler:       _RobotService_StatusStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api/v1/robot.proto",
}
