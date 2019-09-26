// package: tetrate.api.liaison.v1alpha1
// file: virtual_service.proto

import * as virtual_service_pb from "./virtual_service_pb";
import {grpc} from "@improbable-eng/grpc-web";

type VirtualServicesGetVirtualService = {
  readonly methodName: string;
  readonly service: typeof VirtualServices;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof virtual_service_pb.GetVirtualServiceRequest;
  readonly responseType: typeof virtual_service_pb.VirtualService;
};

type VirtualServicesUpdateVirtualService = {
  readonly methodName: string;
  readonly service: typeof VirtualServices;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof virtual_service_pb.UpdateVirtualServiceRequest;
  readonly responseType: typeof virtual_service_pb.VirtualService;
};

export class VirtualServices {
  static readonly serviceName: string;
  static readonly GetVirtualService: VirtualServicesGetVirtualService;
  static readonly UpdateVirtualService: VirtualServicesUpdateVirtualService;
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

export class VirtualServicesClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getVirtualService(
    requestMessage: virtual_service_pb.GetVirtualServiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: virtual_service_pb.VirtualService|null) => void
  ): UnaryResponse;
  getVirtualService(
    requestMessage: virtual_service_pb.GetVirtualServiceRequest,
    callback: (error: ServiceError|null, responseMessage: virtual_service_pb.VirtualService|null) => void
  ): UnaryResponse;
  updateVirtualService(
    requestMessage: virtual_service_pb.UpdateVirtualServiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: virtual_service_pb.VirtualService|null) => void
  ): UnaryResponse;
  updateVirtualService(
    requestMessage: virtual_service_pb.UpdateVirtualServiceRequest,
    callback: (error: ServiceError|null, responseMessage: virtual_service_pb.VirtualService|null) => void
  ): UnaryResponse;
}

