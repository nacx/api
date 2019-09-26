// package: tetrate.api.liaison.v1alpha1
// file: service.proto

var service_pb = require("./service_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Services = (function () {
  function Services() {}
  Services.serviceName = "tetrate.api.liaison.v1alpha1.Services";
  return Services;
}());

Services.ListServices = {
  methodName: "ListServices",
  service: Services,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.ListServicesRequest,
  responseType: service_pb.ListServicesResponse
};

Services.GetService = {
  methodName: "GetService",
  service: Services,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.GetServiceRequest,
  responseType: service_pb.Service
};

Services.CreateService = {
  methodName: "CreateService",
  service: Services,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.CreateServiceRequest,
  responseType: service_pb.Service
};

Services.UpdateService = {
  methodName: "UpdateService",
  service: Services,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.UpdateServiceRequest,
  responseType: service_pb.Service
};

Services.DeleteService = {
  methodName: "DeleteService",
  service: Services,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.DeleteServiceRequest,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Services = Services;

function ServicesClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ServicesClient.prototype.listServices = function listServices(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Services.ListServices, {
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

ServicesClient.prototype.getService = function getService(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Services.GetService, {
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

ServicesClient.prototype.createService = function createService(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Services.CreateService, {
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

ServicesClient.prototype.updateService = function updateService(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Services.UpdateService, {
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

ServicesClient.prototype.deleteService = function deleteService(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Services.DeleteService, {
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

exports.ServicesClient = ServicesClient;

