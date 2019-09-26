// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: endpoint.proto

package v1

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

// Validate checks the field values on Endpoint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Endpoint) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Service

	// no validation rules for Tenant

	// no validation rules for Workspace

	// no validation rules for Address

	// no validation rules for Ports

	// no validation rules for Labels

	// no validation rules for Locality

	// no validation rules for Weight

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

// Validate checks the field values on CreateEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Service

	// no validation rules for Tenant

	// no validation rules for Workspace

	// no validation rules for Address

	// no validation rules for Ports

	// no validation rules for Labels

	// no validation rules for Locality

	// no validation rules for Weight

	return nil
}

// CreateEndpointRequestValidationError is the validation error returned by
// CreateEndpointRequest.Validate if the designated constraints aren't met.
type CreateEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEndpointRequestValidationError) ErrorName() string {
	return "CreateEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEndpointRequestValidationError{}

// Validate checks the field values on GetEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Workspace

	// no validation rules for Tenant

	return nil
}

// GetEndpointRequestValidationError is the validation error returned by
// GetEndpointRequest.Validate if the designated constraints aren't met.
type GetEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEndpointRequestValidationError) ErrorName() string {
	return "GetEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEndpointRequestValidationError{}

// Validate checks the field values on ListWorkspaceEndpointRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListWorkspaceEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Tenant

	// no validation rules for Workspace

	return nil
}

// ListWorkspaceEndpointRequestValidationError is the validation error returned
// by ListWorkspaceEndpointRequest.Validate if the designated constraints
// aren't met.
type ListWorkspaceEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWorkspaceEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWorkspaceEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWorkspaceEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWorkspaceEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWorkspaceEndpointRequestValidationError) ErrorName() string {
	return "ListWorkspaceEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWorkspaceEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWorkspaceEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWorkspaceEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWorkspaceEndpointRequestValidationError{}

// Validate checks the field values on ListServiceEndpointRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListServiceEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	// no validation rules for Tenant

	// no validation rules for Workspace

	return nil
}

// ListServiceEndpointRequestValidationError is the validation error returned
// by ListServiceEndpointRequest.Validate if the designated constraints aren't met.
type ListServiceEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListServiceEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListServiceEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListServiceEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListServiceEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListServiceEndpointRequestValidationError) ErrorName() string {
	return "ListServiceEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListServiceEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListServiceEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListServiceEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListServiceEndpointRequestValidationError{}

// Validate checks the field values on ListServiceSubsetEndpointRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListServiceSubsetEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	// no validation rules for Tenant

	// no validation rules for Workspace

	// no validation rules for Subset

	return nil
}

// ListServiceSubsetEndpointRequestValidationError is the validation error
// returned by ListServiceSubsetEndpointRequest.Validate if the designated
// constraints aren't met.
type ListServiceSubsetEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListServiceSubsetEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListServiceSubsetEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListServiceSubsetEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListServiceSubsetEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListServiceSubsetEndpointRequestValidationError) ErrorName() string {
	return "ListServiceSubsetEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListServiceSubsetEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListServiceSubsetEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListServiceSubsetEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListServiceSubsetEndpointRequestValidationError{}

// Validate checks the field values on ListEndpointResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListEndpointResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEndpointResponseValidationError{
					field:  fmt.Sprintf("Endpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListEndpointResponseValidationError is the validation error returned by
// ListEndpointResponse.Validate if the designated constraints aren't met.
type ListEndpointResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEndpointResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEndpointResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEndpointResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEndpointResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEndpointResponseValidationError) ErrorName() string {
	return "ListEndpointResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListEndpointResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEndpointResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEndpointResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEndpointResponseValidationError{}

// Validate checks the field values on UpdateEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Workspace

	// no validation rules for Tenant

	// no validation rules for Address

	// no validation rules for Ports

	// no validation rules for Labels

	// no validation rules for Locality

	// no validation rules for Weight

	return nil
}

// UpdateEndpointRequestValidationError is the validation error returned by
// UpdateEndpointRequest.Validate if the designated constraints aren't met.
type UpdateEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEndpointRequestValidationError) ErrorName() string {
	return "UpdateEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEndpointRequestValidationError{}

// Validate checks the field values on DeleteEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteEndpointRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Workspace

	// no validation rules for Tenant

	return nil
}

// DeleteEndpointRequestValidationError is the validation error returned by
// DeleteEndpointRequest.Validate if the designated constraints aren't met.
type DeleteEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEndpointRequestValidationError) ErrorName() string {
	return "DeleteEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEndpointRequestValidationError{}
