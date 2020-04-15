// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/v1/workspace.proto

// A Workspace carves a chunk of the cluster resources owned by a
// tenant into an isolated configuration domain.
//
// The following example claims `ns1` and `ns2` namespaces across all
// clusters owned by the tenant `mycompany`.
//
// ```yaml
// apiVersion: api.tsb.tetrate.io/v1
// kind: Workspace
// metadata:
//   name: w1
//   tenant: mycompany
// spec:
//   namespacesSelector:
//     names:
//     - */ns1
//     - */ns2
//```
//
// The following example claims `ns1` namespace only from the `c1`
// cluster and claims all namespaces from the `c2` cluster.
//
// ```yaml
// apiVersion: api.tsb.tetrate.io/v1
// kind: Workspace
// metadata:
//   name: w1
//   tenant: mycompany
// spec:
//   namespaceSelector:
//     names:
//     - c1/ns1
//     - c2/*
//```
//

package v1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/tetrateio/api/tsb/types/v1"
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

// A Kubernetes cluster managing both pods and VMs.
type Workspace struct {
	// Set of namespaces owned exclusively by this workspace. A
	// workspace can own all namespaces of a cluster or a set of
	// namespaces across any cluster or a set of namespaces in a
	// specific cluster. Use `*/*` to claim all cluster resources under
	// the tenant.
	NamespaceSelector *v1.NamespaceSelector `protobuf:"bytes,1,opt,name=namespace_selector,json=namespaceSelector,proto3" json:"namespace_selector,omitempty"`
	// If set to true, allows Gateways in the workspace to route to
	// services in other workspaces. Set this to true for workspaces
	// owning cluster-wide gateways shared by multiple teams.
	Privileged           *types.BoolValue `protobuf:"bytes,2,opt,name=privileged,proto3" json:"privileged,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fe8d5dd3750e975, []int{0}
}
func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(m, src)
}
func (m *Workspace) XXX_Size() int {
	return m.Size()
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetNamespaceSelector() *v1.NamespaceSelector {
	if m != nil {
		return m.NamespaceSelector
	}
	return nil
}

func (m *Workspace) GetPrivileged() *types.BoolValue {
	if m != nil {
		return m.Privileged
	}
	return nil
}

func init() {
	proto.RegisterType((*Workspace)(nil), "tetrateio.api.tsb.v1.Workspace")
}

func init() { proto.RegisterFile("tsb/v1/workspace.proto", fileDescriptor_7fe8d5dd3750e975) }

var fileDescriptor_7fe8d5dd3750e975 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0x41, 0xc1, 0xee, 0x45, 0x8b, 0x68, 0xe9, 0xa1, 0x2b, 0x9e, 0xbc, 0x6c, 0x42,
	0xf5, 0xe6, 0xcd, 0x3e, 0x80, 0x87, 0x15, 0x14, 0xbc, 0x2c, 0x49, 0x3b, 0xdb, 0x0d, 0x66, 0x77,
	0x42, 0x92, 0x66, 0xf1, 0x15, 0x7c, 0x1c, 0x9f, 0xc2, 0xa3, 0x8f, 0xb0, 0xf4, 0x31, 0x3c, 0xc9,
	0xa6, 0xad, 0x88, 0xde, 0x66, 0xf8, 0xe7, 0x9f, 0xf9, 0xe6, 0x8f, 0xcf, 0x9c, 0x15, 0xcc, 0x17,
	0x6c, 0x8b, 0xe6, 0xc5, 0x6a, 0x5e, 0x01, 0xd5, 0x06, 0x1d, 0x26, 0xa7, 0x0e, 0x9c, 0xe1, 0x0e,
	0x24, 0x52, 0xae, 0x25, 0x75, 0x56, 0x50, 0x5f, 0x64, 0xd3, 0x06, 0xb1, 0x51, 0xc0, 0xb8, 0x96,
	0x6c, 0x29, 0x41, 0xd5, 0x0b, 0x01, 0x2b, 0xee, 0x25, 0x9a, 0xde, 0x96, 0xe5, 0xc3, 0x40, 0xe8,
	0x44, 0xbb, 0x64, 0x5b, 0xc3, 0xb5, 0x06, 0x63, 0x07, 0xfd, 0xdc, 0x73, 0x25, 0x6b, 0xee, 0x80,
	0x8d, 0xc5, 0x20, 0xa4, 0x7b, 0x0e, 0xf7, 0xaa, 0xc1, 0xee, 0x69, 0x42, 0xd1, 0x2b, 0x97, 0xef,
	0x24, 0x3e, 0x7a, 0x1a, 0xe9, 0x12, 0x19, 0x27, 0x1b, 0xbe, 0x86, 0xd0, 0x2c, 0x2c, 0x28, 0xa8,
	0x1c, 0x9a, 0x94, 0x5c, 0x90, 0xab, 0xc9, 0xf5, 0x8c, 0xfe, 0x87, 0xee, 0x37, 0xf9, 0x82, 0xde,
	0x8f, 0xae, 0x87, 0xc1, 0x54, 0x4e, 0x76, 0x77, 0xd1, 0x57, 0x79, 0xf0, 0x46, 0xa2, 0x63, 0x32,
	0x3f, 0xd9, 0xfc, 0xd5, 0x93, 0xdb, 0x38, 0xd6, 0x46, 0x7a, 0xa9, 0xa0, 0x81, 0x3a, 0x8d, 0xc2,
	0x89, 0x8c, 0xf6, 0x0f, 0xd2, 0xf1, 0x41, 0x5a, 0x22, 0xaa, 0x47, 0xae, 0x5a, 0x98, 0xff, 0x9a,
	0x2e, 0x67, 0x1f, 0x5d, 0x4e, 0x3e, 0xbb, 0x9c, 0xec, 0xba, 0x9c, 0x3c, 0x4f, 0x1b, 0xe9, 0x56,
	0xad, 0xa0, 0x15, 0xae, 0xd9, 0x0f, 0x62, 0x08, 0xb1, 0x4f, 0x5f, 0x1c, 0x86, 0x75, 0x37, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xb9, 0xd9, 0x01, 0x8e, 0x01, 0x00, 0x00,
}

func (m *Workspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Workspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Workspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Privileged != nil {
		{
			size, err := m.Privileged.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkspace(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.NamespaceSelector != nil {
		{
			size, err := m.NamespaceSelector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkspace(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkspace(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkspace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Workspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NamespaceSelector != nil {
		l = m.NamespaceSelector.Size()
		n += 1 + l + sovWorkspace(uint64(l))
	}
	if m.Privileged != nil {
		l = m.Privileged.Size()
		n += 1 + l + sovWorkspace(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkspace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkspace(x uint64) (n int) {
	return sovWorkspace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Workspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkspace
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
			return fmt.Errorf("proto: Workspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Workspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NamespaceSelector == nil {
				m.NamespaceSelector = &v1.NamespaceSelector{}
			}
			if err := m.NamespaceSelector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileged", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Privileged == nil {
				m.Privileged = &types.BoolValue{}
			}
			if err := m.Privileged.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkspace
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkspace
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
func skipWorkspace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
				return 0, ErrInvalidLengthWorkspace
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWorkspace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkspace
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
				next, err := skipWorkspace(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWorkspace
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
	ErrInvalidLengthWorkspace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkspace   = fmt.Errorf("proto: integer overflow")
)