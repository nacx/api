// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit.proto

package auditv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A system log describing something that happened in the system
type AuditLog struct {
	// Time when the audit log was generated
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Log severity (INFO, WARN, ERROR...)
	Severity string `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	// The kind of the audit log (PolicyAssigned, ServiceOrphaned, etc)
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// Audit log details
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Person who triggered the audit log, or "SYSTEM" if the log was automatically
	// triggered by the system
	TriggeredBy string `protobuf:"bytes,5,opt,name=triggered_by,json=triggeredBy,proto3" json:"triggered_by,omitempty"`
	// Key value pairs with additional information for the audit log
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
	return xxx_messageInfo_AuditLog.Unmarshal(m, b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
}
func (m *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(m, src)
}
func (m *AuditLog) XXX_Size() int {
	return xxx_messageInfo_AuditLog.Size(m)
}
func (m *AuditLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLog.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLog proto.InternalMessageInfo

func (m *AuditLog) GetCreateTime() *timestamp.Timestamp {
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

// Request to get the audit logs
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
	return xxx_messageInfo_ListAuditLogsRequest.Unmarshal(m, b)
}
func (m *ListAuditLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuditLogsRequest.Marshal(b, m, deterministic)
}
func (m *ListAuditLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuditLogsRequest.Merge(m, src)
}
func (m *ListAuditLogsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAuditLogsRequest.Size(m)
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

// The list of audit logs
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
	return xxx_messageInfo_ListAuditLogsResponse.Unmarshal(m, b)
}
func (m *ListAuditLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuditLogsResponse.Marshal(b, m, deterministic)
}
func (m *ListAuditLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuditLogsResponse.Merge(m, src)
}
func (m *ListAuditLogsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAuditLogsResponse.Size(m)
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
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x86, 0x99, 0x64, 0xb3, 0x6e, 0x6a, 0x56, 0x5d, 0x9a, 0x28, 0xc3, 0x20, 0x1a, 0x73, 0x0a,
	0x2a, 0xdd, 0x4c, 0xf6, 0x22, 0x2e, 0x7b, 0x70, 0xc1, 0xdb, 0x22, 0x32, 0x8a, 0x07, 0x2f, 0xa1,
	0x93, 0x94, 0x6d, 0xb3, 0xc9, 0xf4, 0xd8, 0x5d, 0x33, 0x30, 0x37, 0xf1, 0x0d, 0xc4, 0x47, 0xf2,
	0x11, 0x7c, 0x05, 0x1f, 0x44, 0xa6, 0x3b, 0x13, 0x34, 0x04, 0xf1, 0x56, 0xf5, 0xd7, 0x5f, 0x5f,
	0x57, 0x35, 0x05, 0xb1, 0xac, 0x56, 0x9a, 0x78, 0x69, 0x0d, 0x19, 0x36, 0x22, 0x24, 0x2b, 0x09,
	0xb9, 0x2c, 0x35, 0x0f, 0x85, 0x3a, 0x4b, 0x1f, 0x28, 0x63, 0xd4, 0x1a, 0x85, 0x2c, 0xb5, 0x90,
	0x45, 0x61, 0x48, 0x92, 0x36, 0x85, 0x0b, 0x3d, 0xe9, 0xa3, 0x6d, 0xd5, 0x67, 0x8b, 0xea, 0xa3,
	0x20, 0xbd, 0x41, 0x47, 0x72, 0x53, 0x06, 0xc3, 0xe4, 0x47, 0x0f, 0x4e, 0x5e, 0xb6, 0xac, 0x6b,
	0xa3, 0xd8, 0x05, 0xc4, 0x4b, 0x8b, 0x92, 0x70, 0xde, 0xda, 0x92, 0x68, 0x1c, 0x4d, 0xe3, 0x59,
	0xca, 0x03, 0x83, 0x77, 0x0c, 0xfe, 0xae, 0x63, 0xe4, 0x10, 0xec, 0xad, 0xc0, 0x52, 0x38, 0x71,
	0x58, 0xa3, 0xd5, 0xd4, 0x24, 0xbd, 0x71, 0x34, 0x1d, 0xe6, 0xbb, 0x9c, 0x31, 0x38, 0xba, 0xd1,
	0xc5, 0x2a, 0xe9, 0x7b, 0xdd, 0xc7, 0x2c, 0x81, 0x5b, 0x1b, 0x74, 0x4e, 0x2a, 0x4c, 0x8e, 0xbc,
	0xdc, 0xa5, 0xec, 0x31, 0x9c, 0x92, 0xd5, 0x4a, 0xa1, 0xc5, 0xd5, 0x7c, 0xd1, 0x24, 0x03, 0x5f,
	0x8e, 0x77, 0xda, 0x55, 0xc3, 0x5e, 0x03, 0x94, 0xd6, 0x94, 0x68, 0x49, 0xa3, 0x4b, 0x8e, 0xc7,
	0xfd, 0x69, 0x3c, 0xe3, 0xfc, 0xd0, 0x07, 0xf1, 0x6e, 0x3b, 0xfe, 0x66, 0xd7, 0xf0, 0xaa, 0x20,
	0xdb, 0xe4, 0x7f, 0x10, 0xd2, 0x4b, 0xb8, 0xbb, 0x57, 0x66, 0x67, 0xd0, 0xbf, 0xc1, 0xc6, 0x7f,
	0xc2, 0x30, 0x6f, 0x43, 0x36, 0x82, 0x41, 0x2d, 0xd7, 0x15, 0x6e, 0xd7, 0x0b, 0xc9, 0x8b, 0xde,
	0xf3, 0x68, 0xf2, 0x0c, 0x46, 0xd7, 0xda, 0x51, 0xf7, 0x94, 0xcb, 0xf1, 0x73, 0x85, 0x8e, 0xda,
	0x8e, 0xa5, 0xa9, 0x0a, 0xf2, 0x94, 0x41, 0x1e, 0x92, 0xc9, 0x7b, 0xb8, 0xb7, 0xe7, 0x76, 0xa5,
	0x29, 0x1c, 0xb2, 0x4b, 0x00, 0x3f, 0xf6, 0x7c, 0x6d, 0x94, 0x4b, 0x22, 0xbf, 0xd5, 0xc3, 0x7f,
	0x6f, 0x95, 0x0f, 0x65, 0x87, 0x99, 0x7d, 0x8b, 0xe0, 0xd4, 0xeb, 0x6f, 0xd1, 0xd6, 0x7a, 0x89,
	0xec, 0x4b, 0x04, 0xb7, 0xff, 0x7a, 0x89, 0x3d, 0x39, 0x4c, 0x3b, 0x34, 0x7c, 0xfa, 0xf4, 0xbf,
	0xbc, 0x61, 0xf4, 0xc9, 0xfd, 0xaf, 0x3f, 0x7f, 0x7d, 0xef, 0x9d, 0xb1, 0x3b, 0xa2, 0xce, 0x44,
	0x3b, 0xbe, 0xf0, 0x0d, 0x57, 0xe7, 0x1f, 0x32, 0xa5, 0xe9, 0x53, 0xb5, 0xe0, 0x4b, 0xb3, 0x11,
	0x5b, 0xa0, 0x36, 0x5d, 0x14, 0xce, 0xb6, 0x75, 0x8a, 0x3a, 0xbb, 0xf0, 0x41, 0x9d, 0x2d, 0x8e,
	0xfd, 0xa9, 0x9d, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x76, 0x01, 0xf0, 0x1b, 0xff, 0x02, 0x00,
	0x00,
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
	// List audit logs. If no 'count' parameter has been specified, the last 25
	// audit logs are returned
	ListAuditLogs(ctx context.Context, in *ListAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error)
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

// AuditServiceServer is the server API for AuditService service.
type AuditServiceServer interface {
	// List audit logs. If no 'count' parameter has been specified, the last 25
	// audit logs are returned
	ListAuditLogs(context.Context, *ListAuditLogsRequest) (*ListAuditLogsResponse, error)
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

var _AuditService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.audit.v1.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuditLogs",
			Handler:    _AuditService_ListAuditLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit.proto",
}
