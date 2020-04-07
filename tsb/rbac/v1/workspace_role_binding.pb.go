// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/rbac/v1/workspace_role_binding.proto

// `WorkspaceRoleBinding` is an assignment of roles to a set of users or
// teams to access resources under a Workspace.  The user or team
// information is obtained from an LDAP server that should have been
// configured as part of Service Bridge installation. Note that a
// `WorkspaceRoleBinding` can be created or modified only by users who
// have `set_rbac` permission on the Workspace.
//
// The following example assigns the `workspace-admin` role to users
// `alice`, `bob`, and members of the `t1` team for all workspace `w1`
// owned by the tenant `mycompany`.
//
// ```yaml
// apiVersion: rbac.tsb.tetrate.io/v1
// kind: WorkspaceRoleBinding
// metadata:
//   name: workspace-g1
//   tenant: mycompany
//   workspace: w1
// spec:
//   role: workspace-admin
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// `WorkspaceRoleBinding` assigns permissions to users of a workspace.
type WorkspaceRoleBinding struct {
	// The role being assinged to the subjects.
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// The set of users who are being assigned the role.
	Subjects             []*Subject `protobuf:"bytes,2,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WorkspaceRoleBinding) Reset()         { *m = WorkspaceRoleBinding{} }
func (m *WorkspaceRoleBinding) String() string { return proto.CompactTextString(m) }
func (*WorkspaceRoleBinding) ProtoMessage()    {}
func (*WorkspaceRoleBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_5940e6db42ff5bbf, []int{0}
}
func (m *WorkspaceRoleBinding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkspaceRoleBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkspaceRoleBinding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkspaceRoleBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceRoleBinding.Merge(m, src)
}
func (m *WorkspaceRoleBinding) XXX_Size() int {
	return m.Size()
}
func (m *WorkspaceRoleBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceRoleBinding.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceRoleBinding proto.InternalMessageInfo

func (m *WorkspaceRoleBinding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *WorkspaceRoleBinding) GetSubjects() []*Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkspaceRoleBinding)(nil), "tetrateio.api.tsb.rbac.v1.WorkspaceRoleBinding")
}

func init() {
	proto.RegisterFile("tsb/rbac/v1/workspace_role_binding.proto", fileDescriptor_5940e6db42ff5bbf)
}

var fileDescriptor_5940e6db42ff5bbf = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x4f, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0x96, 0x43, 0x41, 0xc5, 0x5d, 0x50, 0x85, 0xa0, 0x74, 0x08, 0x55, 0xc5, 0x90, 0xc9, 0x56,
	0xcb, 0x09, 0xf0, 0xca, 0x16, 0x06, 0x24, 0x96, 0xc8, 0x4e, 0x1e, 0xa9, 0xc1, 0xf4, 0x45, 0xf6,
	0x6b, 0x38, 0x00, 0x37, 0xe0, 0x54, 0x8c, 0x1c, 0xa1, 0xca, 0x31, 0x98, 0x50, 0x12, 0x5a, 0x75,
	0x61, 0x7b, 0xfa, 0xde, 0xf7, 0xcb, 0x13, 0x0a, 0x46, 0x7a, 0xa3, 0x73, 0x59, 0x2f, 0xe4, 0x3b,
	0xfa, 0xd7, 0x50, 0xe9, 0x1c, 0x32, 0x8f, 0x0e, 0x32, 0x63, 0xd7, 0x85, 0x5d, 0x97, 0xa2, 0xf2,
	0x48, 0x38, 0xbe, 0x22, 0x20, 0xaf, 0x09, 0x2c, 0x0a, 0x5d, 0x59, 0x41, 0xc1, 0x88, 0x56, 0x27,
	0xea, 0xc5, 0xf4, 0xba, 0x44, 0x2c, 0x1d, 0x48, 0x5d, 0x59, 0xf9, 0x6c, 0xc1, 0x15, 0x99, 0x81,
	0x95, 0xae, 0x2d, 0xfa, 0x5e, 0x3b, 0xbd, 0xac, 0xb5, 0xb3, 0x85, 0x26, 0x90, 0xbb, 0xe3, 0xef,
	0x71, 0x71, 0x18, 0xdf, 0x86, 0xf6, 0xf8, 0xfc, 0x83, 0xf1, 0xf3, 0xc7, 0x5d, 0x9b, 0x14, 0x1d,
	0xa8, 0xbe, 0xcb, 0x38, 0xe6, 0x83, 0x96, 0x36, 0x61, 0x33, 0x96, 0x9c, 0x2a, 0xbe, 0xbd, 0x8b,
	0x7e, 0xd4, 0xc0, 0x47, 0x67, 0x2c, 0xed, 0xf0, 0xf1, 0x3d, 0x1f, 0x86, 0x8d, 0x79, 0x81, 0x9c,
	0xc2, 0x24, 0x9a, 0x1d, 0x25, 0xa3, 0xe5, 0x5c, 0xfc, 0x5b, 0x5c, 0x3c, 0xf4, 0x54, 0x35, 0xea,
	0x7c, 0x8e, 0x3f, 0x59, 0x34, 0x64, 0xe9, 0xde, 0x40, 0x2d, 0xbf, 0x9a, 0x98, 0x7d, 0x37, 0x31,
	0xdb, 0x36, 0x31, 0x7b, 0xba, 0x29, 0x2d, 0xad, 0x36, 0x46, 0xe4, 0xf8, 0x26, 0xf7, 0x96, 0xdd,
	0xe6, 0x83, 0x11, 0xe6, 0xa4, 0x1b, 0x70, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x59, 0x90,
	0xf4, 0x59, 0x01, 0x00, 0x00,
}

func (m *WorkspaceRoleBinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkspaceRoleBinding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkspaceRoleBinding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintWorkspaceRoleBinding(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintWorkspaceRoleBinding(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkspaceRoleBinding(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkspaceRoleBinding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkspaceRoleBinding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovWorkspaceRoleBinding(uint64(l))
	}
	if len(m.Subjects) > 0 {
		for _, e := range m.Subjects {
			l = e.Size()
			n += 1 + l + sovWorkspaceRoleBinding(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkspaceRoleBinding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkspaceRoleBinding(x uint64) (n int) {
	return sovWorkspaceRoleBinding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkspaceRoleBinding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkspaceRoleBinding
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
			return fmt.Errorf("proto: WorkspaceRoleBinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkspaceRoleBinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspaceRoleBinding
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
				return ErrInvalidLengthWorkspaceRoleBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspaceRoleBinding
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
					return ErrIntOverflowWorkspaceRoleBinding
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
				return ErrInvalidLengthWorkspaceRoleBinding
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspaceRoleBinding
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
			skippy, err := skipWorkspaceRoleBinding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkspaceRoleBinding
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkspaceRoleBinding
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
func skipWorkspaceRoleBinding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkspaceRoleBinding
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
					return 0, ErrIntOverflowWorkspaceRoleBinding
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
					return 0, ErrIntOverflowWorkspaceRoleBinding
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
				return 0, ErrInvalidLengthWorkspaceRoleBinding
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWorkspaceRoleBinding
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkspaceRoleBinding
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
				next, err := skipWorkspaceRoleBinding(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWorkspaceRoleBinding
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
	ErrInvalidLengthWorkspaceRoleBinding = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkspaceRoleBinding   = fmt.Errorf("proto: integer overflow")
)
