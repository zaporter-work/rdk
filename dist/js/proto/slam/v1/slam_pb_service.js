// package: proto.slam.v1
// file: proto/slam/v1/slam.proto

var proto_slam_v1_slam_pb = require("../../../proto/slam/v1/slam_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var SlamService = (function () {
  function SlamService() {}
  SlamService.serviceName = "proto.slam.v1.SlamService";
  return SlamService;
}());

SlamService.Save = {
  methodName: "Save",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.SaveRequest,
  responseType: proto_slam_v1_slam_pb.SaveResponse
};

SlamService.Stats = {
  methodName: "Stats",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.StatsRequest,
  responseType: proto_slam_v1_slam_pb.StatsResponse
};

SlamService.Calibrate = {
  methodName: "Calibrate",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.CalibrateRequest,
  responseType: proto_slam_v1_slam_pb.CalibrateResponse
};

SlamService.MoveRobot = {
  methodName: "MoveRobot",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.MoveRobotRequest,
  responseType: proto_slam_v1_slam_pb.MoveRobotResponse
};

SlamService.MoveRobotForward = {
  methodName: "MoveRobotForward",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.MoveRobotForwardRequest,
  responseType: proto_slam_v1_slam_pb.MoveRobotForwardResponse
};

SlamService.MoveRobotBackward = {
  methodName: "MoveRobotBackward",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.MoveRobotBackwardRequest,
  responseType: proto_slam_v1_slam_pb.MoveRobotBackwardResponse
};

SlamService.TurnRobotTo = {
  methodName: "TurnRobotTo",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.TurnRobotToRequest,
  responseType: proto_slam_v1_slam_pb.TurnRobotToResponse
};

SlamService.UpdateRobotDeviceOffset = {
  methodName: "UpdateRobotDeviceOffset",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetRequest,
  responseType: proto_slam_v1_slam_pb.UpdateRobotDeviceOffsetResponse
};

SlamService.StartLidar = {
  methodName: "StartLidar",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.StartLidarRequest,
  responseType: proto_slam_v1_slam_pb.StartLidarResponse
};

SlamService.StopLidar = {
  methodName: "StopLidar",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.StopLidarRequest,
  responseType: proto_slam_v1_slam_pb.StopLidarResponse
};

SlamService.GetLidarSeed = {
  methodName: "GetLidarSeed",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.GetLidarSeedRequest,
  responseType: proto_slam_v1_slam_pb.GetLidarSeedResponse
};

SlamService.SetLidarSeed = {
  methodName: "SetLidarSeed",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.SetLidarSeedRequest,
  responseType: proto_slam_v1_slam_pb.SetLidarSeedResponse
};

SlamService.SetClientZoom = {
  methodName: "SetClientZoom",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.SetClientZoomRequest,
  responseType: proto_slam_v1_slam_pb.SetClientZoomResponse
};

SlamService.SetClientLidarViewMode = {
  methodName: "SetClientLidarViewMode",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.SetClientLidarViewModeRequest,
  responseType: proto_slam_v1_slam_pb.SetClientLidarViewModeResponse
};

SlamService.SetClientClickMode = {
  methodName: "SetClientClickMode",
  service: SlamService,
  requestStream: false,
  responseStream: false,
  requestType: proto_slam_v1_slam_pb.SetClientClickModeRequest,
  responseType: proto_slam_v1_slam_pb.SetClientClickModeResponse
};

exports.SlamService = SlamService;

function SlamServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

SlamServiceClient.prototype.save = function save(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.Save, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.stats = function stats(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.Stats, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.calibrate = function calibrate(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.Calibrate, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.moveRobot = function moveRobot(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.MoveRobot, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.moveRobotForward = function moveRobotForward(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.MoveRobotForward, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.moveRobotBackward = function moveRobotBackward(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.MoveRobotBackward, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.turnRobotTo = function turnRobotTo(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.TurnRobotTo, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.updateRobotDeviceOffset = function updateRobotDeviceOffset(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.UpdateRobotDeviceOffset, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.startLidar = function startLidar(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.StartLidar, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.stopLidar = function stopLidar(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.StopLidar, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.getLidarSeed = function getLidarSeed(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.GetLidarSeed, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.setLidarSeed = function setLidarSeed(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.SetLidarSeed, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.setClientZoom = function setClientZoom(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.SetClientZoom, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.setClientLidarViewMode = function setClientLidarViewMode(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.SetClientLidarViewMode, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SlamServiceClient.prototype.setClientClickMode = function setClientClickMode(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SlamService.SetClientClickMode, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.SlamServiceClient = SlamServiceClient;

