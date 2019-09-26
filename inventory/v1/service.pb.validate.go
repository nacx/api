// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service.proto

package invv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Service) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Kind

	// no validation rules for Status

	for idx, item := range m.GetPorts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceValidationError{
					field:  fmt.Sprintf("Ports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCluster() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceValidationError{
					field:  fmt.Sprintf("Cluster[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetWorkload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "Workload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLastModified()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "LastModified",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on ListServicesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListServicesRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListServicesRequestValidationError is the validation error returned by
// ListServicesRequest.Validate if the designated constraints aren't met.
type ListServicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListServicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListServicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListServicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListServicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListServicesRequestValidationError) ErrorName() string {
	return "ListServicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListServicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListServicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListServicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListServicesRequestValidationError{}

// Validate checks the field values on ListServicesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListServicesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListServicesResponseValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListServicesResponseValidationError is the validation error returned by
// ListServicesResponse.Validate if the designated constraints aren't met.
type ListServicesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListServicesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListServicesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListServicesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListServicesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListServicesResponseValidationError) ErrorName() string {
	return "ListServicesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListServicesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListServicesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListServicesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListServicesResponseValidationError{}

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Metadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Namespace

	// no validation rules for Domain

	// no validation rules for Version

	// no validation rules for Labels

	// no validation rules for Annotation

	// no validation rules for Ou

	return nil
}

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

// Validate checks the field values on Cluster with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Cluster) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterValidationError{
					field:  fmt.Sprintf("Endpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetLastModified()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterValidationError{
				field:  "LastModified",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ClusterValidationError is the validation error returned by Cluster.Validate
// if the designated constraints aren't met.
type ClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterValidationError) ErrorName() string { return "ClusterValidationError" }

// Error satisfies the builtin error interface
func (e ClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterValidationError{}

// Validate checks the field values on Workload with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Workload) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Provider

	// no validation rules for Region

	// no validation rules for Node

	return nil
}

// WorkloadValidationError is the validation error returned by
// Workload.Validate if the designated constraints aren't met.
type WorkloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkloadValidationError) ErrorName() string { return "WorkloadValidationError" }

// Error satisfies the builtin error interface
func (e WorkloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkloadValidationError{}

// Validate checks the field values on Endpoint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Endpoint) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EndpointValidationError{
				field:  "Port",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Address.(type) {

	case *Endpoint_Ipv4:
		// no validation rules for Ipv4

	case *Endpoint_Ipv6:
		// no validation rules for Ipv6

	case *Endpoint_DnsName:
		// no validation rules for DnsName

	}

	return nil
}

// EndpointValidationError is the validation error returned by
// Endpoint.Validate if the designated constraints aren't met.
type EndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointValidationError) ErrorName() string { return "EndpointValidationError" }

// Error satisfies the builtin error interface
func (e EndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointValidationError{}

// Validate checks the field values on PortSpec with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PortSpec) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Protocol

	// no validation rules for ExternalPort

	// no validation rules for ContainerPort

	return nil
}

// PortSpecValidationError is the validation error returned by
// PortSpec.Validate if the designated constraints aren't met.
type PortSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PortSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PortSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PortSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PortSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PortSpecValidationError) ErrorName() string { return "PortSpecValidationError" }

// Error satisfies the builtin error interface
func (e PortSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPortSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PortSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PortSpecValidationError{}

// Validate checks the field values on ServiceBinding with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ServiceBinding) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Etag

	// no validation rules for Policy

	return nil
}

// ServiceBindingValidationError is the validation error returned by
// ServiceBinding.Validate if the designated constraints aren't met.
type ServiceBindingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceBindingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceBindingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceBindingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceBindingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceBindingValidationError) ErrorName() string { return "ServiceBindingValidationError" }

// Error satisfies the builtin error interface
func (e ServiceBindingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceBinding.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceBindingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceBindingValidationError{}

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Policy) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Description

	return nil
}

// PolicyValidationError is the validation error returned by Policy.Validate if
// the designated constraints aren't met.
type PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyValidationError) ErrorName() string { return "PolicyValidationError" }

// Error satisfies the builtin error interface
func (e PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyValidationError{}

// Validate checks the field values on BindServiceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BindServiceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	if v, ok := interface{}(m.GetBinding()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BindServiceRequestValidationError{
				field:  "Binding",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BindServiceRequestValidationError is the validation error returned by
// BindServiceRequest.Validate if the designated constraints aren't met.
type BindServiceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BindServiceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BindServiceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BindServiceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BindServiceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BindServiceRequestValidationError) ErrorName() string {
	return "BindServiceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BindServiceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBindServiceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BindServiceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BindServiceRequestValidationError{}

// Validate checks the field values on ServicePolicyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServicePolicyRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	// no validation rules for Policy

	return nil
}

// ServicePolicyRequestValidationError is the validation error returned by
// ServicePolicyRequest.Validate if the designated constraints aren't met.
type ServicePolicyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServicePolicyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServicePolicyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServicePolicyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServicePolicyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServicePolicyRequestValidationError) ErrorName() string {
	return "ServicePolicyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ServicePolicyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServicePolicyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServicePolicyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServicePolicyRequestValidationError{}

// Validate checks the field values on ListBindingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListBindingsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	return nil
}

// ListBindingsRequestValidationError is the validation error returned by
// ListBindingsRequest.Validate if the designated constraints aren't met.
type ListBindingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBindingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBindingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBindingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBindingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBindingsRequestValidationError) ErrorName() string {
	return "ListBindingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListBindingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBindingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBindingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBindingsRequestValidationError{}

// Validate checks the field values on ListBindingsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListBindingsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetBindings() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBindingsResponseValidationError{
					field:  fmt.Sprintf("Bindings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListBindingsResponseValidationError is the validation error returned by
// ListBindingsResponse.Validate if the designated constraints aren't met.
type ListBindingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBindingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBindingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBindingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBindingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBindingsResponseValidationError) ErrorName() string {
	return "ListBindingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListBindingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBindingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBindingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBindingsResponseValidationError{}

// Validate checks the field values on ListPoliciesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPoliciesRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListPoliciesRequestValidationError is the validation error returned by
// ListPoliciesRequest.Validate if the designated constraints aren't met.
type ListPoliciesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPoliciesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPoliciesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPoliciesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPoliciesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPoliciesRequestValidationError) ErrorName() string {
	return "ListPoliciesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPoliciesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPoliciesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPoliciesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPoliciesRequestValidationError{}

// Validate checks the field values on ListPoliciesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPoliciesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPolicies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPoliciesResponseValidationError{
					field:  fmt.Sprintf("Policies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListPoliciesResponseValidationError is the validation error returned by
// ListPoliciesResponse.Validate if the designated constraints aren't met.
type ListPoliciesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPoliciesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPoliciesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPoliciesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPoliciesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPoliciesResponseValidationError) ErrorName() string {
	return "ListPoliciesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPoliciesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPoliciesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPoliciesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPoliciesResponseValidationError{}