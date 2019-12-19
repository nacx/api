// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: settings.proto

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

// Operator
//
// Possible operators to "judge" the current metric value against a particular threshold value.
type Operator int32

const (
	Operator_EQ Operator = 0
	Operator_GT Operator = 1
	Operator_LT Operator = 2
)

var Operator_name = map[int32]string{
	0: "EQ",
	1: "GT",
	2: "LT",
}

var Operator_value = map[string]int32{
	"EQ": 0,
	"GT": 1,
	"LT": 2,
}

func (x Operator) String() string {
	return proto.EnumName(Operator_name, int32(x))
}

func (Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{0}
}

// Settings
//
// A set of alert rules.
type Settings struct {
	Rules                map[string]*Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return m.Size()
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetRules() map[string]*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Rule
//
// An alert rule defines the condition when an alert should be triggered.
type Rule struct {
	// Short ID for this rule.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Metric name to be watched.
	// clang-format off
	MetricName string `protobuf:"bytes,2,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// The threshold value of a metric; depends on the operator, it triggers alert if the metric value
	// is not match.
	// clang-format off
	Threshold uint32 `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// The operator. E.g. greater-than, lower-than and equal.
	Operator Operator `protobuf:"varint,4,opt,name=operator,proto3,enum=tetrate.api.tcc.spm.alert.v1.Operator" json:"operator,omitempty"`
	// The period of an alert rule should be checked.
	Period *types.Duration `protobuf:"bytes,5,opt,name=period,proto3" json:"period,omitempty"`
	// The number of "violation" in a certain period.
	Count uint32 `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	// Keep silence for some amount of time after an alarm is triggered. By default, the value is the
	// same as period.
	SilencePeriod *types.Duration `protobuf:"bytes,7,opt,name=silence_period,json=silencePeriod,proto3" json:"silence_period,omitempty"`
	// Friendly name for a rule.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Whether this rule is enabled or disabled. Default to false or disabled
	Enabled bool `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Included services in this rule. This will be translated into include-names in skywalking
	// configuration response.
	// By default, spmserver will add all services in this application.
	IncludeServices []string `protobuf:"bytes,10,rep,name=include_services,json=includeServices,proto3" json:"include_services,omitempty"`
	// Included services in this rule. This will be translated into exclude-names in skywalking
	// configuration response.
	// By default, spmserver will omit this from response.
	ExcludeServices      []string `protobuf:"bytes,11,rep,name=exclude_services,json=excludeServices,proto3" json:"exclude_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7cab62fa432213, []int{1}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return m.Size()
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rule) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *Rule) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *Rule) GetOperator() Operator {
	if m != nil {
		return m.Operator
	}
	return Operator_EQ
}

func (m *Rule) GetPeriod() *types.Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

func (m *Rule) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Rule) GetSilencePeriod() *types.Duration {
	if m != nil {
		return m.SilencePeriod
	}
	return nil
}

func (m *Rule) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Rule) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Rule) GetIncludeServices() []string {
	if m != nil {
		return m.IncludeServices
	}
	return nil
}

func (m *Rule) GetExcludeServices() []string {
	if m != nil {
		return m.ExcludeServices
	}
	return nil
}

func init() {
	proto.RegisterEnum("tetrate.api.tcc.spm.alert.v1.Operator", Operator_name, Operator_value)
	proto.RegisterType((*Settings)(nil), "tetrate.api.tcc.spm.alert.v1.Settings")
	proto.RegisterMapType((map[string]*Rule)(nil), "tetrate.api.tcc.spm.alert.v1.Settings.RulesEntry")
	proto.RegisterType((*Rule)(nil), "tetrate.api.tcc.spm.alert.v1.Rule")
}

func init() { proto.RegisterFile("settings.proto", fileDescriptor_6c7cab62fa432213) }

var fileDescriptor_6c7cab62fa432213 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xbf, 0x8f, 0xd3, 0x3e,
	0x1c, 0x3d, 0x27, 0x6d, 0x9a, 0x38, 0x77, 0xfd, 0x46, 0x5e, 0xbe, 0xa1, 0x42, 0x25, 0x74, 0x40,
	0xb9, 0x4a, 0x38, 0x6a, 0x59, 0x4e, 0x30, 0x80, 0x22, 0x4e, 0xb7, 0x20, 0x7e, 0xf8, 0x6e, 0x42,
	0x48, 0x95, 0x9b, 0x98, 0xd6, 0x22, 0x89, 0x23, 0xc7, 0xa9, 0xe8, 0xbf, 0xc4, 0xc6, 0xca, 0xc4,
	0xc8, 0xc8, 0x9f, 0x80, 0xba, 0x31, 0xf1, 0x2f, 0xa0, 0x24, 0x0e, 0x27, 0x31, 0x94, 0xc9, 0xce,
	0xf3, 0x7b, 0x9f, 0xf7, 0xc9, 0x7b, 0x70, 0x5c, 0x31, 0xa5, 0x78, 0xb1, 0xa9, 0x70, 0x29, 0x85,
	0x12, 0xe8, 0xae, 0x62, 0x4a, 0x52, 0xc5, 0x30, 0x2d, 0x39, 0x56, 0x49, 0x82, 0xab, 0x32, 0xc7,
	0x34, 0x63, 0x52, 0xe1, 0xdd, 0x62, 0x32, 0xdd, 0x08, 0xb1, 0xc9, 0x58, 0xd4, 0x72, 0xd7, 0xf5,
	0xfb, 0x28, 0xad, 0x25, 0x55, 0x5c, 0x14, 0x9d, 0x7a, 0xf2, 0xff, 0x8e, 0x66, 0x3c, 0xa5, 0x8a,
	0x45, 0xfd, 0xa5, 0x7b, 0x98, 0x7d, 0x06, 0xd0, 0xbe, 0xd6, 0x4e, 0xe8, 0x0a, 0x0e, 0x65, 0x9d,
	0xb1, 0xca, 0x07, 0x81, 0x19, 0xba, 0xcb, 0x05, 0x3e, 0xe6, 0x89, 0x7b, 0x19, 0x26, 0x8d, 0xe6,
	0xb2, 0x50, 0x72, 0x4f, 0x3a, 0xfd, 0xe4, 0x1d, 0x84, 0xb7, 0x20, 0xf2, 0xa0, 0xf9, 0x81, 0xed,
	0x7d, 0x10, 0x80, 0xd0, 0x21, 0xcd, 0x15, 0x5d, 0xc0, 0xe1, 0x8e, 0x66, 0x35, 0xf3, 0x8d, 0x00,
	0x84, 0xee, 0x72, 0x76, 0xdc, 0xa8, 0x19, 0x45, 0x3a, 0xc1, 0x63, 0xe3, 0x02, 0xcc, 0x7e, 0x99,
	0x70, 0xd0, 0x60, 0x68, 0x0c, 0x0d, 0x9e, 0xea, 0xb9, 0x06, 0x4f, 0xd1, 0x02, 0xba, 0x39, 0x53,
	0x92, 0x27, 0xab, 0x82, 0xe6, 0xdd, 0x70, 0x27, 0xf6, 0xbe, 0xfc, 0xfc, 0x6a, 0xba, 0xd2, 0xf1,
	0xac, 0x78, 0xb8, 0x6a, 0xd6, 0x23, 0xb0, 0x23, 0xbd, 0xa4, 0x39, 0x43, 0xe7, 0xd0, 0x51, 0x5b,
	0xc9, 0xaa, 0xad, 0xc8, 0x52, 0xdf, 0x0c, 0x40, 0x78, 0x16, 0xbb, 0x8d, 0xc0, 0x9a, 0x0f, 0xfc,
	0x34, 0x3c, 0x21, 0xb7, 0xaf, 0x28, 0x86, 0xb6, 0x28, 0x99, 0xa4, 0x4a, 0x48, 0x7f, 0x10, 0x80,
	0x70, 0xbc, 0x7c, 0x70, 0x7c, 0xef, 0x57, 0x9a, 0x4d, 0xfe, 0xe8, 0xd0, 0x53, 0x68, 0x95, 0x4c,
	0x72, 0x91, 0xfa, 0xc3, 0xf6, 0xcf, 0xef, 0xe0, 0xae, 0x38, 0xdc, 0x17, 0x87, 0x9f, 0xeb, 0xe2,
	0xe2, 0xd3, 0x66, 0x8d, 0xd1, 0x27, 0x30, 0x98, 0x18, 0x36, 0x20, 0x5a, 0x86, 0xee, 0xc1, 0x61,
	0x22, 0xea, 0x42, 0xf9, 0x56, 0xbb, 0xab, 0xd3, 0x90, 0x06, 0x73, 0x23, 0x3c, 0x21, 0x1d, 0x8e,
	0x9e, 0xc1, 0x71, 0xc5, 0x33, 0x56, 0x24, 0x6c, 0xa5, 0x9d, 0x46, 0xff, 0x70, 0x22, 0x67, 0x5a,
	0xf0, 0xba, 0xb3, 0xb8, 0x0f, 0x4f, 0x53, 0x5e, 0x95, 0x19, 0xdd, 0x77, 0x31, 0xda, 0x6d, 0xbe,
	0xae, 0xc6, 0xda, 0xd4, 0x7c, 0x38, 0x62, 0x05, 0x5d, 0x67, 0x2c, 0xf5, 0x9d, 0x00, 0x84, 0x36,
	0xe9, 0x3f, 0xd1, 0x39, 0xf4, 0x78, 0x91, 0x64, 0x75, 0xca, 0x56, 0x15, 0x93, 0x3b, 0x9e, 0xb0,
	0xca, 0x87, 0x81, 0x19, 0x3a, 0xe4, 0x3f, 0x8d, 0x5f, 0x6b, 0xb8, 0xa1, 0xb2, 0x8f, 0x7f, 0x51,
	0xdd, 0x8e, 0xaa, 0xf1, 0x9e, 0x3a, 0x9f, 0x41, 0xbb, 0x0f, 0x13, 0x59, 0xd0, 0xb8, 0x7c, 0xe3,
	0x9d, 0x34, 0xe7, 0xd5, 0x8d, 0x07, 0x9a, 0xf3, 0xc5, 0x8d, 0x67, 0xc4, 0x4f, 0xbe, 0x1d, 0xa6,
	0xe0, 0xfb, 0x61, 0x0a, 0x7e, 0x1c, 0xa6, 0xe0, 0xed, 0xc3, 0x0d, 0x57, 0xdb, 0x7a, 0x8d, 0x13,
	0x91, 0x47, 0xba, 0x24, 0x2e, 0xfa, 0x5b, 0x44, 0x4b, 0x1e, 0x55, 0x65, 0x1e, 0xb5, 0x55, 0x45,
	0xbb, 0xc5, 0xda, 0x6a, 0x53, 0x79, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x95, 0xfa, 0x58, 0xe5,
	0x76, 0x03, 0x00, 0x00,
}

func (m *Settings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Settings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Settings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Rules) > 0 {
		for k := range m.Rules {
			v := m.Rules[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintSettings(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintSettings(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintSettings(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ExcludeServices) > 0 {
		for iNdEx := len(m.ExcludeServices) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ExcludeServices[iNdEx])
			copy(dAtA[i:], m.ExcludeServices[iNdEx])
			i = encodeVarintSettings(dAtA, i, uint64(len(m.ExcludeServices[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.IncludeServices) > 0 {
		for iNdEx := len(m.IncludeServices) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IncludeServices[iNdEx])
			copy(dAtA[i:], m.IncludeServices[iNdEx])
			i = encodeVarintSettings(dAtA, i, uint64(len(m.IncludeServices[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintSettings(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x42
	}
	if m.SilencePeriod != nil {
		{
			size, err := m.SilencePeriod.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Count != 0 {
		i = encodeVarintSettings(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x30
	}
	if m.Period != nil {
		{
			size, err := m.Period.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSettings(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Operator != 0 {
		i = encodeVarintSettings(dAtA, i, uint64(m.Operator))
		i--
		dAtA[i] = 0x20
	}
	if m.Threshold != 0 {
		i = encodeVarintSettings(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MetricName) > 0 {
		i -= len(m.MetricName)
		copy(dAtA[i:], m.MetricName)
		i = encodeVarintSettings(dAtA, i, uint64(len(m.MetricName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSettings(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSettings(dAtA []byte, offset int, v uint64) int {
	offset -= sovSettings(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Settings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for k, v := range m.Rules {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovSettings(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovSettings(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovSettings(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	l = len(m.MetricName)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.Threshold != 0 {
		n += 1 + sovSettings(uint64(m.Threshold))
	}
	if m.Operator != 0 {
		n += 1 + sovSettings(uint64(m.Operator))
	}
	if m.Period != nil {
		l = m.Period.Size()
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovSettings(uint64(m.Count))
	}
	if m.SilencePeriod != nil {
		l = m.SilencePeriod.Size()
		n += 1 + l + sovSettings(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.IncludeServices) > 0 {
		for _, s := range m.IncludeServices {
			l = len(s)
			n += 1 + l + sovSettings(uint64(l))
		}
	}
	if len(m.ExcludeServices) > 0 {
		for _, s := range m.ExcludeServices {
			l = len(s)
			n += 1 + l + sovSettings(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSettings(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSettings(x uint64) (n int) {
	return sovSettings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Settings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: Settings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Settings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rules == nil {
				m.Rules = make(map[string]*Rule)
			}
			var mapkey string
			var mapvalue *Rule
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSettings
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
							return ErrIntOverflowSettings
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
						return ErrInvalidLengthSettings
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthSettings
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSettings
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthSettings
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthSettings
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Rule{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSettings(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthSettings
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Rules[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSettings
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
func (m *Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			m.Operator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operator |= Operator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Period == nil {
				m.Period = &types.Duration{}
			}
			if err := m.Period.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SilencePeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SilencePeriod == nil {
				m.SilencePeriod = &types.Duration{}
			}
			if err := m.SilencePeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncludeServices", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncludeServices = append(m.IncludeServices, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExcludeServices", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExcludeServices = append(m.ExcludeServices, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSettings
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
func skipSettings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSettings
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
					return 0, ErrIntOverflowSettings
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
					return 0, ErrIntOverflowSettings
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
				return 0, ErrInvalidLengthSettings
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSettings
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSettings
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
				next, err := skipSettings(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSettings
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
	ErrInvalidLengthSettings = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSettings   = fmt.Errorf("proto: integer overflow")
)
