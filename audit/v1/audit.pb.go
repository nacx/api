// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: audit.proto

package auditv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
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

// AuditLog
//
// A system log describing something that happened in the system.
type AuditLog struct {
	// Time when the audit log was generated.
	CreateTime *types.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Log severity (INFO, WARN, ERROR...).
	Severity string `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	// The kind of the audit log (PolicyAssigned, ServiceOrphaned, etc).
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// Audit log details.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Person who triggered the audit log, or "SYSTEM" if the log was automatically
	// triggered by the system.
	TriggeredBy string `protobuf:"bytes,5,opt,name=triggered_by,json=triggeredBy,proto3" json:"triggered_by,omitempty"`
	// Key value pairs with additional information for the audit log.
	Properties           map[string]string `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuditLog) Reset()         { *m = AuditLog{} }
func (m *AuditLog) String() string { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()    {}
func (*AuditLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{0}
}
func (m *AuditLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(m, src)
}
func (m *AuditLog) XXX_Size() int {
	return m.Size()
}
func (m *AuditLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLog.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLog proto.InternalMessageInfo

func (m *AuditLog) GetCreateTime() *types.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *AuditLog) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AuditLog) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *AuditLog) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AuditLog) GetTriggeredBy() string {
	if m != nil {
		return m.TriggeredBy
	}
	return ""
}

func (m *AuditLog) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

// Request to get the audit logs.
type ListAuditLogsRequest struct {
	// Number of audit logs to retrieve. By default is 25.
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAuditLogsRequest) Reset()         { *m = ListAuditLogsRequest{} }
func (m *ListAuditLogsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAuditLogsRequest) ProtoMessage()    {}
func (*ListAuditLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{1}
}
func (m *ListAuditLogsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAuditLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAuditLogsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAuditLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuditLogsRequest.Merge(m, src)
}
func (m *ListAuditLogsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListAuditLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuditLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuditLogsRequest proto.InternalMessageInfo

func (m *ListAuditLogsRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// The list of audit logs.
type ListAuditLogsResponse struct {
	AuditLogs            []*AuditLog `protobuf:"bytes,1,rep,name=audit_logs,json=auditLogs,proto3" json:"audit_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListAuditLogsResponse) Reset()         { *m = ListAuditLogsResponse{} }
func (m *ListAuditLogsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAuditLogsResponse) ProtoMessage()    {}
func (*ListAuditLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{2}
}
func (m *ListAuditLogsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAuditLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAuditLogsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAuditLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuditLogsResponse.Merge(m, src)
}
func (m *ListAuditLogsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListAuditLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuditLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuditLogsResponse proto.InternalMessageInfo

func (m *ListAuditLogsResponse) GetAuditLogs() []*AuditLog {
	if m != nil {
		return m.AuditLogs
	}
	return nil
}

func init() {
	proto.RegisterType((*AuditLog)(nil), "tetrate.api.audit.v1.AuditLog")
	proto.RegisterMapType((map[string]string)(nil), "tetrate.api.audit.v1.AuditLog.PropertiesEntry")
	proto.RegisterType((*ListAuditLogsRequest)(nil), "tetrate.api.audit.v1.ListAuditLogsRequest")
	proto.RegisterType((*ListAuditLogsResponse)(nil), "tetrate.api.audit.v1.ListAuditLogsResponse")
}

func init() { proto.RegisterFile("audit.proto", fileDescriptor_5594839dd8e38a1b) }

var fileDescriptor_5594839dd8e38a1b = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3f, 0x8b, 0x13, 0x41,
	0x14, 0x67, 0x92, 0xcb, 0x79, 0x79, 0x39, 0x35, 0x8c, 0xf1, 0x5c, 0x57, 0x8d, 0x61, 0x45, 0x08,
	0xe7, 0x31, 0x4b, 0x62, 0x23, 0x27, 0x87, 0x18, 0x38, 0x6c, 0x0e, 0x91, 0x3d, 0xb1, 0xb0, 0x09,
	0x93, 0xe4, 0xb9, 0x0e, 0x97, 0xec, 0xac, 0x33, 0xb3, 0x0b, 0xdb, 0x89, 0xa5, 0xad, 0xdf, 0xc6,
	0xca, 0xd2, 0x52, 0xf0, 0x0b, 0x48, 0xb0, 0xb1, 0xb7, 0xb3, 0x91, 0x9d, 0xcd, 0x06, 0x5d, 0x83,
	0xda, 0xbd, 0x37, 0xbf, 0x3f, 0xef, 0xcf, 0xcc, 0x40, 0x8b, 0x27, 0x33, 0x61, 0x58, 0xac, 0xa4,
	0x91, 0xb4, 0x63, 0xd0, 0x28, 0x6e, 0x90, 0xf1, 0x58, 0xb0, 0x02, 0x48, 0x07, 0xee, 0xf5, 0x50,
	0xca, 0x70, 0x8e, 0x3e, 0x8f, 0x85, 0xcf, 0xa3, 0x48, 0x1a, 0x6e, 0x84, 0x8c, 0x74, 0xa1, 0x71,
	0xaf, 0xad, 0x50, 0x9b, 0x4d, 0x92, 0x17, 0x3e, 0x2e, 0x62, 0x93, 0xad, 0xc0, 0x9b, 0x55, 0xd0,
	0x88, 0x05, 0x6a, 0xc3, 0x17, 0xf1, 0x8a, 0x70, 0x25, 0xe5, 0x73, 0x31, 0xe3, 0x06, 0xfd, 0x32,
	0x28, 0x00, 0xef, 0x47, 0x0d, 0x76, 0x1e, 0xe6, 0x1d, 0x9c, 0xc8, 0x90, 0x3e, 0x82, 0xd6, 0x54,
	0x21, 0x37, 0x38, 0xce, 0xf5, 0x0e, 0xe9, 0x91, 0x7e, 0x6b, 0xe8, 0xb2, 0xc2, 0x9c, 0x95, 0xe6,
	0xec, 0x69, 0x69, 0x3e, 0x82, 0xf7, 0xdf, 0x3e, 0xd4, 0x1b, 0x6f, 0x49, 0xad, 0x4d, 0x02, 0x28,
	0xa4, 0x39, 0x48, 0x6f, 0xc3, 0x8e, 0xc6, 0x14, 0x95, 0x30, 0x99, 0x53, 0xeb, 0x91, 0x7e, 0x73,
	0xd4, 0xcc, 0x99, 0x5b, 0x2a, 0x27, 0xae, 0x21, 0x7a, 0x03, 0xb6, 0xce, 0x44, 0x34, 0x73, 0xea,
	0x55, 0x8a, 0x3d, 0xa6, 0xb7, 0xe0, 0xdc, 0x02, 0xb5, 0xe6, 0x21, 0x3a, 0x5b, 0x55, 0x46, 0x89,
	0xd0, 0x03, 0xd8, 0x35, 0x4a, 0x84, 0x21, 0x2a, 0x9c, 0x8d, 0x27, 0x99, 0xd3, 0xa8, 0x32, 0x5b,
	0x6b, 0x78, 0x94, 0xd1, 0xc7, 0x00, 0xb1, 0x92, 0x31, 0x2a, 0x23, 0x50, 0x3b, 0xdb, 0xbd, 0x7a,
	0xbf, 0x35, 0x64, 0x6c, 0xd3, 0x75, 0xb0, 0x72, 0x2b, 0xec, 0xc9, 0x5a, 0x70, 0x1c, 0x19, 0x95,
	0x05, 0xbf, 0x38, 0xb8, 0x47, 0x70, 0xb1, 0x02, 0xd3, 0x36, 0xd4, 0xcf, 0x30, 0xb3, 0xcb, 0x6b,
	0x06, 0x79, 0x48, 0x3b, 0xd0, 0x48, 0xf9, 0x3c, 0xc1, 0x62, 0x15, 0x41, 0x91, 0x1c, 0xd6, 0xee,
	0x11, 0xef, 0x00, 0x3a, 0x27, 0x42, 0x9b, 0xb2, 0x94, 0x0e, 0xf0, 0x55, 0x82, 0xda, 0xe4, 0x8a,
	0xa9, 0x4c, 0x22, 0x63, 0x5d, 0x1a, 0x41, 0x91, 0x78, 0xcf, 0xe0, 0x72, 0x85, 0xad, 0x63, 0x19,
	0x69, 0xa4, 0x47, 0x00, 0xb6, 0xed, 0xf1, 0x5c, 0x86, 0xda, 0x21, 0x76, 0xaa, 0xee, 0xdf, 0xa7,
	0x0a, 0x9a, 0xbc, 0xb4, 0x19, 0x7e, 0x27, 0xb0, 0x6b, 0xcf, 0x4f, 0x51, 0xa5, 0x62, 0x8a, 0xf4,
	0x35, 0x81, 0xf3, 0xbf, 0x55, 0xa2, 0xfb, 0x9b, 0xdd, 0x36, 0x35, 0xef, 0xde, 0xf9, 0x2f, 0x6e,
	0xd1, 0xba, 0xb7, 0xf7, 0xe6, 0xf3, 0xd7, 0x77, 0xb5, 0x36, 0xbd, 0xe0, 0xa7, 0x03, 0x3f, 0x6f,
	0xdf, 0xb7, 0x02, 0xca, 0x61, 0xf7, 0x14, 0xa3, 0xd9, 0xfa, 0x69, 0xfe, 0x63, 0x1c, 0x77, 0xef,
	0x8f, 0x57, 0x7a, 0x9c, 0xff, 0x0f, 0xef, 0xaa, 0xf5, 0xbf, 0xe4, 0x55, 0xfc, 0x0f, 0xc9, 0xfe,
	0xe8, 0xc1, 0xc7, 0x65, 0x97, 0x7c, 0x5a, 0x76, 0xc9, 0x97, 0x65, 0x97, 0x3c, 0x1f, 0x84, 0xc2,
	0xbc, 0x4c, 0x26, 0x6c, 0x2a, 0x17, 0xfe, 0xaa, 0x94, 0x90, 0x65, 0x54, 0xfc, 0xc9, 0x5c, 0xe8,
	0xa7, 0x83, 0xfb, 0x36, 0x48, 0x07, 0x93, 0x6d, 0x5b, 0xeb, 0xee, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x98, 0xf6, 0x54, 0xfe, 0xdc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuditServiceClient is the client API for AuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditServiceClient interface {
	// List audit logs. If no 'count' parameter has been specified, the last 25 audit logs are
	// returned.
	ListAuditLogs(ctx context.Context, in *ListAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error)
	// Send a log to the audit system.
	SendAuditLog(ctx context.Context, in *AuditLog, opts ...grpc.CallOption) (*types.Empty, error)
}

type auditServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuditServiceClient(cc *grpc.ClientConn) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) ListAuditLogs(ctx context.Context, in *ListAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error) {
	out := new(ListAuditLogsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.audit.v1.AuditService/ListAuditLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) SendAuditLog(ctx context.Context, in *AuditLog, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/tetrate.api.audit.v1.AuditService/SendAuditLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServiceServer is the server API for AuditService service.
type AuditServiceServer interface {
	// List audit logs. If no 'count' parameter has been specified, the last 25 audit logs are
	// returned.
	ListAuditLogs(context.Context, *ListAuditLogsRequest) (*ListAuditLogsResponse, error)
	// Send a log to the audit system.
	SendAuditLog(context.Context, *AuditLog) (*types.Empty, error)
}

// UnimplementedAuditServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuditServiceServer struct {
}

func (*UnimplementedAuditServiceServer) ListAuditLogs(ctx context.Context, req *ListAuditLogsRequest) (*ListAuditLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuditLogs not implemented")
}
func (*UnimplementedAuditServiceServer) SendAuditLog(ctx context.Context, req *AuditLog) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAuditLog not implemented")
}

func RegisterAuditServiceServer(s *grpc.Server, srv AuditServiceServer) {
	s.RegisterService(&_AuditService_serviceDesc, srv)
}

func _AuditService_ListAuditLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuditLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).ListAuditLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.audit.v1.AuditService/ListAuditLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).ListAuditLogs(ctx, req.(*ListAuditLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_SendAuditLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).SendAuditLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.audit.v1.AuditService/SendAuditLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).SendAuditLog(ctx, req.(*AuditLog))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuditService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.audit.v1.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuditLogs",
			Handler:    _AuditService_ListAuditLogs_Handler,
		},
		{
			MethodName: "SendAuditLog",
			Handler:    _AuditService_SendAuditLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit.proto",
}

func (m *AuditLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuditLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuditLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Properties) > 0 {
		for k := range m.Properties {
			v := m.Properties[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintAudit(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintAudit(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintAudit(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TriggeredBy) > 0 {
		i -= len(m.TriggeredBy)
		copy(dAtA[i:], m.TriggeredBy)
		i = encodeVarintAudit(dAtA, i, uint64(len(m.TriggeredBy)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintAudit(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintAudit(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Severity) > 0 {
		i -= len(m.Severity)
		copy(dAtA[i:], m.Severity)
		i = encodeVarintAudit(dAtA, i, uint64(len(m.Severity)))
		i--
		dAtA[i] = 0x12
	}
	if m.CreateTime != nil {
		{
			size, err := m.CreateTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAudit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListAuditLogsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAuditLogsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAuditLogsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Count != 0 {
		i = encodeVarintAudit(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListAuditLogsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAuditLogsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAuditLogsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AuditLogs) > 0 {
		for iNdEx := len(m.AuditLogs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuditLogs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAudit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAudit(dAtA []byte, offset int, v uint64) int {
	offset -= sovAudit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuditLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreateTime != nil {
		l = m.CreateTime.Size()
		n += 1 + l + sovAudit(uint64(l))
	}
	l = len(m.Severity)
	if l > 0 {
		n += 1 + l + sovAudit(uint64(l))
	}
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovAudit(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovAudit(uint64(l))
	}
	l = len(m.TriggeredBy)
	if l > 0 {
		n += 1 + l + sovAudit(uint64(l))
	}
	if len(m.Properties) > 0 {
		for k, v := range m.Properties {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovAudit(uint64(len(k))) + 1 + len(v) + sovAudit(uint64(len(v)))
			n += mapEntrySize + 1 + sovAudit(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListAuditLogsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovAudit(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListAuditLogsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuditLogs) > 0 {
		for _, e := range m.AuditLogs {
			l = e.Size()
			n += 1 + l + sovAudit(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAudit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAudit(x uint64) (n int) {
	return sovAudit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuditLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudit
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
			return fmt.Errorf("proto: AuditLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuditLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateTime == nil {
				m.CreateTime = &types.Timestamp{}
			}
			if err := m.CreateTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Severity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggeredBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TriggeredBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Properties == nil {
				m.Properties = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAudit
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAudit
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthAudit
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthAudit
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAudit
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthAudit
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthAudit
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipAudit(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthAudit
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Properties[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAudit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAudit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAudit
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
func (m *ListAuditLogsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudit
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
			return fmt.Errorf("proto: ListAuditLogsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAuditLogsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAudit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAudit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAudit
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
func (m *ListAuditLogsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudit
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
			return fmt.Errorf("proto: ListAuditLogsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAuditLogsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuditLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuditLogs = append(m.AuditLogs, &AuditLog{})
			if err := m.AuditLogs[len(m.AuditLogs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAudit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAudit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAudit
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
func skipAudit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAudit
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
					return 0, ErrIntOverflowAudit
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
					return 0, ErrIntOverflowAudit
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
				return 0, ErrInvalidLengthAudit
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAudit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAudit
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
				next, err := skipAudit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAudit
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
	ErrInvalidLengthAudit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAudit   = fmt.Errorf("proto: integer overflow")
)
