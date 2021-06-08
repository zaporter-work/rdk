// package: proto.slam.v1
// file: proto/slam/v1/slam.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../../google/api/annotations_pb";

export class SaveRequest extends jspb.Message {
  getFile(): string;
  setFile(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SaveRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SaveRequest): SaveRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SaveRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SaveRequest;
  static deserializeBinaryFromReader(message: SaveRequest, reader: jspb.BinaryReader): SaveRequest;
}

export namespace SaveRequest {
  export type AsObject = {
    file: string,
  }
}

export class SaveResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SaveResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SaveResponse): SaveResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SaveResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SaveResponse;
  static deserializeBinaryFromReader(message: SaveResponse, reader: jspb.BinaryReader): SaveResponse;
}

export namespace SaveResponse {
  export type AsObject = {
  }
}

export class StatsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StatsRequest): StatsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatsRequest;
  static deserializeBinaryFromReader(message: StatsRequest, reader: jspb.BinaryReader): StatsRequest;
}

export namespace StatsRequest {
  export type AsObject = {
  }
}

export class StatsResponse extends jspb.Message {
  hasBasePosition(): boolean;
  clearBasePosition(): void;
  getBasePosition(): BasePosition | undefined;
  setBasePosition(value?: BasePosition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StatsResponse): StatsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatsResponse;
  static deserializeBinaryFromReader(message: StatsResponse, reader: jspb.BinaryReader): StatsResponse;
}

export namespace StatsResponse {
  export type AsObject = {
    basePosition?: BasePosition.AsObject,
  }
}

export class BasePosition extends jspb.Message {
  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BasePosition.AsObject;
  static toObject(includeInstance: boolean, msg: BasePosition): BasePosition.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BasePosition, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BasePosition;
  static deserializeBinaryFromReader(message: BasePosition, reader: jspb.BinaryReader): BasePosition;
}

export namespace BasePosition {
  export type AsObject = {
    x: number,
    y: number,
  }
}

export class CalibrateRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CalibrateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CalibrateRequest): CalibrateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CalibrateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CalibrateRequest;
  static deserializeBinaryFromReader(message: CalibrateRequest, reader: jspb.BinaryReader): CalibrateRequest;
}

export namespace CalibrateRequest {
  export type AsObject = {
  }
}

export class CalibrateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CalibrateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CalibrateResponse): CalibrateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CalibrateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CalibrateResponse;
  static deserializeBinaryFromReader(message: CalibrateResponse, reader: jspb.BinaryReader): CalibrateResponse;
}

export namespace CalibrateResponse {
  export type AsObject = {
  }
}

export class MoveRobotRequest extends jspb.Message {
  getDirection(): DirectionMap[keyof DirectionMap];
  setDirection(value: DirectionMap[keyof DirectionMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveRobotRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MoveRobotRequest): MoveRobotRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveRobotRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveRobotRequest;
  static deserializeBinaryFromReader(message: MoveRobotRequest, reader: jspb.BinaryReader): MoveRobotRequest;
}

export namespace MoveRobotRequest {
  export type AsObject = {
    direction: DirectionMap[keyof DirectionMap],
  }
}

export class MoveRobotResponse extends jspb.Message {
  hasNewPosition(): boolean;
  clearNewPosition(): void;
  getNewPosition(): BasePosition | undefined;
  setNewPosition(value?: BasePosition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveRobotResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MoveRobotResponse): MoveRobotResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveRobotResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveRobotResponse;
  static deserializeBinaryFromReader(message: MoveRobotResponse, reader: jspb.BinaryReader): MoveRobotResponse;
}

export namespace MoveRobotResponse {
  export type AsObject = {
    newPosition?: BasePosition.AsObject,
  }
}

export class MoveRobotForwardRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveRobotForwardRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MoveRobotForwardRequest): MoveRobotForwardRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveRobotForwardRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveRobotForwardRequest;
  static deserializeBinaryFromReader(message: MoveRobotForwardRequest, reader: jspb.BinaryReader): MoveRobotForwardRequest;
}

export namespace MoveRobotForwardRequest {
  export type AsObject = {
  }
}

export class MoveRobotForwardResponse extends jspb.Message {
  hasNewPosition(): boolean;
  clearNewPosition(): void;
  getNewPosition(): BasePosition | undefined;
  setNewPosition(value?: BasePosition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveRobotForwardResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MoveRobotForwardResponse): MoveRobotForwardResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveRobotForwardResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveRobotForwardResponse;
  static deserializeBinaryFromReader(message: MoveRobotForwardResponse, reader: jspb.BinaryReader): MoveRobotForwardResponse;
}

export namespace MoveRobotForwardResponse {
  export type AsObject = {
    newPosition?: BasePosition.AsObject,
  }
}

export class MoveRobotBackwardRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveRobotBackwardRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MoveRobotBackwardRequest): MoveRobotBackwardRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveRobotBackwardRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveRobotBackwardRequest;
  static deserializeBinaryFromReader(message: MoveRobotBackwardRequest, reader: jspb.BinaryReader): MoveRobotBackwardRequest;
}

export namespace MoveRobotBackwardRequest {
  export type AsObject = {
  }
}

export class MoveRobotBackwardResponse extends jspb.Message {
  hasNewPosition(): boolean;
  clearNewPosition(): void;
  getNewPosition(): BasePosition | undefined;
  setNewPosition(value?: BasePosition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveRobotBackwardResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MoveRobotBackwardResponse): MoveRobotBackwardResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveRobotBackwardResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveRobotBackwardResponse;
  static deserializeBinaryFromReader(message: MoveRobotBackwardResponse, reader: jspb.BinaryReader): MoveRobotBackwardResponse;
}

export namespace MoveRobotBackwardResponse {
  export type AsObject = {
    newPosition?: BasePosition.AsObject,
  }
}

export class TurnRobotToRequest extends jspb.Message {
  getDirection(): DirectionMap[keyof DirectionMap];
  setDirection(value: DirectionMap[keyof DirectionMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TurnRobotToRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TurnRobotToRequest): TurnRobotToRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TurnRobotToRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TurnRobotToRequest;
  static deserializeBinaryFromReader(message: TurnRobotToRequest, reader: jspb.BinaryReader): TurnRobotToRequest;
}

export namespace TurnRobotToRequest {
  export type AsObject = {
    direction: DirectionMap[keyof DirectionMap],
  }
}

export class TurnRobotToResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TurnRobotToResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TurnRobotToResponse): TurnRobotToResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TurnRobotToResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TurnRobotToResponse;
  static deserializeBinaryFromReader(message: TurnRobotToResponse, reader: jspb.BinaryReader): TurnRobotToResponse;
}

export namespace TurnRobotToResponse {
  export type AsObject = {
  }
}

export class UpdateRobotDeviceOffsetRequest extends jspb.Message {
  getOffsetIndex(): number;
  setOffsetIndex(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): DeviceOffset | undefined;
  setOffset(value?: DeviceOffset): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRobotDeviceOffsetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRobotDeviceOffsetRequest): UpdateRobotDeviceOffsetRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRobotDeviceOffsetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRobotDeviceOffsetRequest;
  static deserializeBinaryFromReader(message: UpdateRobotDeviceOffsetRequest, reader: jspb.BinaryReader): UpdateRobotDeviceOffsetRequest;
}

export namespace UpdateRobotDeviceOffsetRequest {
  export type AsObject = {
    offsetIndex: number,
    offset?: DeviceOffset.AsObject,
  }
}

export class DeviceOffset extends jspb.Message {
  getAngle(): number;
  setAngle(value: number): void;

  getDistanceX(): number;
  setDistanceX(value: number): void;

  getDistanceY(): number;
  setDistanceY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeviceOffset.AsObject;
  static toObject(includeInstance: boolean, msg: DeviceOffset): DeviceOffset.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeviceOffset, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeviceOffset;
  static deserializeBinaryFromReader(message: DeviceOffset, reader: jspb.BinaryReader): DeviceOffset;
}

export namespace DeviceOffset {
  export type AsObject = {
    angle: number,
    distanceX: number,
    distanceY: number,
  }
}

export class UpdateRobotDeviceOffsetResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRobotDeviceOffsetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRobotDeviceOffsetResponse): UpdateRobotDeviceOffsetResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRobotDeviceOffsetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRobotDeviceOffsetResponse;
  static deserializeBinaryFromReader(message: UpdateRobotDeviceOffsetResponse, reader: jspb.BinaryReader): UpdateRobotDeviceOffsetResponse;
}

export namespace UpdateRobotDeviceOffsetResponse {
  export type AsObject = {
  }
}

export class StartLidarRequest extends jspb.Message {
  getDeviceNumber(): number;
  setDeviceNumber(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartLidarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StartLidarRequest): StartLidarRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartLidarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartLidarRequest;
  static deserializeBinaryFromReader(message: StartLidarRequest, reader: jspb.BinaryReader): StartLidarRequest;
}

export namespace StartLidarRequest {
  export type AsObject = {
    deviceNumber: number,
  }
}

export class StartLidarResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartLidarResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StartLidarResponse): StartLidarResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartLidarResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartLidarResponse;
  static deserializeBinaryFromReader(message: StartLidarResponse, reader: jspb.BinaryReader): StartLidarResponse;
}

export namespace StartLidarResponse {
  export type AsObject = {
  }
}

export class StopLidarRequest extends jspb.Message {
  getDeviceNumber(): number;
  setDeviceNumber(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StopLidarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StopLidarRequest): StopLidarRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StopLidarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StopLidarRequest;
  static deserializeBinaryFromReader(message: StopLidarRequest, reader: jspb.BinaryReader): StopLidarRequest;
}

export namespace StopLidarRequest {
  export type AsObject = {
    deviceNumber: number,
  }
}

export class StopLidarResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StopLidarResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StopLidarResponse): StopLidarResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StopLidarResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StopLidarResponse;
  static deserializeBinaryFromReader(message: StopLidarResponse, reader: jspb.BinaryReader): StopLidarResponse;
}

export namespace StopLidarResponse {
  export type AsObject = {
  }
}

export class GetLidarSeedRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLidarSeedRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetLidarSeedRequest): GetLidarSeedRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetLidarSeedRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLidarSeedRequest;
  static deserializeBinaryFromReader(message: GetLidarSeedRequest, reader: jspb.BinaryReader): GetLidarSeedRequest;
}

export namespace GetLidarSeedRequest {
  export type AsObject = {
  }
}

export class GetLidarSeedResponse extends jspb.Message {
  clearSeedsList(): void;
  getSeedsList(): Array<string>;
  setSeedsList(value: Array<string>): void;
  addSeeds(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLidarSeedResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetLidarSeedResponse): GetLidarSeedResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetLidarSeedResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLidarSeedResponse;
  static deserializeBinaryFromReader(message: GetLidarSeedResponse, reader: jspb.BinaryReader): GetLidarSeedResponse;
}

export namespace GetLidarSeedResponse {
  export type AsObject = {
    seedsList: Array<string>,
  }
}

export class SetLidarSeedRequest extends jspb.Message {
  getDeviceNumber(): number;
  setDeviceNumber(value: number): void;

  getSeed(): number;
  setSeed(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetLidarSeedRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetLidarSeedRequest): SetLidarSeedRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetLidarSeedRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetLidarSeedRequest;
  static deserializeBinaryFromReader(message: SetLidarSeedRequest, reader: jspb.BinaryReader): SetLidarSeedRequest;
}

export namespace SetLidarSeedRequest {
  export type AsObject = {
    deviceNumber: number,
    seed: number,
  }
}

export class SetLidarSeedResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetLidarSeedResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SetLidarSeedResponse): SetLidarSeedResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetLidarSeedResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetLidarSeedResponse;
  static deserializeBinaryFromReader(message: SetLidarSeedResponse, reader: jspb.BinaryReader): SetLidarSeedResponse;
}

export namespace SetLidarSeedResponse {
  export type AsObject = {
  }
}

export class SetClientZoomRequest extends jspb.Message {
  getZoom(): number;
  setZoom(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetClientZoomRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetClientZoomRequest): SetClientZoomRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetClientZoomRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetClientZoomRequest;
  static deserializeBinaryFromReader(message: SetClientZoomRequest, reader: jspb.BinaryReader): SetClientZoomRequest;
}

export namespace SetClientZoomRequest {
  export type AsObject = {
    zoom: number,
  }
}

export class SetClientZoomResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetClientZoomResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SetClientZoomResponse): SetClientZoomResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetClientZoomResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetClientZoomResponse;
  static deserializeBinaryFromReader(message: SetClientZoomResponse, reader: jspb.BinaryReader): SetClientZoomResponse;
}

export namespace SetClientZoomResponse {
  export type AsObject = {
  }
}

export class SetClientLidarViewModeRequest extends jspb.Message {
  getMode(): LidarViewModeMap[keyof LidarViewModeMap];
  setMode(value: LidarViewModeMap[keyof LidarViewModeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetClientLidarViewModeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetClientLidarViewModeRequest): SetClientLidarViewModeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetClientLidarViewModeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetClientLidarViewModeRequest;
  static deserializeBinaryFromReader(message: SetClientLidarViewModeRequest, reader: jspb.BinaryReader): SetClientLidarViewModeRequest;
}

export namespace SetClientLidarViewModeRequest {
  export type AsObject = {
    mode: LidarViewModeMap[keyof LidarViewModeMap],
  }
}

export class SetClientLidarViewModeResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetClientLidarViewModeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SetClientLidarViewModeResponse): SetClientLidarViewModeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetClientLidarViewModeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetClientLidarViewModeResponse;
  static deserializeBinaryFromReader(message: SetClientLidarViewModeResponse, reader: jspb.BinaryReader): SetClientLidarViewModeResponse;
}

export namespace SetClientLidarViewModeResponse {
  export type AsObject = {
  }
}

export class SetClientClickModeRequest extends jspb.Message {
  getMode(): ClickModeMap[keyof ClickModeMap];
  setMode(value: ClickModeMap[keyof ClickModeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetClientClickModeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetClientClickModeRequest): SetClientClickModeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetClientClickModeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetClientClickModeRequest;
  static deserializeBinaryFromReader(message: SetClientClickModeRequest, reader: jspb.BinaryReader): SetClientClickModeRequest;
}

export namespace SetClientClickModeRequest {
  export type AsObject = {
    mode: ClickModeMap[keyof ClickModeMap],
  }
}

export class SetClientClickModeResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetClientClickModeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SetClientClickModeResponse): SetClientClickModeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetClientClickModeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetClientClickModeResponse;
  static deserializeBinaryFromReader(message: SetClientClickModeResponse, reader: jspb.BinaryReader): SetClientClickModeResponse;
}

export namespace SetClientClickModeResponse {
  export type AsObject = {
  }
}

export interface DirectionMap {
  DIRECTION_UNSPECIFIED: 0;
  DIRECTION_UP: 1;
  DIRECTION_RIGHT: 2;
  DIRECTION_DOWN: 3;
  DIRECTION_LEFT: 4;
}

export const Direction: DirectionMap;

export interface LidarViewModeMap {
  LIDAR_VIEW_MODE_UNSPECIFIED: 0;
  LIDAR_VIEW_MODE_STORED: 1;
  LIDAR_VIEW_MODE_LIVE: 2;
}

export const LidarViewMode: LidarViewModeMap;

export interface ClickModeMap {
  CLICK_MODE_UNSPECIFIED: 0;
  CLICK_MODE_MOVE: 1;
  CLICK_MODE_INFO: 2;
}

export const ClickMode: ClickModeMap;

