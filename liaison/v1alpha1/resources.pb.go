// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

package v1alpha1 // import "github.com/tetrateio/tetrate/api/liaison/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"
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

// A representation of a stored resource.
type Resource struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_71915ca02609450d, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// A request message to upload a resource in YAML format.
type SetResourceRequest struct {
	// Resource representation in YAML bytes.
	Yaml                 []byte   `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResourceRequest) Reset()         { *m = SetResourceRequest{} }
func (m *SetResourceRequest) String() string { return proto.CompactTextString(m) }
func (*SetResourceRequest) ProtoMessage()    {}
func (*SetResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_71915ca02609450d, []int{1}
}
func (m *SetResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResourceRequest.Unmarshal(m, b)
}
func (m *SetResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResourceRequest.Marshal(b, m, deterministic)
}
func (dst *SetResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResourceRequest.Merge(dst, src)
}
func (m *SetResourceRequest) XXX_Size() int {
	return xxx_messageInfo_SetResourceRequest.Size(m)
}
func (m *SetResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetResourceRequest proto.InternalMessageInfo

func (m *SetResourceRequest) GetYaml() []byte {
	if m != nil {
		return m.Yaml
	}
	return nil
}

// Request to get an uploaded resource.
type GetResourceRequest struct {
	// The resource key.
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourceRequest) Reset()         { *m = GetResourceRequest{} }
func (m *GetResourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourceRequest) ProtoMessage()    {}
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_71915ca02609450d, []int{2}
}
func (m *GetResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourceRequest.Unmarshal(m, b)
}
func (m *GetResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourceRequest.Marshal(b, m, deterministic)
}
func (dst *GetResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceRequest.Merge(dst, src)
}
func (m *GetResourceRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourceRequest.Size(m)
}
func (m *GetResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceRequest proto.InternalMessageInfo

func (m *GetResourceRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "tetrate.api.liaison.v1alpha1.Resource")
	proto.RegisterType((*SetResourceRequest)(nil), "tetrate.api.liaison.v1alpha1.SetResourceRequest")
	proto.RegisterType((*GetResourceRequest)(nil), "tetrate.api.liaison.v1alpha1.GetResourceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourcesServiceClient is the client API for ResourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourcesServiceClient interface {
	// SetResource sets a resource from YAML bytes. When the resource exists it
	// updates the corresponding resource.
	SetResource(ctx context.Context, in *SetResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	// GetResource gets a resource by the provided key.
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error)
}

type resourcesServiceClient struct {
	cc *grpc.ClientConn
}

func NewResourcesServiceClient(cc *grpc.ClientConn) ResourcesServiceClient {
	return &resourcesServiceClient{cc}
}

func (c *resourcesServiceClient) SetResource(ctx context.Context, in *SetResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/tetrate.api.liaison.v1alpha1.ResourcesService/SetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesServiceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/tetrate.api.liaison.v1alpha1.ResourcesService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcesServiceServer is the server API for ResourcesService service.
type ResourcesServiceServer interface {
	// SetResource sets a resource from YAML bytes. When the resource exists it
	// updates the corresponding resource.
	SetResource(context.Context, *SetResourceRequest) (*Resource, error)
	// GetResource gets a resource by the provided key.
	GetResource(context.Context, *GetResourceRequest) (*Resource, error)
}

func RegisterResourcesServiceServer(s *grpc.Server, srv ResourcesServiceServer) {
	s.RegisterService(&_ResourcesService_serviceDesc, srv)
}

func _ResourcesService_SetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServiceServer).SetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.liaison.v1alpha1.ResourcesService/SetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServiceServer).SetResource(ctx, req.(*SetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourcesService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.liaison.v1alpha1.ResourcesService/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServiceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourcesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.liaison.v1alpha1.ResourcesService",
	HandlerType: (*ResourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetResource",
			Handler:    _ResourcesService_SetResource_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _ResourcesService_GetResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources.proto",
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor_resources_71915ca02609450d) }

var fileDescriptor_resources_71915ca02609450d = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0x49, 0x2d,
	0x29, 0x4a, 0x2c, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0xcb, 0xc9, 0x4c, 0xcc, 0x2c, 0xce, 0xcf,
	0xd3, 0x2b, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0x83, 0xea, 0x95, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20,
	0x12, 0x4a, 0x32, 0x5c, 0x1c, 0x41, 0x50, 0x7b, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x0d, 0x2e, 0xa1, 0xe0, 0xd4, 0x12, 0x98,
	0x82, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xca, 0xc4, 0xdc, 0x1c,
	0xb0, 0x42, 0x9e, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x4b, 0xc8, 0x1d, 0x53, 0x25, 0x86, 0x89, 0x46,
	0x2b, 0x99, 0xb8, 0x04, 0x60, 0xaa, 0x8a, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xda,
	0x19, 0xb9, 0xb8, 0x91, 0xec, 0x11, 0x32, 0xd0, 0xc3, 0xe7, 0x55, 0x3d, 0x4c, 0x27, 0x49, 0xa9,
	0xe1, 0xd7, 0x01, 0x53, 0xae, 0x24, 0xd7, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x09, 0x25, 0x61, 0x7d,
	0x98, 0x9c, 0x3e, 0x3c, 0x98, 0xad, 0x18, 0xb5, 0x84, 0xba, 0x18, 0xb9, 0xb8, 0xdd, 0x89, 0x77,
	0x89, 0x3b, 0xf9, 0x2e, 0x51, 0x04, 0xbb, 0x44, 0x5a, 0x48, 0x12, 0x8b, 0x4b, 0xf4, 0xab, 0xb3,
	0x53, 0x2b, 0x6b, 0x9d, 0x8c, 0xa3, 0x0c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xa1, 0xc6, 0x66, 0xe6, 0xc3, 0x58, 0xe0, 0xa8, 0x86, 0x5a, 0x00, 0x37, 0x24, 0x89,
	0x0d, 0x1c, 0xaf, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x13, 0x98, 0xad, 0x3f, 0x02,
	0x00, 0x00,
}
