// package: tetrate.api.liaison.v1alpha1
// file: virtual_service.proto

var virtual_service_pb = require("./virtual_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var VirtualServices = (function () {
  function VirtualServices() {}
  VirtualServices.serviceName = "tetrate.api.liaison.v1alpha1.VirtualServices";
  return VirtualServices;
}());

VirtualServices.GetVirtualService = {
  methodName: "GetVirtualService",
  service: VirtualServices,
  requestStream: false,
  responseStream: false,
  requestType: virtual_service_pb.GetVirtualServiceRequest,
  responseType: virtual_service_pb.VirtualService
};

VirtualServices.UpdateVirtualService = {
  methodName: "UpdateVirtualService",
  service: VirtualServices,
  requestStream: false,
  responseStream: false,
  requestType: virtual_service_pb.UpdateVirtualServiceRequest,
  responseType: virtual_service_pb.VirtualService
};

exports.VirtualServices = VirtualServices;

function VirtualServicesClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

VirtualServicesClient.prototype.getVirtualService = function getVirtualService(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(VirtualServices.GetVirtualService, {
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

VirtualServicesClient.prototype.updateVirtualService = function updateVirtualService(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(VirtualServices.UpdateVirtualService, {
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

exports.VirtualServicesClient = VirtualServicesClient;

