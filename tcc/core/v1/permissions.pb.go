// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permissions.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/core/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Permissions_Operation int32

const (
	Permissions_READ  Permissions_Operation = 0
	Permissions_WRITE Permissions_Operation = 1
)

var Permissions_Operation_name = map[int32]string{
	0: "READ",
	1: "WRITE",
}
var Permissions_Operation_value = map[string]int32{
	"READ":  0,
	"WRITE": 1,
}

func (x Permissions_Operation) String() string {
	return proto.EnumName(Permissions_Operation_name, int32(x))
}
func (Permissions_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_permissions_7449c8b98190d0c2, []int{0, 0}
}

type Permissions struct {
	// team : operation
	Values               map[string]Permissions_Operation `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=tetrate.api.tcc.core.v1.Permissions_Operation"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_permissions_7449c8b98190d0c2, []int{0}
}
func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (dst *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(dst, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetValues() map[string]Permissions_Operation {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Permissions)(nil), "tetrate.api.tcc.core.v1.Permissions")
	proto.RegisterMapType((map[string]Permissions_Operation)(nil), "tetrate.api.tcc.core.v1.Permissions.ValuesEntry")
	proto.RegisterEnum("tetrate.api.tcc.core.v1.Permissions_Operation", Permissions_Operation_name, Permissions_Operation_value)
}

func init() { proto.RegisterFile("permissions.proto", fileDescriptor_permissions_7449c8b98190d0c2) }

var fileDescriptor_permissions_7449c8b98190d0c2 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0x2d, 0xca,
	0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2f,
	0x49, 0x2d, 0x29, 0x4a, 0x2c, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x49, 0x4e, 0xd6, 0x4b,
	0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x33, 0x54, 0x7a, 0xc2, 0xc8, 0xc5, 0x1d, 0x80, 0x50, 0x2e, 0xe4,
	0xc1, 0xc5, 0x56, 0x96, 0x98, 0x53, 0x9a, 0x5a, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x64,
	0xa0, 0x87, 0x43, 0xa7, 0x1e, 0x92, 0x2e, 0xbd, 0x30, 0xb0, 0x16, 0xd7, 0xbc, 0x92, 0xa2, 0xca,
	0x20, 0xa8, 0x7e, 0xa9, 0x4c, 0x2e, 0x6e, 0x24, 0x61, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0xc8, 0x85, 0x8b, 0x15, 0xac, 0x54, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x8f, 0x28, 0x9b, 0xfc, 0x0b, 0x52, 0x8b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x82, 0x20, 0x9a, 0xad, 0x98, 0x2c, 0x18, 0x95, 0x14, 0xb8, 0x38, 0xe1, 0xe2, 0x42,
	0x1c, 0x5c, 0x2c, 0x41, 0xae, 0x8e, 0x2e, 0x02, 0x0c, 0x42, 0x9c, 0x5c, 0xac, 0xe1, 0x41, 0x9e,
	0x21, 0xae, 0x02, 0x8c, 0x4e, 0x7a, 0x51, 0x3a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x50, 0x8b, 0x32, 0xf3, 0x61, 0x2c, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0x92, 0xe4,
	0x64, 0x7d, 0x90, 0x95, 0xfa, 0x65, 0x86, 0x49, 0x6c, 0xe0, 0x60, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xbe, 0x75, 0xcb, 0x36, 0x4b, 0x01, 0x00, 0x00,
}
