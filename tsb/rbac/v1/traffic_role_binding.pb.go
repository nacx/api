// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/rbac/v1/traffic_role_binding.proto

// `TrafficRoleBinding` is an assignment of roles to a set of users or
// teams to access resources under a Traffic group.  The user or team
// information is obtained from an LDAP server that should have been
// configured as part of Service Bridge installation. Note that a
// `TrafficRoleBinding` can be created or modified only by users who
// have `set_rbac` permission on the Traffic group.
//
// The following example assigns the `traffic-admin` role to users
// `alice`, `bob`, and members of the `traffic-ops` team for
// traffic group `g1` under workspace `w1` owned by the tenant
// `mycompany`.
//
// ```yaml
// apiVersion: rbac.tsb.tetrate.io/v1
// kind: TrafficRoleBinding
// metadata:
//   name: traffic-g1
//   tenant: mycompany
//   workspace: w1
//   group: g1
// spec:
//   role: traffic-admin
//   subjects:
//   - user: alice
//   - user: bob
//   - team: traffic-ops
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

// `TrafficRoleBinding` assigns permissions to users of traffic groups.
type TrafficRoleBinding struct {
	// The role being assinged to the subjects.
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// The set of users who are being assigned the role.
	Subjects             []*Subject `protobuf:"bytes,2,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TrafficRoleBinding) Reset()         { *m = TrafficRoleBinding{} }
func (m *TrafficRoleBinding) String() string { return proto.CompactTextString(m) }
func (*TrafficRoleBinding) ProtoMessage()    {}
func (*TrafficRoleBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_955456607c92aa7f, []int{0}
}
func (m *TrafficRoleBinding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficRoleBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficRoleBinding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficRoleBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficRoleBinding.Merge(m, src)
}
func (m *TrafficRoleBinding) XXX_Size() int {
	return m.Size()
}
func (m *TrafficRoleBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficRoleBinding.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficRoleBinding proto.InternalMessageInfo

func (m *TrafficRoleBinding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TrafficRoleBinding) GetSubjects() []*Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func init() {
	proto.RegisterType((*TrafficRoleBinding)(nil), "tetrateio.api.tsb.rbac.v1.TrafficRoleBinding")
}

func init() {
	proto.RegisterFile("tsb/rbac/v1/traffic_role_binding.proto", fileDescriptor_955456607c92aa7f)
}

var fileDescriptor_955456607c92aa7f = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xe5, 0x50, 0x50, 0x71, 0x17, 0x94, 0x01, 0x4a, 0x87, 0x50, 0x55, 0x08, 0x75, 0xb2,
	0xd5, 0x72, 0x02, 0xbc, 0xb2, 0x05, 0x26, 0x96, 0xc8, 0x4e, 0x9c, 0xf4, 0x21, 0xd3, 0x17, 0x39,
	0xaf, 0xd9, 0xb9, 0x02, 0xa7, 0x62, 0xe4, 0x08, 0x55, 0x8e, 0xc1, 0x84, 0x12, 0x43, 0xd5, 0x85,
	0xcd, 0x7a, 0xfe, 0xfe, 0xfc, 0x3e, 0x7e, 0x47, 0x8d, 0x91, 0xde, 0xe8, 0x5c, 0xb6, 0x2b, 0x49,
	0x5e, 0x97, 0x25, 0xe4, 0x99, 0x47, 0x67, 0x33, 0x03, 0xdb, 0x02, 0xb6, 0x95, 0xa8, 0x3d, 0x12,
	0xc6, 0xd7, 0x64, 0xc9, 0x6b, 0xb2, 0x80, 0x42, 0xd7, 0x20, 0xa8, 0x31, 0xa2, 0x77, 0x89, 0x76,
	0x35, 0xbb, 0xa9, 0x10, 0x2b, 0x67, 0xa5, 0xae, 0x41, 0x96, 0x60, 0x5d, 0x91, 0x19, 0xbb, 0xd1,
	0x2d, 0xa0, 0x0f, 0xde, 0xd9, 0x55, 0xab, 0x1d, 0x14, 0x9a, 0xac, 0xfc, 0x7b, 0xfc, 0x7e, 0x5c,
	0x1e, 0x97, 0xf7, 0xa5, 0xe1, 0xbe, 0x78, 0x67, 0x3c, 0x7e, 0x0e, 0x2c, 0x29, 0x3a, 0xab, 0x02,
	0x49, 0x9c, 0xf0, 0x51, 0x2f, 0x9a, 0xb2, 0x39, 0x5b, 0x9e, 0x2b, 0xbe, 0x7f, 0x88, 0xbe, 0xd5,
	0xc8, 0x47, 0x17, 0x2c, 0x1d, 0xee, 0xf1, 0x23, 0x1f, 0x37, 0x3b, 0xf3, 0x6a, 0x73, 0x6a, 0xa6,
	0xd1, 0xfc, 0x64, 0x39, 0x59, 0x2f, 0xc4, 0xbf, 0xd8, 0xe2, 0x29, 0x48, 0xd5, 0x64, 0xc8, 0x39,
	0xfd, 0x60, 0xd1, 0x98, 0xa5, 0x87, 0x00, 0xb5, 0xfe, 0xec, 0x12, 0xf6, 0xd5, 0x25, 0x6c, 0xdf,
	0x25, 0xec, 0xe5, 0xb6, 0x02, 0xda, 0xec, 0x8c, 0xc8, 0xf1, 0x4d, 0x1e, 0x22, 0x87, 0xc5, 0x47,
	0x13, 0xcc, 0xd9, 0x80, 0x7f, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x47, 0x97, 0x8d, 0x92, 0x55,
	0x01, 0x00, 0x00,
}

func (m *TrafficRoleBinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficRoleBinding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrafficRoleBinding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintTrafficRoleBinding(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintTrafficRoleBinding(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrafficRoleBinding(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrafficRoleBinding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TrafficRoleBinding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovTrafficRoleBinding(uint64(l))
	}
	if len(m.Subjects) > 0 {
		for _, e := range m.Subjects {
			l = e.Size()
			n += 1 + l + sovTrafficRoleBinding(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTrafficRoleBinding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrafficRoleBinding(x uint64) (n int) {
	return sovTrafficRoleBinding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrafficRoleBinding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficRoleBinding
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
			return fmt.Errorf("proto: TrafficRoleBinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrafficRoleBinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficRoleBinding
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
				return ErrInvalidLengthTrafficRoleBinding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficRoleBinding
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
					return ErrIntOverflowTrafficRoleBinding
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
				return ErrInvalidLengthTrafficRoleBinding
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficRoleBinding
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
			skippy, err := skipTrafficRoleBinding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficRoleBinding
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficRoleBinding
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
func skipTrafficRoleBinding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrafficRoleBinding
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
					return 0, ErrIntOverflowTrafficRoleBinding
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
					return 0, ErrIntOverflowTrafficRoleBinding
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
				return 0, ErrInvalidLengthTrafficRoleBinding
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTrafficRoleBinding
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrafficRoleBinding
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
				next, err := skipTrafficRoleBinding(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTrafficRoleBinding
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
	ErrInvalidLengthTrafficRoleBinding = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrafficRoleBinding   = fmt.Errorf("proto: integer overflow")
)
