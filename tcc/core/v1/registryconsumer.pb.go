// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registryconsumer.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/core/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

// Please do not over engineer this anything beyond a glorified
// file copy.
type RegistryData struct {
	// currently accepted values are kubernetes or f5
	Registrytype string `protobuf:"bytes,1,opt,name=registrytype,proto3" json:"registrytype,omitempty"`
	Tenant       string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// A unique identifier for the k8s/f5 cluster
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// A huge json dump of all services, endpoints.
	// For k8s, its literally kubectl get services --all-namespaces -o json
	// and kubectl get po --all-namespaces -o json
	Payload              string   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryData) Reset()         { *m = RegistryData{} }
func (m *RegistryData) String() string { return proto.CompactTextString(m) }
func (*RegistryData) ProtoMessage()    {}
func (*RegistryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_registryconsumer_c7007a092ca970f2, []int{0}
}
func (m *RegistryData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryData.Unmarshal(m, b)
}
func (m *RegistryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryData.Marshal(b, m, deterministic)
}
func (dst *RegistryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryData.Merge(dst, src)
}
func (m *RegistryData) XXX_Size() int {
	return xxx_messageInfo_RegistryData.Size(m)
}
func (m *RegistryData) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryData.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryData proto.InternalMessageInfo

func (m *RegistryData) GetRegistrytype() string {
	if m != nil {
		return m.Registrytype
	}
	return ""
}

func (m *RegistryData) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *RegistryData) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RegistryData) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistryData)(nil), "tetrate.api.tcc.core.v1.RegistryData")
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
	Upload(ctx context.Context, in *RegistryData, opts ...grpc.CallOption) (*empty.Empty, error)
}

type registryConsumerClient struct {
	cc *grpc.ClientConn
}

func NewRegistryConsumerClient(cc *grpc.ClientConn) RegistryConsumerClient {
	return &registryConsumerClient{cc}
}

func (c *registryConsumerClient) Upload(ctx context.Context, in *RegistryData, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.core.v1.RegistryConsumer/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryConsumerServer is the server API for RegistryConsumer service.
type RegistryConsumerServer interface {
	Upload(context.Context, *RegistryData) (*empty.Empty, error)
}

func RegisterRegistryConsumerServer(s *grpc.Server, srv RegistryConsumerServer) {
	s.RegisterService(&_RegistryConsumer_serviceDesc, srv)
}

func _RegistryConsumer_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryConsumerServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.core.v1.RegistryConsumer/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryConsumerServer).Upload(ctx, req.(*RegistryData))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryConsumer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.tcc.core.v1.RegistryConsumer",
	HandlerType: (*RegistryConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _RegistryConsumer_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registryconsumer.proto",
}

func init() {
	proto.RegisterFile("registryconsumer.proto", fileDescriptor_registryconsumer_c7007a092ca970f2)
}

var fileDescriptor_registryconsumer_c7007a092ca970f2 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x10, 0xc7, 0xe9, 0x7e, 0x1f, 0x2b, 0x86, 0x3d, 0x48, 0x0e, 0x6b, 0x59, 0x3d, 0x48, 0x41, 0x10,
	0x91, 0x19, 0x56, 0x6f, 0x1e, 0x3c, 0xb8, 0x7a, 0xf0, 0xba, 0xe0, 0xc5, 0x5b, 0x36, 0xc6, 0x1a,
	0xd8, 0x4d, 0x42, 0x3a, 0x2d, 0x94, 0x52, 0x0f, 0x3e, 0x82, 0x3e, 0x84, 0x0f, 0xe4, 0x2b, 0xf8,
	0x20, 0xd2, 0x26, 0x85, 0x55, 0xf0, 0x36, 0xff, 0xfc, 0x08, 0xf3, 0xfb, 0x0f, 0x9b, 0x7a, 0x95,
	0xeb, 0x82, 0x7c, 0x2d, 0xad, 0x29, 0xca, 0x8d, 0xf2, 0xe0, 0xbc, 0x25, 0xcb, 0xf7, 0x49, 0x91,
	0x17, 0xa4, 0x40, 0x38, 0x0d, 0x24, 0x25, 0x48, 0xeb, 0x15, 0x54, 0xf3, 0xd9, 0x61, 0x6e, 0x6d,
	0xbe, 0x56, 0x28, 0x9c, 0x46, 0x61, 0x8c, 0x25, 0x41, 0xda, 0x9a, 0x22, 0x7c, 0x9b, 0x1d, 0x44,
	0xda, 0xa7, 0x55, 0xf9, 0x84, 0x6a, 0xe3, 0xa8, 0x0e, 0x30, 0x7b, 0x61, 0x93, 0x65, 0xdc, 0x76,
	0x23, 0x48, 0xf0, 0x8c, 0x4d, 0x86, 0xed, 0x54, 0x3b, 0x95, 0x26, 0x47, 0xc9, 0xc9, 0xee, 0xf2,
	0xc7, 0x1b, 0x9f, 0xb2, 0x31, 0x29, 0x23, 0x0c, 0xa5, 0xa3, 0x9e, 0xc6, 0xc4, 0x53, 0xb6, 0x23,
	0xd7, 0x65, 0x41, 0xca, 0xa7, 0xff, 0x7a, 0x30, 0xc4, 0x8e, 0x38, 0x51, 0xaf, 0xad, 0x78, 0x4c,
	0xff, 0x07, 0x12, 0xe3, 0xf9, 0x47, 0xc2, 0xf6, 0x06, 0x81, 0x45, 0xac, 0xcb, 0xdf, 0x12, 0x36,
	0xbe, 0x77, 0x1d, 0xe7, 0xc7, 0xf0, 0x47, 0x69, 0xd8, 0xd6, 0x9e, 0x4d, 0x21, 0x94, 0x84, 0xa1,
	0x24, 0xdc, 0x76, 0x25, 0xb3, 0xbb, 0xd7, 0xcf, 0xaf, 0xf7, 0xd1, 0x22, 0xbb, 0xc2, 0x6a, 0x8e,
	0x41, 0xb3, 0xc0, 0x26, 0x0c, 0x2d, 0xfe, 0xbe, 0x33, 0x36, 0xdb, 0x3d, 0x5b, 0x6c, 0xa2, 0x7f,
	0x7b, 0x99, 0x9c, 0x5e, 0xc3, 0xc3, 0x59, 0xae, 0xe9, 0xb9, 0x5c, 0x81, 0xb4, 0x1b, 0x8c, 0x56,
	0xda, 0x0e, 0x53, 0x7f, 0x7c, 0x92, 0x12, 0x3b, 0x3f, 0xac, 0xe6, 0xab, 0x71, 0xaf, 0x72, 0xf1,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x45, 0x72, 0x9d, 0xce, 0x01, 0x00, 0x00,
}
