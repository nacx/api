// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Tetrate Service definition
type ServiceMessage struct {
	// Service FQDN, kind
	Name string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind ServiceKind `protobuf:"varint,2,opt,name=kind,proto3,enum=tetrate.api.inventory.v1alpha1.ServiceKind" json:"kind,omitempty"`
	// Service specific metadata
	Metadata *Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Status
	Status ServiceStatus `protobuf:"varint,4,opt,name=status,proto3,enum=tetrate.api.inventory.v1alpha1.ServiceStatus" json:"status,omitempty"`
	// Cluster information, endpoints
	Cluster []*Cluster `protobuf:"bytes,5,rep,name=cluster,proto3" json:"cluster,omitempty"`
	// Cloud workload information
	Workload *Workload `protobuf:"bytes,6,opt,name=workload,proto3" json:"workload,omitempty"`
	// last modified
	LastModified         *timestamp.Timestamp `protobuf:"bytes,7,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ServiceMessage) Reset()         { *m = ServiceMessage{} }
func (m *ServiceMessage) String() string { return proto.CompactTextString(m) }
func (*ServiceMessage) ProtoMessage()    {}
func (*ServiceMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5ca993bfceebbd4f, []int{0}
}
func (m *ServiceMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceMessage.Unmarshal(m, b)
}
func (m *ServiceMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceMessage.Marshal(b, m, deterministic)
}
func (dst *ServiceMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMessage.Merge(dst, src)
}
func (m *ServiceMessage) XXX_Size() int {
	return xxx_messageInfo_ServiceMessage.Size(m)
}
func (m *ServiceMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMessage proto.InternalMessageInfo

func (m *ServiceMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceMessage) GetKind() ServiceKind {
	if m != nil {
		return m.Kind
	}
	return ServiceKind_SERVICE
}

func (m *ServiceMessage) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ServiceMessage) GetStatus() ServiceStatus {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_UNHEALTHY
}

func (m *ServiceMessage) GetCluster() []*Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ServiceMessage) GetWorkload() *Workload {
	if m != nil {
		return m.Workload
	}
	return nil
}

func (m *ServiceMessage) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type ServiceResponse struct {
	RespMsg              []*ServiceMessage `protobuf:"bytes,1,rep,name=respMsg,proto3" json:"respMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceResponse) Reset()         { *m = ServiceResponse{} }
func (m *ServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceResponse) ProtoMessage()    {}
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5ca993bfceebbd4f, []int{1}
}
func (m *ServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceResponse.Unmarshal(m, b)
}
func (m *ServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceResponse.Marshal(b, m, deterministic)
}
func (dst *ServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceResponse.Merge(dst, src)
}
func (m *ServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceResponse.Size(m)
}
func (m *ServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceResponse proto.InternalMessageInfo

func (m *ServiceResponse) GetRespMsg() []*ServiceMessage {
	if m != nil {
		return m.RespMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceMessage)(nil), "tetrate.api.inventory.v1alpha1.ServiceMessage")
	proto.RegisterType((*ServiceResponse)(nil), "tetrate.api.inventory.v1alpha1.ServiceResponse")
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
	GetServices(ctx context.Context, in *ServiceMessage, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetServices(ctx context.Context, in *ServiceMessage, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1alpha1.Service/getServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	GetServices(context.Context, *ServiceMessage) (*ServiceResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1alpha1.Service/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetServices(ctx, req.(*ServiceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.inventory.v1alpha1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getServices",
			Handler:    _Service_GetServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_5ca993bfceebbd4f) }

var fileDescriptor_service_5ca993bfceebbd4f = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8a, 0xd4, 0x40,
	0x14, 0xc6, 0xc9, 0x74, 0xdb, 0x3d, 0x56, 0x4f, 0x46, 0x28, 0x37, 0x45, 0x10, 0x0d, 0xbd, 0x31,
	0x20, 0x56, 0xe8, 0x78, 0x80, 0xc1, 0x7f, 0x20, 0x48, 0x36, 0x35, 0x82, 0xa0, 0x0b, 0x79, 0xd3,
	0x79, 0x13, 0x8b, 0x49, 0xaa, 0x42, 0xea, 0x75, 0xc4, 0xad, 0x1e, 0xc0, 0x85, 0xd7, 0xf1, 0x16,
	0x5e, 0xc1, 0x83, 0xc8, 0x54, 0x2a, 0xb3, 0x74, 0xe2, 0x2e, 0x8f, 0xfa, 0x7e, 0x1f, 0xef, 0x7b,
	0x1f, 0x61, 0xb1, 0xc3, 0x7e, 0xd0, 0x7b, 0x94, 0x5d, 0x6f, 0xc9, 0xf2, 0x87, 0x84, 0xd4, 0x03,
	0xa1, 0x84, 0x4e, 0x4b, 0x6d, 0x06, 0x34, 0x64, 0xfb, 0xaf, 0x72, 0xd8, 0x41, 0xd3, 0x7d, 0x86,
	0x5d, 0xf2, 0xa0, 0xb6, 0xb6, 0x6e, 0x30, 0x87, 0x4e, 0xe7, 0x60, 0x8c, 0x25, 0x20, 0x6d, 0x8d,
	0x1b, 0xe9, 0xe4, 0x51, 0x78, 0xf5, 0xd3, 0xc5, 0xe1, 0x32, 0x27, 0xdd, 0xa2, 0x23, 0x68, 0xbb,
	0x20, 0x38, 0xd9, 0xdb, 0xb6, 0xb5, 0x26, 0x4c, 0xf1, 0xbe, 0x39, 0x38, 0xc2, 0x7e, 0x1c, 0xb7,
	0xbf, 0x16, 0xec, 0xf4, 0x7c, 0xdc, 0xa6, 0x44, 0xe7, 0xa0, 0x46, 0xce, 0xd9, 0xd2, 0x40, 0x8b,
	0x22, 0x4a, 0xa3, 0xec, 0xae, 0xf2, 0xdf, 0xfc, 0x8c, 0x2d, 0xaf, 0xb4, 0xa9, 0xc4, 0x51, 0x1a,
	0x65, 0xa7, 0xc5, 0x13, 0xf9, 0xef, 0x8d, 0x65, 0x70, 0x7c, 0xab, 0x4d, 0xa5, 0x3c, 0xc8, 0x5f,
	0xb1, 0xe3, 0x16, 0x09, 0x2a, 0x20, 0x10, 0x8b, 0x34, 0xca, 0x36, 0x45, 0x76, 0x9b, 0x49, 0x19,
	0xf4, 0xea, 0x86, 0xe4, 0xaf, 0xd9, 0xca, 0x11, 0xd0, 0xc1, 0x89, 0xa5, 0x5f, 0xe4, 0xe9, 0xcc,
	0x45, 0xce, 0x3d, 0xa4, 0x02, 0xcc, 0x9f, 0xb3, 0x75, 0xb8, 0x82, 0xb8, 0x93, 0x2e, 0xb2, 0x4d,
	0xf1, 0xf8, 0x36, 0x9f, 0x97, 0xa3, 0x5c, 0x4d, 0xdc, 0x75, 0x9e, 0x2f, 0xb6, 0xbf, 0x6a, 0x2c,
	0x54, 0x62, 0x35, 0x2f, 0xcf, 0xfb, 0xa0, 0x57, 0x37, 0x24, 0x3f, 0x63, 0x71, 0x03, 0x8e, 0x3e,
	0xb5, 0xb6, 0xd2, 0x97, 0x1a, 0x2b, 0xb1, 0xf6, 0x56, 0x89, 0x1c, 0x3b, 0x95, 0x53, 0xa7, 0xf2,
	0xdd, 0xd4, 0xa9, 0x3a, 0xb9, 0x06, 0xca, 0xa0, 0xdf, 0x7e, 0x64, 0xf7, 0x42, 0x44, 0x85, 0xae,
	0xb3, 0xc6, 0x21, 0x7f, 0xc3, 0xd6, 0x3d, 0xba, 0xae, 0x74, 0xb5, 0x88, 0x7c, 0x38, 0x39, 0xf3,
	0x48, 0xa1, 0x7f, 0x35, 0xe1, 0xc5, 0x8f, 0x88, 0xad, 0xc3, 0x1b, 0xff, 0x1e, 0xb1, 0x4d, 0x8d,
	0x14, 0x46, 0xc7, 0xff, 0xd3, 0x34, 0xc9, 0x67, 0xea, 0xa7, 0x18, 0xdb, 0xfb, 0xdf, 0x7e, 0xff,
	0xf9, 0x79, 0x14, 0xf3, 0x4d, 0x3e, 0xec, 0xf2, 0xf0, 0xbf, 0xbc, 0x60, 0x1f, 0x8e, 0x27, 0xe0,
	0x62, 0xe5, 0x8f, 0xf3, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x93, 0xca, 0x76, 0x4d,
	0x03, 0x00, 0x00,
}
