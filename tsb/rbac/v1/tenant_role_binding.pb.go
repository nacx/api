// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tenant_role_binding.proto

// `TenantRoleBinding` is an assignment of roles to a set of users or
// teams to access resources under a Tenant.  The user or team
// information is obtained from an LDAP server that should have been
// configured as part of Service Bridge installation. Note that a
// `TenantRoleBinding` can be created or modified only by users who
// have `set_rbac` permission on the Tenant.
//
// The following example assigns the `tenant-admin` role to users
// `alice`, `bob`, and members of the `t1` team owned by the tenant
// `mycompany`.
//
// ```yaml
// apiVersion: rbac.tsb.tetrate.io/v1
// kind: TenantRoleBinding
// metadata:
//   name: tenant-g1
//   tenant: mycompany
// spec:
//   role: tenant-admin
//   subjects:
//   - user: alice
//   - user: bob
//   - team: t1
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// `TenantRoleBinding` assigns permissions to users under a tenant.
type TenantRoleBinding struct {
	// The role being assinged to the subjects.
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// The set of users who are being assigned the role.
	Subjects             []*Subject `protobuf:"bytes,2,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TenantRoleBinding) Reset()         { *m = TenantRoleBinding{} }
func (m *TenantRoleBinding) String() string { return proto.CompactTextString(m) }
func (*TenantRoleBinding) ProtoMessage()    {}
func (*TenantRoleBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5e9e0251d3ea9fe, []int{0}
}
func (m *TenantRoleBinding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TenantRoleBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TenantRoleBinding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TenantRoleBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantRoleBinding.Merge(m, src)
}
func (m *TenantRoleBinding) XXX_Size() int {
	return m.Size()
}
func (m *TenantRoleBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantRoleBinding.DiscardUnknown(m)
}

var xxx_messageInfo_TenantRoleBinding proto.InternalMessageInfo

func (m *TenantRoleBinding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TenantRoleBinding) GetSubjects() []*Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func init() {
	proto.RegisterType((*TenantRoleBinding)(nil), "tetrate.api.tsb.rbac.v1.TenantRoleBinding")
}

func init() { proto.RegisterFile("tenant_role_binding.proto", fileDescriptor_f5e9e0251d3ea9fe) }

var fileDescriptor_f5e9e0251d3ea9fe = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x85, 0xb5, 0x4e, 0x40, 0x61, 0x01, 0x09, 0xdc, 0x24, 0xa4, 0x30, 0x16, 0x55, 0x0a, 0xb4,
	0xab, 0x40, 0x47, 0x87, 0x5b, 0x3a, 0x43, 0x45, 0x63, 0xed, 0xda, 0x83, 0xb3, 0x68, 0xf1, 0x58,
	0xeb, 0x89, 0x0f, 0xc0, 0x11, 0x38, 0x0e, 0x15, 0x25, 0x25, 0x47, 0x88, 0xdc, 0x71, 0x0b, 0xe4,
	0x3f, 0xba, 0x74, 0x4f, 0x7a, 0x6f, 0xbe, 0x37, 0x8f, 0x5f, 0x10, 0x14, 0xaa, 0xa0, 0xc4, 0xa1,
	0x85, 0x44, 0x9b, 0x22, 0x33, 0x45, 0x2e, 0x4a, 0x87, 0x84, 0xfe, 0x9c, 0x80, 0x9c, 0x22, 0x10,
	0xaa, 0x34, 0x82, 0x2a, 0x2d, 0x9c, 0x56, 0xa9, 0xa8, 0xd7, 0xcb, 0xcb, 0x1c, 0x31, 0xb7, 0x20,
	0x55, 0x69, 0xe4, 0x8b, 0x01, 0x9b, 0x25, 0x1a, 0x36, 0xaa, 0x36, 0xe8, 0xfa, 0xcb, 0xe5, 0xbc,
	0x56, 0xd6, 0x64, 0x8a, 0x40, 0x8e, 0x62, 0x30, 0x78, 0x5b, 0xd3, 0xeb, 0xab, 0x77, 0xc6, 0xcf,
	0x9f, 0xba, 0xf2, 0x18, 0x2d, 0x44, 0x7d, 0xb5, 0x1f, 0xf2, 0x69, 0x9b, 0x59, 0xb0, 0x90, 0xad,
	0x8e, 0xa2, 0x93, 0xdd, 0xbd, 0xf7, 0xf9, 0xfb, 0x35, 0x99, 0x3a, 0xef, 0x8c, 0xc5, 0x9d, 0xe3,
	0x3f, 0xf0, 0x59, 0xb5, 0xd5, 0xaf, 0x90, 0x52, 0xb5, 0xf0, 0xc2, 0xc9, 0xea, 0xf8, 0x26, 0x14,
	0x7b, 0x3e, 0x15, 0x8f, 0x7d, 0x30, 0x3a, 0x1d, 0x38, 0x07, 0x1f, 0xcc, 0x9b, 0xb1, 0xf8, 0x1f,
	0x10, 0xdd, 0x7d, 0x37, 0x01, 0xfb, 0x69, 0x02, 0xb6, 0x6b, 0x02, 0xf6, 0x7c, 0x9d, 0x1b, 0xda,
	0x6c, 0xb5, 0x48, 0xf1, 0x4d, 0x0e, 0x48, 0x83, 0xa3, 0xea, 0xe6, 0x52, 0xa5, 0x65, 0x0b, 0x97,
	0xf5, 0x5a, 0x1f, 0x76, 0x3b, 0x6e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x44, 0xf1, 0x31,
	0x43, 0x01, 0x00, 0x00,
}

func (m *TenantRoleBinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TenantRoleBinding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TenantRoleBinding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Subjects) > 0 {
		for iNdEx := len(m.Subjects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subjects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTenantRoleBinding(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintTenantRoleBinding(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTenantRoleBinding(dAtA []byte, offset int, v uint64) int {
	offset -= sovTenantRoleBinding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TenantRoleBinding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovTenantRoleBinding(uint64(l))
	}
	if len(m.Subjects) > 0 {
		for _, e := range m.Subjects {
			l = e.Size()
			n += 1 + l + sovTenantRoleBinding(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTenantRoleBinding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTenantRoleBinding(x uint64) (n int) {
	return sovTenantRoleBinding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TenantRoleBinding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTenantRoleBinding
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
			return fmt.Errorf("proto: TenantRoleBinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TenantRoleBinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTenantRoleBinding
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
				return ErrInvalidLengthTenantRoleBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTenantRoleBinding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subjects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTenantRoleBinding
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
				return ErrInvalidLengthTenantRoleBinding
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTenantRoleBinding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subjects = append(m.Subjects, &Subject{})
			if err := m.Subjects[len(m.Subjects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTenantRoleBinding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTenantRoleBinding
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTenantRoleBinding
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
func skipTenantRoleBinding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTenantRoleBinding
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
					return 0, ErrIntOverflowTenantRoleBinding
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
					return 0, ErrIntOverflowTenantRoleBinding
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
				return 0, ErrInvalidLengthTenantRoleBinding
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTenantRoleBinding
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTenantRoleBinding
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
				next, err := skipTenantRoleBinding(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTenantRoleBinding
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
	ErrInvalidLengthTenantRoleBinding = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTenantRoleBinding   = fmt.Errorf("proto: integer overflow")
)
