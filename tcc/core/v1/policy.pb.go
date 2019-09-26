// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/core/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Policy
//
// A policy defines the set of subjects that can access a resource and under which conditions
// that access is granted.
type Policy struct {
	// The list of bindings configures the different access profiles that are allowed on the resource
	// configured by the policy.
	Bindings             []*Binding `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a7089d2b9fb12102, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetBindings() []*Binding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

// Binding
//
// A binding associates a role with a set of subjects.
// Bindings are used to configure policies, where different roles can be assigned to different sets of
// subjects to configure a fine-grained access control to the resource protected by the policy.
type Binding struct {
	// The role that defines the permissions that will be granted to the target resource
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// The set of subjects that will be allowed to access the target resource with the permissions
	// defined by the role
	Subjects             []string `protobuf:"bytes,2,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Binding) Reset()         { *m = Binding{} }
func (m *Binding) String() string { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()    {}
func (*Binding) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a7089d2b9fb12102, []int{1}
}
func (m *Binding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Binding.Unmarshal(m, b)
}
func (m *Binding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Binding.Marshal(b, m, deterministic)
}
func (dst *Binding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binding.Merge(dst, src)
}
func (m *Binding) XXX_Size() int {
	return xxx_messageInfo_Binding.Size(m)
}
func (m *Binding) XXX_DiscardUnknown() {
	xxx_messageInfo_Binding.DiscardUnknown(m)
}

var xxx_messageInfo_Binding proto.InternalMessageInfo

func (m *Binding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Binding) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "tetrate.api.tcc.core.v1.Policy")
	proto.RegisterType((*Binding)(nil), "tetrate.api.tcc.core.v1.Binding")
}

func init() { proto.RegisterFile("policy.proto", fileDescriptor_policy_a7089d2b9fb12102) }

var fileDescriptor_policy_a7089d2b9fb12102 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0xcf, 0xc9,
	0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2f, 0x49, 0x2d, 0x29, 0x4a, 0x2c,
	0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x49, 0x4e, 0xd6, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x2b,
	0x33, 0x94, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x3a,
	0x94, 0x82, 0xb8, 0xd8, 0x02, 0xc0, 0x26, 0x08, 0x79, 0x70, 0x71, 0x24, 0x65, 0xe6, 0xa5, 0x64,
	0xe6, 0xa5, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x29, 0xe8, 0xe1, 0x30, 0x4e, 0xcf,
	0x09, 0xa2, 0xd0, 0x89, 0x6b, 0xd7, 0xcb, 0x03, 0xcc, 0xac, 0x93, 0x18, 0x99, 0x38, 0x18, 0x83,
	0xe0, 0xba, 0x95, 0xc2, 0xb8, 0xd8, 0xa1, 0x0a, 0x84, 0x64, 0xb9, 0x58, 0x8a, 0xf2, 0x73, 0x52,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x38, 0x41, 0xca, 0x59, 0x8a, 0x98, 0x04, 0x18, 0x83,
	0xc0, 0xc2, 0x42, 0x3a, 0x5c, 0x1c, 0xc5, 0xa5, 0x49, 0x59, 0xa9, 0xc9, 0x25, 0xc5, 0x12, 0x4c,
	0x0a, 0xcc, 0x1a, 0x9c, 0x4e, 0x02, 0x20, 0x25, 0xdc, 0x93, 0x18, 0x39, 0x38, 0x18, 0x95, 0x20,
	0x2a, 0xe1, 0x2a, 0x9c, 0xf4, 0xa2, 0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xa1, 0x6e, 0xcb, 0xcc, 0x87, 0xb1, 0xf4, 0x13, 0x0b, 0x32, 0xf5, 0x4b, 0x92, 0x93,
	0xf5, 0x41, 0xae, 0xd4, 0x2f, 0x33, 0x4c, 0x62, 0x03, 0x7b, 0xd1, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0xe6, 0x0a, 0x0a, 0x24, 0x01, 0x00, 0x00,
}
