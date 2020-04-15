// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/rbac/v1/role.proto

// `Role` is a named collection of permissions that can be assigned to
// any user or team in the system. The set of actions that can be
// performed by a user, such as the ability to create, delete, or
// update configuration will depend on the permissions associated with
// the user's role. Roles are tenant level resources that are defined
// once. `RoleBindings` in each configuration group will bind a user
// to a specific role defined apriori.
//
// The following example declares a `workspace-admin` role with the
// ability to create, delete configurations and the ability to set
// RBAC policies on the groups within the workspace.
//
// ```yaml
// apiVersion: rbac.tsb.tetrate.io/v1
// kind: Role
// metadata:
//   name: role1
//   tenant: mycompany
// spec:
//   rules:
//   - types:
//     - apiGroup: api.tsb.tetrate.io/v1
//       kinds:
//       - workspaceSetting
//     permissions:
//     - create
//     - read
//     - delete
//     - write
//     - set_rbac
//   - types:
//     - apiGroup: traffic.tsb.tetrate.io/v1
//     - apiGroup: security.tsb.tetrate.io/v1
//     - apiGroup: gateway.tsb.tetrate.io/v1
//     permissions:
//     - create
//     - delete
//     - read
//     - write
//     - set_rbac
//```
//

package v1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// A permission defines an action that can be performed on a
// resource. By default access to resources is denied unless an
// explicit permission grants access to perform an operation against
// it.
type Permission int32

const (
	// Default value to designate no value was explicitly set for the permission.
	Permission_INVALID Permission = 0
	// The read permission grants read-only access to the resource.
	Permission_READ Permission = 1
	// The write permission allows the subject to modify an existing resource.
	Permission_WRITE Permission = 2
	// The create permission allows subjects to create child resources on the resource.
	Permission_CREATE Permission = 3
	// The delete permission grants permissions to delete the resource.
	Permission_DELETE Permission = 4
	// The set-rbac permission allows subjects to manage the access policies for the resources.
	Permission_SET_RBAC Permission = 5
)

var Permission_name = map[int32]string{
	0: "INVALID",
	1: "READ",
	2: "WRITE",
	3: "CREATE",
	4: "DELETE",
	5: "SET_RBAC",
}

var Permission_value = map[string]int32{
	"INVALID":  0,
	"READ":     1,
	"WRITE":    2,
	"CREATE":   3,
	"DELETE":   4,
	"SET_RBAC": 5,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_016dff8c57a66e4c, []int{0}
}

// `Role` is a named collection of permissions that can be assigned to
// any user or team in the system.
type Role struct {
	// A set of rules that define the permissions associated with each API group.
	Rules                []*Role_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_016dff8c57a66e4c, []int{0}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Role.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return m.Size()
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetRules() []*Role_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// The type of API resource for which the role is being created.
type Role_ResourceType struct {
	// A specific API group such as traffic.tsb.tetrate.io/v1.
	ApiGroup string `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Specific kinds of APIs under the API group. If omitted, the
	// role will apply to all kinds under the group.
	Kinds                []string `protobuf:"bytes,2,rep,name=kinds,proto3" json:"kinds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role_ResourceType) Reset()         { *m = Role_ResourceType{} }
func (m *Role_ResourceType) String() string { return proto.CompactTextString(m) }
func (*Role_ResourceType) ProtoMessage()    {}
func (*Role_ResourceType) Descriptor() ([]byte, []int) {
	return fileDescriptor_016dff8c57a66e4c, []int{0, 0}
}
func (m *Role_ResourceType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Role_ResourceType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Role_ResourceType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Role_ResourceType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role_ResourceType.Merge(m, src)
}
func (m *Role_ResourceType) XXX_Size() int {
	return m.Size()
}
func (m *Role_ResourceType) XXX_DiscardUnknown() {
	xxx_messageInfo_Role_ResourceType.DiscardUnknown(m)
}

var xxx_messageInfo_Role_ResourceType proto.InternalMessageInfo

func (m *Role_ResourceType) GetApiGroup() string {
	if m != nil {
		return m.ApiGroup
	}
	return ""
}

func (m *Role_ResourceType) GetKinds() []string {
	if m != nil {
		return m.Kinds
	}
	return nil
}

// A rule defines the set of api groups
type Role_Rule struct {
	// The set of API groups and the api Kinds within the group on which this rule is applicable.
	Types []*Role_ResourceType `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	// The set of actions allowed for these APIs.
	Permissions          []Permission `protobuf:"varint,2,rep,packed,name=permissions,proto3,enum=tetrateio.api.tsb.rbac.v1.Permission" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Role_Rule) Reset()         { *m = Role_Rule{} }
func (m *Role_Rule) String() string { return proto.CompactTextString(m) }
func (*Role_Rule) ProtoMessage()    {}
func (*Role_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_016dff8c57a66e4c, []int{0, 1}
}
func (m *Role_Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Role_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Role_Rule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Role_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role_Rule.Merge(m, src)
}
func (m *Role_Rule) XXX_Size() int {
	return m.Size()
}
func (m *Role_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Role_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Role_Rule proto.InternalMessageInfo

func (m *Role_Rule) GetTypes() []*Role_ResourceType {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Role_Rule) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// Subject identifies a user or a team under a tenant. Roles are
// assigned to subjects for specific resources in the system.
type Subject struct {
	// Types that are valid to be assigned to S:
	//	*Subject_User
	//	*Subject_Team
	S                    isSubject_S `protobuf_oneof:"s"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Subject) Reset()         { *m = Subject{} }
func (m *Subject) String() string { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()    {}
func (*Subject) Descriptor() ([]byte, []int) {
	return fileDescriptor_016dff8c57a66e4c, []int{1}
}
func (m *Subject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subject.Merge(m, src)
}
func (m *Subject) XXX_Size() int {
	return m.Size()
}
func (m *Subject) XXX_DiscardUnknown() {
	xxx_messageInfo_Subject.DiscardUnknown(m)
}

var xxx_messageInfo_Subject proto.InternalMessageInfo

type isSubject_S interface {
	isSubject_S()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Subject_User struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}
type Subject_Team struct {
	Team string `protobuf:"bytes,2,opt,name=team,proto3,oneof"`
}

func (*Subject_User) isSubject_S() {}
func (*Subject_Team) isSubject_S() {}

func (m *Subject) GetS() isSubject_S {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *Subject) GetUser() string {
	if x, ok := m.GetS().(*Subject_User); ok {
		return x.User
	}
	return ""
}

func (m *Subject) GetTeam() string {
	if x, ok := m.GetS().(*Subject_Team); ok {
		return x.Team
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Subject) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Subject_User)(nil),
		(*Subject_Team)(nil),
	}
}

func init() {
	proto.RegisterEnum("tetrateio.api.tsb.rbac.v1.Permission", Permission_name, Permission_value)
	proto.RegisterType((*Role)(nil), "tetrateio.api.tsb.rbac.v1.Role")
	proto.RegisterType((*Role_ResourceType)(nil), "tetrateio.api.tsb.rbac.v1.Role.ResourceType")
	proto.RegisterType((*Role_Rule)(nil), "tetrateio.api.tsb.rbac.v1.Role.Rule")
	proto.RegisterType((*Subject)(nil), "tetrateio.api.tsb.rbac.v1.Subject")
}

func init() { proto.RegisterFile("tsb/rbac/v1/role.proto", fileDescriptor_016dff8c57a66e4c) }

var fileDescriptor_016dff8c57a66e4c = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xde, 0x49, 0x93, 0xdd, 0xf4, 0x74, 0x91, 0x30, 0x2c, 0x5a, 0x7b, 0x51, 0xcb, 0xb2, 0x62,
	0x11, 0x49, 0xd8, 0x7a, 0xe7, 0x5d, 0xb2, 0x0d, 0x5a, 0x58, 0x7f, 0x98, 0x0d, 0x0a, 0xde, 0x94,
	0x49, 0x7b, 0xec, 0x8e, 0xa6, 0x3b, 0xc3, 0xcc, 0xa4, 0xb0, 0xaf, 0xe1, 0x7b, 0xf8, 0x1e, 0x82,
	0x37, 0x3e, 0xc2, 0xd2, 0xc7, 0xe8, 0x95, 0x24, 0xa1, 0xb5, 0x5e, 0xa8, 0x77, 0xe7, 0xe7, 0x3b,
	0xdf, 0xf7, 0x1d, 0xf8, 0xe0, 0xbe, 0x35, 0x79, 0xa4, 0x73, 0x3e, 0x8b, 0x56, 0xe7, 0x91, 0x96,
	0x05, 0x86, 0x4a, 0x4b, 0x2b, 0xe9, 0x43, 0x8b, 0x56, 0x73, 0x8b, 0x42, 0x86, 0x5c, 0x89, 0xd0,
	0x9a, 0x3c, 0xac, 0x50, 0xe1, 0xea, 0xbc, 0xf7, 0x68, 0x21, 0xe5, 0xa2, 0xc0, 0x88, 0x2b, 0x11,
	0x7d, 0x12, 0x58, 0xcc, 0xa7, 0x39, 0x5e, 0xf3, 0x95, 0x90, 0xba, 0xb9, 0xed, 0x3d, 0x58, 0xf1,
	0x42, 0xcc, 0xb9, 0xc5, 0x68, 0x5b, 0x34, 0x8b, 0xd3, 0x1f, 0x0e, 0xb8, 0x4c, 0x16, 0x48, 0xc7,
	0xe0, 0xe9, 0xb2, 0x40, 0xd3, 0x25, 0x83, 0xd6, 0xb0, 0x33, 0x3a, 0x0b, 0xff, 0xaa, 0x16, 0x56,
	0xf8, 0x90, 0x95, 0x05, 0x26, 0xfe, 0x26, 0xf1, 0xbe, 0x12, 0xc7, 0x27, 0xac, 0x39, 0xee, 0xbd,
	0x86, 0x63, 0x86, 0x46, 0x96, 0x7a, 0x86, 0xd9, 0xad, 0x42, 0xfa, 0x04, 0xda, 0x5c, 0x89, 0xe9,
	0x42, 0xcb, 0x52, 0x75, 0xc9, 0x80, 0x0c, 0xdb, 0x09, 0xdc, 0xc5, 0xce, 0x26, 0x71, 0xb5, 0x13,
	0x10, 0xe6, 0x73, 0x25, 0x5e, 0x56, 0x3b, 0x7a, 0x02, 0xde, 0x17, 0x71, 0x33, 0x37, 0x5d, 0x67,
	0xd0, 0x1a, 0xb6, 0x59, 0xd3, 0xf4, 0xbe, 0x11, 0x70, 0x2b, 0x21, 0xfa, 0x16, 0x3c, 0x7b, 0xab,
	0x76, 0xee, 0x9e, 0xfd, 0xd7, 0xdd, 0x9e, 0x89, 0xa4, 0x53, 0x2b, 0x6e, 0x8d, 0xd6, 0x3c, 0x34,
	0x83, 0x8e, 0x42, 0xbd, 0x14, 0xc6, 0x08, 0x79, 0xd3, 0xa8, 0xde, 0x1b, 0x3d, 0xfe, 0x07, 0xed,
	0xbb, 0x1d, 0xfa, 0x4f, 0xbe, 0x7d, 0x9a, 0xd3, 0x17, 0x70, 0x74, 0x55, 0xe6, 0x9f, 0x71, 0x66,
	0xe9, 0x09, 0xb8, 0xa5, 0x41, 0xdd, 0x3c, 0xfd, 0xea, 0x80, 0xd5, 0x5d, 0x35, 0xb5, 0xc8, 0x97,
	0x5d, 0x67, 0x3b, 0xad, 0xba, 0xa4, 0x05, 0xc4, 0x3c, 0xcd, 0x00, 0x7e, 0x6b, 0xd0, 0x0e, 0x1c,
	0x4d, 0xde, 0xbc, 0x8f, 0x2f, 0x27, 0xe3, 0xe0, 0x80, 0xfa, 0xe0, 0xb2, 0x34, 0x1e, 0x07, 0x84,
	0xb6, 0xc1, 0xfb, 0xc0, 0x26, 0x59, 0x1a, 0x38, 0x14, 0xe0, 0xf0, 0x82, 0xa5, 0x71, 0x96, 0x06,
	0xad, 0xaa, 0x1e, 0xa7, 0x97, 0x69, 0x96, 0x06, 0x2e, 0x3d, 0x06, 0xff, 0x2a, 0xcd, 0xa6, 0x2c,
	0x89, 0x2f, 0x02, 0x2f, 0x19, 0x7d, 0x5f, 0xf7, 0xc9, 0xcf, 0x75, 0x9f, 0xdc, 0xad, 0xfb, 0xe4,
	0xe3, 0xd9, 0x42, 0xd8, 0xeb, 0x32, 0x0f, 0x67, 0x72, 0x19, 0xed, 0x5e, 0xad, 0x53, 0xb3, 0x97,
	0xb9, 0xfc, 0xb0, 0x8e, 0xc6, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xeb, 0xbf, 0x8e,
	0x89, 0x02, 0x00, 0x00,
}

func (m *Role) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Role) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Role) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRole(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Role_ResourceType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Role_ResourceType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Role_ResourceType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Kinds) > 0 {
		for iNdEx := len(m.Kinds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Kinds[iNdEx])
			copy(dAtA[i:], m.Kinds[iNdEx])
			i = encodeVarintRole(dAtA, i, uint64(len(m.Kinds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ApiGroup) > 0 {
		i -= len(m.ApiGroup)
		copy(dAtA[i:], m.ApiGroup)
		i = encodeVarintRole(dAtA, i, uint64(len(m.ApiGroup)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Role_Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Role_Rule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Role_Rule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Permissions) > 0 {
		dAtA2 := make([]byte, len(m.Permissions)*10)
		var j1 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintRole(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Types) > 0 {
		for iNdEx := len(m.Types) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Types[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRole(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Subject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.S != nil {
		{
			size := m.S.Size()
			i -= size
			if _, err := m.S.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Subject_User) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *Subject_User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.User)
	copy(dAtA[i:], m.User)
	i = encodeVarintRole(dAtA, i, uint64(len(m.User)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *Subject_Team) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *Subject_Team) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Team)
	copy(dAtA[i:], m.Team)
	i = encodeVarintRole(dAtA, i, uint64(len(m.Team)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintRole(dAtA []byte, offset int, v uint64) int {
	offset -= sovRole(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Role) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovRole(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Role_ResourceType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ApiGroup)
	if l > 0 {
		n += 1 + l + sovRole(uint64(l))
	}
	if len(m.Kinds) > 0 {
		for _, s := range m.Kinds {
			l = len(s)
			n += 1 + l + sovRole(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Role_Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Types) > 0 {
		for _, e := range m.Types {
			l = e.Size()
			n += 1 + l + sovRole(uint64(l))
		}
	}
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovRole(uint64(e))
		}
		n += 1 + sovRole(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Subject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.S != nil {
		n += m.S.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Subject_User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	n += 1 + l + sovRole(uint64(l))
	return n
}
func (m *Subject_Team) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Team)
	n += 1 + l + sovRole(uint64(l))
	return n
}

func sovRole(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRole(x uint64) (n int) {
	return sovRole(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Role) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Role: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Role: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &Role_Rule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Role_ResourceType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiGroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiGroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kinds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kinds = append(m.Kinds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Role_Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Types = append(m.Types, &Role_ResourceType{})
			if err := m.Types[len(m.Types)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Permission
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRole
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Permission(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRole
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRole
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthRole
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]Permission, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Permission
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRole
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Permission(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Subject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Subject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = &Subject_User{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Team", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = &Subject_Team{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRole(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRole
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRole
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRole
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRole
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRole
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRole
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRole(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRole
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRole = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRole   = fmt.Errorf("proto: integer overflow")
)