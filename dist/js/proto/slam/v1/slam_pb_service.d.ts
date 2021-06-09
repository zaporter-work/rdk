// package: proto.slam.v1
// file: proto/slam/v1/slam.proto

import * as proto_slam_v1_slam_pb from "../../../proto/slam/v1/slam_pb";
import {grpc} from "@improbable-eng/grpc-web";

type SlamServiceSave = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.SaveRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.SaveResponse;
};

type SlamServiceStats = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.StatsRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.StatsResponse;
};

type SlamServiceCalibrate = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.CalibrateRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.CalibrateResponse;
};

type SlamServiceMoveRobot = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.MoveRobotRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.MoveRobotResponse;
};

type SlamServiceMoveRobotForward = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.MoveRobotForwardRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.MoveRobotForwardResponse;
};

type SlamServiceMoveRobotBackward = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.MoveRobotBackwardRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.MoveRobotBackwardResponse;
};

type SlamServiceTurnRobotTo = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.TurnRobotToRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.TurnRobotToResponse;
};

type SlamServiceUpdateRobotDeviceOffset = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetResponse;
};

type SlamServiceStartLidar = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.StartLidarRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.StartLidarResponse;
};

type SlamServiceStopLidar = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.StopLidarRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.StopLidarResponse;
};

type SlamServiceGetLidarSeed = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.GetLidarSeedRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.GetLidarSeedResponse;
};

type SlamServiceSetLidarSeed = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.SetLidarSeedRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.SetLidarSeedResponse;
};

type SlamServiceSetClientZoom = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.SetClientZoomRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.SetClientZoomResponse;
};

type SlamServiceSetClientLidarViewMode = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.SetClientLidarViewModeRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.SetClientLidarViewModeResponse;
};

type SlamServiceSetClientClickMode = {
  readonly methodName: string;
  readonly service: typeof SlamService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_slam_v1_slam_pb.SetClientClickModeRequest;
  readonly responseType: typeof proto_slam_v1_slam_pb.SetClientClickModeResponse;
};

export class SlamService {
  static readonly serviceName: string;
  static readonly Save: SlamServiceSave;
  static readonly Stats: SlamServiceStats;
  static readonly Calibrate: SlamServiceCalibrate;
  static readonly MoveRobot: SlamServiceMoveRobot;
  static readonly MoveRobotForward: SlamServiceMoveRobotForward;
  static readonly MoveRobotBackward: SlamServiceMoveRobotBackward;
  static readonly TurnRobotTo: SlamServiceTurnRobotTo;
  static readonly UpdateRobotDeviceOffset: SlamServiceUpdateRobotDeviceOffset;
  static readonly StartLidar: SlamServiceStartLidar;
  static readonly StopLidar: SlamServiceStopLidar;
  static readonly GetLidarSeed: SlamServiceGetLidarSeed;
  static readonly SetLidarSeed: SlamServiceSetLidarSeed;
  static readonly SetClientZoom: SlamServiceSetClientZoom;
  static readonly SetClientLidarViewMode: SlamServiceSetClientLidarViewMode;
  static readonly SetClientClickMode: SlamServiceSetClientClickMode;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class SlamServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  save(
    requestMessage: proto_slam_v1_slam_pb.SaveRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SaveResponse|null) => void
  ): UnaryResponse;
  save(
    requestMessage: proto_slam_v1_slam_pb.SaveRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SaveResponse|null) => void
  ): UnaryResponse;
  stats(
    requestMessage: proto_slam_v1_slam_pb.StatsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.StatsResponse|null) => void
  ): UnaryResponse;
  stats(
    requestMessage: proto_slam_v1_slam_pb.StatsRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.StatsResponse|null) => void
  ): UnaryResponse;
  calibrate(
    requestMessage: proto_slam_v1_slam_pb.CalibrateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.CalibrateResponse|null) => void
  ): UnaryResponse;
  calibrate(
    requestMessage: proto_slam_v1_slam_pb.CalibrateRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.CalibrateResponse|null) => void
  ): UnaryResponse;
  moveRobot(
    requestMessage: proto_slam_v1_slam_pb.MoveRobotRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.MoveRobotResponse|null) => void
  ): UnaryResponse;
  moveRobot(
    requestMessage: proto_slam_v1_slam_pb.MoveRobotRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.MoveRobotResponse|null) => void
  ): UnaryResponse;
  moveRobotForward(
    requestMessage: proto_slam_v1_slam_pb.MoveRobotForwardRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.MoveRobotForwardResponse|null) => void
  ): UnaryResponse;
  moveRobotForward(
    requestMessage: proto_slam_v1_slam_pb.MoveRobotForwardRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.MoveRobotForwardResponse|null) => void
  ): UnaryResponse;
  moveRobotBackward(
    requestMessage: proto_slam_v1_slam_pb.MoveRobotBackwardRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.MoveRobotBackwardResponse|null) => void
  ): UnaryResponse;
  moveRobotBackward(
    requestMessage: proto_slam_v1_slam_pb.MoveRobotBackwardRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.MoveRobotBackwardResponse|null) => void
  ): UnaryResponse;
  turnRobotTo(
    requestMessage: proto_slam_v1_slam_pb.TurnRobotToRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.TurnRobotToResponse|null) => void
  ): UnaryResponse;
  turnRobotTo(
    requestMessage: proto_slam_v1_slam_pb.TurnRobotToRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.TurnRobotToResponse|null) => void
  ): UnaryResponse;
  updateRobotDeviceOffset(
    requestMessage: proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetResponse|null) => void
  ): UnaryResponse;
  updateRobotDeviceOffset(
    requestMessage: proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetResponse|null) => void
  ): UnaryResponse;
  startLidar(
    requestMessage: proto_slam_v1_slam_pb.StartLidarRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.StartLidarResponse|null) => void
  ): UnaryResponse;
  startLidar(
    requestMessage: proto_slam_v1_slam_pb.StartLidarRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.StartLidarResponse|null) => void
  ): UnaryResponse;
  stopLidar(
    requestMessage: proto_slam_v1_slam_pb.StopLidarRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.StopLidarResponse|null) => void
  ): UnaryResponse;
  stopLidar(
    requestMessage: proto_slam_v1_slam_pb.StopLidarRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.StopLidarResponse|null) => void
  ): UnaryResponse;
  getLidarSeed(
    requestMessage: proto_slam_v1_slam_pb.GetLidarSeedRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.GetLidarSeedResponse|null) => void
  ): UnaryResponse;
  getLidarSeed(
    requestMessage: proto_slam_v1_slam_pb.GetLidarSeedRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.GetLidarSeedResponse|null) => void
  ): UnaryResponse;
  setLidarSeed(
    requestMessage: proto_slam_v1_slam_pb.SetLidarSeedRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetLidarSeedResponse|null) => void
  ): UnaryResponse;
  setLidarSeed(
    requestMessage: proto_slam_v1_slam_pb.SetLidarSeedRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetLidarSeedResponse|null) => void
  ): UnaryResponse;
  setClientZoom(
    requestMessage: proto_slam_v1_slam_pb.SetClientZoomRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetClientZoomResponse|null) => void
  ): UnaryResponse;
  setClientZoom(
    requestMessage: proto_slam_v1_slam_pb.SetClientZoomRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetClientZoomResponse|null) => void
  ): UnaryResponse;
  setClientLidarViewMode(
    requestMessage: proto_slam_v1_slam_pb.SetClientLidarViewModeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetClientLidarViewModeResponse|null) => void
  ): UnaryResponse;
  setClientLidarViewMode(
    requestMessage: proto_slam_v1_slam_pb.SetClientLidarViewModeRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetClientLidarViewModeResponse|null) => void
  ): UnaryResponse;
  setClientClickMode(
    requestMessage: proto_slam_v1_slam_pb.SetClientClickModeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetClientClickModeResponse|null) => void
  ): UnaryResponse;
  setClientClickMode(
    requestMessage: proto_slam_v1_slam_pb.SetClientClickModeRequest,
    callback: (error: ServiceError|null, responseMessage: proto_slam_v1_slam_pb.SetClientClickModeResponse|null) => void
  ): UnaryResponse;
}

