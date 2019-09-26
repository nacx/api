// package: tetrate.api.liaison.v1alpha1
// file: resources.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as validate_validate_pb from "./validate/validate_pb";

export class Resource extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Resource.AsObject;
  static toObject(includeInstance: boolean, msg: Resource): Resource.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Resource, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Resource;
  static deserializeBinaryFromReader(message: Resource, reader: jspb.BinaryReader): Resource;
}

export namespace Resource {
  export type AsObject = {
    key: string,
  }
}

export class SetResourceRequest extends jspb.Message {
  getYaml(): Uint8Array | string;
  getYaml_asU8(): Uint8Array;
  getYaml_asB64(): string;
  setYaml(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetResourceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetResourceRequest): SetResourceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetResourceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetResourceRequest;
  static deserializeBinaryFromReader(message: SetResourceRequest, reader: jspb.BinaryReader): SetResourceRequest;
}

export namespace SetResourceRequest {
  export type AsObject = {
    yaml: Uint8Array | string,
  }
}

export class GetResourceRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetResourceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetResourceRequest): GetResourceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetResourceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetResourceRequest;
  static deserializeBinaryFromReader(message: GetResourceRequest, reader: jspb.BinaryReader): GetResourceRequest;
}

export namespace GetResourceRequest {
  export type AsObject = {
    key: string,
  }
}

