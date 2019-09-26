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
	return fileDescriptor_client_settings_80aca818d7073565, []int{0, 0}
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
	return fileDescriptor_client_settings_80aca818d7073565, []int{0}
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
	return fileDescriptor_client_settings_80aca818d7073565, []int{1}
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
	proto.RegisterFile("client_settings.proto", fileDescriptor_client_settings_80aca818d7073565)
}

var fileDescriptor_client_settings_80aca818d7073565 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0x86, 0x6b, 0x42, 0x20, 0x99, 0x00, 0x4d, 0xb7, 0xad, 0x48, 0x52, 0xa9, 0x42, 0x39, 0x71,
	0xa8, 0xd6, 0x02, 0x0e, 0x3d, 0x56, 0x0d, 0x44, 0x4d, 0x44, 0x29, 0x95, 0x63, 0x5a, 0xa9, 0x17,
	0x6b, 0xb3, 0x4c, 0xc3, 0x0a, 0xc7, 0xbb, 0x5d, 0x8f, 0x5d, 0xf9, 0x01, 0xfa, 0x9a, 0x7d, 0x96,
	0x2a, 0x6b, 0x3b, 0x50, 0x21, 0xc4, 0x6d, 0x47, 0xb3, 0xff, 0x37, 0xff, 0xfc, 0x03, 0xaf, 0x65,
	0xac, 0x30, 0xa1, 0x28, 0x45, 0x22, 0x95, 0x2c, 0x52, 0x6e, 0xac, 0x26, 0xcd, 0xf6, 0x09, 0xc9,
	0x0a, 0x42, 0x2e, 0x8c, 0xe2, 0x24, 0x25, 0x97, 0xda, 0x22, 0xcf, 0x8f, 0x06, 0x6f, 0x17, 0x5a,
	0x2f, 0x62, 0xf4, 0xdd, 0xb7, 0x79, 0xf6, 0xd3, 0xbf, 0xce, 0xac, 0x20, 0xa5, 0x93, 0x52, 0xf8,
	0xb0, 0xff, 0xdb, 0x0a, 0x63, 0xd0, 0x56, 0xe0, 0xc1, 0x7e, 0x2e, 0x62, 0x75, 0x2d, 0x08, 0xfd,
	0xfa, 0x51, 0x36, 0x86, 0x7f, 0x1b, 0xb0, 0x77, 0xea, 0xbc, 0xcc, 0x2a, 0x2b, 0xec, 0x1c, 0x5e,
	0xdd, 0x10, 0x99, 0xc8, 0xe2, 0xaf, 0x0c, 0x53, 0x8a, 0x48, 0x2d, 0x51, 0x67, 0xd4, 0xf3, 0x0e,
	0xbc, 0xc3, 0xce, 0x71, 0x9f, 0x97, 0xa3, 0x78, 0x3d, 0x8a, 0x9f, 0x55, 0x56, 0x02, 0xb6, 0x92,
	0x05, 0xa5, 0x2a, 0x2c, 0x45, 0x6c, 0x0c, 0x3b, 0x15, 0x8c, 0xac, 0xc2, 0xb4, 0xb7, 0xe1, 0x20,
	0x43, 0xfe, 0xc8, 0xa2, 0x7c, 0x12, 0x86, 0x5f, 0x03, 0x24, 0x5b, 0x04, 0x9d, 0x92, 0xe6, 0x64,
	0x6c, 0x0a, 0x2f, 0x49, 0x9a, 0x48, 0xea, 0x24, 0x41, 0x79, 0x67, 0xa9, 0xf1, 0x94, 0xa5, 0x17,
	0x24, 0xcd, 0x69, 0x29, 0xaa, 0x1d, 0x7d, 0x80, 0xdd, 0x15, 0xea, 0x16, 0xd1, 0x88, 0x58, 0xe5,
	0xd8, 0xdb, 0x74, 0x90, 0xc1, 0x03, 0xc8, 0x48, 0xeb, 0xf8, 0x9b, 0x88, 0x33, 0x0c, 0x76, 0x48,
	0x9a, 0xf3, 0xfa, 0x3f, 0x4b, 0xe1, 0x8d, 0x54, 0x56, 0x66, 0x8a, 0xa2, 0xb9, 0x45, 0x71, 0x8b,
	0x36, 0x4a, 0x31, 0x49, 0x15, 0xa9, 0x5c, 0x51, 0xd1, 0x6b, 0x1e, 0x78, 0x87, 0x7b, 0xc7, 0x27,
	0x8f, 0x6e, 0xf8, 0x7f, 0xda, 0x7c, 0x76, 0x27, 0x0d, 0xfa, 0x15, 0x77, 0x54, 0x62, 0xef, 0xb5,
	0x86, 0xef, 0xa1, 0x73, 0xaf, 0x64, 0x6d, 0x68, 0x5e, 0x7d, 0x99, 0x8d, 0xc3, 0xee, 0x33, 0xb6,
	0x0d, 0x8d, 0xcf, 0x97, 0xdf, 0xbb, 0x1e, 0x03, 0xd8, 0xba, 0x18, 0x9f, 0x4d, 0xaf, 0x2e, 0xba,
	0x1b, 0xac, 0x05, 0x9b, 0x93, 0xe9, 0xa7, 0x49, 0xb7, 0x31, 0xfc, 0xe3, 0x41, 0x7b, 0x1d, 0x2a,
	0x1b, 0x40, 0x4b, 0x10, 0xe1, 0xd2, 0x50, 0xea, 0xee, 0xd9, 0x0c, 0xd6, 0x35, 0xfb, 0x08, 0xcf,
	0x0d, 0xda, 0x88, 0x6c, 0xb1, 0xce, 0x77, 0xe3, 0xa9, 0x7c, 0x77, 0x0d, 0xda, 0xd0, 0x16, 0x75,
	0xb6, 0x7d, 0x68, 0xad, 0x0e, 0x5d, 0x44, 0x3a, 0x71, 0xb7, 0x69, 0x07, 0xdb, 0xae, 0xbe, 0x4c,
	0x46, 0xfc, 0xc7, 0xbb, 0x85, 0xa2, 0x9b, 0x6c, 0xce, 0xa5, 0x5e, 0xfa, 0x55, 0x38, 0x4a, 0xd7,
	0x2f, 0x5f, 0x18, 0xe5, 0x93, 0x94, 0xfe, 0x2a, 0x26, 0x3f, 0x3f, 0x9a, 0x6f, 0xb9, 0x61, 0x27,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x67, 0xba, 0x44, 0x2a, 0x03, 0x00, 0x00,
}
