// package: tetrate.api.liaison.v1alpha1
// file: service.proto

import * as service_pb from "./service_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ServicesListServices = {
  readonly methodName: string;
  readonly service: typeof Services;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.ListServicesRequest;
  readonly responseType: typeof service_pb.ListServicesResponse;
};

type ServicesGetService = {
  readonly methodName: string;
  readonly service: typeof Services;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.GetServiceRequest;
  readonly responseType: typeof service_pb.Service;
};

type ServicesCreateService = {
  readonly methodName: string;
  readonly service: typeof Services;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.CreateServiceRequest;
  readonly responseType: typeof service_pb.Service;
};

type ServicesUpdateService = {
  readonly methodName: string;
  readonly service: typeof Services;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.UpdateServiceRequest;
  readonly responseType: typeof service_pb.Service;
};

type ServicesDeleteService = {
  readonly methodName: string;
  readonly service: typeof Services;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.DeleteServiceRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class Services {
  static readonly serviceName: string;
  static readonly ListServices: ServicesListServices;
  static readonly GetService: ServicesGetService;
  static readonly CreateService: ServicesCreateService;
  static readonly UpdateService: ServicesUpdateService;
  static readonly DeleteService: ServicesDeleteService;
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

export class ServicesClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listServices(
    requestMessage: service_pb.ListServicesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_pb.ListServicesResponse|null) => void
  ): UnaryResponse;
  listServices(
    requestMessage: service_pb.ListServicesRequest,
    callback: (error: ServiceError|null, responseMessage: service_pb.ListServicesResponse|null) => void
  ): UnaryResponse;
  getService(
    requestMessage: service_pb.GetServiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_pb.Service|null) => void
  ): UnaryResponse;
  getService(
    requestMessage: service_pb.GetServiceRequest,
    callback: (error: ServiceError|null, responseMessage: service_pb.Service|null) => void
  ): UnaryResponse;
  createService(
    requestMessage: service_pb.CreateServiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_pb.Service|null) => void
  ): UnaryResponse;
  createService(
    requestMessage: service_pb.CreateServiceRequest,
    callback: (error: ServiceError|null, responseMessage: service_pb.Service|null) => void
  ): UnaryResponse;
  updateService(
    requestMessage: service_pb.UpdateServiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_pb.Service|null) => void
  ): UnaryResponse;
  updateService(
    requestMessage: service_pb.UpdateServiceRequest,
    callback: (error: ServiceError|null, responseMessage: service_pb.Service|null) => void
  ): UnaryResponse;
  deleteService(
    requestMessage: service_pb.DeleteServiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  deleteService(
    requestMessage: service_pb.DeleteServiceRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
}

