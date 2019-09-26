// package: tetrate.api.liaison.v1alpha1
// file: service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as validate_validate_pb from "./validate/validate_pb";

export class Service extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Service.AsObject;
  static toObject(includeInstance: boolean, msg: Service): Service.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Service, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Service;
  static deserializeBinaryFromReader(message: Service, reader: jspb.BinaryReader): Service;
}

export namespace Service {
  export type AsObject = {
    name: string,
    displayName: string,
  }
}

export class ListServicesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListServicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListServicesRequest): ListServicesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListServicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListServicesRequest;
  static deserializeBinaryFromReader(message: ListServicesRequest, reader: jspb.BinaryReader): ListServicesRequest;
}

export namespace ListServicesRequest {
  export type AsObject = {
  }
}

export class ListServicesResponse extends jspb.Message {
  clearServicesList(): void;
  getServicesList(): Array<Service>;
  setServicesList(value: Array<Service>): void;
  addServices(value?: Service, index?: number): Service;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListServicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListServicesResponse): ListServicesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListServicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListServicesResponse;
  static deserializeBinaryFromReader(message: ListServicesResponse, reader: jspb.BinaryReader): ListServicesResponse;
}

export namespace ListServicesResponse {
  export type AsObject = {
    servicesList: Array<Service.AsObject>,
  }
}

export class GetServiceRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetServiceRequest): GetServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetServiceRequest;
  static deserializeBinaryFromReader(message: GetServiceRequest, reader: jspb.BinaryReader): GetServiceRequest;
}

export namespace GetServiceRequest {
  export type AsObject = {
    name: string,
  }
}

export class CreateServiceRequest extends jspb.Message {
  hasService(): boolean;
  clearService(): void;
  getService(): Service | undefined;
  setService(value?: Service): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateServiceRequest): CreateServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateServiceRequest;
  static deserializeBinaryFromReader(message: CreateServiceRequest, reader: jspb.BinaryReader): CreateServiceRequest;
}

export namespace CreateServiceRequest {
  export type AsObject = {
    service?: Service.AsObject,
  }
}

export class UpdateServiceRequest extends jspb.Message {
  hasService(): boolean;
  clearService(): void;
  getService(): Service | undefined;
  setService(value?: Service): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateServiceRequest): UpdateServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateServiceRequest;
  static deserializeBinaryFromReader(message: UpdateServiceRequest, reader: jspb.BinaryReader): UpdateServiceRequest;
}

export namespace UpdateServiceRequest {
  export type AsObject = {
    service?: Service.AsObject,
  }
}

export class DeleteServiceRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteServiceRequest): DeleteServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteServiceRequest;
  static deserializeBinaryFromReader(message: DeleteServiceRequest, reader: jspb.BinaryReader): DeleteServiceRequest;
}

export namespace DeleteServiceRequest {
  export type AsObject = {
    name: string,
  }
}

