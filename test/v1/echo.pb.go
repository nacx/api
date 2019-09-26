// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package testv1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EchoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_007b671ac1918d0b, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (dst *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(dst, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EchoRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type EchoResponse struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 []byte               `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	ServerTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_007b671ac1918d0b, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (dst *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(dst, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EchoResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *EchoResponse) GetServerTime() *timestamp.Timestamp {
	if m != nil {
		return m.ServerTime
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "tetrate.api.test.v1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "tetrate.api.test.v1.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.test.v1.Service/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.test.v1.Service/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.test.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Service_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_echo_007b671ac1918d0b) }

var fileDescriptor_echo_007b671ac1918d0b = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x95, 0xb4, 0x2a, 0xad, 0x5b, 0x10, 0x72, 0x07, 0xa2, 0x08, 0x44, 0xc8, 0x14, 0x31,
	0xd8, 0x4a, 0xd9, 0x80, 0xa9, 0x12, 0x17, 0x08, 0x4c, 0x2c, 0xe0, 0x24, 0x8f, 0xd6, 0xa2, 0x8d,
	0x43, 0xfc, 0x6a, 0x09, 0x10, 0x0b, 0x57, 0xe0, 0x36, 0x4c, 0xdc, 0x81, 0x2b, 0xb0, 0x70, 0x0b,
	0x94, 0xb8, 0x91, 0x2a, 0x81, 0xd8, 0xec, 0xf7, 0xfe, 0x4f, 0xfe, 0x3f, 0x13, 0x02, 0xd9, 0x5c,
	0xb1, 0xb2, 0x52, 0xa8, 0xe8, 0x18, 0x01, 0x2b, 0x81, 0xc0, 0x44, 0x29, 0x19, 0x82, 0x46, 0x66,
	0x62, 0x7f, 0x7f, 0xa6, 0xd4, 0x6c, 0x01, 0x5c, 0x94, 0x92, 0x8b, 0xa2, 0x50, 0x28, 0x50, 0xaa,
	0x42, 0x5b, 0xc4, 0x3f, 0x5c, 0x6f, 0x9b, 0x5b, 0xba, 0xba, 0xe3, 0x28, 0x97, 0xa0, 0x51, 0x2c,
	0xcb, 0x75, 0x60, 0xcf, 0x88, 0x85, 0xcc, 0x05, 0x02, 0x6f, 0x0f, 0x76, 0x11, 0x9e, 0x93, 0xe1,
	0x45, 0x36, 0x57, 0x09, 0x3c, 0xac, 0x40, 0x23, 0xdd, 0x21, 0xae, 0xcc, 0x3d, 0x27, 0x70, 0xa2,
	0x41, 0xe2, 0xca, 0x9c, 0x1e, 0x90, 0x6e, 0xaa, 0xf2, 0x47, 0xcf, 0x0d, 0x9c, 0x68, 0x34, 0x1d,
	0xbc, 0x7f, 0x7f, 0x74, 0xba, 0x4f, 0xee, 0x6e, 0x27, 0x69, 0xc6, 0xa1, 0x22, 0x23, 0x4b, 0xeb,
	0x52, 0x15, 0x1a, 0x7e, 0xe1, 0x74, 0x13, 0xb7, 0x0c, 0x3d, 0x23, 0x43, 0x0d, 0x95, 0x81, 0xea,
	0xa6, 0x2e, 0xe9, 0x75, 0x02, 0x27, 0x1a, 0x4e, 0x7c, 0x66, 0x0d, 0x58, 0x6b, 0xc0, 0xae, 0x5a,
	0x83, 0x84, 0xd8, 0x78, 0x3d, 0x98, 0xdc, 0x93, 0xad, 0x4b, 0xa8, 0x8c, 0xcc, 0x80, 0xde, 0x92,
	0x6e, 0xfd, 0x36, 0x0d, 0xd8, 0x1f, 0xff, 0xc5, 0x36, 0xa4, 0xfc, 0xa3, 0x7f, 0x12, 0xb6, 0x78,
	0x38, 0x7e, 0xfd, 0xfc, 0x7a, 0x73, 0xb7, 0xfd, 0x3e, 0x37, 0x31, 0x7f, 0x96, 0xf9, 0xcb, 0xa9,
	0x73, 0x3c, 0xed, 0x5f, 0xf7, 0xea, 0xb0, 0x89, 0xd3, 0x5e, 0x53, 0xeb, 0xe4, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xc4, 0xb9, 0x57, 0x27, 0xa7, 0x01, 0x00, 0x00,
}
