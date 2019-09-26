// Code generated by protoc-gen-go. DO NOT EDIT.
// source: onboarding.proto

package invv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference to a cloud account to be used in an onboarding request
type CloudRequest struct {
	// Identifier of the cloud account to connect to
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudRequest) Reset()         { *m = CloudRequest{} }
func (m *CloudRequest) String() string { return proto.CompactTextString(m) }
func (*CloudRequest) ProtoMessage()    {}
func (*CloudRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{0}
}

func (m *CloudRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudRequest.Unmarshal(m, b)
}
func (m *CloudRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudRequest.Marshal(b, m, deterministic)
}
func (m *CloudRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudRequest.Merge(m, src)
}
func (m *CloudRequest) XXX_Size() int {
	return xxx_messageInfo_CloudRequest.Size(m)
}
func (m *CloudRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloudRequest proto.InternalMessageInfo

func (m *CloudRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// Requests all existing users in the given cloud account
type ListCloudUsersRequest struct {
	Cloud                *CloudRequest `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListCloudUsersRequest) Reset()         { *m = ListCloudUsersRequest{} }
func (m *ListCloudUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListCloudUsersRequest) ProtoMessage()    {}
func (*ListCloudUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{1}
}

func (m *ListCloudUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudUsersRequest.Unmarshal(m, b)
}
func (m *ListCloudUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListCloudUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudUsersRequest.Merge(m, src)
}
func (m *ListCloudUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListCloudUsersRequest.Size(m)
}
func (m *ListCloudUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudUsersRequest proto.InternalMessageInfo

func (m *ListCloudUsersRequest) GetCloud() *CloudRequest {
	if m != nil {
		return m.Cloud
	}
	return nil
}

// Requests all existing groups in the given cloud account
type ListCloudGroupsRequest struct {
	Cloud                *CloudRequest `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListCloudGroupsRequest) Reset()         { *m = ListCloudGroupsRequest{} }
func (m *ListCloudGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCloudGroupsRequest) ProtoMessage()    {}
func (*ListCloudGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{2}
}

func (m *ListCloudGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudGroupsRequest.Unmarshal(m, b)
}
func (m *ListCloudGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudGroupsRequest.Marshal(b, m, deterministic)
}
func (m *ListCloudGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudGroupsRequest.Merge(m, src)
}
func (m *ListCloudGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCloudGroupsRequest.Size(m)
}
func (m *ListCloudGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudGroupsRequest proto.InternalMessageInfo

func (m *ListCloudGroupsRequest) GetCloud() *CloudRequest {
	if m != nil {
		return m.Cloud
	}
	return nil
}

// Requests all existing roles in the given cloud account
type ListCloudRolesRequest struct {
	Cloud                *CloudRequest `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListCloudRolesRequest) Reset()         { *m = ListCloudRolesRequest{} }
func (m *ListCloudRolesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCloudRolesRequest) ProtoMessage()    {}
func (*ListCloudRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{3}
}

func (m *ListCloudRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudRolesRequest.Unmarshal(m, b)
}
func (m *ListCloudRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudRolesRequest.Marshal(b, m, deterministic)
}
func (m *ListCloudRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudRolesRequest.Merge(m, src)
}
func (m *ListCloudRolesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCloudRolesRequest.Size(m)
}
func (m *ListCloudRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudRolesRequest proto.InternalMessageInfo

func (m *ListCloudRolesRequest) GetCloud() *CloudRequest {
	if m != nil {
		return m.Cloud
	}
	return nil
}

// Requests all existing roles in the given cloud account
type ListCloudRoleBindingsRequest struct {
	Cloud *CloudRequest `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	// Identifier of the resource to get the bindings for
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCloudRoleBindingsRequest) Reset()         { *m = ListCloudRoleBindingsRequest{} }
func (m *ListCloudRoleBindingsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCloudRoleBindingsRequest) ProtoMessage()    {}
func (*ListCloudRoleBindingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{4}
}

func (m *ListCloudRoleBindingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudRoleBindingsRequest.Unmarshal(m, b)
}
func (m *ListCloudRoleBindingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudRoleBindingsRequest.Marshal(b, m, deterministic)
}
func (m *ListCloudRoleBindingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudRoleBindingsRequest.Merge(m, src)
}
func (m *ListCloudRoleBindingsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCloudRoleBindingsRequest.Size(m)
}
func (m *ListCloudRoleBindingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudRoleBindingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudRoleBindingsRequest proto.InternalMessageInfo

func (m *ListCloudRoleBindingsRequest) GetCloud() *CloudRequest {
	if m != nil {
		return m.Cloud
	}
	return nil
}

func (m *ListCloudRoleBindingsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Imports the given groups (and their child groups) from the given cloud account
type ImportGroupsRequest struct {
	Cloud                *CloudRequest `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ImportGroupsRequest) Reset()         { *m = ImportGroupsRequest{} }
func (m *ImportGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ImportGroupsRequest) ProtoMessage()    {}
func (*ImportGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{5}
}

func (m *ImportGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportGroupsRequest.Unmarshal(m, b)
}
func (m *ImportGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportGroupsRequest.Marshal(b, m, deterministic)
}
func (m *ImportGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportGroupsRequest.Merge(m, src)
}
func (m *ImportGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ImportGroupsRequest.Size(m)
}
func (m *ImportGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportGroupsRequest proto.InternalMessageInfo

func (m *ImportGroupsRequest) GetCloud() *CloudRequest {
	if m != nil {
		return m.Cloud
	}
	return nil
}

type ImportGroupsResponse struct {
	// The imported groups
	Groups               []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportGroupsResponse) Reset()         { *m = ImportGroupsResponse{} }
func (m *ImportGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ImportGroupsResponse) ProtoMessage()    {}
func (*ImportGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{6}
}

func (m *ImportGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportGroupsResponse.Unmarshal(m, b)
}
func (m *ImportGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportGroupsResponse.Marshal(b, m, deterministic)
}
func (m *ImportGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportGroupsResponse.Merge(m, src)
}
func (m *ImportGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ImportGroupsResponse.Size(m)
}
func (m *ImportGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportGroupsResponse proto.InternalMessageInfo

func (m *ImportGroupsResponse) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

// Imports the given users from the given cloud account
type ImportUsersRequest struct {
	Cloud                *CloudRequest `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Names                []string      `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ImportUsersRequest) Reset()         { *m = ImportUsersRequest{} }
func (m *ImportUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ImportUsersRequest) ProtoMessage()    {}
func (*ImportUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{7}
}

func (m *ImportUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportUsersRequest.Unmarshal(m, b)
}
func (m *ImportUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportUsersRequest.Marshal(b, m, deterministic)
}
func (m *ImportUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportUsersRequest.Merge(m, src)
}
func (m *ImportUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ImportUsersRequest.Size(m)
}
func (m *ImportUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportUsersRequest proto.InternalMessageInfo

func (m *ImportUsersRequest) GetCloud() *CloudRequest {
	if m != nil {
		return m.Cloud
	}
	return nil
}

func (m *ImportUsersRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type ImportUsersResponse struct {
	// The imported users
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportUsersResponse) Reset()         { *m = ImportUsersResponse{} }
func (m *ImportUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ImportUsersResponse) ProtoMessage()    {}
func (*ImportUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe1fceb30c27431, []int{8}
}

func (m *ImportUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportUsersResponse.Unmarshal(m, b)
}
func (m *ImportUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportUsersResponse.Marshal(b, m, deterministic)
}
func (m *ImportUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportUsersResponse.Merge(m, src)
}
func (m *ImportUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ImportUsersResponse.Size(m)
}
func (m *ImportUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportUsersResponse proto.InternalMessageInfo

func (m *ImportUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*CloudRequest)(nil), "tetrate.api.inventory.v1.CloudRequest")
	proto.RegisterType((*ListCloudUsersRequest)(nil), "tetrate.api.inventory.v1.ListCloudUsersRequest")
	proto.RegisterType((*ListCloudGroupsRequest)(nil), "tetrate.api.inventory.v1.ListCloudGroupsRequest")
	proto.RegisterType((*ListCloudRolesRequest)(nil), "tetrate.api.inventory.v1.ListCloudRolesRequest")
	proto.RegisterType((*ListCloudRoleBindingsRequest)(nil), "tetrate.api.inventory.v1.ListCloudRoleBindingsRequest")
	proto.RegisterType((*ImportGroupsRequest)(nil), "tetrate.api.inventory.v1.ImportGroupsRequest")
	proto.RegisterType((*ImportGroupsResponse)(nil), "tetrate.api.inventory.v1.ImportGroupsResponse")
	proto.RegisterType((*ImportUsersRequest)(nil), "tetrate.api.inventory.v1.ImportUsersRequest")
	proto.RegisterType((*ImportUsersResponse)(nil), "tetrate.api.inventory.v1.ImportUsersResponse")
}

func init() { proto.RegisterFile("onboarding.proto", fileDescriptor_afe1fceb30c27431) }

var fileDescriptor_afe1fceb30c27431 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x49, 0xeb, 0xb6, 0xf4, 0x6c, 0xf1, 0xcf, 0x58, 0x25, 0x84, 0xa2, 0xcb, 0xa0, 0x75,
	0xd5, 0xdd, 0x99, 0xee, 0x56, 0x2b, 0xa8, 0x57, 0xf5, 0x42, 0x44, 0xa1, 0x90, 0x52, 0x2f, 0xbc,
	0xcb, 0x6e, 0x87, 0x74, 0x60, 0x77, 0x26, 0x26, 0x93, 0x80, 0x94, 0xde, 0xf8, 0x0a, 0x5e, 0xea,
	0x85, 0xef, 0xe1, 0x63, 0xf8, 0x0a, 0xde, 0xfb, 0x0a, 0x92, 0xc9, 0x64, 0x9b, 0x2c, 0x4d, 0x13,
	0x21, 0x77, 0x93, 0xe1, 0x7c, 0xe7, 0xfc, 0xf8, 0x72, 0xbe, 0x04, 0x6e, 0x4a, 0x31, 0x91, 0x5e,
	0x78, 0xc2, 0x85, 0x4f, 0x82, 0x50, 0x2a, 0x89, 0x6c, 0xc5, 0x54, 0xe8, 0x29, 0x46, 0xbc, 0x80,
	0x13, 0x2e, 0x12, 0x26, 0x94, 0x0c, 0xbf, 0x90, 0x64, 0xe4, 0x6c, 0xfb, 0x52, 0xfa, 0x33, 0x46,
	0xbd, 0x80, 0x53, 0x4f, 0x08, 0xa9, 0x3c, 0xc5, 0xa5, 0x88, 0x32, 0x9d, 0x03, 0x8a, 0x79, 0xf3,
	0xec, 0x8c, 0xfb, 0xb0, 0xf9, 0x66, 0x26, 0xe3, 0x13, 0x97, 0x7d, 0x8e, 0x59, 0xa4, 0x90, 0x0d,
	0xeb, 0xde, 0x74, 0x2a, 0x63, 0xa1, 0x6c, 0xab, 0x67, 0xf5, 0x37, 0xdc, 0xfc, 0x11, 0x1f, 0xc3,
	0x9d, 0x0f, 0x3c, 0x52, 0xba, 0xfa, 0x38, 0x62, 0x61, 0x94, 0x4b, 0x5e, 0x43, 0x67, 0x9a, 0x5e,
	0x6a, 0x41, 0x77, 0xbc, 0x43, 0xaa, 0xb0, 0x48, 0x71, 0x92, 0x9b, 0x89, 0xf0, 0x47, 0xb8, 0xbb,
	0x68, 0xfb, 0x36, 0x94, 0x71, 0xd0, 0x52, 0xdf, 0x22, 0xae, 0x2b, 0x67, 0xac, 0xa5, 0xb6, 0x01,
	0x6c, 0x97, 0xda, 0x1e, 0x70, 0x91, 0xbe, 0x91, 0x76, 0xba, 0x23, 0x04, 0xd7, 0x84, 0x37, 0x67,
	0xf6, 0x8a, 0xb6, 0x5e, 0x9f, 0xf1, 0x11, 0xdc, 0x7e, 0x37, 0x0f, 0x64, 0xa8, 0xda, 0x74, 0xe7,
	0x10, 0xb6, 0xca, 0x4d, 0xa3, 0x40, 0x8a, 0x88, 0xa1, 0x17, 0xb0, 0xe6, 0xeb, 0x1b, 0xdb, 0xea,
	0xad, 0xf6, 0xbb, 0xe3, 0xfb, 0xd5, 0x6d, 0xb5, 0xd2, 0x35, 0xe5, 0xf8, 0x14, 0x50, 0xd6, 0xb0,
	0xbd, 0xd5, 0x40, 0x5b, 0xd0, 0x49, 0x1d, 0x88, 0xec, 0x95, 0xde, 0x6a, 0x7f, 0xc3, 0xcd, 0x1e,
	0xf0, 0xfb, 0xdc, 0x0f, 0x33, 0xc9, 0x90, 0x3f, 0x83, 0x4e, 0x9c, 0x5e, 0x18, 0xf0, 0x7b, 0xd5,
	0xa3, 0x52, 0x9d, 0x9b, 0x15, 0x8f, 0xff, 0xae, 0xc3, 0xad, 0xc3, 0x45, 0xae, 0x8e, 0x58, 0x98,
	0xf0, 0x29, 0x43, 0x3f, 0x2d, 0xb8, 0xb1, 0xb4, 0x94, 0x68, 0xb7, 0xba, 0xe1, 0xe5, 0xfb, 0xeb,
	0x0c, 0xae, 0x56, 0x94, 0x9d, 0xc7, 0xc3, 0xaf, 0xbf, 0xff, 0x7c, 0x5b, 0x79, 0x84, 0x1e, 0xd2,
	0x64, 0x44, 0x2f, 0xa2, 0x4e, 0xcf, 0xb4, 0x17, 0xc4, 0xa4, 0xf0, 0x9c, 0x66, 0x7e, 0xa3, 0x1f,
	0x16, 0x5c, 0x2f, 0xc7, 0x11, 0xd1, 0x06, 0x84, 0xc5, 0xb7, 0xe3, 0x3c, 0xbd, 0x5a, 0x50, 0xf2,
	0x17, 0x0f, 0x34, 0xdf, 0x0e, 0x7a, 0x50, 0xc3, 0xa7, 0x7d, 0x2d, 0xe3, 0xe9, 0xf8, 0x35, 0xc2,
	0x2b, 0x06, 0xb5, 0x0e, 0xcf, 0xd4, 0xfe, 0x27, 0x5e, 0xa8, 0x59, 0x7e, 0x59, 0x4b, 0x5f, 0x87,
	0x3c, 0xc6, 0x68, 0xbf, 0x21, 0xe5, 0x52, 0xee, 0x9d, 0x71, 0x3d, 0xec, 0x85, 0xc4, 0x30, 0xef,
	0x6b, 0xe6, 0x5d, 0x44, 0x6a, 0x98, 0x27, 0x46, 0x48, 0xcf, 0xd2, 0x04, 0x9c, 0xa7, 0xe6, 0x6e,
	0x16, 0xd3, 0x8b, 0x86, 0xd5, 0xc3, 0x2f, 0xf9, 0x74, 0x38, 0xa4, 0x69, 0x79, 0x79, 0x35, 0x71,
	0xc3, 0xd5, 0xfc, 0x6e, 0x41, 0xb7, 0x90, 0x50, 0x34, 0xa8, 0x1b, 0x57, 0x5a, 0xca, 0x61, 0xc3,
	0x6a, 0xc3, 0x46, 0x35, 0xdb, 0x63, 0xdc, 0x68, 0x2d, 0x5f, 0x5a, 0x4f, 0x0e, 0x9e, 0x7f, 0xda,
	0xf3, 0xb9, 0x3a, 0x8d, 0x27, 0x64, 0x2a, 0xe7, 0xd4, 0xcc, 0xe2, 0x32, 0x3f, 0xe9, 0x5f, 0xe6,
	0x62, 0x2a, 0x4d, 0x46, 0xaf, 0xb8, 0x48, 0x92, 0xd1, 0x64, 0x4d, 0xff, 0x2e, 0xf7, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0xc4, 0x38, 0x60, 0x86, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnboardingServiceClient is the client API for OnboardingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnboardingServiceClient interface {
	// Gets the groups in the given cloud account
	ListCloudGroups(ctx context.Context, in *ListCloudGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// Gets the users in the given cloud account
	ListCloudUsers(ctx context.Context, in *ListCloudUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// Gets the roles in the given cloud account
	ListCloudRoles(ctx context.Context, in *ListCloudRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	// Returns all bindings for the given resource in the given cloud account
	ListCloudRoleBindings(ctx context.Context, in *ListCloudRoleBindingsRequest, opts ...grpc.CallOption) (*ListRoleBindingsResponse, error)
	// Imports the given groups to the inventory service. Groups are imported recursively
	// so for every group passed as an argument, all its child group tree will be imported
	// as well.
	// Note this method will only import groups, not the users.
	ImportGroups(ctx context.Context, in *ImportGroupsRequest, opts ...grpc.CallOption) (*ImportGroupsResponse, error)
	// Imports the given users to the inventory service, and assigns them to the existing groups.
	ImportUsers(ctx context.Context, in *ImportUsersRequest, opts ...grpc.CallOption) (*ImportUsersResponse, error)
}

type onboardingServiceClient struct {
	cc *grpc.ClientConn
}

func NewOnboardingServiceClient(cc *grpc.ClientConn) OnboardingServiceClient {
	return &onboardingServiceClient{cc}
}

func (c *onboardingServiceClient) ListCloudGroups(ctx context.Context, in *ListCloudGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.OnboardingService/ListCloudGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingServiceClient) ListCloudUsers(ctx context.Context, in *ListCloudUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.OnboardingService/ListCloudUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingServiceClient) ListCloudRoles(ctx context.Context, in *ListCloudRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.OnboardingService/ListCloudRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingServiceClient) ListCloudRoleBindings(ctx context.Context, in *ListCloudRoleBindingsRequest, opts ...grpc.CallOption) (*ListRoleBindingsResponse, error) {
	out := new(ListRoleBindingsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.OnboardingService/ListCloudRoleBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingServiceClient) ImportGroups(ctx context.Context, in *ImportGroupsRequest, opts ...grpc.CallOption) (*ImportGroupsResponse, error) {
	out := new(ImportGroupsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.OnboardingService/ImportGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingServiceClient) ImportUsers(ctx context.Context, in *ImportUsersRequest, opts ...grpc.CallOption) (*ImportUsersResponse, error) {
	out := new(ImportUsersResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.OnboardingService/ImportUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnboardingServiceServer is the server API for OnboardingService service.
type OnboardingServiceServer interface {
	// Gets the groups in the given cloud account
	ListCloudGroups(context.Context, *ListCloudGroupsRequest) (*ListGroupsResponse, error)
	// Gets the users in the given cloud account
	ListCloudUsers(context.Context, *ListCloudUsersRequest) (*ListUsersResponse, error)
	// Gets the roles in the given cloud account
	ListCloudRoles(context.Context, *ListCloudRolesRequest) (*ListRolesResponse, error)
	// Returns all bindings for the given resource in the given cloud account
	ListCloudRoleBindings(context.Context, *ListCloudRoleBindingsRequest) (*ListRoleBindingsResponse, error)
	// Imports the given groups to the inventory service. Groups are imported recursively
	// so for every group passed as an argument, all its child group tree will be imported
	// as well.
	// Note this method will only import groups, not the users.
	ImportGroups(context.Context, *ImportGroupsRequest) (*ImportGroupsResponse, error)
	// Imports the given users to the inventory service, and assigns them to the existing groups.
	ImportUsers(context.Context, *ImportUsersRequest) (*ImportUsersResponse, error)
}

func RegisterOnboardingServiceServer(s *grpc.Server, srv OnboardingServiceServer) {
	s.RegisterService(&_OnboardingService_serviceDesc, srv)
}

func _OnboardingService_ListCloudGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).ListCloudGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.OnboardingService/ListCloudGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).ListCloudGroups(ctx, req.(*ListCloudGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingService_ListCloudUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).ListCloudUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.OnboardingService/ListCloudUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).ListCloudUsers(ctx, req.(*ListCloudUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingService_ListCloudRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).ListCloudRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.OnboardingService/ListCloudRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).ListCloudRoles(ctx, req.(*ListCloudRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingService_ListCloudRoleBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudRoleBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).ListCloudRoleBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.OnboardingService/ListCloudRoleBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).ListCloudRoleBindings(ctx, req.(*ListCloudRoleBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingService_ImportGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).ImportGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.OnboardingService/ImportGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).ImportGroups(ctx, req.(*ImportGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardingService_ImportUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).ImportUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.OnboardingService/ImportUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).ImportUsers(ctx, req.(*ImportUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnboardingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.inventory.v1.OnboardingService",
	HandlerType: (*OnboardingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCloudGroups",
			Handler:    _OnboardingService_ListCloudGroups_Handler,
		},
		{
			MethodName: "ListCloudUsers",
			Handler:    _OnboardingService_ListCloudUsers_Handler,
		},
		{
			MethodName: "ListCloudRoles",
			Handler:    _OnboardingService_ListCloudRoles_Handler,
		},
		{
			MethodName: "ListCloudRoleBindings",
			Handler:    _OnboardingService_ListCloudRoleBindings_Handler,
		},
		{
			MethodName: "ImportGroups",
			Handler:    _OnboardingService_ImportGroups_Handler,
		},
		{
			MethodName: "ImportUsers",
			Handler:    _OnboardingService_ImportUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onboarding.proto",
}
