// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package tetrate_api_inventory_v1

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

// ServiceKind represents a service's type; whether it's a one-shot job,
// a traditional service, stateful set, etc.
type Service_Kind int32

const (
	// A Service, a long running job that answers many requests
	Service_SERVICE        Service_Kind = 0
	Service_JOB            Service_Kind = 1
	Service_VIRTUALSERVICE Service_Kind = 2
)

var Service_Kind_name = map[int32]string{
	0: "SERVICE",
	1: "JOB",
	2: "VIRTUALSERVICE",
}
var Service_Kind_value = map[string]int32{
	"SERVICE":        0,
	"JOB":            1,
	"VIRTUALSERVICE": 2,
}

func (x Service_Kind) String() string {
	return proto.EnumName(Service_Kind_name, int32(x))
}
func (Service_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{0, 0}
}

// ServiceStatus represents the health status of this service
type Service_Status int32

const (
	Service_UNHEALTHY Service_Status = 0
	Service_HEALTHY   Service_Status = 1
)

var Service_Status_name = map[int32]string{
	0: "UNHEALTHY",
	1: "HEALTHY",
}
var Service_Status_value = map[string]int32{
	"UNHEALTHY": 0,
	"HEALTHY":   1,
}

func (x Service_Status) String() string {
	return proto.EnumName(Service_Status_name, int32(x))
}
func (Service_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{0, 1}
}

// Transport protocols that are currently supported
type PortSpec_Protocol int32

const (
	PortSpec_TCP   PortSpec_Protocol = 0
	PortSpec_HTTP  PortSpec_Protocol = 1
	PortSpec_GRPC  PortSpec_Protocol = 2
	PortSpec_REDIS PortSpec_Protocol = 3
)

var PortSpec_Protocol_name = map[int32]string{
	0: "TCP",
	1: "HTTP",
	2: "GRPC",
	3: "REDIS",
}
var PortSpec_Protocol_value = map[string]int32{
	"TCP":   0,
	"HTTP":  1,
	"GRPC":  2,
	"REDIS": 3,
}

func (x PortSpec_Protocol) String() string {
	return proto.EnumName(PortSpec_Protocol_name, int32(x))
}
func (PortSpec_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{7, 0}
}

// Tetrate service definition that uniquely identifies services
type Service struct {
	// Unique name of this service to identify it. This name must be a
	// fully qualified one, e.g. `foo.bar.com`, and should be a valid DNS (RFC 1035) name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This Service's kind
	Kind Service_Kind `protobuf:"varint,2,opt,name=kind,proto3,enum=tetrate.api.inventory.v1.Service_Kind" json:"kind,omitempty"`
	// This Service's health status
	Status Service_Status `protobuf:"varint,3,opt,name=status,proto3,enum=tetrate.api.inventory.v1.Service_Status" json:"status,omitempty"`
	// Transport port spec
	Ports []*PortSpec `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	// Service's metadata attributes
	Metadata *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Service's cluster participation
	Cluster []*Cluster `protobuf:"bytes,6,rep,name=cluster,proto3" json:"cluster,omitempty"`
	// Service's workload presence
	Workload             *Workload            `protobuf:"bytes,7,opt,name=workload,proto3" json:"workload,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,8,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetKind() Service_Kind {
	if m != nil {
		return m.Kind
	}
	return Service_SERVICE
}

func (m *Service) GetStatus() Service_Status {
	if m != nil {
		return m.Status
	}
	return Service_UNHEALTHY
}

func (m *Service) GetPorts() []*PortSpec {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetCluster() []*Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *Service) GetWorkload() *Workload {
	if m != nil {
		return m.Workload
	}
	return nil
}

func (m *Service) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

// ListServicesRequest
type ListServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServicesRequest) Reset()         { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()    {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{1}
}
func (m *ListServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesRequest.Unmarshal(m, b)
}
func (m *ListServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesRequest.Marshal(b, m, deterministic)
}
func (dst *ListServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesRequest.Merge(dst, src)
}
func (m *ListServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListServicesRequest.Size(m)
}
func (m *ListServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesRequest proto.InternalMessageInfo

// ListServicesResponse is a collection of Tetrate-discovered services
type ListServicesResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListServicesResponse) Reset()         { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()    {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{2}
}
func (m *ListServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesResponse.Unmarshal(m, b)
}
func (m *ListServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesResponse.Marshal(b, m, deterministic)
}
func (dst *ListServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesResponse.Merge(dst, src)
}
func (m *ListServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListServicesResponse.Size(m)
}
func (m *ListServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesResponse proto.InternalMessageInfo

func (m *ListServicesResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

// Metadata extends object definitions by substaintiating them with
// additional attributes
type Metadata struct {
	// Generic metadata name
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Specific domain associated with this entity e.g. "svc.cluster.local"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Entities can be versioned e.g. "v1alpha1"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Labels represent any "key" ==> "value" attached to entities.
	// These identify and query meaningful attributes of entities
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations represent "key" ==> "value" attached to entities.
	// These are non-identifying information of entities
	Annotation map[string]string `protobuf:"bytes,5,rep,name=annotation,proto3" json:"annotation,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Organizational Unit associated with service e.g. "OU=ApplicationServers"
	Ou                   string   `protobuf:"bytes,6,opt,name=ou,proto3" json:"ou,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{3}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Metadata) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Metadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Metadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Metadata) GetAnnotation() map[string]string {
	if m != nil {
		return m.Annotation
	}
	return nil
}

func (m *Metadata) GetOu() string {
	if m != nil {
		return m.Ou
	}
	return ""
}

// Cluster is a collection of distributed endpoints serving certain
// functionality, that is typically bounded by administrative domain
type Cluster struct {
	// Cluster name e.g. "clusterA.local"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// One of more endpoints serving as part of cluster
	Endpoints            []*Endpoint          `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{4}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (dst *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(dst, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Cluster) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

// Workload represents the location of a specific instance of a service,
// identified by its provider (a public cloud like GCP or AWS or a private cloud),
// location, and node.
type Workload struct {
	// Generic name representing cloud workload
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Workload provider info e.g. AWS, GCP, Azure
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	// Cloud region this service resides e.g. "us-west-2"
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// Specific physical cloud node e.g. "gke-testing-default-pool-842099ce-7rdn"
	Node                 string   `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Workload) Reset()         { *m = Workload{} }
func (m *Workload) String() string { return proto.CompactTextString(m) }
func (*Workload) ProtoMessage()    {}
func (*Workload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{5}
}
func (m *Workload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workload.Unmarshal(m, b)
}
func (m *Workload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workload.Marshal(b, m, deterministic)
}
func (dst *Workload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workload.Merge(dst, src)
}
func (m *Workload) XXX_Size() int {
	return xxx_messageInfo_Workload.Size(m)
}
func (m *Workload) XXX_DiscardUnknown() {
	xxx_messageInfo_Workload.DiscardUnknown(m)
}

var xxx_messageInfo_Workload proto.InternalMessageInfo

func (m *Workload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Workload) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Workload) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Workload) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

// Endpoint represents
type Endpoint struct {
	// Endpoint represents one of IP or DNS name
	//
	// Types that are valid to be assigned to Address:
	//	*Endpoint_Ipv4
	//	*Endpoint_Ipv6
	//	*Endpoint_DnsName
	Address isEndpoint_Address `protobuf_oneof:"address"`
	// Endpoint port for this entity
	Port                 *PortSpec `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{6}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

type isEndpoint_Address interface {
	isEndpoint_Address()
}

type Endpoint_Ipv4 struct {
	Ipv4 []byte `protobuf:"bytes,1,opt,name=ipv4,proto3,oneof"`
}
type Endpoint_Ipv6 struct {
	Ipv6 []byte `protobuf:"bytes,2,opt,name=ipv6,proto3,oneof"`
}
type Endpoint_DnsName struct {
	DnsName string `protobuf:"bytes,3,opt,name=dns_name,json=dnsName,proto3,oneof"`
}

func (*Endpoint_Ipv4) isEndpoint_Address()    {}
func (*Endpoint_Ipv6) isEndpoint_Address()    {}
func (*Endpoint_DnsName) isEndpoint_Address() {}

func (m *Endpoint) GetAddress() isEndpoint_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetIpv4() []byte {
	if x, ok := m.GetAddress().(*Endpoint_Ipv4); ok {
		return x.Ipv4
	}
	return nil
}

func (m *Endpoint) GetIpv6() []byte {
	if x, ok := m.GetAddress().(*Endpoint_Ipv6); ok {
		return x.Ipv6
	}
	return nil
}

func (m *Endpoint) GetDnsName() string {
	if x, ok := m.GetAddress().(*Endpoint_DnsName); ok {
		return x.DnsName
	}
	return ""
}

func (m *Endpoint) GetPort() *PortSpec {
	if m != nil {
		return m.Port
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Endpoint) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Endpoint_OneofMarshaler, _Endpoint_OneofUnmarshaler, _Endpoint_OneofSizer, []interface{}{
		(*Endpoint_Ipv4)(nil),
		(*Endpoint_Ipv6)(nil),
		(*Endpoint_DnsName)(nil),
	}
}

func _Endpoint_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Endpoint)
	// address
	switch x := m.Address.(type) {
	case *Endpoint_Ipv4:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ipv4)
	case *Endpoint_Ipv6:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ipv6)
	case *Endpoint_DnsName:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DnsName)
	case nil:
	default:
		return fmt.Errorf("Endpoint.Address has unexpected type %T", x)
	}
	return nil
}

func _Endpoint_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Endpoint)
	switch tag {
	case 1: // address.ipv4
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Address = &Endpoint_Ipv4{x}
		return true, err
	case 2: // address.ipv6
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Address = &Endpoint_Ipv6{x}
		return true, err
	case 3: // address.dns_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &Endpoint_DnsName{x}
		return true, err
	default:
		return false, nil
	}
}

func _Endpoint_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Endpoint)
	// address
	switch x := m.Address.(type) {
	case *Endpoint_Ipv4:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Ipv4)))
		n += len(x.Ipv4)
	case *Endpoint_Ipv6:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Ipv6)))
		n += len(x.Ipv6)
	case *Endpoint_DnsName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.DnsName)))
		n += len(x.DnsName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// PortSpec represents an endpoint's port details. A single endpoint can have
// one or more port specs.
type PortSpec struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This spec's protocol
	Protocol PortSpec_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=tetrate.api.inventory.v1.PortSpec_Protocol" json:"protocol,omitempty"`
	// Port facing external network
	ExternalPort int32 `protobuf:"varint,3,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	// This service's incoming port redirected to specific target
	TargetPort int32 `protobuf:"varint,4,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	// Container port that services this request
	ContainerPort        int32    `protobuf:"varint,5,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortSpec) Reset()         { *m = PortSpec{} }
func (m *PortSpec) String() string { return proto.CompactTextString(m) }
func (*PortSpec) ProtoMessage()    {}
func (*PortSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_9f2b8006ef59627a, []int{7}
}
func (m *PortSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortSpec.Unmarshal(m, b)
}
func (m *PortSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortSpec.Marshal(b, m, deterministic)
}
func (dst *PortSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortSpec.Merge(dst, src)
}
func (m *PortSpec) XXX_Size() int {
	return xxx_messageInfo_PortSpec.Size(m)
}
func (m *PortSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PortSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PortSpec proto.InternalMessageInfo

func (m *PortSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PortSpec) GetProtocol() PortSpec_Protocol {
	if m != nil {
		return m.Protocol
	}
	return PortSpec_TCP
}

func (m *PortSpec) GetExternalPort() int32 {
	if m != nil {
		return m.ExternalPort
	}
	return 0
}

func (m *PortSpec) GetTargetPort() int32 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

func (m *PortSpec) GetContainerPort() int32 {
	if m != nil {
		return m.ContainerPort
	}
	return 0
}

func init() {
	proto.RegisterType((*Service)(nil), "tetrate.api.inventory.v1.Service")
	proto.RegisterType((*ListServicesRequest)(nil), "tetrate.api.inventory.v1.ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "tetrate.api.inventory.v1.ListServicesResponse")
	proto.RegisterType((*Metadata)(nil), "tetrate.api.inventory.v1.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "tetrate.api.inventory.v1.Metadata.AnnotationEntry")
	proto.RegisterMapType((map[string]string)(nil), "tetrate.api.inventory.v1.Metadata.LabelsEntry")
	proto.RegisterType((*Cluster)(nil), "tetrate.api.inventory.v1.Cluster")
	proto.RegisterType((*Workload)(nil), "tetrate.api.inventory.v1.Workload")
	proto.RegisterType((*Endpoint)(nil), "tetrate.api.inventory.v1.Endpoint")
	proto.RegisterType((*PortSpec)(nil), "tetrate.api.inventory.v1.PortSpec")
	proto.RegisterEnum("tetrate.api.inventory.v1.Service_Kind", Service_Kind_name, Service_Kind_value)
	proto.RegisterEnum("tetrate.api.inventory.v1.Service_Status", Service_Status_name, Service_Status_value)
	proto.RegisterEnum("tetrate.api.inventory.v1.PortSpec_Protocol", PortSpec_Protocol_name, PortSpec_Protocol_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryServiceClient interface {
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
}

type inventoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewInventoryServiceClient(cc *grpc.ClientConn) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/tetrate.api.inventory.v1.InventoryService/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
type InventoryServiceServer interface {
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
}

func RegisterInventoryServiceServer(s *grpc.Server, srv InventoryServiceServer) {
	s.RegisterService(&_InventoryService_serviceDesc, srv)
}

func _InventoryService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tetrate.api.inventory.v1.InventoryService/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InventoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tetrate.api.inventory.v1.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _InventoryService_ListServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_9f2b8006ef59627a) }

var fileDescriptor_service_9f2b8006ef59627a = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0x5d, 0xe7, 0xd3, 0xb9, 0x4d, 0x42, 0x34, 0x14, 0x64, 0x85, 0x95, 0x76, 0x31, 0x1f, 0xaa,
	0x84, 0x70, 0xd8, 0x80, 0xaa, 0x65, 0xd1, 0xc2, 0x76, 0x4b, 0xd8, 0x16, 0xba, 0x4b, 0x34, 0x4d,
	0x17, 0xf1, 0x54, 0x4d, 0xe3, 0x69, 0x65, 0xd5, 0xf1, 0x18, 0xcf, 0xc4, 0xd0, 0x37, 0x04, 0x2f,
	0xbc, 0x23, 0x21, 0x1e, 0x79, 0xe7, 0xe7, 0xf0, 0x17, 0xf8, 0x21, 0x5c, 0x8f, 0x67, 0xdc, 0x02,
	0xcd, 0xa6, 0xfb, 0xe6, 0xb9, 0x73, 0xce, 0x3d, 0x73, 0xe7, 0x9e, 0xb9, 0x86, 0x9e, 0xe4, 0x59,
	0x1e, 0xcd, 0x79, 0x90, 0x66, 0x42, 0x09, 0xe2, 0x29, 0xae, 0x32, 0xa6, 0x78, 0xc0, 0xd2, 0x28,
	0x88, 0x92, 0x9c, 0x27, 0x4a, 0x64, 0x17, 0x41, 0x7e, 0x6f, 0x78, 0xfb, 0x4c, 0x88, 0xb3, 0x98,
	0x8f, 0x70, 0x63, 0xc4, 0x92, 0x44, 0x28, 0xa6, 0x22, 0x91, 0xc8, 0x92, 0x37, 0xbc, 0x63, 0x76,
	0xf5, 0xea, 0x64, 0x79, 0x3a, 0x52, 0xd1, 0x82, 0x4b, 0xc5, 0x16, 0x69, 0x09, 0xf0, 0xff, 0x6c,
	0x40, 0xfb, 0xb0, 0x94, 0x22, 0x04, 0x1a, 0x09, 0x5b, 0x70, 0xcf, 0xb9, 0xeb, 0x6c, 0x75, 0xa8,
	0xfe, 0x26, 0x0f, 0xa0, 0x71, 0x1e, 0x25, 0xa1, 0x57, 0xc3, 0x58, 0x7f, 0xfc, 0x6e, 0xb0, 0xea,
	0x1c, 0x81, 0x49, 0x12, 0x7c, 0x85, 0x68, 0xaa, 0x39, 0xe4, 0x11, 0xb4, 0x50, 0x4a, 0x2d, 0xa5,
	0x57, 0xd7, 0xec, 0xad, 0xf5, 0xec, 0x43, 0x8d, 0xa7, 0x86, 0x47, 0xee, 0x43, 0x33, 0x15, 0x99,
	0x92, 0x5e, 0xe3, 0x6e, 0x7d, 0x6b, 0x63, 0xec, 0xaf, 0x4e, 0x30, 0x45, 0xd8, 0x61, 0xca, 0xe7,
	0xb4, 0x24, 0x90, 0x4f, 0xc1, 0x5d, 0x70, 0xc5, 0x42, 0xa6, 0x98, 0xd7, 0x44, 0xf5, 0x17, 0x92,
	0x9f, 0x1a, 0x24, 0xad, 0x38, 0xe4, 0x13, 0x68, 0xcf, 0xe3, 0xa5, 0x54, 0x3c, 0xf3, 0x5a, 0x5a,
	0xfb, 0xcd, 0xd5, 0xf4, 0xdd, 0x12, 0x48, 0x2d, 0xa3, 0x10, 0xff, 0x5e, 0x64, 0xe7, 0xb1, 0x60,
	0xa1, 0xd7, 0x5e, 0x27, 0xfe, 0x8d, 0x41, 0xd2, 0x8a, 0x43, 0x3e, 0x83, 0x5e, 0xcc, 0xa4, 0x3a,
	0x5e, 0x88, 0x30, 0x3a, 0x8d, 0x78, 0xe8, 0xb9, 0x3a, 0xc9, 0x30, 0x28, 0xbb, 0x19, 0xd8, 0x6e,
	0x06, 0x33, 0xdb, 0x4d, 0xda, 0x2d, 0x08, 0x4f, 0x0d, 0xde, 0xff, 0x00, 0x1a, 0x45, 0x1f, 0xc8,
	0x06, 0x36, 0x77, 0x42, 0x9f, 0xef, 0xef, 0x4e, 0x06, 0xb7, 0x48, 0x1b, 0xea, 0x5f, 0x7e, 0xfd,
	0x78, 0xe0, 0x60, 0x9f, 0xfb, 0xcf, 0xf7, 0xe9, 0xec, 0x68, 0xe7, 0xc0, 0x6e, 0xd6, 0xfc, 0xb7,
	0xa1, 0x55, 0xde, 0x3d, 0xe9, 0x41, 0xe7, 0xe8, 0xd9, 0xde, 0x64, 0xe7, 0x60, 0xb6, 0xf7, 0x2d,
	0xb2, 0x30, 0x85, 0x5d, 0x38, 0xfe, 0x6b, 0xf0, 0xea, 0x41, 0x24, 0x95, 0xe9, 0x96, 0xa4, 0xfc,
	0xbb, 0x25, 0xea, 0xfb, 0x47, 0xb0, 0xf9, 0xef, 0xb0, 0x4c, 0xd1, 0x82, 0x9c, 0x3c, 0x04, 0xd7,
	0xd8, 0x58, 0xa2, 0xa9, 0xd6, 0xdc, 0xa2, 0x61, 0xd3, 0x8a, 0xe2, 0xff, 0x58, 0x07, 0xd7, 0xb6,
	0x86, 0xdc, 0x86, 0x4e, 0x61, 0x48, 0x99, 0xb2, 0xb9, 0x75, 0xe8, 0x65, 0x80, 0xbc, 0x0e, 0xad,
	0x50, 0x2c, 0x58, 0x94, 0x68, 0xa3, 0x76, 0xa8, 0x59, 0x11, 0x0f, 0xda, 0x39, 0xcf, 0x24, 0xbe,
	0x08, 0xed, 0xc1, 0x0e, 0xb5, 0x4b, 0xf2, 0x05, 0xb4, 0x62, 0x76, 0xc2, 0x63, 0xeb, 0xad, 0x60,
	0xbd, 0x3d, 0x82, 0x03, 0x4d, 0x98, 0x24, 0x2a, 0xbb, 0xa0, 0x86, 0x4d, 0x28, 0xc0, 0xe5, 0xb3,
	0x43, 0xab, 0x15, 0xb9, 0xc6, 0x37, 0xc8, 0xb5, 0x53, 0x91, 0xca, 0x7c, 0x57, 0xb2, 0x90, 0x3e,
	0xd4, 0xc4, 0x12, 0x7d, 0x57, 0x1c, 0x18, 0xbf, 0x86, 0x1f, 0xc3, 0xc6, 0x15, 0x69, 0x32, 0x80,
	0xfa, 0x39, 0xbf, 0x30, 0x97, 0x50, 0x7c, 0x92, 0x4d, 0x68, 0xe6, 0x2c, 0x5e, 0x72, 0x53, 0x7d,
	0xb9, 0x78, 0x50, 0xbb, 0xef, 0x0c, 0x1f, 0xc2, 0x2b, 0xff, 0x51, 0x7a, 0x19, 0xba, 0xff, 0x87,
	0x03, 0x6d, 0x63, 0xef, 0x6b, 0xc7, 0xc3, 0x23, 0xe8, 0xf0, 0x24, 0x4c, 0x45, 0x94, 0xe0, 0x23,
	0xad, 0xad, 0x7b, 0xa4, 0x13, 0x03, 0xa5, 0x97, 0xa4, 0xff, 0x7b, 0xbd, 0xfe, 0x92, 0x5e, 0x3f,
	0x05, 0xd7, 0x3e, 0xa1, 0x6b, 0x8f, 0x38, 0x04, 0x17, 0x73, 0xe4, 0x51, 0x88, 0x4f, 0xb9, 0xac,
	0xaf, 0x5a, 0x17, 0xb6, 0xc9, 0xf8, 0xd9, 0xa5, 0x3b, 0xcc, 0x4a, 0xe7, 0x11, 0x21, 0x47, 0x6b,
	0x94, 0x79, 0xf0, 0xdb, 0xff, 0xcd, 0x01, 0xd7, 0x16, 0x80, 0x37, 0xd6, 0x88, 0xd2, 0xfc, 0x23,
	0x2d, 0xd4, 0xdd, 0xbb, 0x45, 0xf5, 0xca, 0x44, 0xb7, 0xb5, 0x8c, 0x8d, 0x6e, 0x93, 0x37, 0xc0,
	0x0d, 0x13, 0x79, 0xac, 0x0f, 0xa6, 0x65, 0x70, 0xa7, 0x8d, 0x91, 0x67, 0xc5, 0xe9, 0xb6, 0xa1,
	0x51, 0x0c, 0x2c, 0xad, 0x74, 0xb3, 0x01, 0xa7, 0xf1, 0x8f, 0x3b, 0xd0, 0x66, 0x61, 0x98, 0x71,
	0x29, 0xfd, 0x5f, 0x6a, 0xe0, 0xda, 0xdd, 0x6b, 0x6f, 0xe0, 0x89, 0xbe, 0x01, 0x25, 0xe6, 0x22,
	0x36, 0x73, 0xfc, 0xbd, 0xf5, 0x3a, 0xc1, 0xd4, 0x50, 0x68, 0x45, 0x26, 0x6f, 0x41, 0x8f, 0xff,
	0x80, 0x56, 0x48, 0x58, 0x7c, 0xac, 0x4f, 0x5d, 0x94, 0xd3, 0xa4, 0x5d, 0x1b, 0x2c, 0xb8, 0xe4,
	0x0e, 0x6c, 0x28, 0x96, 0x9d, 0x71, 0x75, 0x5c, 0x15, 0xd6, 0xa4, 0x50, 0x86, 0x34, 0xe0, 0x1d,
	0xe8, 0xcf, 0x45, 0xa2, 0xf0, 0x79, 0xf2, 0xac, 0xc4, 0x34, 0x35, 0xa6, 0x57, 0x45, 0x0b, 0x98,
	0x3f, 0xc6, 0xaa, 0xac, 0x30, 0x8e, 0xae, 0xd9, 0xee, 0x14, 0xa7, 0x91, 0x0b, 0x8d, 0xbd, 0xd9,
	0x6c, 0x8a, 0x43, 0x0c, 0xbf, 0x9e, 0xd0, 0xe9, 0xee, 0xa0, 0x46, 0x3a, 0xd0, 0xa4, 0x93, 0xcf,
	0xf7, 0x0f, 0x07, 0xf5, 0xf1, 0xef, 0x0e, 0x0c, 0xf6, 0x6d, 0x35, 0xf6, 0xb7, 0xf6, 0xb3, 0x03,
	0xdd, 0xab, 0xe3, 0x89, 0xbc, 0xbf, 0xba, 0xfa, 0x6b, 0xa6, 0xdb, 0x30, 0xb8, 0x29, 0xbc, 0x9c,
	0x7a, 0xfe, 0xe6, 0x4f, 0x7f, 0xfd, 0xfd, 0x6b, 0xad, 0x4f, 0xba, 0xa3, 0xfc, 0xde, 0xc8, 0x0e,
	0xb3, 0x93, 0x96, 0xbe, 0xc5, 0x0f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x37, 0x9f, 0x16,
	0xd9, 0x07, 0x00, 0x00,
}
