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
//   namespacesSelectors:
//   - name: */ns1
//   - name: */ns2
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
//   namespaceSelectors:
//   - name: c1/ns1
//   - name: c2/*
//```
//

package v1

import (
	fmt "fmt"
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
	// specific cluster.
	NamespaceSelectors []*v1.NamespaceSelector `protobuf:"bytes,1,rep,name=namespace_selectors,json=namespaceSelectors,proto3" json:"namespace_selectors,omitempty"`
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

func (m *Workspace) GetNamespaceSelectors() []*v1.NamespaceSelector {
	if m != nil {
		return m.NamespaceSelectors
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
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x96, 0x5b, 0x09, 0x89, 0x74, 0x0b, 0x08, 0x45, 0x19, 0xd2, 0x8a, 0xa9, 0x4b, 0xcf, 0x4a,
	0xd9, 0xd8, 0xc8, 0x03, 0x30, 0x14, 0x09, 0x24, 0x96, 0xca, 0x4e, 0xaf, 0xa9, 0x85, 0xdb, 0xb3,
	0x6c, 0x27, 0x15, 0xef, 0xc4, 0x83, 0x30, 0xf2, 0x08, 0x55, 0x9e, 0x04, 0x35, 0x4e, 0x10, 0x82,
	0xcd, 0xe7, 0xbb, 0xef, 0x37, 0xba, 0xf1, 0x4e, 0xf2, 0x26, 0xe7, 0x47, 0xb2, 0x6f, 0xce, 0x88,
	0x12, 0xc1, 0x58, 0xf2, 0x14, 0x5f, 0x7b, 0xf4, 0x56, 0x78, 0x54, 0x04, 0xc2, 0x28, 0xf0, 0x4e,
	0x42, 0x93, 0xa7, 0xd3, 0x8a, 0xa8, 0xd2, 0xc8, 0x85, 0x51, 0x7c, 0xab, 0x50, 0x6f, 0xd6, 0x12,
	0x77, 0xa2, 0x51, 0x64, 0x03, 0x2c, 0xcd, 0xfa, 0x83, 0x6e, 0x92, 0xf5, 0x96, 0x1f, 0xad, 0x30,
	0x06, 0xad, 0xeb, 0xf7, 0xc9, 0x59, 0xce, 0xbf, 0x1b, 0x74, 0x67, 0xd1, 0xee, 0x11, 0x36, 0xb7,
	0x1f, 0x2c, 0xba, 0x7c, 0x19, 0x4c, 0xc4, 0x65, 0x74, 0x75, 0x10, 0x7b, 0xec, 0x86, 0xb5, 0x43,
	0x8d, 0xa5, 0x27, 0xeb, 0x12, 0x36, 0x1b, 0xcf, 0x27, 0xcb, 0x05, 0xfc, 0x37, 0x17, 0xa8, 0x9a,
	0x1c, 0x1e, 0x07, 0xd8, 0x53, 0x8f, 0x2a, 0xc6, 0xa7, 0x87, 0xd1, 0x2a, 0x3e, 0xfc, 0xfd, 0x77,
	0xf1, 0x7d, 0x14, 0x19, 0xab, 0x1a, 0xa5, 0xb1, 0xc2, 0x4d, 0x32, 0x9a, 0xb1, 0xf9, 0x64, 0x99,
	0x42, 0x48, 0x00, 0x43, 0x02, 0x28, 0x88, 0xf4, 0xb3, 0xd0, 0x35, 0xae, 0x7e, 0x5d, 0x17, 0x8b,
	0xcf, 0x36, 0x63, 0x5f, 0x6d, 0xc6, 0x4e, 0x6d, 0xc6, 0x5e, 0xa7, 0x95, 0xf2, 0xbb, 0x5a, 0x42,
	0x49, 0x7b, 0xfe, 0xe3, 0xad, 0x6b, 0x29, 0xd4, 0x2b, 0x2f, 0x3a, 0xba, 0xbb, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x20, 0x7a, 0x44, 0x8e, 0x6f, 0x01, 0x00, 0x00,
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
	if len(m.NamespaceSelectors) > 0 {
		for iNdEx := len(m.NamespaceSelectors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NamespaceSelectors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWorkspace(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.NamespaceSelectors) > 0 {
		for _, e := range m.NamespaceSelectors {
			l = e.Size()
			n += 1 + l + sovWorkspace(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceSelectors", wireType)
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
			m.NamespaceSelectors = append(m.NamespaceSelectors, &v1.NamespaceSelector{})
			if err := m.NamespaceSelectors[len(m.NamespaceSelectors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
