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

// SlamServiceClient is the client API for SlamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlamServiceClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	Calibrate(ctx context.Context, in *CalibrateRequest, opts ...grpc.CallOption) (*CalibrateResponse, error)
	MoveRobot(ctx context.Context, in *MoveRobotRequest, opts ...grpc.CallOption) (*MoveRobotResponse, error)
	MoveRobotForward(ctx context.Context, in *MoveRobotForwardRequest, opts ...grpc.CallOption) (*MoveRobotForwardResponse, error)
	MoveRobotBackward(ctx context.Context, in *MoveRobotBackwardRequest, opts ...grpc.CallOption) (*MoveRobotBackwardResponse, error)
	TurnRobotTo(ctx context.Context, in *TurnRobotToRequest, opts ...grpc.CallOption) (*TurnRobotToResponse, error)
	UpdateRobotDeviceOffset(ctx context.Context, in *UpdateRobotDeviceOffsetRequest, opts ...grpc.CallOption) (*UpdateRobotDeviceOffsetResponse, error)
	StartLidar(ctx context.Context, in *StartLidarRequest, opts ...grpc.CallOption) (*StartLidarResponse, error)
	StopLidar(ctx context.Context, in *StopLidarRequest, opts ...grpc.CallOption) (*StopLidarResponse, error)
	GetLidarSeed(ctx context.Context, in *GetLidarSeedRequest, opts ...grpc.CallOption) (*GetLidarSeedResponse, error)
	SetLidarSeed(ctx context.Context, in *SetLidarSeedRequest, opts ...grpc.CallOption) (*SetLidarSeedResponse, error)
	SetClientZoom(ctx context.Context, in *SetClientZoomRequest, opts ...grpc.CallOption) (*SetClientZoomResponse, error)
	SetClientLidarViewMode(ctx context.Context, in *SetClientLidarViewModeRequest, opts ...grpc.CallOption) (*SetClientLidarViewModeResponse, error)
	SetClientClickMode(ctx context.Context, in *SetClientClickModeRequest, opts ...grpc.CallOption) (*SetClientClickModeResponse, error)
}

type slamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSlamServiceClient(cc grpc.ClientConnInterface) SlamServiceClient {
	return &slamServiceClient{cc}
}

func (c *slamServiceClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) Calibrate(ctx context.Context, in *CalibrateRequest, opts ...grpc.CallOption) (*CalibrateResponse, error) {
	out := new(CalibrateResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/Calibrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) MoveRobot(ctx context.Context, in *MoveRobotRequest, opts ...grpc.CallOption) (*MoveRobotResponse, error) {
	out := new(MoveRobotResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/MoveRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) MoveRobotForward(ctx context.Context, in *MoveRobotForwardRequest, opts ...grpc.CallOption) (*MoveRobotForwardResponse, error) {
	out := new(MoveRobotForwardResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/MoveRobotForward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) MoveRobotBackward(ctx context.Context, in *MoveRobotBackwardRequest, opts ...grpc.CallOption) (*MoveRobotBackwardResponse, error) {
	out := new(MoveRobotBackwardResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/MoveRobotBackward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) TurnRobotTo(ctx context.Context, in *TurnRobotToRequest, opts ...grpc.CallOption) (*TurnRobotToResponse, error) {
	out := new(TurnRobotToResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/TurnRobotTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) UpdateRobotDeviceOffset(ctx context.Context, in *UpdateRobotDeviceOffsetRequest, opts ...grpc.CallOption) (*UpdateRobotDeviceOffsetResponse, error) {
	out := new(UpdateRobotDeviceOffsetResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/UpdateRobotDeviceOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) StartLidar(ctx context.Context, in *StartLidarRequest, opts ...grpc.CallOption) (*StartLidarResponse, error) {
	out := new(StartLidarResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/StartLidar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) StopLidar(ctx context.Context, in *StopLidarRequest, opts ...grpc.CallOption) (*StopLidarResponse, error) {
	out := new(StopLidarResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/StopLidar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) GetLidarSeed(ctx context.Context, in *GetLidarSeedRequest, opts ...grpc.CallOption) (*GetLidarSeedResponse, error) {
	out := new(GetLidarSeedResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/GetLidarSeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) SetLidarSeed(ctx context.Context, in *SetLidarSeedRequest, opts ...grpc.CallOption) (*SetLidarSeedResponse, error) {
	out := new(SetLidarSeedResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/SetLidarSeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) SetClientZoom(ctx context.Context, in *SetClientZoomRequest, opts ...grpc.CallOption) (*SetClientZoomResponse, error) {
	out := new(SetClientZoomResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/SetClientZoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) SetClientLidarViewMode(ctx context.Context, in *SetClientLidarViewModeRequest, opts ...grpc.CallOption) (*SetClientLidarViewModeResponse, error) {
	out := new(SetClientLidarViewModeResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/SetClientLidarViewMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slamServiceClient) SetClientClickMode(ctx context.Context, in *SetClientClickModeRequest, opts ...grpc.CallOption) (*SetClientClickModeResponse, error) {
	out := new(SetClientClickModeResponse)
	err := c.cc.Invoke(ctx, "/proto.slam.v1.SlamService/SetClientClickMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlamServiceServer is the server API for SlamService service.
// All implementations must embed UnimplementedSlamServiceServer
// for forward compatibility
type SlamServiceServer interface {
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	Calibrate(context.Context, *CalibrateRequest) (*CalibrateResponse, error)
	MoveRobot(context.Context, *MoveRobotRequest) (*MoveRobotResponse, error)
	MoveRobotForward(context.Context, *MoveRobotForwardRequest) (*MoveRobotForwardResponse, error)
	MoveRobotBackward(context.Context, *MoveRobotBackwardRequest) (*MoveRobotBackwardResponse, error)
	TurnRobotTo(context.Context, *TurnRobotToRequest) (*TurnRobotToResponse, error)
	UpdateRobotDeviceOffset(context.Context, *UpdateRobotDeviceOffsetRequest) (*UpdateRobotDeviceOffsetResponse, error)
	StartLidar(context.Context, *StartLidarRequest) (*StartLidarResponse, error)
	StopLidar(context.Context, *StopLidarRequest) (*StopLidarResponse, error)
	GetLidarSeed(context.Context, *GetLidarSeedRequest) (*GetLidarSeedResponse, error)
	SetLidarSeed(context.Context, *SetLidarSeedRequest) (*SetLidarSeedResponse, error)
	SetClientZoom(context.Context, *SetClientZoomRequest) (*SetClientZoomResponse, error)
	SetClientLidarViewMode(context.Context, *SetClientLidarViewModeRequest) (*SetClientLidarViewModeResponse, error)
	SetClientClickMode(context.Context, *SetClientClickModeRequest) (*SetClientClickModeResponse, error)
	mustEmbedUnimplementedSlamServiceServer()
}

// UnimplementedSlamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSlamServiceServer struct {
}

func (UnimplementedSlamServiceServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedSlamServiceServer) Stats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedSlamServiceServer) Calibrate(context.Context, *CalibrateRequest) (*CalibrateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calibrate not implemented")
}
func (UnimplementedSlamServiceServer) MoveRobot(context.Context, *MoveRobotRequest) (*MoveRobotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveRobot not implemented")
}
func (UnimplementedSlamServiceServer) MoveRobotForward(context.Context, *MoveRobotForwardRequest) (*MoveRobotForwardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveRobotForward not implemented")
}
func (UnimplementedSlamServiceServer) MoveRobotBackward(context.Context, *MoveRobotBackwardRequest) (*MoveRobotBackwardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveRobotBackward not implemented")
}
func (UnimplementedSlamServiceServer) TurnRobotTo(context.Context, *TurnRobotToRequest) (*TurnRobotToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TurnRobotTo not implemented")
}
func (UnimplementedSlamServiceServer) UpdateRobotDeviceOffset(context.Context, *UpdateRobotDeviceOffsetRequest) (*UpdateRobotDeviceOffsetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRobotDeviceOffset not implemented")
}
func (UnimplementedSlamServiceServer) StartLidar(context.Context, *StartLidarRequest) (*StartLidarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLidar not implemented")
}
func (UnimplementedSlamServiceServer) StopLidar(context.Context, *StopLidarRequest) (*StopLidarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopLidar not implemented")
}
func (UnimplementedSlamServiceServer) GetLidarSeed(context.Context, *GetLidarSeedRequest) (*GetLidarSeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLidarSeed not implemented")
}
func (UnimplementedSlamServiceServer) SetLidarSeed(context.Context, *SetLidarSeedRequest) (*SetLidarSeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLidarSeed not implemented")
}
func (UnimplementedSlamServiceServer) SetClientZoom(context.Context, *SetClientZoomRequest) (*SetClientZoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetClientZoom not implemented")
}
func (UnimplementedSlamServiceServer) SetClientLidarViewMode(context.Context, *SetClientLidarViewModeRequest) (*SetClientLidarViewModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetClientLidarViewMode not implemented")
}
func (UnimplementedSlamServiceServer) SetClientClickMode(context.Context, *SetClientClickModeRequest) (*SetClientClickModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetClientClickMode not implemented")
}
func (UnimplementedSlamServiceServer) mustEmbedUnimplementedSlamServiceServer() {}

// UnsafeSlamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlamServiceServer will
// result in compilation errors.
type UnsafeSlamServiceServer interface {
	mustEmbedUnimplementedSlamServiceServer()
}

func RegisterSlamServiceServer(s grpc.ServiceRegistrar, srv SlamServiceServer) {
	s.RegisterService(&SlamService_ServiceDesc, srv)
}

func _SlamService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_Calibrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalibrateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).Calibrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/Calibrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).Calibrate(ctx, req.(*CalibrateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_MoveRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).MoveRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/MoveRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).MoveRobot(ctx, req.(*MoveRobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_MoveRobotForward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRobotForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).MoveRobotForward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/MoveRobotForward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).MoveRobotForward(ctx, req.(*MoveRobotForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_MoveRobotBackward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRobotBackwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).MoveRobotBackward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/MoveRobotBackward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).MoveRobotBackward(ctx, req.(*MoveRobotBackwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_TurnRobotTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnRobotToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).TurnRobotTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/TurnRobotTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).TurnRobotTo(ctx, req.(*TurnRobotToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_UpdateRobotDeviceOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRobotDeviceOffsetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).UpdateRobotDeviceOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/UpdateRobotDeviceOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).UpdateRobotDeviceOffset(ctx, req.(*UpdateRobotDeviceOffsetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_StartLidar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLidarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).StartLidar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/StartLidar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).StartLidar(ctx, req.(*StartLidarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_StopLidar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopLidarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).StopLidar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/StopLidar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).StopLidar(ctx, req.(*StopLidarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_GetLidarSeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLidarSeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).GetLidarSeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/GetLidarSeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).GetLidarSeed(ctx, req.(*GetLidarSeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_SetLidarSeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLidarSeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).SetLidarSeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/SetLidarSeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).SetLidarSeed(ctx, req.(*SetLidarSeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_SetClientZoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientZoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).SetClientZoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/SetClientZoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).SetClientZoom(ctx, req.(*SetClientZoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_SetClientLidarViewMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientLidarViewModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).SetClientLidarViewMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/SetClientLidarViewMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).SetClientLidarViewMode(ctx, req.(*SetClientLidarViewModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlamService_SetClientClickMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientClickModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlamServiceServer).SetClientClickMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.slam.v1.SlamService/SetClientClickMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlamServiceServer).SetClientClickMode(ctx, req.(*SetClientClickModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SlamService_ServiceDesc is the grpc.ServiceDesc for SlamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SlamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.slam.v1.SlamService",
	HandlerType: (*SlamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _SlamService_Save_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _SlamService_Stats_Handler,
		},
		{
			MethodName: "Calibrate",
			Handler:    _SlamService_Calibrate_Handler,
		},
		{
			MethodName: "MoveRobot",
			Handler:    _SlamService_MoveRobot_Handler,
		},
		{
			MethodName: "MoveRobotForward",
			Handler:    _SlamService_MoveRobotForward_Handler,
		},
		{
			MethodName: "MoveRobotBackward",
			Handler:    _SlamService_MoveRobotBackward_Handler,
		},
		{
			MethodName: "TurnRobotTo",
			Handler:    _SlamService_TurnRobotTo_Handler,
		},
		{
			MethodName: "UpdateRobotDeviceOffset",
			Handler:    _SlamService_UpdateRobotDeviceOffset_Handler,
		},
		{
			MethodName: "StartLidar",
			Handler:    _SlamService_StartLidar_Handler,
		},
		{
			MethodName: "StopLidar",
			Handler:    _SlamService_StopLidar_Handler,
		},
		{
			MethodName: "GetLidarSeed",
			Handler:    _SlamService_GetLidarSeed_Handler,
		},
		{
			MethodName: "SetLidarSeed",
			Handler:    _SlamService_SetLidarSeed_Handler,
		},
		{
			MethodName: "SetClientZoom",
			Handler:    _SlamService_SetClientZoom_Handler,
		},
		{
			MethodName: "SetClientLidarViewMode",
			Handler:    _SlamService_SetClientLidarViewMode_Handler,
		},
		{
			MethodName: "SetClientClickMode",
			Handler:    _SlamService_SetClientClickMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/slam/v1/slam.proto",
}
