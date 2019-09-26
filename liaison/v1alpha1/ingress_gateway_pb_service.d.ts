// package: tetrate.api.liaison.v1alpha1
// file: ingress_gateway.proto

import * as ingress_gateway_pb from "./ingress_gateway_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type IngressGatewaysListIngressGateways = {
  readonly methodName: string;
  readonly service: typeof IngressGateways;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ingress_gateway_pb.ListIngressGatewaysRequest;
  readonly responseType: typeof ingress_gateway_pb.ListIngressGatewaysResponse;
};

type IngressGatewaysGetIngressGateway = {
  readonly methodName: string;
  readonly service: typeof IngressGateways;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ingress_gateway_pb.GetIngressGatewayRequest;
  readonly responseType: typeof ingress_gateway_pb.IngressGateway;
};

type IngressGatewaysCreateIngressGateway = {
  readonly methodName: string;
  readonly service: typeof IngressGateways;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ingress_gateway_pb.CreateIngressGatewayRequest;
  readonly responseType: typeof ingress_gateway_pb.IngressGateway;
};

type IngressGatewaysUpdateIngressGateway = {
  readonly methodName: string;
  readonly service: typeof IngressGateways;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ingress_gateway_pb.UpdateIngressGatewayRequest;
  readonly responseType: typeof ingress_gateway_pb.IngressGateway;
};

type IngressGatewaysDeleteIngressGateway = {
  readonly methodName: string;
  readonly service: typeof IngressGateways;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ingress_gateway_pb.DeleteIngressGatewayRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class IngressGateways {
  static readonly serviceName: string;
  static readonly ListIngressGateways: IngressGatewaysListIngressGateways;
  static readonly GetIngressGateway: IngressGatewaysGetIngressGateway;
  static readonly CreateIngressGateway: IngressGatewaysCreateIngressGateway;
  static readonly UpdateIngressGateway: IngressGatewaysUpdateIngressGateway;
  static readonly DeleteIngressGateway: IngressGatewaysDeleteIngressGateway;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class IngressGatewaysClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listIngressGateways(
    requestMessage: ingress_gateway_pb.ListIngressGatewaysRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.ListIngressGatewaysResponse|null) => void
  ): UnaryResponse;
  listIngressGateways(
    requestMessage: ingress_gateway_pb.ListIngressGatewaysRequest,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.ListIngressGatewaysResponse|null) => void
  ): UnaryResponse;
  getIngressGateway(
    requestMessage: ingress_gateway_pb.GetIngressGatewayRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.IngressGateway|null) => void
  ): UnaryResponse;
  getIngressGateway(
    requestMessage: ingress_gateway_pb.GetIngressGatewayRequest,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.IngressGateway|null) => void
  ): UnaryResponse;
  createIngressGateway(
    requestMessage: ingress_gateway_pb.CreateIngressGatewayRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.IngressGateway|null) => void
  ): UnaryResponse;
  createIngressGateway(
    requestMessage: ingress_gateway_pb.CreateIngressGatewayRequest,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.IngressGateway|null) => void
  ): UnaryResponse;
  updateIngressGateway(
    requestMessage: ingress_gateway_pb.UpdateIngressGatewayRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.IngressGateway|null) => void
  ): UnaryResponse;
  updateIngressGateway(
    requestMessage: ingress_gateway_pb.UpdateIngressGatewayRequest,
    callback: (error: ServiceError|null, responseMessage: ingress_gateway_pb.IngressGateway|null) => void
  ): UnaryResponse;
  deleteIngressGateway(
    requestMessage: ingress_gateway_pb.DeleteIngressGatewayRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  deleteIngressGateway(
    requestMessage: ingress_gateway_pb.DeleteIngressGatewayRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
}

