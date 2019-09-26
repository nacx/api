// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_settings.proto

package v1 // import "github.com/tetrateio/tetrate/api/tcc/core/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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
	return fileDescriptor_client_settings_0cedc4208e615748, []int{0, 0}
}

// All the timeouts, retries, circuit breakers etc. that consumers
// want to set to gaurd against failures of their dependencies. These
// roughly translate to pieces of istio virtual service, destination
// rules, etc. At runtime, for a given namespace, for the namespace's
// dependencies, we combine the client specified reliability settings
// with the virtual service/dest rule for the dependencies to produce
// an updated virtual service/dest rule per dependent service. All
// these client customized virtual services/dest rules will be private
// to the namespace i.e. will have an exportTo setting '.' . Doing so
// will make Pilot apply these customizations only to the namespace
// concerned without leaking it to all other namespaces.
type ClientSettings struct {
	// Timeout for HTTP requests.
	HttpRequestTimeout *duration.Duration `protobuf:"bytes,1,opt,name=http_request_timeout,json=httpRequestTimeout,proto3" json:"http_request_timeout,omitempty"`
	// Retry policy for HTTP requests.
	HttpRetries *HTTPRetry `protobuf:"bytes,2,opt,name=http_retries,json=httpRetries,proto3" json:"http_retries,omitempty"`
	// These two settings will go to dest rule
	// TCP connection timeout.
	TcpConnectTimeout *duration.Duration `protobuf:"bytes,3,opt,name=tcp_connect_timeout,json=tcpConnectTimeout,proto3" json:"tcp_connect_timeout,omitempty"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *wrappers.BoolValue `protobuf:"bytes,4,opt,name=tcp_keepalive,json=tcpKeepalive,proto3" json:"tcp_keepalive,omitempty"`
	// the sensitivity levels will translate to specific values of
	// dest rule outlier detection.
	CircuitBreakerSensitivity ClientSettings_Sensitivity `protobuf:"varint,5,opt,name=circuit_breaker_sensitivity,json=circuitBreakerSensitivity,proto3,enum=tetrate.api.tcc.core.v1.ClientSettings_Sensitivity" json:"circuit_breaker_sensitivity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                   `json:"-"`
	XXX_unrecognized          []byte                     `json:"-"`
	XXX_sizecache             int32                      `json:"-"`
}

func (m *ClientSettings) Reset()         { *m = ClientSettings{} }
func (m *ClientSettings) String() string { return proto.CompactTextString(m) }
func (*ClientSettings) ProtoMessage()    {}
func (*ClientSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_settings_0cedc4208e615748, []int{0}
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
	// REQUIRED. Number of retries for a given request. The interval
	// between retries will be determined automatically (25ms+). Actual
	// number of retries attempted depends on the httpReqTimeout.
	Attempts int32 `protobuf:"varint,1,opt,name=attempts,proto3" json:"attempts,omitempty"`
	// Timeout per retry attempt for a given request. format: 1h/1m/1s/1ms. MUST BE >=1ms.
	PerTryTimeout *duration.Duration `protobuf:"bytes,2,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
	// Specifies the conditions under which retry takes place.
	// One or more policies can be specified using a ‘,’ delimited list.
	// See the [supported
	// policies](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#x-envoy-retry-on)
	// and
	// [here](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#x-envoy-retry-grpc-on)
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
	return fileDescriptor_client_settings_0cedc4208e615748, []int{1}
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
	proto.RegisterFile("client_settings.proto", fileDescriptor_client_settings_0cedc4208e615748)
}

var fileDescriptor_client_settings_0cedc4208e615748 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xd2, 0xb4, 0xc9, 0xa4, 0x2d, 0x61, 0x01, 0x91, 0x04, 0x09, 0x55, 0x39, 0xf5,
	0x80, 0xd6, 0x6a, 0x7b, 0xe0, 0x88, 0x48, 0x1b, 0x91, 0xa8, 0x94, 0x22, 0xc7, 0x05, 0x89, 0x8b,
	0xb5, 0x59, 0x86, 0x74, 0x55, 0xc7, 0xbb, 0xac, 0xc7, 0x41, 0x7e, 0x00, 0x5e, 0x93, 0x67, 0x41,
	0x5e, 0xdb, 0x69, 0x51, 0x55, 0xf5, 0xb6, 0xa3, 0x9d, 0xff, 0xfb, 0x67, 0xfe, 0x81, 0x97, 0x32,
	0x56, 0x98, 0x50, 0x94, 0x22, 0x91, 0x4a, 0x96, 0x29, 0x37, 0x56, 0x93, 0x66, 0xaf, 0x08, 0xc9,
	0x0a, 0x42, 0x2e, 0x8c, 0xe2, 0x24, 0x25, 0x97, 0xda, 0x22, 0x5f, 0x1f, 0x0d, 0xdf, 0x2c, 0xb5,
	0x5e, 0xc6, 0xe8, 0xbb, 0xb6, 0x45, 0xf6, 0xd3, 0xff, 0x91, 0x59, 0x41, 0x4a, 0x27, 0xa5, 0xf0,
	0xfe, 0xff, 0x6f, 0x2b, 0x8c, 0x41, 0x5b, 0x81, 0x47, 0x7f, 0x9b, 0xb0, 0x7f, 0xea, 0x2c, 0xe7,
	0x95, 0x23, 0x3b, 0x87, 0x17, 0xd7, 0x44, 0x26, 0xb2, 0xf8, 0x2b, 0xc3, 0x94, 0x22, 0x52, 0x2b,
	0xd4, 0x19, 0xf5, 0xbd, 0x03, 0xef, 0xb0, 0x7b, 0x3c, 0xe0, 0x25, 0x91, 0xd7, 0x44, 0x7e, 0x56,
	0x39, 0x06, 0xac, 0x90, 0x05, 0xa5, 0x2a, 0x2c, 0x45, 0x6c, 0x02, 0xbb, 0x15, 0x8c, 0xac, 0xc2,
	0xb4, 0xdf, 0x70, 0x90, 0x11, 0x7f, 0x60, 0x1f, 0x3e, 0x0d, 0xc3, 0x2f, 0x01, 0x92, 0xcd, 0x83,
	0x6e, 0x49, 0x73, 0x32, 0x36, 0x83, 0xe7, 0x24, 0x4d, 0x24, 0x75, 0x92, 0xa0, 0xbc, 0x1d, 0xa9,
	0xf9, 0xd8, 0x48, 0xcf, 0x48, 0x9a, 0xd3, 0x52, 0x54, 0x4f, 0xf4, 0x1e, 0xf6, 0x0a, 0xd4, 0x0d,
	0xa2, 0x11, 0xb1, 0x5a, 0x63, 0x7f, 0xcb, 0x41, 0x86, 0xf7, 0x20, 0x63, 0xad, 0xe3, 0xaf, 0x22,
	0xce, 0x30, 0xd8, 0x25, 0x69, 0xce, 0xeb, 0x7e, 0x96, 0xc2, 0x6b, 0xa9, 0xac, 0xcc, 0x14, 0x45,
	0x0b, 0x8b, 0xe2, 0x06, 0x6d, 0x94, 0x62, 0x92, 0x2a, 0x52, 0x6b, 0x45, 0x79, 0xbf, 0x75, 0xe0,
	0x1d, 0xee, 0x1f, 0x9f, 0x3c, 0xb8, 0xe1, 0xff, 0x69, 0xf3, 0xf9, 0xad, 0x34, 0x18, 0x54, 0xdc,
	0x71, 0x89, 0xbd, 0xf3, 0x35, 0x7a, 0x07, 0xdd, 0x3b, 0x25, 0xeb, 0x40, 0xeb, 0xea, 0xf3, 0x7c,
	0x12, 0xf6, 0x9e, 0xb0, 0x1d, 0x68, 0x7e, 0xba, 0xfc, 0xd6, 0xf3, 0x18, 0xc0, 0xf6, 0xc5, 0xe4,
	0x6c, 0x76, 0x75, 0xd1, 0x6b, 0xb0, 0x36, 0x6c, 0x4d, 0x67, 0x1f, 0xa7, 0xbd, 0xe6, 0xe8, 0x8f,
	0x07, 0x9d, 0x4d, 0xa8, 0x6c, 0x08, 0x6d, 0x41, 0x84, 0x2b, 0x43, 0xa9, 0xbb, 0x67, 0x2b, 0xd8,
	0xd4, 0xec, 0x03, 0x3c, 0x35, 0x68, 0x23, 0xb2, 0xf9, 0x26, 0xdf, 0xc6, 0x63, 0xf9, 0xee, 0x19,
	0xb4, 0xa1, 0xcd, 0xeb, 0x6c, 0x07, 0xd0, 0x2e, 0x0e, 0x9d, 0x47, 0x3a, 0x71, 0xb7, 0xe9, 0x04,
	0x3b, 0xae, 0xbe, 0x4c, 0xc6, 0xfc, 0xfb, 0xdb, 0xa5, 0xa2, 0xeb, 0x6c, 0xc1, 0xa5, 0x5e, 0xf9,
	0x55, 0x38, 0x4a, 0xd7, 0x2f, 0x5f, 0x18, 0xe5, 0x93, 0x94, 0x7e, 0x11, 0x93, 0xbf, 0x3e, 0x5a,
	0x6c, 0x3b, 0xb3, 0x93, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x50, 0xa8, 0x60, 0x11, 0x03,
	0x00, 0x00,
}
