// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_settings.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/core/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import duration "github.com/golang/protobuf/ptypes/duration"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_client_settings_05871f27aeafa0ef, []int{0, 0}
}

// ClientSettings
//
// All the timeouts, retries, circuit breakers etc. that consumers want to set to gaurd against failures of their
// dependencies. These roughly translate to pieces of istio virtual service, destination rules, etc. At runtime, for a
// given namespace, for the namespace's dependencies, we combine the client specified reliability settings with the
// virtual service/dest rule for the dependencies to produce an updated virtual service/dest rule per dependent service.
// All these client customized virtual services/dest rules will be private to the namespace i.e. will have an exportTo
// setting '.' . Doing so will make Pilot apply these customizations only to the namespace concerned without leaking it
// to all other namespaces.
type ClientSettings struct {
	// Timeout for HTTP requests.
	HttpRequestTimeout *duration.Duration `protobuf:"bytes,1,opt,name=http_request_timeout,json=httpRequestTimeout,proto3" json:"http_request_timeout,omitempty"`
	// Retry policy for HTTP requests.
	HttpRetries *HTTPRetry `protobuf:"bytes,2,opt,name=http_retries,json=httpRetries,proto3" json:"http_retries,omitempty"`
	// TCP connection timeout.
	TcpConnectTimeout *duration.Duration `protobuf:"bytes,3,opt,name=tcp_connect_timeout,json=tcpConnectTimeout,proto3" json:"tcp_connect_timeout,omitempty"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *wrappers.BoolValue `protobuf:"bytes,4,opt,name=tcp_keepalive,json=tcpKeepalive,proto3" json:"tcp_keepalive,omitempty"`
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
	return fileDescriptor_client_settings_05871f27aeafa0ef, []int{0}
}
func (m *ClientSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSettings.Unmarshal(m, b)
}
func (m *ClientSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSettings.Marshal(b, m, deterministic)
}
func (dst *ClientSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSettings.Merge(dst, src)
}
func (m *ClientSettings) XXX_Size() int {
	return xxx_messageInfo_ClientSettings.Size(m)
}
func (m *ClientSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSettings proto.InternalMessageInfo

func (m *ClientSettings) GetHttpRequestTimeout() *duration.Duration {
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

func (m *ClientSettings) GetTcpConnectTimeout() *duration.Duration {
	if m != nil {
		return m.TcpConnectTimeout
	}
	return nil
}

func (m *ClientSettings) GetTcpKeepalive() *wrappers.BoolValue {
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
	PerTryTimeout *duration.Duration `protobuf:"bytes,2,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
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
	return fileDescriptor_client_settings_05871f27aeafa0ef, []int{1}
}
func (m *HTTPRetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPRetry.Unmarshal(m, b)
}
func (m *HTTPRetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPRetry.Marshal(b, m, deterministic)
}
func (dst *HTTPRetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPRetry.Merge(dst, src)
}
func (m *HTTPRetry) XXX_Size() int {
	return xxx_messageInfo_HTTPRetry.Size(m)
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

func (m *HTTPRetry) GetPerTryTimeout() *duration.Duration {
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
	proto.RegisterType((*ClientSettings)(nil), "tetrate.api.tcc.core.v1.ClientSettings")
	proto.RegisterType((*HTTPRetry)(nil), "tetrate.api.tcc.core.v1.HTTPRetry")
	proto.RegisterEnum("tetrate.api.tcc.core.v1.ClientSettings_Sensitivity", ClientSettings_Sensitivity_name, ClientSettings_Sensitivity_value)
}

func init() {
	proto.RegisterFile("client_settings.proto", fileDescriptor_client_settings_05871f27aeafa0ef)
}

var fileDescriptor_client_settings_05871f27aeafa0ef = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0xd3, 0x6e, 0xad, 0xbb, 0x8d, 0x62, 0x40, 0x6b, 0x8b, 0x84, 0xa6, 0x4a, 0x48,
	0xbd, 0x40, 0x8e, 0xb6, 0x5d, 0x70, 0x89, 0xe8, 0x56, 0xd1, 0x6a, 0x8c, 0xa1, 0x34, 0x03, 0x89,
	0x9b, 0xc8, 0xf5, 0x0e, 0x9d, 0xb5, 0x34, 0x36, 0xce, 0x49, 0x50, 0xde, 0x82, 0xe7, 0xe1, 0x8a,
	0x27, 0xe1, 0x9e, 0xb7, 0x40, 0x75, 0x92, 0x6e, 0x68, 0x9a, 0x7a, 0x67, 0xeb, 0xf7, 0xff, 0x9d,
	0xdf, 0xff, 0x21, 0xcf, 0x45, 0x24, 0x21, 0xc6, 0x30, 0x01, 0x44, 0x19, 0x2f, 0x12, 0xa6, 0x8d,
	0x42, 0x45, 0xf7, 0x11, 0xd0, 0x70, 0x04, 0xc6, 0xb5, 0x64, 0x28, 0x04, 0x13, 0xca, 0x00, 0xcb,
	0x0e, 0xfb, 0x2f, 0x17, 0x4a, 0x2d, 0x22, 0xf0, 0xec, 0xb3, 0x79, 0xfa, 0xcd, 0xbb, 0x4a, 0x0d,
	0x47, 0xa9, 0xe2, 0xc2, 0x78, 0x5f, 0xff, 0x61, 0xb8, 0xd6, 0x60, 0x4a, 0x70, 0x7f, 0x3f, 0xe3,
	0x91, 0xbc, 0xe2, 0x08, 0x5e, 0x75, 0x28, 0x84, 0xc1, 0x1f, 0x97, 0xec, 0x9d, 0xd8, 0x2c, 0xb3,
	0x32, 0x0a, 0x3d, 0x23, 0xcf, 0xae, 0x11, 0x75, 0x68, 0xe0, 0x7b, 0x0a, 0x09, 0x86, 0x28, 0x97,
	0xa0, 0x52, 0xec, 0x3a, 0x07, 0xce, 0xb0, 0x7d, 0xd4, 0x63, 0xc5, 0x28, 0x56, 0x8d, 0x62, 0xa7,
	0x65, 0x14, 0x9f, 0xae, 0x6c, 0x7e, 0xe1, 0x0a, 0x0a, 0x13, 0x1d, 0x93, 0x9d, 0x12, 0x86, 0x46,
	0x42, 0xd2, 0xad, 0x59, 0xc8, 0x80, 0x3d, 0xf0, 0x51, 0x36, 0x09, 0x82, 0x4f, 0x3e, 0xa0, 0xc9,
	0xfd, 0x76, 0x41, 0xb3, 0x36, 0x3a, 0x25, 0x4f, 0x51, 0xe8, 0x50, 0xa8, 0x38, 0x06, 0x71, 0x1b,
	0xc9, 0xdd, 0x14, 0xe9, 0x09, 0x0a, 0x7d, 0x52, 0x98, 0xaa, 0x44, 0x6f, 0xc9, 0xee, 0x0a, 0x75,
	0x03, 0xa0, 0x79, 0x24, 0x33, 0xe8, 0xd6, 0x2d, 0xa4, 0x7f, 0x0f, 0x32, 0x52, 0x2a, 0xfa, 0xcc,
	0xa3, 0x14, 0xfc, 0x1d, 0x14, 0xfa, 0xac, 0x7a, 0x4f, 0x13, 0xf2, 0x42, 0x48, 0x23, 0x52, 0x89,
	0xe1, 0xdc, 0x00, 0xbf, 0x01, 0x13, 0x26, 0x10, 0x27, 0x12, 0x65, 0x26, 0x31, 0xef, 0x36, 0x0e,
	0x9c, 0xe1, 0xde, 0xd1, 0xf1, 0x83, 0x3f, 0xfc, 0xbf, 0x6d, 0x36, 0xbb, 0xb5, 0xfa, 0xbd, 0x92,
	0x3b, 0x2a, 0xb0, 0x77, 0xa4, 0xc1, 0x1b, 0xd2, 0xbe, 0x73, 0xa5, 0x2d, 0xd2, 0xb8, 0xfc, 0x38,
	0x1b, 0x07, 0x9d, 0x47, 0x74, 0x9b, 0xb8, 0x1f, 0x2e, 0xbe, 0x74, 0x1c, 0x4a, 0xc8, 0xd6, 0xf9,
	0xf8, 0x74, 0x7a, 0x79, 0xde, 0xa9, 0xd1, 0x26, 0xa9, 0x4f, 0xa6, 0xef, 0x27, 0x1d, 0x77, 0xf0,
	0xd3, 0x21, 0xad, 0x75, 0xa9, 0xf4, 0x15, 0x69, 0x72, 0x44, 0x58, 0x6a, 0x4c, 0xec, 0x3e, 0x1b,
	0xa3, 0xd6, 0xaf, 0xbf, 0xbf, 0xdd, 0x7a, 0xbf, 0x36, 0x74, 0xfc, 0xb5, 0x44, 0xdf, 0x91, 0xc7,
	0x1a, 0x4c, 0x88, 0x26, 0x5f, 0x57, 0x5d, 0xdb, 0x54, 0xf5, 0xae, 0x06, 0x13, 0x98, 0xbc, 0xaa,
	0xb9, 0x47, 0x9a, 0xab, 0x9d, 0xe7, 0xa1, 0x8a, 0xed, 0x9a, 0x5a, 0xfe, 0xb6, 0xbd, 0x5f, 0xc4,
	0x23, 0xf6, 0xf5, 0xf5, 0x42, 0xe2, 0x75, 0x3a, 0x67, 0x42, 0x2d, 0xbd, 0xb2, 0x27, 0xa9, 0xaa,
	0x93, 0xc7, 0xb5, 0xf4, 0x50, 0x08, 0x6f, 0xd5, 0x98, 0x97, 0x1d, 0xce, 0xb7, 0xec, 0xb0, 0xe3,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xd8, 0x04, 0x94, 0x35, 0x03, 0x00, 0x00,
}
