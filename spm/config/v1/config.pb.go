// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ConfigurationRequest struct {
	// Logic name of this cluster,
	// in case the remote configuration center implementation support
	// configuration management for multiple clusters.
	ClusterName          string   `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigurationRequest) Reset()         { *m = ConfigurationRequest{} }
func (m *ConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigurationRequest) ProtoMessage()    {}
func (*ConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}
func (m *ConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigurationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationRequest.Merge(m, src)
}
func (m *ConfigurationRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationRequest proto.InternalMessageInfo

func (m *ConfigurationRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type ConfigurationResponse struct {
	// Include all config items.
	// All config name should be not empty,
	// the name is composed by "module name"."provider name"."item name".
	// Each watcher implementor provides this, and it will be notified when the value changed.
	//
	// If the config center wants to set the value to NULL or empty,
	// must set the name with empty value explicitly.
	ConfigTable          []*Config `protobuf:"bytes,1,rep,name=configTable,proto3" json:"configTable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ConfigurationResponse) Reset()         { *m = ConfigurationResponse{} }
func (m *ConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigurationResponse) ProtoMessage()    {}
func (*ConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}
func (m *ConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigurationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationResponse.Merge(m, src)
}
func (m *ConfigurationResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationResponse proto.InternalMessageInfo

func (m *ConfigurationResponse) GetConfigTable() []*Config {
	if m != nil {
		return m.ConfigTable
	}
	return nil
}

type Config struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigurationRequest)(nil), "ConfigurationRequest")
	proto.RegisterType((*ConfigurationResponse)(nil), "ConfigurationResponse")
	proto.RegisterType((*Config)(nil), "Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x7f, 0xff, 0x94, 0x22, 0x1c, 0x26, 0xab, 0x45, 0x15, 0x43, 0x54, 0x65, 0x2a, 0xcb,
	0x8d, 0x08, 0x12, 0x30, 0x97, 0x1d, 0x50, 0x60, 0x62, 0xbb, 0xb1, 0x2e, 0xa9, 0x55, 0x27, 0x36,
	0xb6, 0x13, 0xc4, 0x1b, 0x32, 0xf2, 0x08, 0x28, 0x4f, 0x82, 0x48, 0x84, 0x88, 0xaa, 0x6e, 0xf7,
	0x7c, 0xf6, 0xf1, 0x39, 0xd6, 0xe5, 0x27, 0xd2, 0xd4, 0x2f, 0xaa, 0x04, 0xeb, 0x4c, 0x30, 0xc9,
	0x0d, 0x9f, 0xdd, 0xf6, 0xba, 0x71, 0x18, 0x94, 0xa9, 0x73, 0x7a, 0x6d, 0xc8, 0x07, 0xb1, 0xe4,
	0x91, 0xd4, 0x8d, 0x0f, 0xe4, 0xee, 0xb0, 0xa2, 0x05, 0x5b, 0xb2, 0xd5, 0x71, 0x3e, 0x46, 0xc9,
	0x9a, 0xcf, 0x77, 0x9c, 0xde, 0x9a, 0xda, 0x93, 0x38, 0xe7, 0xd1, 0x10, 0xf1, 0x84, 0x85, 0xfe,
	0xb1, 0x1e, 0xac, 0xa2, 0xec, 0x08, 0x86, 0xcb, 0xf9, 0xf8, 0x2c, 0xc9, 0xf8, 0x74, 0xc0, 0x42,
	0xf0, 0x49, 0xfd, 0x17, 0xd4, 0xcf, 0x62, 0xc6, 0x0f, 0x5b, 0xd4, 0x0d, 0x2d, 0xfe, 0xf7, 0x70,
	0x10, 0xd9, 0xfd, 0x4e, 0xe3, 0x47, 0x72, 0xad, 0x92, 0x24, 0xae, 0xf9, 0x44, 0xa2, 0xd6, 0x62,
	0x0e, 0xfb, 0x3e, 0x74, 0x76, 0x0a, 0x7b, 0xdb, 0x26, 0xff, 0xd6, 0xed, 0x47, 0x17, 0xb3, 0xcf,
	0x2e, 0x66, 0x5f, 0x5d, 0xcc, 0xf8, 0x95, 0x71, 0x25, 0xa0, 0x45, 0xb9, 0x21, 0xf0, 0xdb, 0xf7,
	0x37, 0xd4, 0x5b, 0x55, 0x97, 0x60, 0xd0, 0x82, 0x27, 0xd7, 0x92, 0x03, 0x39, 0x7e, 0xa7, 0x87,
	0x4a, 0xd2, 0x03, 0x7b, 0x86, 0x52, 0x85, 0x4d, 0x53, 0x80, 0x34, 0x55, 0x1a, 0x28, 0x38, 0x0c,
	0xa4, 0xcc, 0xef, 0x94, 0xa2, 0x55, 0xa9, 0xb7, 0x55, 0x3a, 0xb8, 0xd3, 0xf6, 0xa2, 0x98, 0xf6,
	0x1b, 0xb8, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x97, 0xde, 0xee, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationServiceClient is the client API for ConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationServiceClient interface {
	Call(ctx context.Context, in *ConfigurationRequest, opts ...grpc.CallOption) (*ConfigurationResponse, error)
}

type configurationServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationServiceClient(cc *grpc.ClientConn) ConfigurationServiceClient {
	return &configurationServiceClient{cc}
}

func (c *configurationServiceClient) Call(ctx context.Context, in *ConfigurationRequest, opts ...grpc.CallOption) (*ConfigurationResponse, error) {
	out := new(ConfigurationResponse)
	err := c.cc.Invoke(ctx, "/ConfigurationService/call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServiceServer is the server API for ConfigurationService service.
type ConfigurationServiceServer interface {
	Call(context.Context, *ConfigurationRequest) (*ConfigurationResponse, error)
}

// UnimplementedConfigurationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationServiceServer struct {
}

func (*UnimplementedConfigurationServiceServer) Call(ctx context.Context, req *ConfigurationRequest) (*ConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterConfigurationServiceServer(s *grpc.Server, srv ConfigurationServiceServer) {
	s.RegisterService(&_ConfigurationService_serviceDesc, srv)
}

func _ConfigurationService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigurationService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).Call(ctx, req.(*ConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigurationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ConfigurationService",
	HandlerType: (*ConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "call",
			Handler:    _ConfigurationService_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}

func (m *ConfigurationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigurationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigurationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigurationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigurationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigurationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ConfigTable) > 0 {
		for iNdEx := len(m.ConfigTable) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConfigTable[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigurationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigurationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ConfigTable) > 0 {
		for _, e := range m.ConfigTable {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigurationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ConfigurationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigurationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *ConfigurationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ConfigurationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigurationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigTable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigTable = append(m.ConfigTable, &Config{})
			if err := m.ConfigTable[len(m.ConfigTable)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
