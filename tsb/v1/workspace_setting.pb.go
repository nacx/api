// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/v1/workspace_setting.proto

// Workspace Setting allows configuring the default traffic and
// security settings for all the workloads in the namespaces owned by
// the workspace. Any namespace in the workspace that is not part of a
// traffic or security group will use these default settings.
//
// The following example sets the default security policy to accept
// either mutual TLS or plaintext traffic, and only accept connections
// at a sidecar from services within the same namespace. The default
// traffic policy allows unknown traffic from a sidecar to be
// forwarded via an egress gateway `tsb-egress` in the `perimeter`
// namespace in the same cluster.
//
// ```yaml
// apiVersion: api.tsb.io/v1
// kind: WorkspaceSetting
// metadata:
//   name: w1-settings
//   workspace: w1
//   tenant: mycompany
// spec:
//   defaultSecuritySetting:
//     authentication: OPTIONAL
//   defaultTrafficSetting:
//     egressGateway: bookinfo-perimeter/tsb-egress
//```
//

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/tetrateio/api/tsb/security/v1"
	v11 "github.com/tetrateio/api/tsb/traffic/v1"
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

// Default security and traffic settings for all sidecars in the workspace.
type WorkspaceSetting struct {
	// Security settings for all sidecars in this workspace. Sidecars
	// without a specific security group will inherit these settings. If
	// omitted, the following semantics apply:
	//
	// 1. Sidecars will accept connections from clients using Istio
	// Mutual TLS as well as legacy clients using plaintext (i.e. any
	// traffic not using Istio Mutual TLS authentication),
	// i.e. authentication mode defaults to `OPTIONAL`.
	//
	// 2. No authorization will be performed, i.e., authorization mode defaults to `DISABLED`.
	DefaultSecuritySetting *v1.SecuritySetting `protobuf:"bytes,1,opt,name=default_security_setting,json=defaultSecuritySetting,proto3" json:"default_security_setting,omitempty"`
	// Traffic settings for all sidecars in this workspace. Sidecars
	// without a specific traffic group will inherit these settings. If
	// omitted, the following semantics apply:
	//
	// 1. Sidecars will be able to reach any service in the
	// cluster, i.e. reachability mode defaults to `CLUSTER`.
	//
	// 2. Traffic to unknown destinations will be directly routed from
	// the sidecar to the destination.
	DefaultTrafficSetting *v11.TrafficSetting `protobuf:"bytes,2,opt,name=default_traffic_setting,json=defaultTrafficSetting,proto3" json:"default_traffic_setting,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}            `json:"-"`
	XXX_unrecognized      []byte              `json:"-"`
	XXX_sizecache         int32               `json:"-"`
}

func (m *WorkspaceSetting) Reset()         { *m = WorkspaceSetting{} }
func (m *WorkspaceSetting) String() string { return proto.CompactTextString(m) }
func (*WorkspaceSetting) ProtoMessage()    {}
func (*WorkspaceSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_f877c5f6d2837d4b, []int{0}
}
func (m *WorkspaceSetting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkspaceSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkspaceSetting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkspaceSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceSetting.Merge(m, src)
}
func (m *WorkspaceSetting) XXX_Size() int {
	return m.Size()
}
func (m *WorkspaceSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceSetting.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceSetting proto.InternalMessageInfo

func (m *WorkspaceSetting) GetDefaultSecuritySetting() *v1.SecuritySetting {
	if m != nil {
		return m.DefaultSecuritySetting
	}
	return nil
}

func (m *WorkspaceSetting) GetDefaultTrafficSetting() *v11.TrafficSetting {
	if m != nil {
		return m.DefaultTrafficSetting
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkspaceSetting)(nil), "tetrateio.api.tsb.v1.WorkspaceSetting")
}

func init() { proto.RegisterFile("tsb/v1/workspace_setting.proto", fileDescriptor_f877c5f6d2837d4b) }

var fileDescriptor_f877c5f6d2837d4b = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x29, 0x4e, 0xd2,
	0x2f, 0x33, 0xd4, 0x2f, 0xcf, 0x2f, 0xca, 0x2e, 0x2e, 0x48, 0x4c, 0x4e, 0x8d, 0x2f, 0x4e, 0x2d,
	0x29, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x29, 0x49, 0x2d, 0x29,
	0x4a, 0x2c, 0x49, 0xcd, 0xcc, 0xd7, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x29, 0x4e, 0xd2, 0x2b, 0x33,
	0x94, 0x52, 0x01, 0xe9, 0x2a, 0x29, 0x4a, 0x4c, 0x4b, 0xcb, 0x4c, 0x06, 0xe9, 0x86, 0x32, 0x51,
	0xf5, 0x4a, 0xa9, 0x81, 0x54, 0x15, 0xa7, 0x26, 0x97, 0x16, 0x65, 0x96, 0x54, 0x82, 0x94, 0xc1,
	0xd8, 0xa8, 0xea, 0x94, 0x1e, 0x31, 0x72, 0x09, 0x84, 0xc3, 0xec, 0x0f, 0x86, 0x48, 0x09, 0x65,
	0x70, 0x49, 0xa4, 0xa4, 0xa6, 0x25, 0x96, 0xe6, 0x94, 0xc4, 0xa3, 0x6b, 0x93, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0xd2, 0xd3, 0xc3, 0x74, 0x1b, 0x4c, 0xa9, 0x5e, 0x99, 0xa1, 0x5e, 0x30, 0x94,
	0x0d, 0x35, 0x31, 0x48, 0x0c, 0x6a, 0x1e, 0x9a, 0xb8, 0x50, 0x0a, 0x97, 0x38, 0xcc, 0x26, 0x34,
	0x7f, 0x48, 0x30, 0x81, 0x2d, 0xd2, 0xc1, 0x62, 0x11, 0x54, 0x25, 0xc8, 0x9e, 0x10, 0x08, 0x13,
	0x66, 0x8d, 0x28, 0xd4, 0x30, 0x54, 0x61, 0x27, 0xdd, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x4a, 0x3e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0x6e, 0xb8, 0x7e, 0x62, 0x41, 0xa6, 0x3e, 0x24, 0x3e, 0x92, 0xd8, 0xc0, 0x41,
	0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x4c, 0x94, 0xcf, 0xa0, 0x01, 0x00, 0x00,
}

func (m *WorkspaceSetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkspaceSetting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkspaceSetting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DefaultTrafficSetting != nil {
		{
			size, err := m.DefaultTrafficSetting.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkspaceSetting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.DefaultSecuritySetting != nil {
		{
			size, err := m.DefaultSecuritySetting.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkspaceSetting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkspaceSetting(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkspaceSetting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkspaceSetting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultSecuritySetting != nil {
		l = m.DefaultSecuritySetting.Size()
		n += 1 + l + sovWorkspaceSetting(uint64(l))
	}
	if m.DefaultTrafficSetting != nil {
		l = m.DefaultTrafficSetting.Size()
		n += 1 + l + sovWorkspaceSetting(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkspaceSetting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkspaceSetting(x uint64) (n int) {
	return sovWorkspaceSetting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkspaceSetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkspaceSetting
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
			return fmt.Errorf("proto: WorkspaceSetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkspaceSetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultSecuritySetting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspaceSetting
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
				return ErrInvalidLengthWorkspaceSetting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspaceSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultSecuritySetting == nil {
				m.DefaultSecuritySetting = &v1.SecuritySetting{}
			}
			if err := m.DefaultSecuritySetting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultTrafficSetting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspaceSetting
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
				return ErrInvalidLengthWorkspaceSetting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspaceSetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultTrafficSetting == nil {
				m.DefaultTrafficSetting = &v11.TrafficSetting{}
			}
			if err := m.DefaultTrafficSetting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkspaceSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkspaceSetting
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkspaceSetting
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
func skipWorkspaceSetting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkspaceSetting
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
					return 0, ErrIntOverflowWorkspaceSetting
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
					return 0, ErrIntOverflowWorkspaceSetting
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
				return 0, ErrInvalidLengthWorkspaceSetting
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWorkspaceSetting
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkspaceSetting
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
				next, err := skipWorkspaceSetting(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWorkspaceSetting
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
	ErrInvalidLengthWorkspaceSetting = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkspaceSetting   = fmt.Errorf("proto: integer overflow")
)
