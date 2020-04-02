// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: q/rbac/v1/permissions.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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

// Permission
//
// A permission defines an action that can be performed on a resource. By default access to
// resources is denied unless an explicit permission grants access to perform an operation against
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
	// The set-policy permission allows subjects to manage the access policies for the resources.
	Permission_SET_POLICY Permission = 5
)

var Permission_name = map[int32]string{
	0: "INVALID",
	1: "READ",
	2: "WRITE",
	3: "CREATE",
	4: "DELETE",
	5: "SET_POLICY",
}

var Permission_value = map[string]int32{
	"INVALID":    0,
	"READ":       1,
	"WRITE":      2,
	"CREATE":     3,
	"DELETE":     4,
	"SET_POLICY": 5,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bdfc77e40898da1, []int{0}
}

// RequiredPermission
//
// Configures the sets of permissions that are required to invoke the method where this option is
// applied.
type RequiredPermission struct {
	// The required set of permissions. The full name of each permission (such as ReadApplication)
	// will be inferred from the name of the method where this option is applied.
	Permissions []Permission `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=tetrateio.api.q.rbac.v1.Permission" json:"permissions,omitempty"`
	// Set of raw permission names values. Only use this if the method being protected does not follow
	// the common naming convention and the proper name of the permission cannot be inferred just by
	// using the Permission enum and the method name.
	RawPermissions []string `protobuf:"bytes,2,rep,name=raw_permissions,json=rawPermissions,proto3" json:"raw_permissions,omitempty"`
	// When this flag is set to true, the permission checks will not be made at the API surface.
	// This is usually needed when there is not an explicit set of permissions that can be
	// preconfigured for the API methods, so the access control checks will be implemented at runtime
	// by the application.
	// The default value is 'false' and will only be taken into account if the permission properties
	// are empty. If any permission is set, this flag will be ignored.
	DeferPermissionCheckToApplication bool     `protobuf:"varint,3,opt,name=defer_permission_check_to_application,json=deferPermissionCheckToApplication,proto3" json:"defer_permission_check_to_application,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *RequiredPermission) Reset()         { *m = RequiredPermission{} }
func (m *RequiredPermission) String() string { return proto.CompactTextString(m) }
func (*RequiredPermission) ProtoMessage()    {}
func (*RequiredPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bdfc77e40898da1, []int{0}
}
func (m *RequiredPermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequiredPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequiredPermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequiredPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredPermission.Merge(m, src)
}
func (m *RequiredPermission) XXX_Size() int {
	return m.Size()
}
func (m *RequiredPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredPermission.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredPermission proto.InternalMessageInfo

func (m *RequiredPermission) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *RequiredPermission) GetRawPermissions() []string {
	if m != nil {
		return m.RawPermissions
	}
	return nil
}

func (m *RequiredPermission) GetDeferPermissionCheckToApplication() bool {
	if m != nil {
		return m.DeferPermissionCheckToApplication
	}
	return false
}

var E_DefaultRequires = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*RequiredPermission)(nil),
	Field:         2000,
	Name:          "tetrateio.api.q.rbac.v1.default_requires",
	Tag:           "bytes,2000,opt,name=default_requires",
	Filename:      "q/rbac/v1/permissions.proto",
}

var E_Requires = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*RequiredPermission)(nil),
	Field:         2000,
	Name:          "tetrateio.api.q.rbac.v1.requires",
	Tag:           "bytes,2000,opt,name=requires",
	Filename:      "q/rbac/v1/permissions.proto",
}

func init() {
	proto.RegisterEnum("tetrateio.api.q.rbac.v1.Permission", Permission_name, Permission_value)
	proto.RegisterType((*RequiredPermission)(nil), "tetrateio.api.q.rbac.v1.RequiredPermission")
	proto.RegisterExtension(E_DefaultRequires)
	proto.RegisterExtension(E_Requires)
}

func init() { proto.RegisterFile("q/rbac/v1/permissions.proto", fileDescriptor_8bdfc77e40898da1) }

var fileDescriptor_8bdfc77e40898da1 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0xf1, 0xba, 0x8d, 0xee, 0xab, 0xd4, 0x46, 0xbe, 0x30, 0x81, 0x54, 0xc2, 0x10, 0xa2,
	0x02, 0xc9, 0x66, 0xe3, 0xb6, 0x5b, 0x68, 0x7d, 0x88, 0x54, 0x58, 0xe4, 0x45, 0x43, 0x70, 0x89,
	0x9c, 0xc4, 0x6d, 0x2c, 0xba, 0xd9, 0x75, 0x9c, 0x8e, 0x9f, 0xc8, 0x71, 0x27, 0xce, 0xa8, 0xbf,
	0x04, 0x35, 0x1d, 0x69, 0xa4, 0x69, 0x27, 0x6e, 0x91, 0xf2, 0x7c, 0xef, 0xe3, 0xef, 0xb5, 0xe1,
	0xc5, 0x92, 0xda, 0x54, 0x64, 0x74, 0x75, 0x4a, 0x8d, 0xb4, 0xd7, 0xaa, 0x2c, 0x95, 0xbe, 0x29,
	0x89, 0xb1, 0xda, 0x69, 0xfc, 0xcc, 0x49, 0x67, 0x85, 0x93, 0x4a, 0x13, 0x61, 0x14, 0x59, 0x92,
	0x0d, 0x4a, 0x56, 0xa7, 0xcf, 0xfd, 0xb9, 0xd6, 0xf3, 0x85, 0xa4, 0x35, 0x96, 0x56, 0x33, 0x9a,
	0xcb, 0x32, 0xb3, 0xca, 0x38, 0x6d, 0xb7, 0xa3, 0x27, 0xbf, 0x11, 0x60, 0x2e, 0x97, 0x95, 0xb2,
	0x32, 0x8f, 0x9a, 0x60, 0xcc, 0xa0, 0xd7, 0xd2, 0x1c, 0x23, 0xbf, 0x33, 0xea, 0x9f, 0xbd, 0x26,
	0x8f, 0x78, 0xc8, 0x6e, 0x92, 0xb7, 0xe7, 0xf0, 0x5b, 0x18, 0x58, 0x71, 0x9b, 0xb4, 0xa3, 0xf6,
	0xfc, 0xce, 0xe8, 0x88, 0xf7, 0xad, 0xb8, 0x8d, 0x5a, 0x60, 0x04, 0x6f, 0x72, 0x39, 0x93, 0xb6,
	0x85, 0x26, 0x59, 0x21, 0xb3, 0x1f, 0x89, 0xd3, 0x89, 0x30, 0x66, 0xa1, 0x32, 0xe1, 0x94, 0xbe,
	0x39, 0xee, 0xf8, 0x68, 0xd4, 0xe5, 0xaf, 0x6a, 0x78, 0x17, 0x30, 0xde, 0xa0, 0xb1, 0x0e, 0x76,
	0xe0, 0xbb, 0x2b, 0x80, 0xd6, 0x3e, 0x3d, 0x78, 0x1a, 0x7e, 0xb9, 0x0a, 0xa6, 0xe1, 0xc4, 0x7b,
	0x82, 0xbb, 0xb0, 0xcf, 0x59, 0x30, 0xf1, 0x10, 0x3e, 0x82, 0x83, 0xaf, 0x3c, 0x8c, 0x99, 0xb7,
	0x87, 0x01, 0x0e, 0xc7, 0x9c, 0x05, 0x31, 0xf3, 0x3a, 0x9b, 0xef, 0x09, 0x9b, 0xb2, 0x98, 0x79,
	0xfb, 0xb8, 0x0f, 0x70, 0xc9, 0xe2, 0x24, 0xba, 0x98, 0x86, 0xe3, 0x6f, 0xde, 0xc1, 0xf9, 0x4f,
	0xf0, 0x72, 0x39, 0x13, 0xd5, 0xc2, 0x25, 0x76, 0xdb, 0x5b, 0x89, 0x5f, 0x92, 0x6d, 0xcf, 0xe4,
	0x5f, 0xcf, 0xe4, 0x52, 0xda, 0x95, 0xca, 0xe4, 0x85, 0x71, 0xf5, 0xd2, 0x77, 0x03, 0x1f, 0x8d,
	0x7a, 0x67, 0xef, 0x1f, 0x2d, 0xf0, 0xe1, 0x15, 0xf0, 0xc1, 0xbd, 0xe6, 0xfe, 0x57, 0x79, 0x5e,
	0x40, 0xb7, 0x31, 0x0e, 0x1f, 0x18, 0x3f, 0x4b, 0x57, 0xe8, 0xfc, 0x7f, 0x84, 0x4d, 0xfa, 0xa7,
	0x0f, 0xbf, 0xd6, 0x43, 0x74, 0xb7, 0x1e, 0xa2, 0x3f, 0xeb, 0x21, 0xfa, 0x7e, 0x32, 0x57, 0xae,
	0xa8, 0x52, 0x92, 0xe9, 0x6b, 0xda, 0xc4, 0x51, 0x61, 0x14, 0x6d, 0xde, 0x64, 0x7a, 0x58, 0x9f,
	0xe3, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x55, 0xdb, 0x94, 0xa7, 0x02, 0x00, 0x00,
}

func (m *RequiredPermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequiredPermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequiredPermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DeferPermissionCheckToApplication {
		i--
		if m.DeferPermissionCheckToApplication {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.RawPermissions) > 0 {
		for iNdEx := len(m.RawPermissions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RawPermissions[iNdEx])
			copy(dAtA[i:], m.RawPermissions[iNdEx])
			i = encodeVarintPermissions(dAtA, i, uint64(len(m.RawPermissions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
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
		i = encodeVarintPermissions(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPermissions(dAtA []byte, offset int, v uint64) int {
	offset -= sovPermissions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequiredPermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovPermissions(uint64(e))
		}
		n += 1 + sovPermissions(uint64(l)) + l
	}
	if len(m.RawPermissions) > 0 {
		for _, s := range m.RawPermissions {
			l = len(s)
			n += 1 + l + sovPermissions(uint64(l))
		}
	}
	if m.DeferPermissionCheckToApplication {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPermissions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPermissions(x uint64) (n int) {
	return sovPermissions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequiredPermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: RequiredPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequiredPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v Permission
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPermissions
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
						return ErrIntOverflowPermissions
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
					return ErrInvalidLengthPermissions
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPermissions
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
							return ErrIntOverflowPermissions
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPermissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPermissions = append(m.RawPermissions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeferPermissionCheckToApplication", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DeferPermissionCheckToApplication = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissions
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func skipPermissions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermissions
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
					return 0, ErrIntOverflowPermissions
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
					return 0, ErrIntOverflowPermissions
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
				return 0, ErrInvalidLengthPermissions
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPermissions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPermissions
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
				next, err := skipPermissions(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPermissions
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
	ErrInvalidLengthPermissions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermissions   = fmt.Errorf("proto: integer overflow")
)
