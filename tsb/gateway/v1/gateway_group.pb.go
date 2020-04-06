// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/gateway/v1/gateway_group.proto

// Gateway Groups allow grouping the gateways in a set of namespaces
// owned by its parent workspace. Gateway related configurations can
// then be applied on the group to control the behavior of these
// gateways. The group can be in one of two modes: `TETRATE` and
// `ISTIO`. `TETRATE` mode is a minimalistic mode that allows users to
// quickly configure the most commonly used features in the service
// mesh using Tetrate specific APIs, while the `ISTIO` mode provides
// more flexibility for power users by allowing them to configure the
// gateways's traffic and security properties using a restricted
// subset of Istio Networking and Security APIs.
//
// The following example creates a gateway group for the gateways in
// `ns1`, `ns2` and `ns3` namespaces owned by its parent workspace
// `w1` under tenant `mycompany`
//
// ```yaml
// apiVersion: gateway.tsb.tetrate.io/v1
// kind: Group
// metadata:
//   name: g1
//   workspace: w1
//   tenant: mycompany
// spec:
//   namespaceSelector:
//     names:
//     - */ns1
//     - */ns2
//     - */ns3
//   configMode: TETRATE
//```
//
// Under the hood, Service Bridge translates these minimalistic
// settings into Istio APIs such as `Gateway`, `VirtualService`,
// `AuthorizationPolicy`, etc. for the namespaces managed by the
// gateway group. These APIs are then pushed to the Istio control
// planes of clusters where the workspace is applicable.
//
// It is possible to create a gateway group for namespaces in a
// specific cluster as long as the parent workspace owns those
// namespaces in that cluster. For example,
//
// ```yaml
// apiVersion: gateway.tsb.tetrate.io/v1
// kind: Group
// metadata:
//   name: g1
//   workspace: w1
//   tenant: mycompany
// spec:
//   namespaceSelector:
//     names:
//     - c1/ns1 # pick ns1 namespace only from c1 cluster
//     - */ns2
//     - */ns3
//   configMode: TETRATE
//```
//
// In the `ISTIO` mode, it is possible to directly attach Istio
// `VirtualService`, and `Gateway` to the gateway group. These
// configurations will be validated for correctness and conflict free
// operations and then pushed to the appropriate Istio control planes.
//
// The following example declares a `Gateway` and a `VirtualService`
// for a specific workload in the `ns1` namespace:
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: ingress
//   namespace: ns1
//   annotations:
//     tsb.tetrate.io/tenant: mycompany
//     tsb.tetrate.io/workspace: w1
//     tsb.tetrate.io/gatewayGroup: g1
// spec:
//   selector:
//       app: my-ingress-gateway
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
// ```
//
// and the associated `VirtualService`
//
//```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: ingress-rule
//   namespace: ns1
//   annotations:
//     tsb.tetrate.io/tenant: mycompany
//     tsb.tetrate.io/workspace: w1
//     tsb.tetrate.io/gatewayGroup: g1
// spec:
//   hosts:
//   - uk.bookinfo.com
//   - eu.bookinfo.com
//   gateways:
//   - ns1/ingress # Has to bind to the same gateway
//   http:
//   - route:
//     - destination:
//         port:
//           number: 7777
//         host: reviews.ns1.svc.cluster.local
//```
//
// The namespace where the Istio APIs are applied will need to be part
// of the parent gateway group. In addition, each API object will need
// to have annotations to indicate the tenant, workspace and the
// gateway group to which it belongs to.
//
// TODO(rshriram): Document the list of fields in Istio APIs that are
// disallowed.

package v1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

// A gateway group manages the gateways in a group of namespaces owned
// by the parent workspace.
type Group struct {
	// Set of namespaces owned exclusively by this group. If omitted,
	// applies to all resources owned by the workspace. Use `*/*` to
	// claim all cluster resources under the workspace.
	NamespaceSelector *v1.NamespaceSelector `protobuf:"bytes,1,opt,name=namespace_selector,json=namespaceSelector,proto3" json:"namespace_selector,omitempty"`
	// The Configuration types that will be added to this
	// group. `TETRATE` mode indicates that configurations added to this
	// group will use Tetrate APIs such as `IngressGateway`. `ISTIO`
	// mode indicates that configurations added to this group will use
	// Istio Networking APIs such as `Gateway` and
	// `VirtualService`. Defaults to `TETRATE` mode.
	ConfigMode           v1.ConfigMode `protobuf:"varint,2,opt,name=config_mode,json=configMode,proto3,enum=tetrateio.api.tsb.types.v1.ConfigMode" json:"config_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2fab70b32ccea63, []int{0}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetNamespaceSelector() *v1.NamespaceSelector {
	if m != nil {
		return m.NamespaceSelector
	}
	return nil
}

func (m *Group) GetConfigMode() v1.ConfigMode {
	if m != nil {
		return m.ConfigMode
	}
	return v1.ConfigMode_TETRATE
}

func init() {
	proto.RegisterType((*Group)(nil), "tetrateio.api.tsb.gateway.v1.Group")
}

func init() { proto.RegisterFile("tsb/gateway/v1/gateway_group.proto", fileDescriptor_c2fab70b32ccea63) }

var fileDescriptor_c2fab70b32ccea63 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe5, 0x4a, 0xfd, 0x86, 0x54, 0xfa, 0x04, 0x99, 0xaa, 0x0a, 0x95, 0xaa, 0x03, 0x74,
	0xa9, 0xad, 0x14, 0xf1, 0x00, 0x0d, 0x43, 0x27, 0x18, 0xca, 0xc6, 0x12, 0x39, 0xce, 0xad, 0x6b,
	0x29, 0xc9, 0xb5, 0x1c, 0xd7, 0x90, 0x57, 0xe0, 0x95, 0x78, 0x01, 0x46, 0x1e, 0xa1, 0xca, 0x63,
	0x30, 0xa1, 0xfc, 0x69, 0x05, 0x54, 0x62, 0xfb, 0x49, 0xe7, 0x9e, 0xe3, 0x73, 0xaf, 0xbd, 0xa9,
	0x2d, 0x62, 0x26, 0xb9, 0x85, 0x67, 0x5e, 0x32, 0x17, 0x1c, 0x30, 0x92, 0x06, 0x77, 0x9a, 0x6a,
	0x83, 0x16, 0xfd, 0x0b, 0x0b, 0xd6, 0x70, 0x0b, 0x0a, 0x29, 0xd7, 0x8a, 0xda, 0x22, 0xa6, 0xdd,
	0x18, 0x75, 0xc1, 0xe8, 0x52, 0x22, 0xca, 0x14, 0x18, 0xd7, 0x8a, 0x6d, 0x14, 0xa4, 0x49, 0x14,
	0xc3, 0x96, 0x3b, 0x85, 0xa6, 0xb5, 0x8f, 0x96, 0x52, 0xd9, 0xed, 0x2e, 0xa6, 0x02, 0x33, 0x06,
	0xb9, 0xc3, 0x52, 0x1b, 0x7c, 0x29, 0x59, 0x23, 0x8a, 0xb9, 0x84, 0x7c, 0xee, 0x78, 0xaa, 0x12,
	0x6e, 0x81, 0x9d, 0x40, 0x17, 0x31, 0xac, 0x5b, 0xda, 0x52, 0x43, 0x51, 0x77, 0x6c, 0xa0, 0x55,
	0xa6, 0x6f, 0xc4, 0xeb, 0xaf, 0xea, 0xae, 0xbe, 0xf2, 0xfc, 0x9c, 0x67, 0x50, 0x68, 0x2e, 0x20,
	0x2a, 0x20, 0x05, 0x61, 0xd1, 0x0c, 0xc9, 0x84, 0xcc, 0x06, 0x8b, 0x39, 0x3d, 0x5d, 0xa1, 0x4d,
	0x71, 0x01, 0x7d, 0x38, 0xb8, 0x1e, 0x3b, 0x53, 0x38, 0xd8, 0x2f, 0x7b, 0x9f, 0x61, 0xff, 0x95,
	0xf4, 0xce, 0xc8, 0xfa, 0x3c, 0xff, 0xad, 0xfb, 0x2b, 0x6f, 0x20, 0x30, 0xdf, 0x28, 0x19, 0x65,
	0x98, 0xc0, 0xb0, 0x37, 0x21, 0xb3, 0xff, 0x8b, 0xab, 0xbf, 0xde, 0xb8, 0x6b, 0xc6, 0xef, 0x31,
	0x81, 0xb5, 0x27, 0x8e, 0x1c, 0xde, 0xbe, 0x57, 0x63, 0xf2, 0x51, 0x8d, 0xc9, 0xbe, 0x1a, 0x93,
	0xa7, 0xeb, 0x6f, 0x87, 0x3a, 0x66, 0x35, 0x77, 0xfd, 0xf9, 0x49, 0xf1, 0xbf, 0x66, 0xf7, 0x9b,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0xd0, 0xb4, 0xc2, 0xbd, 0x01, 0x00, 0x00,
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConfigMode != 0 {
		i = encodeVarintGatewayGroup(dAtA, i, uint64(m.ConfigMode))
		i--
		dAtA[i] = 0x10
	}
	if m.NamespaceSelector != nil {
		{
			size, err := m.NamespaceSelector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGatewayGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGatewayGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGatewayGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NamespaceSelector != nil {
		l = m.NamespaceSelector.Size()
		n += 1 + l + sovGatewayGroup(uint64(l))
	}
	if m.ConfigMode != 0 {
		n += 1 + sovGatewayGroup(uint64(m.ConfigMode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGatewayGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGatewayGroup(x uint64) (n int) {
	return sovGatewayGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayGroup
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
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayGroup
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
				return ErrInvalidLengthGatewayGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGatewayGroup
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigMode", wireType)
			}
			m.ConfigMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfigMode |= v1.ConfigMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGatewayGroup
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
func skipGatewayGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatewayGroup
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
					return 0, ErrIntOverflowGatewayGroup
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
					return 0, ErrIntOverflowGatewayGroup
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
				return 0, ErrInvalidLengthGatewayGroup
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGatewayGroup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGatewayGroup
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
				next, err := skipGatewayGroup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGatewayGroup
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
	ErrInvalidLengthGatewayGroup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatewayGroup   = fmt.Errorf("proto: integer overflow")
)
