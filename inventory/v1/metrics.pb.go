// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

package invv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request to retrieve metrics for a given service. Duration is rolling in minutes,
// as in last n minutes. Start/end are Unix timestamps in seconds. Any combination of
// duration/start/end may be provided in any order. filter unused for now
type GetMetricsRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Duration             int32    `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Start                uint64   `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  uint64   `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	Filter               string   `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMetricsRequest) Reset()         { *m = GetMetricsRequest{} }
func (m *GetMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMetricsRequest) ProtoMessage()    {}
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

func (m *GetMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricsRequest.Unmarshal(m, b)
}
func (m *GetMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricsRequest.Marshal(b, m, deterministic)
}
func (m *GetMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricsRequest.Merge(m, src)
}
func (m *GetMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMetricsRequest.Size(m)
}
func (m *GetMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricsRequest proto.InternalMessageInfo

func (m *GetMetricsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GetMetricsRequest) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *GetMetricsRequest) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *GetMetricsRequest) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *GetMetricsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Single service response row with value, and corresponding Unix timestamp in seconds
type MetricsResponse struct {
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsResponse) Reset()         { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()    {}
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}

func (m *MetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsResponse.Unmarshal(m, b)
}
func (m *MetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsResponse.Marshal(b, m, deterministic)
}
func (m *MetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsResponse.Merge(m, src)
}
func (m *MetricsResponse) XXX_Size() int {
	return xxx_messageInfo_MetricsResponse.Size(m)
}
func (m *MetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsResponse proto.InternalMessageInfo

func (m *MetricsResponse) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MetricsResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// GetMetricsResponse returns multiple response records for a given service
type GetMetricsResponse struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricsResponse      []*MetricsResponse `protobuf:"bytes,2,rep,name=metricsResponse,proto3" json:"metricsResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetMetricsResponse) Reset()         { *m = GetMetricsResponse{} }
func (m *GetMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetricsResponse) ProtoMessage()    {}
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{2}
}

func (m *GetMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetricsResponse.Unmarshal(m, b)
}
func (m *GetMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetricsResponse.Marshal(b, m, deterministic)
}
func (m *GetMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetricsResponse.Merge(m, src)
}
func (m *GetMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetricsResponse.Size(m)
}
func (m *GetMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetricsResponse proto.InternalMessageInfo

func (m *GetMetricsResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetMetricsResponse) GetMetricsResponse() []*MetricsResponse {
	if m != nil {
		return m.MetricsResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*GetMetricsRequest)(nil), "tetrate.api.inventory.v1.GetMetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "tetrate.api.inventory.v1.MetricsResponse")
	proto.RegisterType((*GetMetricsResponse)(nil), "tetrate.api.inventory.v1.GetMetricsResponse")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x99, 0xfc, 0x68, 0x9b, 0x27, 0x5a, 0x7d, 0x88, 0x2c, 0x31, 0xd5, 0xb0, 0x2a, 0x44,
	0x94, 0x5d, 0xd2, 0xe2, 0xc9, 0x9b, 0x50, 0xd6, 0x83, 0x5e, 0xb6, 0x37, 0x6f, 0xd3, 0xf4, 0x99,
	0x0c, 0xec, 0xce, 0xac, 0xb3, 0x2f, 0x03, 0x45, 0x7a, 0xf1, 0xa4, 0x22, 0x5e, 0xfc, 0xa7, 0xbc,
	0xfb, 0x2f, 0xf8, 0x87, 0x48, 0x66, 0x27, 0xe9, 0x5a, 0x29, 0x78, 0xc9, 0xed, 0xbd, 0xe1, 0xcd,
	0xe7, 0xfb, 0x99, 0x19, 0x06, 0x6e, 0x96, 0xc4, 0x56, 0xcd, 0xea, 0xa4, 0xb2, 0x86, 0x0d, 0x46,
	0x4c, 0x6c, 0x25, 0x53, 0x22, 0x2b, 0x95, 0x28, 0xed, 0x48, 0xb3, 0xb1, 0xe7, 0x89, 0x9b, 0x0e,
	0x47, 0x73, 0x63, 0xe6, 0x05, 0xa5, 0xb2, 0x52, 0xa9, 0xd4, 0xda, 0xb0, 0x64, 0x65, 0x74, 0xd8,
	0x17, 0x7f, 0x11, 0x70, 0x27, 0x23, 0x7e, 0xdb, 0xc0, 0x72, 0xfa, 0xb0, 0xa4, 0x9a, 0x31, 0x82,
	0xdd, 0x9a, 0xac, 0x53, 0x33, 0x8a, 0xc4, 0x58, 0x4c, 0x06, 0xf9, 0xba, 0xc5, 0x21, 0xec, 0x9d,
	0x2d, 0xad, 0x47, 0x44, 0x9d, 0xb1, 0x98, 0xf4, 0xf3, 0x4d, 0x8f, 0x77, 0xa1, 0x5f, 0xb3, 0xb4,
	0x1c, 0x75, 0xc7, 0x62, 0xd2, 0xcb, 0x9b, 0x06, 0x6f, 0x43, 0x97, 0xf4, 0x59, 0xd4, 0xf3, 0x6b,
	0xab, 0x12, 0xef, 0xc1, 0xce, 0x7b, 0x55, 0x30, 0xd9, 0xa8, 0xef, 0xe1, 0xa1, 0x8b, 0x8f, 0x61,
	0x7f, 0xe3, 0x51, 0x57, 0x46, 0xd7, 0xb4, 0x42, 0x3a, 0x59, 0x2c, 0xc9, 0x67, 0x89, 0xbc, 0x69,
	0x70, 0x04, 0x03, 0x56, 0x25, 0xd5, 0x2c, 0xcb, 0x2a, 0x84, 0x5d, 0x2e, 0xc4, 0x17, 0x80, 0xed,
	0x13, 0x05, 0x12, 0x42, 0x4f, 0xcb, 0x72, 0x7d, 0x1e, 0x5f, 0xe3, 0x09, 0xec, 0x97, 0x7f, 0x8f,
	0x45, 0x9d, 0x71, 0x77, 0x72, 0xe3, 0xf0, 0x69, 0x72, 0xdd, 0x75, 0x26, 0x57, 0xb8, 0xf9, 0x55,
	0xc2, 0xe1, 0xcf, 0x1e, 0xdc, 0x0a, 0x43, 0x27, 0xe1, 0xd2, 0x3e, 0x0b, 0xd8, 0xcd, 0x88, 0x73,
	0xc9, 0x84, 0xcf, 0xae, 0x47, 0xff, 0xf3, 0x0e, 0xc3, 0xe7, 0xff, 0x37, 0xdc, 0x24, 0xc7, 0xf1,
	0xa7, 0x5f, 0xbf, 0x7f, 0x74, 0x46, 0x38, 0x4c, 0xdd, 0x34, 0x0d, 0x5a, 0xe9, 0x6a, 0x77, 0xfa,
	0x31, 0x3c, 0xdf, 0x05, 0x7e, 0x15, 0xb0, 0x97, 0x11, 0x1f, 0x5b, 0x6b, 0xec, 0x36, 0x5d, 0x1e,
	0x79, 0x97, 0x03, 0xbc, 0xdf, 0x76, 0xa1, 0x55, 0x6a, 0x4b, 0xe6, 0xbb, 0x00, 0xc8, 0x88, 0xdf,
	0x48, 0x26, 0x3d, 0x3b, 0xdf, 0xa6, 0xce, 0x13, 0xaf, 0xf3, 0x10, 0x0f, 0xda, 0x3a, 0x45, 0x93,
	0xdb, 0x12, 0xfa, 0x26, 0x60, 0x90, 0x11, 0xbf, 0x26, 0x59, 0xf0, 0x62, 0x9b, 0x3e, 0x8f, 0xbd,
	0xcf, 0x03, 0x1c, 0xb5, 0x7d, 0x16, 0x3e, 0xf6, 0x52, 0xe7, 0xd5, 0x8b, 0x77, 0x47, 0x73, 0xc5,
	0x8b, 0xe5, 0x69, 0x32, 0x33, 0x65, 0x1a, 0xf8, 0xca, 0xac, 0x2b, 0xff, 0xa5, 0x37, 0x49, 0xa9,
	0x9b, 0xbe, 0x54, 0xda, 0xb9, 0xe9, 0xe9, 0x8e, 0xff, 0xda, 0x47, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x05, 0x14, 0x5f, 0xb9, 0x23, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Get average rate of requests or number of requests per second (such as QPS)
	// to a given service in a specified rolling duration, or a duration range.
	// If range provided, it overrides duration
	GetRate(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Get average error rate or number of failed requests per second (such as non 200s)
	// for a given service in the specified rolling duration, or a duration range
	GetError(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Get average request duration in seconds for a given service in the specified
	// rolling duration or a duration range. This represents latency at server-side
	// as recorded by mesh
	GetLatency(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Get health of a given service from the list of active workloads for a given service.
	// Value 1 or greater indicates healthy service. Duration ignored
	GetHealth(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) GetRate(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.MetricsService/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetError(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.MetricsService/GetError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetLatency(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.MetricsService/GetLatency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetHealth(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.MetricsService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// Get average rate of requests or number of requests per second (such as QPS)
	// to a given service in a specified rolling duration, or a duration range.
	// If range provided, it overrides duration
	GetRate(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Get average error rate or number of failed requests per second (such as non 200s)
	// for a given service in the specified rolling duration, or a duration range
	GetError(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Get average request duration in seconds for a given service in the specified
	// rolling duration or a duration range. This represents latency at server-side
	// as recorded by mesh
	GetLatency(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Get health of a given service from the list of active workloads for a given service.
	// Value 1 or greater indicates healthy service. Duration ignored
	GetHealth(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.MetricsService/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetRate(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.MetricsService/GetError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetError(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetLatency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetLatency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.MetricsService/GetLatency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetLatency(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.MetricsService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetHealth(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.inventory.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _MetricsService_GetRate_Handler,
		},
		{
			MethodName: "GetError",
			Handler:    _MetricsService_GetError_Handler,
		},
		{
			MethodName: "GetLatency",
			Handler:    _MetricsService_GetLatency_Handler,
		},
		{
			MethodName: "GetHealth",
			Handler:    _MetricsService_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics.proto",
}
