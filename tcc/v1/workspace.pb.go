// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workspace.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/v1"

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

type Workspace struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// The services (or namespaces) that this endpoints in this
	// workspace depend upon for proper operation. Must be of the form
	// `tenant/123/workspace/456/service/foo.com`, or
	// `tenant/123/workspace/456` . If omitted, its assumed that
	// endpoints in this workspace depend only on other services in the
	// same workspace as the endpoint.
	Dependencies         []string `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{0}
}
func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workspace.Unmarshal(m, b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
}
func (dst *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(dst, src)
}
func (m *Workspace) XXX_Size() int {
	return xxx_messageInfo_Workspace.Size(m)
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Workspace) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *Workspace) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type CreateWorkspaceRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// The services (or namespaces) that this endpoints in this
	// workspace depend upon for proper operation. Must be of the form
	// `tenant/123/workspace/456/service/foo.com`, or
	// `tenant/123/workspace/456` . If omitted, its assumed that
	// endpoints in this workspace depend only on other services in the
	// same workspace as the endpoint.
	Dependencies         []string `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWorkspaceRequest) Reset()         { *m = CreateWorkspaceRequest{} }
func (m *CreateWorkspaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkspaceRequest) ProtoMessage()    {}
func (*CreateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{1}
}
func (m *CreateWorkspaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkspaceRequest.Unmarshal(m, b)
}
func (m *CreateWorkspaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkspaceRequest.Marshal(b, m, deterministic)
}
func (dst *CreateWorkspaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkspaceRequest.Merge(dst, src)
}
func (m *CreateWorkspaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkspaceRequest.Size(m)
}
func (m *CreateWorkspaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkspaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkspaceRequest proto.InternalMessageInfo

func (m *CreateWorkspaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateWorkspaceRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *CreateWorkspaceRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type UpdateWorkspaceRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// The services (or namespaces) that this endpoints in this
	// workspace depend upon for proper operation. Must be of the form
	// `tenant/123/workspace/456/service/foo.com`, or
	// `tenant/123/workspace/456` . If omitted, its assumed that
	// endpoints in this workspace depend only on other services in the
	// same workspace as the endpoint.
	Dependencies         []string `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWorkspaceRequest) Reset()         { *m = UpdateWorkspaceRequest{} }
func (m *UpdateWorkspaceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWorkspaceRequest) ProtoMessage()    {}
func (*UpdateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{2}
}
func (m *UpdateWorkspaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWorkspaceRequest.Unmarshal(m, b)
}
func (m *UpdateWorkspaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWorkspaceRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateWorkspaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWorkspaceRequest.Merge(dst, src)
}
func (m *UpdateWorkspaceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWorkspaceRequest.Size(m)
}
func (m *UpdateWorkspaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWorkspaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWorkspaceRequest proto.InternalMessageInfo

func (m *UpdateWorkspaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateWorkspaceRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *UpdateWorkspaceRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type GetWorkspaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant               string   `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkspaceRequest) Reset()         { *m = GetWorkspaceRequest{} }
func (m *GetWorkspaceRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkspaceRequest) ProtoMessage()    {}
func (*GetWorkspaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{3}
}
func (m *GetWorkspaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkspaceRequest.Unmarshal(m, b)
}
func (m *GetWorkspaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkspaceRequest.Marshal(b, m, deterministic)
}
func (dst *GetWorkspaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkspaceRequest.Merge(dst, src)
}
func (m *GetWorkspaceRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkspaceRequest.Size(m)
}
func (m *GetWorkspaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkspaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkspaceRequest proto.InternalMessageInfo

func (m *GetWorkspaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetWorkspaceRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

type ListWorkspacesRequest struct {
	Tenant               string   `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkspacesRequest) Reset()         { *m = ListWorkspacesRequest{} }
func (m *ListWorkspacesRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkspacesRequest) ProtoMessage()    {}
func (*ListWorkspacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{4}
}
func (m *ListWorkspacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkspacesRequest.Unmarshal(m, b)
}
func (m *ListWorkspacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkspacesRequest.Marshal(b, m, deterministic)
}
func (dst *ListWorkspacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkspacesRequest.Merge(dst, src)
}
func (m *ListWorkspacesRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkspacesRequest.Size(m)
}
func (m *ListWorkspacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkspacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkspacesRequest proto.InternalMessageInfo

func (m *ListWorkspacesRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

type ListWorkspacesResponse struct {
	Workspaces           []*Workspace `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListWorkspacesResponse) Reset()         { *m = ListWorkspacesResponse{} }
func (m *ListWorkspacesResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkspacesResponse) ProtoMessage()    {}
func (*ListWorkspacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{5}
}
func (m *ListWorkspacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkspacesResponse.Unmarshal(m, b)
}
func (m *ListWorkspacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkspacesResponse.Marshal(b, m, deterministic)
}
func (dst *ListWorkspacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkspacesResponse.Merge(dst, src)
}
func (m *ListWorkspacesResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkspacesResponse.Size(m)
}
func (m *ListWorkspacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkspacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkspacesResponse proto.InternalMessageInfo

func (m *ListWorkspacesResponse) GetWorkspaces() []*Workspace {
	if m != nil {
		return m.Workspaces
	}
	return nil
}

type DeleteWorkspaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tenant               string   `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteWorkspaceRequest) Reset()         { *m = DeleteWorkspaceRequest{} }
func (m *DeleteWorkspaceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteWorkspaceRequest) ProtoMessage()    {}
func (*DeleteWorkspaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workspace_8c32a6b65efea193, []int{6}
}
func (m *DeleteWorkspaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWorkspaceRequest.Unmarshal(m, b)
}
func (m *DeleteWorkspaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWorkspaceRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteWorkspaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWorkspaceRequest.Merge(dst, src)
}
func (m *DeleteWorkspaceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteWorkspaceRequest.Size(m)
}
func (m *DeleteWorkspaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWorkspaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWorkspaceRequest proto.InternalMessageInfo

func (m *DeleteWorkspaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteWorkspaceRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func init() {
	proto.RegisterType((*Workspace)(nil), "tetrate.api.tcc.v1.Workspace")
	proto.RegisterType((*CreateWorkspaceRequest)(nil), "tetrate.api.tcc.v1.CreateWorkspaceRequest")
	proto.RegisterType((*UpdateWorkspaceRequest)(nil), "tetrate.api.tcc.v1.UpdateWorkspaceRequest")
	proto.RegisterType((*GetWorkspaceRequest)(nil), "tetrate.api.tcc.v1.GetWorkspaceRequest")
	proto.RegisterType((*ListWorkspacesRequest)(nil), "tetrate.api.tcc.v1.ListWorkspacesRequest")
	proto.RegisterType((*ListWorkspacesResponse)(nil), "tetrate.api.tcc.v1.ListWorkspacesResponse")
	proto.RegisterType((*DeleteWorkspaceRequest)(nil), "tetrate.api.tcc.v1.DeleteWorkspaceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkspacesClient is the client API for Workspaces service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkspacesClient interface {
	// admin only: Typically, create namespace fails if the namespace already exists in a service
	// registry import from systems like kubernetes or consul.
	// Such namespaces cannot be created/deleted. Only updates can be made.
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error)
	// anyone: Get properties of a namespace. If namespace name is not provided,
	// lists properties of all namespaces in the system.
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error)
	// admin or worksapce owner: Update properties of a namespace
	UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error)
	ListWorkspaces(ctx context.Context, in *ListWorkspacesRequest, opts ...grpc.CallOption) (*ListWorkspacesResponse, error)
	// admin only: Typically, delete namespace fails if the namespace exists in a service
	// registry import from systems like kubernetes or consul.
	// Such namespaces cannot be created/deleted. Only updates can be made.
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type workspacesClient struct {
	cc *grpc.ClientConn
}

func NewWorkspacesClient(cc *grpc.ClientConn) WorkspacesClient {
	return &workspacesClient{cc}
}

func (c *workspacesClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error) {
	out := new(Workspace)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.v1.Workspaces/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error) {
	out := new(Workspace)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.v1.Workspaces/GetWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesClient) UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error) {
	out := new(Workspace)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.v1.Workspaces/UpdateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesClient) ListWorkspaces(ctx context.Context, in *ListWorkspacesRequest, opts ...grpc.CallOption) (*ListWorkspacesResponse, error) {
	out := new(ListWorkspacesResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.v1.Workspaces/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tetrate.api.tcc.v1.Workspaces/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspacesServer is the server API for Workspaces service.
type WorkspacesServer interface {
	// admin only: Typically, create namespace fails if the namespace already exists in a service
	// registry import from systems like kubernetes or consul.
	// Such namespaces cannot be created/deleted. Only updates can be made.
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*Workspace, error)
	// anyone: Get properties of a namespace. If namespace name is not provided,
	// lists properties of all namespaces in the system.
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*Workspace, error)
	// admin or worksapce owner: Update properties of a namespace
	UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*Workspace, error)
	ListWorkspaces(context.Context, *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	// admin only: Typically, delete namespace fails if the namespace exists in a service
	// registry import from systems like kubernetes or consul.
	// Such namespaces cannot be created/deleted. Only updates can be made.
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*empty.Empty, error)
}

func RegisterWorkspacesServer(s *grpc.Server, srv WorkspacesServer) {
	s.RegisterService(&_Workspaces_serviceDesc, srv)
}

func _Workspaces_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.v1.Workspaces/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspaces_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.v1.Workspaces/GetWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspaces_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.v1.Workspaces/UpdateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServer).UpdateWorkspace(ctx, req.(*UpdateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspaces_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.v1.Workspaces/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServer).ListWorkspaces(ctx, req.(*ListWorkspacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspaces_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.tcc.v1.Workspaces/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workspaces_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.tcc.v1.Workspaces",
	HandlerType: (*WorkspacesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspace",
			Handler:    _Workspaces_CreateWorkspace_Handler,
		},
		{
			MethodName: "GetWorkspace",
			Handler:    _Workspaces_GetWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _Workspaces_UpdateWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspaces",
			Handler:    _Workspaces_ListWorkspaces_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _Workspaces_DeleteWorkspace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace.proto",
}

func init() { proto.RegisterFile("workspace.proto", fileDescriptor_workspace_8c32a6b65efea193) }

var fileDescriptor_workspace_8c32a6b65efea193 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x99, 0x46, 0x0b, 0x79, 0x2d, 0x06, 0x46, 0x5c, 0xc2, 0xfa, 0x87, 0x65, 0x40, 0x93,
	0x2e, 0x38, 0x63, 0xea, 0x4d, 0xf0, 0xa0, 0x56, 0xbc, 0x78, 0x0a, 0x48, 0x41, 0x4f, 0x93, 0xcd,
	0x6b, 0xba, 0xd8, 0xcc, 0x8c, 0x3b, 0x93, 0x88, 0x94, 0x5e, 0x44, 0xaf, 0x7a, 0xf0, 0x83, 0x79,
	0xf0, 0x2b, 0xf8, 0x41, 0xa4, 0xb3, 0xd9, 0x6e, 0xb3, 0x0e, 0x75, 0x21, 0xd8, 0x53, 0xe6, 0x4f,
	0xde, 0x79, 0x7e, 0xfb, 0xbc, 0x0f, 0x2f, 0xf4, 0x3e, 0xea, 0xe2, 0xbd, 0x35, 0x32, 0x43, 0x6e,
	0x0a, 0xed, 0x34, 0xa5, 0x0e, 0x5d, 0x21, 0x1d, 0x72, 0x69, 0x72, 0xee, 0xb2, 0x8c, 0x2f, 0x47,
	0xf1, 0xed, 0x99, 0xd6, 0xb3, 0x23, 0x14, 0xd2, 0xe4, 0x42, 0x2a, 0xa5, 0x9d, 0x74, 0xb9, 0x56,
	0xb6, 0xac, 0x88, 0x6f, 0xad, 0x6e, 0xfd, 0x6e, 0xb2, 0x78, 0x27, 0x70, 0x6e, 0xdc, 0xa7, 0xf2,
	0x92, 0xbd, 0x85, 0xee, 0x41, 0xa5, 0x40, 0x29, 0x5c, 0x51, 0x72, 0x8e, 0x7d, 0x92, 0x90, 0x61,
	0x77, 0xec, 0xd7, 0x34, 0x82, 0x6d, 0x87, 0x4a, 0x2a, 0xd7, 0xdf, 0xf2, 0xa7, 0xab, 0x1d, 0x65,
	0xb0, 0x33, 0x45, 0x83, 0x6a, 0x8a, 0x2a, 0xcb, 0xd1, 0xf6, 0x3b, 0x49, 0x67, 0xd8, 0x1d, 0xaf,
	0x9d, 0xb1, 0x43, 0x88, 0x9e, 0x17, 0x28, 0x1d, 0x9e, 0x49, 0x8c, 0xf1, 0xc3, 0x02, 0xad, 0xfb,
	0x1f, 0x4a, 0xaf, 0xcd, 0xf4, 0x32, 0x94, 0x9e, 0xc2, 0x8d, 0x97, 0xe8, 0x36, 0x91, 0x61, 0x02,
	0x6e, 0xbe, 0xca, 0x6d, 0xfd, 0x86, 0xad, 0x1e, 0xa9, 0x0b, 0xc8, 0x5a, 0xc1, 0x01, 0x44, 0xcd,
	0x02, 0x6b, 0xb4, 0xb2, 0x48, 0x9f, 0x00, 0x9c, 0x05, 0xc4, 0xf6, 0x49, 0xd2, 0x19, 0x5e, 0xdb,
	0xbb, 0xc3, 0xff, 0x8e, 0x08, 0xaf, 0x81, 0xcf, 0x15, 0xb0, 0x7d, 0x88, 0xf6, 0xf1, 0x08, 0x37,
	0xb3, 0x6d, 0xef, 0xe7, 0x55, 0x80, 0x9a, 0x8d, 0x7e, 0x23, 0xd0, 0x6b, 0xb4, 0x9d, 0xa6, 0x21,
	0xa6, 0x70, 0x36, 0xe2, 0x8b, 0xf9, 0xd9, 0xc3, 0xcf, 0xbf, 0x7e, 0xff, 0xd8, 0x4a, 0xd9, 0x3d,
	0xb1, 0x1c, 0x89, 0xfa, 0x5b, 0x44, 0x09, 0x23, 0x8e, 0xcb, 0xdf, 0x13, 0x71, 0x7c, 0xca, 0x7c,
	0xf2, 0x98, 0xa4, 0xf4, 0x2b, 0x81, 0x9d, 0xf3, 0x3d, 0xa3, 0x83, 0x90, 0x42, 0xa0, 0xab, 0xff,
	0x42, 0x79, 0xe0, 0x51, 0x06, 0xb4, 0x1d, 0x8a, 0x37, 0xa6, 0x91, 0xd2, 0xb0, 0x31, 0xe1, 0x28,
	0xb7, 0x34, 0x26, 0x6e, 0x6f, 0xcc, 0x77, 0x02, 0xd7, 0xd7, 0x83, 0x45, 0x77, 0x43, 0x1a, 0xc1,
	0xb4, 0xc6, 0x69, 0x9b, 0xbf, 0x96, 0x39, 0x65, 0xf7, 0x3d, 0x5b, 0x42, 0xef, 0x5e, 0xcc, 0x46,
	0xbf, 0x10, 0xe8, 0x35, 0x12, 0x19, 0xb6, 0x28, 0x1c, 0xdb, 0x38, 0xe2, 0xe5, 0xb0, 0xe3, 0xd5,
	0xb0, 0xe3, 0x2f, 0x4e, 0x87, 0x5d, 0xd5, 0xa9, 0xb4, 0x9d, 0x37, 0xcf, 0x76, 0xdf, 0x0c, 0x66,
	0xb9, 0x3b, 0x5c, 0x4c, 0x78, 0xa6, 0xe7, 0x62, 0x25, 0x9f, 0xeb, 0x6a, 0xe5, 0x07, 0xad, 0xcb,
	0x32, 0xb1, 0x1c, 0x4d, 0xb6, 0xbd, 0xd2, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0xed,
	0xa0, 0x99, 0xa9, 0x05, 0x00, 0x00,
}
