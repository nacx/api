// package: tetrate.api.liaison.v1alpha1
// file: resources.proto

var resources_pb = require("./resources_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ResourcesService = (function () {
  function ResourcesService() {}
  ResourcesService.serviceName = "tetrate.api.liaison.v1alpha1.ResourcesService";
  return ResourcesService;
}());

ResourcesService.SetResource = {
  methodName: "SetResource",
  service: ResourcesService,
  requestStream: false,
  responseStream: false,
  requestType: resources_pb.SetResourceRequest,
  responseType: resources_pb.Resource
};

ResourcesService.GetResource = {
  methodName: "GetResource",
  service: ResourcesService,
  requestStream: false,
  responseStream: false,
  requestType: resources_pb.GetResourceRequest,
  responseType: resources_pb.Resource
};

exports.ResourcesService = ResourcesService;

function ResourcesServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ResourcesServiceClient.prototype.setResource = function setResource(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ResourcesService.SetResource, {
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

ResourcesServiceClient.prototype.getResource = function getResource(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ResourcesService.GetResource, {
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

exports.ResourcesServiceClient = ResourcesServiceClient;

