// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/v1/user.proto

// `User` represents a user that has been loaded from a configured
// Identity Provider (IdP) that can log into the paltform.
// Currently, users are automatically synchronized by TSB from a
// configured LDAP server.
//
// The following example creates a user named `john` under the tenant
// `mycompany` .
//
// ```yaml
// apiVersion: api.tsb.tetrate.io/v1
// kind: User
// metadata:
//   name: john
//   tenant: mycompany
// spec:
//   username: john
//   firstName: John
//   lastName: Doe
//   displayName: John Doe
//   email: john.doe@acme.com
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

// `User` represents a user from the identity Provider that is allowed to log into
// the platform.
type User struct {
	// The username used in the login credentials.
	LoginName string `protobuf:"bytes,1,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	// The first name of the user.
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// The last name of the user, if any.
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// A human friendly display name for the user.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Email for the user where alerts and other notifications will be sent.
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_37ac9e5e3510b07c, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "tetrateio.api.tsb.v1.User")
}

func init() { proto.RegisterFile("tsb/v1/user.proto", fileDescriptor_37ac9e5e3510b07c) }

var fileDescriptor_37ac9e5e3510b07c = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0x31, 0x4e, 0xeb, 0x40,
	0x10, 0x06, 0x60, 0xad, 0x9f, 0xf3, 0x84, 0x07, 0x0a, 0xb0, 0x22, 0x11, 0x05, 0xe1, 0x00, 0x15,
	0x14, 0x78, 0x15, 0x71, 0x02, 0x7c, 0x00, 0x8a, 0x48, 0x34, 0x34, 0xd1, 0x18, 0x4f, 0x9c, 0x91,
	0xd6, 0x59, 0x6b, 0x77, 0x63, 0x89, 0x0b, 0x71, 0x16, 0x4a, 0x8e, 0x10, 0xf9, 0x18, 0x54, 0x28,
	0xe3, 0x84, 0x6e, 0xf7, 0xff, 0xfe, 0x29, 0x66, 0xe0, 0x22, 0xf8, 0x52, 0x77, 0x73, 0xbd, 0xf5,
	0xe4, 0xf2, 0xd6, 0xd9, 0x60, 0xd3, 0x71, 0xa0, 0xe0, 0x30, 0x10, 0xdb, 0x1c, 0x5b, 0xce, 0x83,
	0x2f, 0xf3, 0x6e, 0x3e, 0x9d, 0xd5, 0xd6, 0xd6, 0x86, 0x34, 0xb6, 0xac, 0x57, 0x4c, 0xa6, 0x5a,
	0x96, 0xb4, 0xc6, 0x8e, 0xed, 0x61, 0x6c, 0x7a, 0xd9, 0xa1, 0xe1, 0x0a, 0x03, 0xe9, 0xe3, 0x63,
	0x80, 0xbb, 0x4f, 0x05, 0xf1, 0xab, 0x27, 0x97, 0x3e, 0x00, 0x18, 0x5b, 0xf3, 0x66, 0xb9, 0xc1,
	0x86, 0x26, 0xea, 0x46, 0xdd, 0x27, 0x05, 0xec, 0x9e, 0xa3, 0x9f, 0x22, 0x76, 0xd1, 0xb9, 0x5a,
	0x24, 0xa2, 0x2f, 0xd8, 0x50, 0x7a, 0x0d, 0xb0, 0x62, 0xe7, 0xc3, 0x50, 0x8d, 0xf6, 0xd5, 0x45,
	0x22, 0x89, 0xf0, 0x15, 0x24, 0x06, 0x8f, 0xfa, 0x4f, 0xf4, 0x64, 0x1f, 0x08, 0xde, 0xc2, 0x59,
	0xc5, 0xbe, 0x35, 0xf8, 0x31, 0x78, 0x2c, 0x7e, 0x7a, 0xc8, 0xa4, 0x32, 0x86, 0x11, 0x35, 0xc8,
	0x66, 0x32, 0x12, 0x1b, 0x3e, 0xc5, 0xe3, 0x57, 0x9f, 0xa9, 0xef, 0x3e, 0x53, 0xbb, 0x3e, 0x53,
	0x6f, 0xb3, 0x9a, 0xc3, 0x7a, 0x5b, 0xe6, 0xef, 0xb6, 0xd1, 0x7f, 0x17, 0x91, 0xf5, 0x87, 0x93,
	0x95, 0xff, 0x65, 0xbd, 0xa7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xe4, 0x25, 0x5b, 0x43,
	0x01, 0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LoginName) > 0 {
		i -= len(m.LoginName)
		copy(dAtA[i:], m.LoginName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.LoginName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LoginName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoginName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
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
				next, err := skipUser(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthUser
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
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)
