// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doc.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("doc.proto", fileDescriptor_37cb16cf10c66117) }

var fileDescriptor_37cb16cf10c66117 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0x02, 0x41,
	0x14, 0x45, 0xb3, 0xab, 0x81, 0x38, 0x8d, 0x64, 0x13, 0x35, 0xa1, 0x20, 0x13, 0x4b, 0x02, 0x33,
	0x82, 0x54, 0xd0, 0x88, 0xdb, 0x48, 0xa5, 0xc1, 0x2d, 0x8c, 0xdd, 0x30, 0x3c, 0x76, 0xc7, 0xe8,
	0xbc, 0xc9, 0xcc, 0x5b, 0xd6, 0x2f, 0xf0, 0x03, 0x2c, 0xfd, 0x22, 0x4a, 0x3f, 0xc1, 0xf0, 0x25,
	0xc6, 0x15, 0xb5, 0xb0, 0xbb, 0xb9, 0xf7, 0xe5, 0xbe, 0x9b, 0xc3, 0x0e, 0x96, 0xa8, 0x85, 0xf3,
	0x48, 0x98, 0x9c, 0x10, 0x90, 0x57, 0x04, 0x42, 0x39, 0x23, 0x48, 0x6b, 0xa1, 0xd1, 0x83, 0x58,
	0x0f, 0xda, 0xbd, 0x3a, 0xd7, 0xfd, 0x1c, 0x6c, 0x3f, 0x54, 0x2a, 0xcf, 0xc1, 0x4b, 0x74, 0x64,
	0xd0, 0x06, 0xa9, 0xac, 0x45, 0x52, 0xb5, 0xfe, 0xae, 0xb9, 0x7c, 0x89, 0x5f, 0xa7, 0x9b, 0x28,
	0xb9, 0x63, 0x47, 0x59, 0x9a, 0xf2, 0x14, 0xed, 0xca, 0xe4, 0xa5, 0xaf, 0x4f, 0xf8, 0xf4, 0x66,
	0x76, 0x3a, 0x61, 0xad, 0x6c, 0xf7, 0xc8, 0x60, 0x8f, 0xcf, 0xac, 0x16, 0xc9, 0x71, 0x41, 0xe4,
	0xc2, 0x58, 0xca, 0xaa, 0xaa, 0x04, 0xfd, 0xa6, 0xed, 0x43, 0x63, 0x57, 0x78, 0xf1, 0x67, 0x0c,
	0xe3, 0xf5, 0xa0, 0xdd, 0xcc, 0xd2, 0xf4, 0xea, 0xfa, 0x36, 0xeb, 0xc6, 0x51, 0x3c, 0x6c, 0x29,
	0xe7, 0x1e, 0x8d, 0xae, 0xeb, 0xe5, 0x43, 0x40, 0x3b, 0xfe, 0xe7, 0xcc, 0x27, 0x6c, 0x6f, 0x74,
	0x36, 0x4a, 0x46, 0xac, 0x3b, 0x07, 0x2a, 0xbd, 0x85, 0x25, 0xaf, 0x0a, 0xb0, 0x9c, 0x0a, 0xe0,
	0x1e, 0x02, 0x96, 0x5e, 0x03, 0x5f, 0x22, 0x04, 0x6e, 0x91, 0x38, 0x3c, 0x9b, 0x40, 0x22, 0x69,
	0xb0, 0xfd, 0xb7, 0x38, 0x6a, 0x6e, 0xb6, 0x9d, 0xe8, 0x7d, 0xdb, 0x89, 0x3e, 0xb6, 0x9d, 0xe8,
	0xbe, 0x97, 0x1b, 0x2a, 0xca, 0x85, 0xd0, 0xf8, 0x24, 0x77, 0xa3, 0x0c, 0xfe, 0x28, 0xa9, 0x9c,
	0x91, 0xa4, 0xb5, 0xfc, 0xc2, 0x26, 0xd7, 0x83, 0x45, 0xa3, 0xe6, 0x71, 0xfe, 0x19, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0x45, 0x00, 0x6d, 0x63, 0x01, 0x00, 0x00,
}