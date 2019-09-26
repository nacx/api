// package: tetrate.api.liaison.v1alpha1
// file: ingress_gateway.proto

var ingress_gateway_pb = require("./ingress_gateway_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var IngressGateways = (function () {
  function IngressGateways() {}
  IngressGateways.serviceName = "tetrate.api.liaison.v1alpha1.IngressGateways";
  return IngressGateways;
}());

IngressGateways.ListIngressGateways = {
  methodName: "ListIngressGateways",
  service: IngressGateways,
  requestStream: false,
  responseStream: false,
  requestType: ingress_gateway_pb.ListIngressGatewaysRequest,
  responseType: ingress_gateway_pb.ListIngressGatewaysResponse
};

IngressGateways.GetIngressGateway = {
  methodName: "GetIngressGateway",
  service: IngressGateways,
  requestStream: false,
  responseStream: false,
  requestType: ingress_gateway_pb.GetIngressGatewayRequest,
  responseType: ingress_gateway_pb.IngressGateway
};

IngressGateways.CreateIngressGateway = {
  methodName: "CreateIngressGateway",
  service: IngressGateways,
  requestStream: false,
  responseStream: false,
  requestType: ingress_gateway_pb.CreateIngressGatewayRequest,
  responseType: ingress_gateway_pb.IngressGateway
};

IngressGateways.UpdateIngressGateway = {
  methodName: "UpdateIngressGateway",
  service: IngressGateways,
  requestStream: false,
  responseStream: false,
  requestType: ingress_gateway_pb.UpdateIngressGatewayRequest,
  responseType: ingress_gateway_pb.IngressGateway
};

IngressGateways.DeleteIngressGateway = {
  methodName: "DeleteIngressGateway",
  service: IngressGateways,
  requestStream: false,
  responseStream: false,
  requestType: ingress_gateway_pb.DeleteIngressGatewayRequest,
  responseType: google_protobuf_empty_pb.Empty
};

exports.IngressGateways = IngressGateways;

function IngressGatewaysClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

IngressGatewaysClient.prototype.listIngressGateways = function listIngressGateways(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IngressGateways.ListIngressGateways, {
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

IngressGatewaysClient.prototype.getIngressGateway = function getIngressGateway(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IngressGateways.GetIngressGateway, {
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

IngressGatewaysClient.prototype.createIngressGateway = function createIngressGateway(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IngressGateways.CreateIngressGateway, {
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

IngressGatewaysClient.prototype.updateIngressGateway = function updateIngressGateway(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IngressGateways.UpdateIngressGateway, {
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

IngressGatewaysClient.prototype.deleteIngressGateway = function deleteIngressGateway(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IngressGateways.DeleteIngressGateway, {
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

exports.IngressGatewaysClient = IngressGatewaysClient;

