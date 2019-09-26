// package: tetrate.api.liaison.v1alpha1
// file: virtual_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as networking_v1alpha3_virtual_service_pb from "./networking/v1alpha3/virtual_service_pb";
import * as validate_validate_pb from "./validate/validate_pb";

export class VirtualService extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): networking_v1alpha3_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: networking_v1alpha3_virtual_service_pb.VirtualService): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VirtualService.AsObject;
  static toObject(includeInstance: boolean, msg: VirtualService): VirtualService.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VirtualService, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VirtualService;
  static deserializeBinaryFromReader(message: VirtualService, reader: jspb.BinaryReader): VirtualService;
}

export namespace VirtualService {
  export type AsObject = {
    name: string,
    displayName: string,
    virtualService?: networking_v1alpha3_virtual_service_pb.VirtualService.AsObject,
  }
}

export class GetVirtualServiceRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVirtualServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetVirtualServiceRequest): GetVirtualServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetVirtualServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVirtualServiceRequest;
  static deserializeBinaryFromReader(message: GetVirtualServiceRequest, reader: jspb.BinaryReader): GetVirtualServiceRequest;
}

export namespace GetVirtualServiceRequest {
  export type AsObject = {
    name: string,
  }
}

export class UpdateVirtualServiceRequest extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): VirtualService | undefined;
  setVirtualService(value?: VirtualService): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateVirtualServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateVirtualServiceRequest): UpdateVirtualServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateVirtualServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateVirtualServiceRequest;
  static deserializeBinaryFromReader(message: UpdateVirtualServiceRequest, reader: jspb.BinaryReader): UpdateVirtualServiceRequest;
}

export namespace UpdateVirtualServiceRequest {
  export type AsObject = {
    virtualService?: VirtualService.AsObject,
  }
}

