// package: tetrate.api.liaison.v1alpha1
// file: ingress_gateway.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as networking_v1alpha3_gateway_pb from "./networking/v1alpha3/gateway_pb";
import * as validate_validate_pb from "./validate/validate_pb";

export class IngressGateway extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  getZone(): string;
  setZone(value: string): void;

  hasGateway(): boolean;
  clearGateway(): void;
  getGateway(): networking_v1alpha3_gateway_pb.Gateway | undefined;
  setGateway(value?: networking_v1alpha3_gateway_pb.Gateway): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IngressGateway.AsObject;
  static toObject(includeInstance: boolean, msg: IngressGateway): IngressGateway.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IngressGateway, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IngressGateway;
  static deserializeBinaryFromReader(message: IngressGateway, reader: jspb.BinaryReader): IngressGateway;
}

export namespace IngressGateway {
  export type AsObject = {
    name: string,
    displayName: string,
    zone: string,
    gateway?: networking_v1alpha3_gateway_pb.Gateway.AsObject,
  }
}

export class ListIngressGatewaysRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListIngressGatewaysRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListIngressGatewaysRequest): ListIngressGatewaysRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListIngressGatewaysRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListIngressGatewaysRequest;
  static deserializeBinaryFromReader(message: ListIngressGatewaysRequest, reader: jspb.BinaryReader): ListIngressGatewaysRequest;
}

export namespace ListIngressGatewaysRequest {
  export type AsObject = {
  }
}

export class ListIngressGatewaysResponse extends jspb.Message {
  clearIngressGatewaysList(): void;
  getIngressGatewaysList(): Array<IngressGateway>;
  setIngressGatewaysList(value: Array<IngressGateway>): void;
  addIngressGateways(value?: IngressGateway, index?: number): IngressGateway;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListIngressGatewaysResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListIngressGatewaysResponse): ListIngressGatewaysResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListIngressGatewaysResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListIngressGatewaysResponse;
  static deserializeBinaryFromReader(message: ListIngressGatewaysResponse, reader: jspb.BinaryReader): ListIngressGatewaysResponse;
}

export namespace ListIngressGatewaysResponse {
  export type AsObject = {
    ingressGatewaysList: Array<IngressGateway.AsObject>,
  }
}

export class GetIngressGatewayRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetIngressGatewayRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetIngressGatewayRequest): GetIngressGatewayRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetIngressGatewayRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetIngressGatewayRequest;
  static deserializeBinaryFromReader(message: GetIngressGatewayRequest, reader: jspb.BinaryReader): GetIngressGatewayRequest;
}

export namespace GetIngressGatewayRequest {
  export type AsObject = {
    name: string,
  }
}

export class CreateIngressGatewayRequest extends jspb.Message {
  hasIngressGateway(): boolean;
  clearIngressGateway(): void;
  getIngressGateway(): IngressGateway | undefined;
  setIngressGateway(value?: IngressGateway): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateIngressGatewayRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateIngressGatewayRequest): CreateIngressGatewayRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateIngressGatewayRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateIngressGatewayRequest;
  static deserializeBinaryFromReader(message: CreateIngressGatewayRequest, reader: jspb.BinaryReader): CreateIngressGatewayRequest;
}

export namespace CreateIngressGatewayRequest {
  export type AsObject = {
    ingressGateway?: IngressGateway.AsObject,
  }
}

export class UpdateIngressGatewayRequest extends jspb.Message {
  hasIngressGateway(): boolean;
  clearIngressGateway(): void;
  getIngressGateway(): IngressGateway | undefined;
  setIngressGateway(value?: IngressGateway): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateIngressGatewayRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateIngressGatewayRequest): UpdateIngressGatewayRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateIngressGatewayRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateIngressGatewayRequest;
  static deserializeBinaryFromReader(message: UpdateIngressGatewayRequest, reader: jspb.BinaryReader): UpdateIngressGatewayRequest;
}

export namespace UpdateIngressGatewayRequest {
  export type AsObject = {
    ingressGateway?: IngressGateway.AsObject,
  }
}

export class DeleteIngressGatewayRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteIngressGatewayRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteIngressGatewayRequest): DeleteIngressGatewayRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteIngressGatewayRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteIngressGatewayRequest;
  static deserializeBinaryFromReader(message: DeleteIngressGatewayRequest, reader: jspb.BinaryReader): DeleteIngressGatewayRequest;
}

export namespace DeleteIngressGatewayRequest {
  export type AsObject = {
    name: string,
  }
}

