// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client_settings.proto

package v1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Available sensitivity levels.
type ClientSettings_Sensitivity int32

const (
	ClientSettings_UNSET  ClientSettings_Sensitivity = 0
	ClientSettings_LOW    ClientSettings_Sensitivity = 1
	ClientSettings_MEDIUM ClientSettings_Sensitivity = 2
	ClientSettings_HIGH   ClientSettings_Sensitivity = 3
)

var ClientSettings_Sensitivity_name = map[int32]string{
	0: "UNSET",
	1: "LOW",
	2: "MEDIUM",
	3: "HIGH",
}

var ClientSettings_Sensitivity_value = map[string]int32{
	"UNSET":  0,
	"LOW":    1,
	"MEDIUM": 2,
	"HIGH":   3,
}

func (x ClientSettings_Sensitivity) String() string {
	return proto.EnumName(ClientSettings_Sensitivity_name, int32(x))
}

func (ClientSettings_Sensitivity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9eff91b6df44f07c, []int{0, 0}
}

// ClientSettings
//
// All the timeouts, retries, circuit breakers etc. that consumers want to set to guard against failures of their
// dependencies. These roughly translate to pieces of istio virtual service, destination rules, etc. At runtime, for a
// given namespace, for the namespace's dependencies, we combine the client specified reliability settings with the
// virtual service/dest rule for the dependencies to produce an updated virtual service/dest rule per dependent service.
// All these client customized virtual services/dest rules will be private to the namespace i.e. will have an exportTo
// setting '.' . Doing so will make Pilot apply these customizations only to the namespace concerned without leaking it
// to all other namespaces.
type ClientSettings struct {
	// Timeout for HTTP requests.
	HttpRequestTimeout *types.Duration `protobuf:"bytes,1,opt,name=http_request_timeout,json=httpRequestTimeout,proto3" json:"http_request_timeout,omitempty"`
	// Retry policy for HTTP requests.
	HttpRetries *HTTPRetry `protobuf:"bytes,2,opt,name=http_retries,json=httpRetries,proto3" json:"http_retries,omitempty"`
	// TCP connection timeout.
	TcpConnectTimeout *types.Duration `protobuf:"bytes,3,opt,name=tcp_connect_timeout,json=tcpConnectTimeout,proto3" json:"tcp_connect_timeout,omitempty"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *types.BoolValue `protobuf:"bytes,4,opt,name=tcp_keepalive,json=tcpKeepalive,proto3" json:"tcp_keepalive,omitempty"`
	// The sensitivity levels will translate to specific values of dest rule outlier detection.
	CircuitBreakerSensitivity ClientSettings_Sensitivity `protobuf:"varint,5,opt,name=circuit_breaker_sensitivity,json=circuitBreakerSensitivity,proto3,enum=tetrate.api.tcc.core.v1.ClientSettings_Sensitivity" json:"circuit_breaker_sensitivity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                   `json:"-"`
	XXX_unrecognized          []byte                     `json:"-"`
	XXX_sizecache             int32                      `json:"-"`
}

func (m *ClientSettings) Reset()         { *m = ClientSettings{} }
func (m *ClientSettings) String() string { return proto.CompactTextString(m) }
func (*ClientSettings) ProtoMessage()    {}
func (*ClientSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eff91b6df44f07c, []int{0}
}
func (m *ClientSettings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientSettings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSettings.Merge(m, src)
}
func (m *ClientSettings) XXX_Size() int {
	return m.Size()
}
func (m *ClientSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSettings proto.InternalMessageInfo

func (m *ClientSettings) GetHttpRequestTimeout() *types.Duration {
	if m != nil {
		return m.HttpRequestTimeout
	}
	return nil
}

func (m *ClientSettings) GetHttpRetries() *HTTPRetry {
	if m != nil {
		return m.HttpRetries
	}
	return nil
}

func (m *ClientSettings) GetTcpConnectTimeout() *types.Duration {
	if m != nil {
		return m.TcpConnectTimeout
	}
	return nil
}

func (m *ClientSettings) GetTcpKeepalive() *types.BoolValue {
	if m != nil {
		return m.TcpKeepalive
	}
	return nil
}

func (m *ClientSettings) GetCircuitBreakerSensitivity() ClientSettings_Sensitivity {
	if m != nil {
		return m.CircuitBreakerSensitivity
	}
	return ClientSettings_UNSET
}

type HTTPRetry struct {
	// Number of retries for a given request. The interval between retries will be determined automatically (25ms+).
	// Actual number of retries attempted depends on the httpReqTimeout.
	Attempts int32 `protobuf:"varint,1,opt,name=attempts,proto3" json:"attempts,omitempty"`
	// Timeout per retry attempt for a given request. format: 1h/1m/1s/1ms. MUST BE >=1ms.
	PerTryTimeout *types.Duration `protobuf:"bytes,2,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
	// Specifies the conditions under which retry takes place. One or more policies can be specified using a ‘,’
	// delimited list. See the
	// [supported policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#x-envoy-retry-on)
	// and [here](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#x-envoy-retry-grpc-on)
	// for more details.
	RetryOn              string   `protobuf:"bytes,3,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPRetry) Reset()         { *m = HTTPRetry{} }
func (m *HTTPRetry) String() string { return proto.CompactTextString(m) }
func (*HTTPRetry) ProtoMessage()    {}
func (*HTTPRetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eff91b6df44f07c, []int{1}
}
func (m *HTTPRetry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPRetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HTTPRetry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HTTPRetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPRetry.Merge(m, src)
}
func (m *HTTPRetry) XXX_Size() int {
	return m.Size()
}
func (m *HTTPRetry) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPRetry.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPRetry proto.InternalMessageInfo

func (m *HTTPRetry) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return 0
}

func (m *HTTPRetry) GetPerTryTimeout() *types.Duration {
	if m != nil {
		return m.PerTryTimeout
	}
	return nil
}

func (m *HTTPRetry) GetRetryOn() string {
	if m != nil {
		return m.RetryOn
	}
	return ""
}

func init() {
	proto.RegisterEnum("tetrate.api.tcc.core.v1.ClientSettings_Sensitivity", ClientSettings_Sensitivity_name, ClientSettings_Sensitivity_value)
	proto.RegisterType((*ClientSettings)(nil), "tetrate.api.tcc.core.v1.ClientSettings")
	proto.RegisterType((*HTTPRetry)(nil), "tetrate.api.tcc.core.v1.HTTPRetry")
}

func init() { proto.RegisterFile("client_settings.proto", fileDescriptor_9eff91b6df44f07c) }

var fileDescriptor_9eff91b6df44f07c = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x3f, 0xc7, 0x49, 0x9b, 0x4c, 0xda, 0x7e, 0x61, 0x00, 0x35, 0x09, 0x52, 0x54, 0x45,
	0x42, 0xca, 0x02, 0x8d, 0xd5, 0x76, 0x81, 0xc4, 0x06, 0x91, 0x36, 0x22, 0x51, 0x29, 0x45, 0x8e,
	0x0b, 0x12, 0x1b, 0x6b, 0x32, 0xbd, 0xa4, 0xa3, 0x3a, 0x9e, 0x61, 0x7c, 0x6d, 0xe4, 0xb7, 0xe0,
	0x79, 0x58, 0xb1, 0x64, 0xc9, 0x13, 0x20, 0x94, 0x1d, 0x6f, 0x81, 0xe2, 0x3f, 0x69, 0x51, 0x55,
	0x75, 0x37, 0xa3, 0x33, 0xe7, 0x77, 0xcf, 0x9c, 0x4b, 0x1e, 0x8b, 0x40, 0x42, 0x88, 0x7e, 0x04,
	0x88, 0x32, 0x9c, 0x47, 0x4c, 0x1b, 0x85, 0x8a, 0xee, 0x22, 0xa0, 0xe1, 0x08, 0x8c, 0x6b, 0xc9,
	0x50, 0x08, 0x26, 0x94, 0x01, 0x96, 0xec, 0x77, 0x7b, 0x73, 0xa5, 0xe6, 0x01, 0x38, 0xd9, 0xb3,
	0x59, 0xfc, 0xc9, 0xb9, 0x88, 0x0d, 0x47, 0xa9, 0xc2, 0xdc, 0x78, 0x5b, 0xff, 0x62, 0xb8, 0xd6,
	0x60, 0x0a, 0x70, 0x77, 0x37, 0xe1, 0x81, 0xbc, 0xe0, 0x08, 0x4e, 0x79, 0xc8, 0x85, 0xfe, 0x2f,
	0x9b, 0xec, 0x1c, 0x65, 0x59, 0xa6, 0x45, 0x14, 0x7a, 0x42, 0x1e, 0x5d, 0x22, 0x6a, 0xdf, 0xc0,
	0xe7, 0x18, 0x22, 0xf4, 0x51, 0x2e, 0x40, 0xc5, 0xd8, 0xb6, 0xf6, 0xac, 0x41, 0xf3, 0xa0, 0xc3,
	0xf2, 0x51, 0xac, 0x1c, 0xc5, 0x8e, 0x8b, 0x28, 0x2e, 0x5d, 0xd9, 0xdc, 0xdc, 0xe5, 0xe5, 0x26,
	0x3a, 0x22, 0x5b, 0x05, 0x0c, 0x8d, 0x84, 0xa8, 0x5d, 0xc9, 0x20, 0x7d, 0x76, 0xc7, 0x47, 0xd9,
	0xd8, 0xf3, 0xde, 0xb9, 0x80, 0x26, 0x75, 0x9b, 0x39, 0x2d, 0xb3, 0xd1, 0x09, 0x79, 0x88, 0x42,
	0xfb, 0x42, 0x85, 0x21, 0x88, 0xeb, 0x48, 0xf6, 0x7d, 0x91, 0x1e, 0xa0, 0xd0, 0x47, 0xb9, 0xa9,
	0x4c, 0xf4, 0x92, 0x6c, 0xaf, 0x50, 0x57, 0x00, 0x9a, 0x07, 0x32, 0x81, 0x76, 0x35, 0x83, 0x74,
	0x6f, 0x41, 0x86, 0x4a, 0x05, 0xef, 0x79, 0x10, 0x83, 0xbb, 0x85, 0x42, 0x9f, 0x94, 0xef, 0x69,
	0x44, 0x9e, 0x08, 0x69, 0x44, 0x2c, 0xd1, 0x9f, 0x19, 0xe0, 0x57, 0x60, 0xfc, 0x08, 0xc2, 0x48,
	0xa2, 0x4c, 0x24, 0xa6, 0xed, 0xda, 0x9e, 0x35, 0xd8, 0x39, 0x38, 0xbc, 0xf3, 0x87, 0xff, 0xb6,
	0xcd, 0xa6, 0xd7, 0x56, 0xb7, 0x53, 0x70, 0x87, 0x39, 0xf6, 0x86, 0xd4, 0x7f, 0x4e, 0x9a, 0x37,
	0xae, 0xb4, 0x41, 0x6a, 0xe7, 0x6f, 0xa7, 0x23, 0xaf, 0xf5, 0x1f, 0xdd, 0x24, 0xf6, 0x9b, 0xb3,
	0x0f, 0x2d, 0x8b, 0x12, 0xb2, 0x71, 0x3a, 0x3a, 0x9e, 0x9c, 0x9f, 0xb6, 0x2a, 0xb4, 0x4e, 0xaa,
	0xe3, 0xc9, 0xeb, 0x71, 0xcb, 0xee, 0x7f, 0xb5, 0x48, 0x63, 0x5d, 0x2a, 0x7d, 0x4a, 0xea, 0x1c,
	0x11, 0x16, 0x1a, 0xa3, 0x6c, 0x9f, 0xb5, 0x61, 0xe3, 0xdb, 0x9f, 0xef, 0x76, 0xb5, 0x5b, 0x19,
	0x58, 0xee, 0x5a, 0xa2, 0xaf, 0xc8, 0xff, 0x1a, 0x8c, 0x8f, 0x26, 0x5d, 0x57, 0x5d, 0xb9, 0xaf,
	0xea, 0x6d, 0x0d, 0xc6, 0x33, 0x69, 0x59, 0x73, 0x87, 0xd4, 0x57, 0x3b, 0x4f, 0x7d, 0x15, 0x66,
	0x6b, 0x6a, 0xb8, 0x9b, 0xd9, 0xfd, 0x2c, 0x1c, 0xbe, 0xf8, 0xb1, 0xec, 0x59, 0x3f, 0x97, 0x3d,
	0xeb, 0xf7, 0xb2, 0x67, 0x7d, 0x7c, 0x36, 0x97, 0x78, 0x19, 0xcf, 0x98, 0x50, 0x0b, 0xa7, 0xe8,
	0x4c, 0xaa, 0xf2, 0xe4, 0x70, 0x2d, 0x1d, 0x14, 0xc2, 0x59, 0xb5, 0xe7, 0x24, 0xfb, 0xb3, 0x8d,
	0x6c, 0xf0, 0xe1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0xd2, 0x08, 0x50, 0x41, 0x03, 0x00,
	0x00,
}

func (m *ClientSettings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientSettings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientSettings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CircuitBreakerSensitivity != 0 {
		i = encodeVarintClientSettings(dAtA, i, uint64(m.CircuitBreakerSensitivity))
		i--
		dAtA[i] = 0x28
	}
	if m.TcpKeepalive != nil {
		{
			size, err := m.TcpKeepalive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClientSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TcpConnectTimeout != nil {
		{
			size, err := m.TcpConnectTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClientSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HttpRetries != nil {
		{
			size, err := m.HttpRetries.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClientSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.HttpRequestTimeout != nil {
		{
			size, err := m.HttpRequestTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClientSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HTTPRetry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPRetry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HTTPRetry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RetryOn) > 0 {
		i -= len(m.RetryOn)
		copy(dAtA[i:], m.RetryOn)
		i = encodeVarintClientSettings(dAtA, i, uint64(len(m.RetryOn)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PerTryTimeout != nil {
		{
			size, err := m.PerTryTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClientSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Attempts != 0 {
		i = encodeVarintClientSettings(dAtA, i, uint64(m.Attempts))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClientSettings(dAtA []byte, offset int, v uint64) int {
	offset -= sovClientSettings(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientSettings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpRequestTimeout != nil {
		l = m.HttpRequestTimeout.Size()
		n += 1 + l + sovClientSettings(uint64(l))
	}
	if m.HttpRetries != nil {
		l = m.HttpRetries.Size()
		n += 1 + l + sovClientSettings(uint64(l))
	}
	if m.TcpConnectTimeout != nil {
		l = m.TcpConnectTimeout.Size()
		n += 1 + l + sovClientSettings(uint64(l))
	}
	if m.TcpKeepalive != nil {
		l = m.TcpKeepalive.Size()
		n += 1 + l + sovClientSettings(uint64(l))
	}
	if m.CircuitBreakerSensitivity != 0 {
		n += 1 + sovClientSettings(uint64(m.CircuitBreakerSensitivity))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HTTPRetry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Attempts != 0 {
		n += 1 + sovClientSettings(uint64(m.Attempts))
	}
	if m.PerTryTimeout != nil {
		l = m.PerTryTimeout.Size()
		n += 1 + l + sovClientSettings(uint64(l))
	}
	l = len(m.RetryOn)
	if l > 0 {
		n += 1 + l + sovClientSettings(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClientSettings(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClientSettings(x uint64) (n int) {
	return sovClientSettings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientSettings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientSettings
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
			return fmt.Errorf("proto: ClientSettings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientSettings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpRequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
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
				return ErrInvalidLengthClientSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpRequestTimeout == nil {
				m.HttpRequestTimeout = &types.Duration{}
			}
			if err := m.HttpRequestTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpRetries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
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
				return ErrInvalidLengthClientSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpRetries == nil {
				m.HttpRetries = &HTTPRetry{}
			}
			if err := m.HttpRetries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpConnectTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
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
				return ErrInvalidLengthClientSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TcpConnectTimeout == nil {
				m.TcpConnectTimeout = &types.Duration{}
			}
			if err := m.TcpConnectTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpKeepalive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
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
				return ErrInvalidLengthClientSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TcpKeepalive == nil {
				m.TcpKeepalive = &types.BoolValue{}
			}
			if err := m.TcpKeepalive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CircuitBreakerSensitivity", wireType)
			}
			m.CircuitBreakerSensitivity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CircuitBreakerSensitivity |= ClientSettings_Sensitivity(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClientSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientSettings
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClientSettings
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
func (m *HTTPRetry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientSettings
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
			return fmt.Errorf("proto: HTTPRetry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPRetry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attempts", wireType)
			}
			m.Attempts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Attempts |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerTryTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
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
				return ErrInvalidLengthClientSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PerTryTimeout == nil {
				m.PerTryTimeout = &types.Duration{}
			}
			if err := m.PerTryTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetryOn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSettings
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
				return ErrInvalidLengthClientSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetryOn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientSettings
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClientSettings
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
func skipClientSettings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientSettings
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
					return 0, ErrIntOverflowClientSettings
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
					return 0, ErrIntOverflowClientSettings
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
				return 0, ErrInvalidLengthClientSettings
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthClientSettings
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClientSettings
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
				next, err := skipClientSettings(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthClientSettings
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
	ErrInvalidLengthClientSettings = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientSettings   = fmt.Errorf("proto: integer overflow")
)
