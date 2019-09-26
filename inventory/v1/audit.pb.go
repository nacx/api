// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit.proto

package invv1 // import "github.com/tetrateio/tetrate/api/inventory/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_audit_eeac6809193aade9, []int{0}
}
func (m *AuditLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLog.Unmarshal(m, b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
}
func (dst *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(dst, src)
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
	return fileDescriptor_audit_eeac6809193aade9, []int{1}
}
func (m *ListAuditLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuditLogsRequest.Unmarshal(m, b)
}
func (m *ListAuditLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuditLogsRequest.Marshal(b, m, deterministic)
}
func (dst *ListAuditLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuditLogsRequest.Merge(dst, src)
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

// Request to get the audit logs for a given service
type ListServiceAuditLogsRequest struct {
	// The service to get the audit logs for
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Number of audit logs to retrieve. By default is 25.
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServiceAuditLogsRequest) Reset()         { *m = ListServiceAuditLogsRequest{} }
func (m *ListServiceAuditLogsRequest) String() string { return proto.CompactTextString(m) }
func (*ListServiceAuditLogsRequest) ProtoMessage()    {}
func (*ListServiceAuditLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_audit_eeac6809193aade9, []int{2}
}
func (m *ListServiceAuditLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceAuditLogsRequest.Unmarshal(m, b)
}
func (m *ListServiceAuditLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceAuditLogsRequest.Marshal(b, m, deterministic)
}
func (dst *ListServiceAuditLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceAuditLogsRequest.Merge(dst, src)
}
func (m *ListServiceAuditLogsRequest) XXX_Size() int {
	return xxx_messageInfo_ListServiceAuditLogsRequest.Size(m)
}
func (m *ListServiceAuditLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceAuditLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceAuditLogsRequest proto.InternalMessageInfo

func (m *ListServiceAuditLogsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ListServiceAuditLogsRequest) GetCount() int32 {
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
	return fileDescriptor_audit_eeac6809193aade9, []int{3}
}
func (m *ListAuditLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuditLogsResponse.Unmarshal(m, b)
}
func (m *ListAuditLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuditLogsResponse.Marshal(b, m, deterministic)
}
func (dst *ListAuditLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuditLogsResponse.Merge(dst, src)
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
	proto.RegisterType((*AuditLog)(nil), "tetrate.api.inventory.v1.AuditLog")
	proto.RegisterMapType((map[string]string)(nil), "tetrate.api.inventory.v1.AuditLog.PropertiesEntry")
	proto.RegisterType((*ListAuditLogsRequest)(nil), "tetrate.api.inventory.v1.ListAuditLogsRequest")
	proto.RegisterType((*ListServiceAuditLogsRequest)(nil), "tetrate.api.inventory.v1.ListServiceAuditLogsRequest")
	proto.RegisterType((*ListAuditLogsResponse)(nil), "tetrate.api.inventory.v1.ListAuditLogsResponse")
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
	// List audit logs for the given service. If no 'count' parameter has been specified,
	// the last 25 audit logs are returned
	ListServiceAuditLogs(ctx context.Context, in *ListServiceAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error)
}

type auditServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuditServiceClient(cc *grpc.ClientConn) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) ListAuditLogs(ctx context.Context, in *ListAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error) {
	out := new(ListAuditLogsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.AuditService/ListAuditLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) ListServiceAuditLogs(ctx context.Context, in *ListServiceAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error) {
	out := new(ListAuditLogsResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.AuditService/ListServiceAuditLogs", in, out, opts...)
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
	// List audit logs for the given service. If no 'count' parameter has been specified,
	// the last 25 audit logs are returned
	ListServiceAuditLogs(context.Context, *ListServiceAuditLogsRequest) (*ListAuditLogsResponse, error)
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
		FullMethod: "/tetrate.api.inventory.v1.AuditService/ListAuditLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).ListAuditLogs(ctx, req.(*ListAuditLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_ListServiceAuditLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceAuditLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).ListServiceAuditLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.AuditService/ListServiceAuditLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).ListServiceAuditLogs(ctx, req.(*ListServiceAuditLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuditService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.inventory.v1.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuditLogs",
			Handler:    _AuditService_ListAuditLogs_Handler,
		},
		{
			MethodName: "ListServiceAuditLogs",
			Handler:    _AuditService_ListServiceAuditLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit.proto",
}

func init() { proto.RegisterFile("audit.proto", fileDescriptor_audit_eeac6809193aade9) }

var fileDescriptor_audit_eeac6809193aade9 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0x9d, 0xa6, 0x6d, 0xae, 0xfb, 0x7d, 0x54, 0xa3, 0x82, 0x46, 0x06, 0x89, 0xe0, 0x55,
	0x16, 0x68, 0x46, 0x49, 0x55, 0x09, 0x51, 0xb1, 0x68, 0x25, 0x76, 0x45, 0x42, 0x86, 0x55, 0x37,
	0xd1, 0x24, 0xb9, 0x98, 0x51, 0x93, 0x19, 0x33, 0x33, 0xb6, 0x64, 0x21, 0x36, 0xac, 0xd8, 0xf3,
	0x02, 0x3c, 0x08, 0x7b, 0x1e, 0x80, 0x57, 0xe0, 0x41, 0x90, 0xc7, 0x76, 0x68, 0x03, 0xe1, 0x67,
	0x77, 0xcf, 0xfd, 0x39, 0xe7, 0xde, 0x63, 0x0f, 0x44, 0xa2, 0x58, 0x48, 0xc7, 0x72, 0xa3, 0x9d,
	0x26, 0xd4, 0xa1, 0x33, 0xc2, 0x21, 0x13, 0xb9, 0x64, 0x52, 0x95, 0xa8, 0x9c, 0x36, 0x15, 0x2b,
	0xc7, 0xf1, 0xbd, 0x4c, 0xeb, 0x6c, 0x89, 0x5c, 0xe4, 0x92, 0x0b, 0xa5, 0xb4, 0x13, 0x4e, 0x6a,
	0x65, 0x9b, 0xb9, 0xf8, 0x7e, 0x5b, 0xf5, 0x68, 0x56, 0xbc, 0xe2, 0x4e, 0xae, 0xd0, 0x3a, 0xb1,
	0xca, 0x9b, 0x86, 0xe4, 0x4b, 0x08, 0xfb, 0x67, 0xb5, 0xd0, 0x85, 0xce, 0xc8, 0x29, 0x44, 0x73,
	0x83, 0xc2, 0xe1, 0xb4, 0x6e, 0xa3, 0xc1, 0x30, 0x18, 0x45, 0x93, 0x98, 0x35, 0x1c, 0xac, 0xe3,
	0x60, 0x2f, 0x3b, 0x8e, 0x14, 0x9a, 0xf6, 0x3a, 0x41, 0x62, 0xd8, 0xb7, 0x58, 0xa2, 0x91, 0xae,
	0xa2, 0xe1, 0x30, 0x18, 0x0d, 0xd2, 0x35, 0x26, 0x04, 0x76, 0xae, 0xa4, 0x5a, 0xd0, 0x9e, 0xcf,
	0xfb, 0x98, 0x50, 0xd8, 0x5b, 0xa1, 0xb5, 0x22, 0x43, 0xba, 0xe3, 0xd3, 0x1d, 0x24, 0x0f, 0xe0,
	0xc0, 0x19, 0x99, 0x65, 0x68, 0x70, 0x31, 0x9d, 0x55, 0xb4, 0xef, 0xcb, 0xd1, 0x3a, 0x77, 0x5e,
	0x91, 0x14, 0x20, 0x37, 0x3a, 0x47, 0xe3, 0x24, 0x5a, 0xba, 0x3b, 0xec, 0x8d, 0xa2, 0xc9, 0x84,
	0x6d, 0x33, 0x89, 0x75, 0x17, 0xb2, 0xe7, 0xeb, 0xa1, 0xa7, 0xca, 0x99, 0x2a, 0xbd, 0xc6, 0x12,
	0x3f, 0x81, 0x5b, 0x1b, 0x65, 0x72, 0x08, 0xbd, 0x2b, 0xac, 0xbc, 0x11, 0x83, 0xb4, 0x0e, 0xc9,
	0x11, 0xf4, 0x4b, 0xb1, 0x2c, 0xb0, 0x3d, 0xb1, 0x01, 0x8f, 0xc3, 0x47, 0x41, 0xf2, 0x10, 0x8e,
	0x2e, 0xa4, 0x75, 0x9d, 0x94, 0x4d, 0xf1, 0x4d, 0x81, 0xd6, 0xd5, 0x13, 0x73, 0x5d, 0x28, 0xe7,
	0x59, 0xfa, 0x69, 0x03, 0x92, 0x67, 0x70, 0xb7, 0xee, 0x7e, 0x81, 0xa6, 0x94, 0x73, 0xfc, 0x69,
	0x88, 0xc2, 0x9e, 0x6d, 0x4a, 0xad, 0x78, 0x07, 0x7f, 0xd0, 0x85, 0xd7, 0xe9, 0x2e, 0xe1, 0xf6,
	0x86, 0xb8, 0xcd, 0xb5, 0xb2, 0x48, 0xce, 0x00, 0xfc, 0x7f, 0x34, 0x5d, 0xea, 0xcc, 0xd2, 0xc0,
	0x1b, 0x95, 0xfc, 0xd9, 0xa8, 0x74, 0x20, 0x3a, 0xaa, 0xc9, 0xe7, 0x10, 0x0e, 0x7c, 0xbe, 0x5d,
	0x96, 0x7c, 0x08, 0xe0, 0xbf, 0x1b, 0x6a, 0x84, 0x6d, 0x67, 0xfc, 0x95, 0x27, 0x31, 0xff, 0xeb,
	0xfe, 0xe6, 0x8c, 0xe4, 0xce, 0xfb, 0xaf, 0xdf, 0x3e, 0x86, 0x87, 0xe4, 0x7f, 0x5e, 0x8e, 0x79,
	0x7d, 0x0a, 0xf7, 0xfb, 0x91, 0x4f, 0x41, 0xe3, 0xfa, 0xa6, 0x8f, 0xe4, 0xe4, 0xf7, 0x0a, 0x5b,
	0x7c, 0xff, 0xf7, 0xc5, 0x86, 0x7e, 0xb1, 0x98, 0xd0, 0x9b, 0x8b, 0xf1, 0xb7, 0xed, 0xf7, 0x7a,
	0x77, 0x7e, 0x72, 0x79, 0x9c, 0x49, 0xf7, 0xba, 0x98, 0xb1, 0xb9, 0x5e, 0xf1, 0x96, 0x5e, 0xea,
	0x2e, 0xf2, 0x0f, 0x77, 0x2d, 0xc4, 0xcb, 0xf1, 0xa9, 0x54, 0x65, 0x39, 0x9e, 0xed, 0xfa, 0xe7,
	0x76, 0xfc, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x11, 0x47, 0xb8, 0xf4, 0x07, 0x04, 0x00, 0x00,
}
