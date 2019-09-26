// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registryconsumer.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/core/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RegistryUpdateResourcesRequest request object.
type RegistryUpdateResourcesRequest struct {
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Tenant where this service is residing e.g Tetrate.
	Tenant string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// environment are higher-level isolations e.g. prod, staging, development, etc.
	Environment string `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	// Unique cluster acting as a source.
	Cluster string `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// registrytype indicates source of entry. Currently accepted values "kubernetes", "f5".
	Registrytype string `protobuf:"bytes,5,opt,name=registrytype,proto3" json:"registrytype,omitempty"`
	// JSON bulk payload of all physical resources.
	Payload              string   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryUpdateResourcesRequest) Reset()         { *m = RegistryUpdateResourcesRequest{} }
func (m *RegistryUpdateResourcesRequest) String() string { return proto.CompactTextString(m) }
func (*RegistryUpdateResourcesRequest) ProtoMessage()    {}
func (*RegistryUpdateResourcesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registryconsumer_8b98d8d607db6684, []int{0}
}
func (m *RegistryUpdateResourcesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryUpdateResourcesRequest.Unmarshal(m, b)
}
func (m *RegistryUpdateResourcesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryUpdateResourcesRequest.Marshal(b, m, deterministic)
}
func (dst *RegistryUpdateResourcesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryUpdateResourcesRequest.Merge(dst, src)
}
func (m *RegistryUpdateResourcesRequest) XXX_Size() int {
	return xxx_messageInfo_RegistryUpdateResourcesRequest.Size(m)
}
func (m *RegistryUpdateResourcesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryUpdateResourcesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryUpdateResourcesRequest proto.InternalMessageInfo

func (m *RegistryUpdateResourcesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RegistryUpdateResourcesRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *RegistryUpdateResourcesRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *RegistryUpdateResourcesRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RegistryUpdateResourcesRequest) GetRegistrytype() string {
	if m != nil {
		return m.Registrytype
	}
	return ""
}

func (m *RegistryUpdateResourcesRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

// RegistryConsumerResponse response returned by server.
type RegistryConsumerResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryConsumerResponse) Reset()         { *m = RegistryConsumerResponse{} }
func (m *RegistryConsumerResponse) String() string { return proto.CompactTextString(m) }
func (*RegistryConsumerResponse) ProtoMessage()    {}
func (*RegistryConsumerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_registryconsumer_8b98d8d607db6684, []int{1}
}
func (m *RegistryConsumerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryConsumerResponse.Unmarshal(m, b)
}
func (m *RegistryConsumerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryConsumerResponse.Marshal(b, m, deterministic)
}
func (dst *RegistryConsumerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryConsumerResponse.Merge(dst, src)
}
func (m *RegistryConsumerResponse) XXX_Size() int {
	return xxx_messageInfo_RegistryConsumerResponse.Size(m)
}
func (m *RegistryConsumerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryConsumerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryConsumerResponse proto.InternalMessageInfo

func (m *RegistryConsumerResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistryUpdateResourcesRequest)(nil), "tetrate.api.tcc.core.v1.RegistryUpdateResourcesRequest")
	proto.RegisterType((*RegistryConsumerResponse)(nil), "tetrate.api.tcc.core.v1.RegistryConsumerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryConsumerClient is the client API for RegistryConsumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryConsumerClient interface {
	// UpdateResources to be called periodically with a bulk payload. API detects the delta between last call, and
	// determines resources to create, modify, delete
	UpdateResources(ctx context.Context, in *RegistryUpdateResourcesRequest, opts ...grpc.CallOption) (*RegistryConsumerResponse, error)
}

type registryConsumerClient struct {
	cc *grpc.ClientConn
}

func NewRegistryConsumerClient(cc *grpc.ClientConn) RegistryConsumerClient {
	return &registryConsumerClient{cc}
}

func (c *registryConsumerClient) UpdateResources(ctx context.Context, in *RegistryUpdateResourcesRequest, opts ...grpc.CallOption) (*RegistryConsumerResponse, error) {
	out := new(RegistryConsumerResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.core.v1.RegistryConsumer/UpdateResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryConsumerServer is the server API for RegistryConsumer service.
type RegistryConsumerServer interface {
	// UpdateResources to be called periodically with a bulk payload. API detects the delta between last call, and
	// determines resources to create, modify, delete
	UpdateResources(context.Context, *RegistryUpdateResourcesRequest) (*RegistryConsumerResponse, error)
}

func RegisterRegistryConsumerServer(s *grpc.Server, srv RegistryConsumerServer) {
	s.RegisterService(&_RegistryConsumer_serviceDesc, srv)
}

func _RegistryConsumer_UpdateResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryUpdateResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryConsumerServer).UpdateResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.core.v1.RegistryConsumer/UpdateResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryConsumerServer).UpdateResources(ctx, req.(*RegistryUpdateResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryConsumer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.tcc.core.v1.RegistryConsumer",
	HandlerType: (*RegistryConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateResources",
			Handler:    _RegistryConsumer_UpdateResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registryconsumer.proto",
}

func init() {
	proto.RegisterFile("registryconsumer.proto", fileDescriptor_registryconsumer_8b98d8d607db6684)
}

var fileDescriptor_registryconsumer_8b98d8d607db6684 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0xaa, 0xd4, 0x30,
	0x14, 0x86, 0x69, 0xe7, 0x3a, 0x32, 0xb9, 0x82, 0x43, 0x16, 0xf7, 0x96, 0x41, 0x44, 0xeb, 0x46,
	0x54, 0x12, 0x7a, 0x15, 0x14, 0x97, 0xd7, 0x37, 0x08, 0xb8, 0x71, 0x97, 0xc9, 0x1c, 0x6a, 0xb0,
	0x4d, 0x62, 0x72, 0x5a, 0x18, 0x4a, 0x17, 0xba, 0x72, 0xef, 0xdb, 0xb8, 0xf2, 0x1d, 0x7c, 0x05,
	0x37, 0x82, 0x0f, 0x21, 0x9d, 0xa6, 0x50, 0x2b, 0xea, 0xee, 0x9c, 0xfe, 0xff, 0x9f, 0x1e, 0xbe,
	0x73, 0xc8, 0x85, 0x87, 0x52, 0x07, 0xf4, 0x47, 0x65, 0x4d, 0x68, 0x6a, 0xf0, 0xcc, 0x79, 0x8b,
	0x96, 0x5e, 0x22, 0xa0, 0x97, 0x08, 0x4c, 0x3a, 0xcd, 0x50, 0x29, 0xa6, 0xac, 0x07, 0xd6, 0x16,
	0xbb, 0x3b, 0xa5, 0xb5, 0x65, 0x05, 0x5c, 0x3a, 0xcd, 0xa5, 0x31, 0x16, 0x25, 0x6a, 0x6b, 0xc2,
	0x18, 0xdb, 0x5d, 0xb6, 0xb2, 0xd2, 0x07, 0x89, 0xc0, 0xa7, 0x62, 0x14, 0xf2, 0x4f, 0x29, 0xb9,
	0x2b, 0xe2, 0xaf, 0x5e, 0xbb, 0x41, 0x10, 0x10, 0x6c, 0xe3, 0x15, 0x04, 0x01, 0xef, 0x1b, 0x08,
	0x48, 0x2f, 0xc8, 0xda, 0x49, 0x0f, 0x06, 0xb3, 0xe4, 0x5e, 0xf2, 0x70, 0x23, 0x62, 0x47, 0xef,
	0x93, 0x35, 0x82, 0x91, 0x06, 0xb3, 0x74, 0xf8, 0x7e, 0xbd, 0xf9, 0xf2, 0xe3, 0xeb, 0xea, 0xcc,
	0xa7, 0xdb, 0x44, 0x44, 0x81, 0x3e, 0x26, 0xe7, 0x60, 0x5a, 0xed, 0xad, 0xa9, 0x87, 0xfc, 0x6a,
	0xe9, 0x9b, 0xab, 0xf4, 0x01, 0xb9, 0xa9, 0xaa, 0x26, 0x20, 0xf8, 0xec, 0x6c, 0x69, 0x9c, 0x14,
	0xfa, 0x8c, 0xdc, 0x9a, 0xc8, 0xe0, 0xd1, 0x41, 0x76, 0xe3, 0xe4, 0xdc, 0x0e, 0xce, 0x73, 0xbf,
	0x11, 0xab, 0x77, 0x2f, 0x82, 0x48, 0xdb, 0x5a, 0xfc, 0xe6, 0x1a, 0x9e, 0x76, 0xf2, 0x58, 0x59,
	0x79, 0xc8, 0xd6, 0x7f, 0x3c, 0x1d, 0x95, 0xfc, 0x8a, 0x64, 0x13, 0x89, 0x57, 0x11, 0xba, 0x80,
	0xe0, 0xac, 0x09, 0x30, 0x30, 0x08, 0x28, 0xb1, 0x09, 0x13, 0x83, 0xb1, 0xbb, 0xfa, 0x90, 0x92,
	0xed, 0x32, 0x44, 0x7f, 0x26, 0xe4, 0xf6, 0x82, 0x25, 0x7d, 0xce, 0xfe, 0xb2, 0x38, 0xf6, 0x6f,
	0xfa, 0xbb, 0xe2, 0xbf, 0xc1, 0xe5, 0xb0, 0x79, 0xfd, 0xf1, 0xdb, 0xf7, 0xcf, 0x69, 0x99, 0xef,
	0x79, 0x5b, 0xf0, 0x71, 0x13, 0x81, 0x77, 0x63, 0xd1, 0xf3, 0x19, 0xf2, 0xc0, 0xbb, 0x59, 0xd7,
	0xf3, 0x08, 0x39, 0xf0, 0x2e, 0x56, 0x3d, 0x5f, 0x9e, 0x21, 0xef, 0xe6, 0x60, 0xfb, 0x97, 0xc9,
	0xa3, 0x6b, 0xf6, 0xe6, 0x49, 0xa9, 0xf1, 0x6d, 0xb3, 0x67, 0xca, 0xd6, 0x3c, 0x4e, 0xab, 0xed,
	0x54, 0x9d, 0x2e, 0x12, 0x95, 0xe2, 0xc3, 0xdc, 0xbc, 0x2d, 0xf6, 0xeb, 0xd3, 0xe5, 0x3d, 0xfd,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x05, 0xff, 0x49, 0x8e, 0xe3, 0x02, 0x00, 0x00,
}
