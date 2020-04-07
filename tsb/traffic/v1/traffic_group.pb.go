// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsb/traffic/v1/traffic_group.proto

// Traffic Groups allow grouping the sidecars in a set of namespaces
// owned by its parent workspace. Networking and routing related
// configurations can then be applied on the group to control the
// behavior of these sidecars. The group can be in one of two modes:
// `TETRATE` and `ISTIO`. `TETRATE` mode is a minimalistic mode that
// allows users to quickly configure the most commonly used features
// in the service mesh using Tetrate specific APIs, while the `ISTIO`
// mode provides more flexibility for power users by allowing them to
// configure the sidecar behavior using a restricted subset of Istio
// Networking APIs.
//
// The following example creates a traffic group for the sidecars in
// `ns1`, `ns2` and `ns3` namespaces owned by its parent workspace
// `w1` under tenant `mycompany` and sets up a `TrafficSetting`
// defining the resilience properties for sidecars in these
// namespaces.
//
// ```yaml
// apiVersion: traffic.tsb.tetrate.io/v1
// kind: Group
// metadata:
//   name: t1
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
// And the associated traffic settings for the sidecars in the group
//
// ```yaml
// apiVersion: traffic.tsb.tetrate.io/v1
// kind: TrafficSetting
// metadata:
//   name: defaults
//   group: t1
//   workspace: w1
//   tenant: mycompany
// spec:
//   resilience:
//     circuitBreakerSensitivity: MEDIUM
//```
//
// Under the hood, Service Bridge translates these minimalistic
// settings into Istio APIs such as `Sidecar`, `DestinationRule`,
// etc. for the namespaces managed by the traffic group. These APIs
// are then pushed to the Istio control planes of clusters where the
// workspace is applicable.
//
// It is possible to create a traffic group for namespaces in a
// specific cluster as long as the parent workspace owns those
// namespaces in that cluster. For example,
//
// ```yaml
// apiVersion: traffic.tsb.tetrate.io/v1
// kind: Group
// metadata:
//   name: t1
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
// In the `ISTIO` mode, it is possible to directly attach Istio APIs
// such as `VirtualService`, `DestinationRule`, and `Sidecar` to the
// traffic group. These configurations will be validated for
// correctness and conflict free operations and then pushed to the
// appropriate Istio control planes.
//
// The following example declares a `DestinationRule` with two
// subsets, for the `ratings` service in the `ns1` namespace:
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: ratings-subsets
//   namespace: ns1
//   annotations:
//     tsb.tetrate.io/tenant: mycompany
//     tsb.tetrate.io/workspace: w1
//     tsb.tetrate.io/trafficGroup: t1
// spec:
//   host: ratings.ns1.svc.cluster.local
//   subsets:
//   - name: stableversion
//     labels:
//       app: ratings
//       env: prod
//   - name: testversion
//     labels:
//       app: ratings
//       env: uat
// ```
//
// The namespace where the Istio APIs are applied will need to be part
// of the parent traffic group. In addition, each API object will need
// to have annotations to indicate the tenant, workspace and the
// traffic group to which it belongs to.
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

// A traffic group manages the routing properties of sidecars in a
// group of namespaces owned by the parent workspace.
type Group struct {
	// Set of namespaces owned exclusively by this group. If omitted,
	// applies to all resources owned by the workspace. Use `*/*` to
	// claim all cluster resources under the workspace.
	NamespaceSelector *v1.NamespaceSelector `protobuf:"bytes,1,opt,name=namespace_selector,json=namespaceSelector,proto3" json:"namespace_selector,omitempty"`
	// The Configuration types that will be added to this
	// group. `TETRATE` mode indicates that configurations added to this
	// group will use Tetrate APIs such as `TrafficSetting` and
	// `ServiceRoute`. `ISTIO` mode indicates that configurations added
	// to this group will use Istio Networking APIs such as
	// `VirtualService`, `DestinationRule`, and `Sidecar`. Defaults to
	// `TETRATE` mode.
	ConfigMode           v1.ConfigMode `protobuf:"varint,2,opt,name=config_mode,json=configMode,proto3,enum=tetrateio.api.tsb.types.v1.ConfigMode" json:"config_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_311e16eb711559ee, []int{0}
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
	proto.RegisterType((*Group)(nil), "tetrateio.api.tsb.traffic.v1.Group")
}

func init() { proto.RegisterFile("tsb/traffic/v1/traffic_group.proto", fileDescriptor_311e16eb711559ee) }

var fileDescriptor_311e16eb711559ee = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x99, 0x40, 0xff, 0x45, 0x02, 0x3f, 0x9a, 0x8d, 0xa1, 0x48, 0x2c, 0x5d, 0x68, 0x37,
	0xce, 0x90, 0x8a, 0x0f, 0x60, 0x5c, 0x74, 0xa5, 0x8b, 0xba, 0x73, 0x13, 0x26, 0x93, 0x9b, 0x74,
	0x20, 0xc9, 0x1d, 0x26, 0xd3, 0x01, 0x5f, 0xc1, 0x57, 0xf2, 0x05, 0x5c, 0xfa, 0x08, 0x25, 0x8f,
	0xe1, 0x4a, 0x3a, 0x4d, 0x0a, 0x8a, 0xb8, 0xfb, 0xe0, 0x9e, 0x73, 0xcf, 0xbd, 0xc7, 0x9f, 0x9b,
	0x2e, 0x67, 0x46, 0xf3, 0xb2, 0x94, 0x82, 0xd9, 0x64, 0xc4, 0xac, 0xd2, 0xb8, 0x55, 0x54, 0x69,
	0x34, 0x18, 0x9e, 0x1b, 0x30, 0x9a, 0x1b, 0x90, 0x48, 0xb9, 0x92, 0xd4, 0x74, 0x39, 0x1d, 0x64,
	0xd4, 0x26, 0xd3, 0x33, 0xcb, 0x6b, 0x59, 0x70, 0x03, 0x6c, 0x84, 0x83, 0x6d, 0x7a, 0x51, 0x21,
	0x56, 0x35, 0x30, 0xae, 0x24, 0x2b, 0x25, 0xd4, 0x45, 0x96, 0xc3, 0x86, 0x5b, 0x89, 0x7a, 0x10,
	0x44, 0x2e, 0xfb, 0x45, 0x41, 0xe7, 0x92, 0xf7, 0x70, 0x98, 0xcc, 0xdf, 0x88, 0x3f, 0x59, 0xed,
	0x2f, 0x08, 0xa5, 0x1f, 0xb6, 0xbc, 0x81, 0x4e, 0x71, 0x01, 0x59, 0x07, 0x35, 0x08, 0x83, 0x3a,
	0x22, 0x33, 0xb2, 0x08, 0x96, 0xd7, 0xf4, 0x97, 0xc3, 0xdc, 0x16, 0x9b, 0xd0, 0xc7, 0xd1, 0xf5,
	0x34, 0x98, 0xd2, 0x60, 0x77, 0xe7, 0x7d, 0xa6, 0x93, 0x57, 0xe2, 0x9d, 0x90, 0xf5, 0x69, 0xfb,
	0x73, 0x1e, 0xae, 0xfc, 0x40, 0x60, 0x5b, 0xca, 0x2a, 0x6b, 0xb0, 0x80, 0xc8, 0x9b, 0x91, 0xc5,
	0xff, 0xe5, 0xe5, 0x5f, 0x19, 0xf7, 0x4e, 0xfe, 0x80, 0x05, 0xac, 0x7d, 0x71, 0xe4, 0xf4, 0xf6,
	0xbd, 0x8f, 0xc9, 0x47, 0x1f, 0x93, 0x5d, 0x1f, 0x93, 0xe7, 0xab, 0x4a, 0x9a, 0xcd, 0x36, 0xa7,
	0x02, 0x1b, 0x76, 0xdc, 0xe5, 0x4a, 0xf9, 0x5e, 0x7d, 0xfe, 0xcf, 0xfd, 0x7e, 0xf3, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x58, 0xa4, 0x74, 0x92, 0x93, 0x01, 0x00, 0x00,
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
		i = encodeVarintTrafficGroup(dAtA, i, uint64(m.ConfigMode))
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
			i = encodeVarintTrafficGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrafficGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrafficGroup(v)
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
		n += 1 + l + sovTrafficGroup(uint64(l))
	}
	if m.ConfigMode != 0 {
		n += 1 + sovTrafficGroup(uint64(m.ConfigMode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTrafficGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrafficGroup(x uint64) (n int) {
	return sovTrafficGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficGroup
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
					return ErrIntOverflowTrafficGroup
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
				return ErrInvalidLengthTrafficGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficGroup
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
					return ErrIntOverflowTrafficGroup
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
			skippy, err := skipTrafficGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficGroup
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
func skipTrafficGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrafficGroup
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
					return 0, ErrIntOverflowTrafficGroup
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
					return 0, ErrIntOverflowTrafficGroup
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
				return 0, ErrInvalidLengthTrafficGroup
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTrafficGroup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrafficGroup
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
				next, err := skipTrafficGroup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTrafficGroup
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
	ErrInvalidLengthTrafficGroup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrafficGroup   = fmt.Errorf("proto: integer overflow")
)
