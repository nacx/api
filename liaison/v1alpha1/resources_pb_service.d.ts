// package: tetrate.api.liaison.v1alpha1
// file: resources.proto

import * as resources_pb from "./resources_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ResourcesServiceSetResource = {
  readonly methodName: string;
  readonly service: typeof ResourcesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof resources_pb.SetResourceRequest;
  readonly responseType: typeof resources_pb.Resource;
};

type ResourcesServiceGetResource = {
  readonly methodName: string;
  readonly service: typeof ResourcesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof resources_pb.GetResourceRequest;
  readonly responseType: typeof resources_pb.Resource;
};

export class ResourcesService {
  static readonly serviceName: string;
  static readonly SetResource: ResourcesServiceSetResource;
  static readonly GetResource: ResourcesServiceGetResource;
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

export class ResourcesServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  setResource(
    requestMessage: resources_pb.SetResourceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: resources_pb.Resource|null) => void
  ): UnaryResponse;
  setResource(
    requestMessage: resources_pb.SetResourceRequest,
    callback: (error: ServiceError|null, responseMessage: resources_pb.Resource|null) => void
  ): UnaryResponse;
  getResource(
    requestMessage: resources_pb.GetResourceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: resources_pb.Resource|null) => void
  ): UnaryResponse;
  getResource(
    requestMessage: resources_pb.GetResourceRequest,
    callback: (error: ServiceError|null, responseMessage: resources_pb.Resource|null) => void
  ): UnaryResponse;
}

