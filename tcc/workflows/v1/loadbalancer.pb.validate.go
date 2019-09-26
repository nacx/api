// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: loadbalancer.proto

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

// Validate checks the field values on LoadBalancerWorkflowUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LoadBalancerWorkflowUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Loadbalancer

	// no validation rules for Service

	return nil
}

// LoadBalancerWorkflowUserRequestValidationError is the validation error
// returned by LoadBalancerWorkflowUserRequest.Validate if the designated
// constraints aren't met.
type LoadBalancerWorkflowUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoadBalancerWorkflowUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoadBalancerWorkflowUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoadBalancerWorkflowUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoadBalancerWorkflowUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoadBalancerWorkflowUserRequestValidationError) ErrorName() string {
	return "LoadBalancerWorkflowUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LoadBalancerWorkflowUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoadBalancerWorkflowUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoadBalancerWorkflowUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoadBalancerWorkflowUserRequestValidationError{}

// Validate checks the field values on Status with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Status) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	// no validation rules for Approved

	// no validation rules for Published

	// no validation rules for Note

	return nil
}

// StatusValidationError is the validation error returned by Status.Validate if
// the designated constraints aren't met.
type StatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusValidationError) ErrorName() string { return "StatusValidationError" }

// Error satisfies the builtin error interface
func (e StatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusValidationError{}

// Validate checks the field values on GetStatus with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	return nil
}

// GetStatusValidationError is the validation error returned by
// GetStatus.Validate if the designated constraints aren't met.
type GetStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStatusValidationError) ErrorName() string { return "GetStatusValidationError" }

// Error satisfies the builtin error interface
func (e GetStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStatusValidationError{}

// Validate checks the field values on ListRequests with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListRequests) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListRequestsValidationError is the validation error returned by
// ListRequests.Validate if the designated constraints aren't met.
type ListRequestsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestsValidationError) ErrorName() string { return "ListRequestsValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequests.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestsValidationError{}

// Validate checks the field values on ListPendingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPendingResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPendingRequests() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPendingResponseValidationError{
					field:  fmt.Sprintf("PendingRequests[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListPendingResponseValidationError is the validation error returned by
// ListPendingResponse.Validate if the designated constraints aren't met.
type ListPendingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPendingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPendingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPendingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPendingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPendingResponseValidationError) ErrorName() string {
	return "ListPendingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPendingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPendingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPendingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPendingResponseValidationError{}

// Validate checks the field values on ApproveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ApproveRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	return nil
}

// ApproveRequestValidationError is the validation error returned by
// ApproveRequest.Validate if the designated constraints aren't met.
type ApproveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApproveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApproveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApproveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApproveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApproveRequestValidationError) ErrorName() string { return "ApproveRequestValidationError" }

// Error satisfies the builtin error interface
func (e ApproveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApproveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApproveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApproveRequestValidationError{}

// Validate checks the field values on DenyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DenyRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	return nil
}

// DenyRequestValidationError is the validation error returned by
// DenyRequest.Validate if the designated constraints aren't met.
type DenyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DenyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DenyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DenyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DenyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DenyRequestValidationError) ErrorName() string { return "DenyRequestValidationError" }

// Error satisfies the builtin error interface
func (e DenyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDenyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DenyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DenyRequestValidationError{}

// Validate checks the field values on CancelRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CancelRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	return nil
}

// CancelRequestValidationError is the validation error returned by
// CancelRequest.Validate if the designated constraints aren't met.
type CancelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelRequestValidationError) ErrorName() string { return "CancelRequestValidationError" }

// Error satisfies the builtin error interface
func (e CancelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelRequestValidationError{}

// Validate checks the field values on ApproveResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ApproveResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	// no validation rules for Note

	return nil
}

// ApproveResponseValidationError is the validation error returned by
// ApproveResponse.Validate if the designated constraints aren't met.
type ApproveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApproveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApproveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApproveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApproveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApproveResponseValidationError) ErrorName() string { return "ApproveResponseValidationError" }

// Error satisfies the builtin error interface
func (e ApproveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApproveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApproveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApproveResponseValidationError{}

// Validate checks the field values on DenyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DenyResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	// no validation rules for Note

	return nil
}

// DenyResponseValidationError is the validation error returned by
// DenyResponse.Validate if the designated constraints aren't met.
type DenyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DenyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DenyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DenyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DenyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DenyResponseValidationError) ErrorName() string { return "DenyResponseValidationError" }

// Error satisfies the builtin error interface
func (e DenyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDenyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DenyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DenyResponseValidationError{}

// Validate checks the field values on CancelResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CancelResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Requestid

	// no validation rules for Note

	return nil
}

// CancelResponseValidationError is the validation error returned by
// CancelResponse.Validate if the designated constraints aren't met.
type CancelResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelResponseValidationError) ErrorName() string { return "CancelResponseValidationError" }

// Error satisfies the builtin error interface
func (e CancelResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelResponseValidationError{}

// Validate checks the field values on LoadBalancerWorkflowOwnerRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *LoadBalancerWorkflowOwnerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Loadbalancer

	// no validation rules for Service

	if v, ok := interface{}(m.GetTls()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoadBalancerWorkflowOwnerRequestValidationError{
				field:  "Tls",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LoadBalancerWorkflowOwnerRequestValidationError is the validation error
// returned by LoadBalancerWorkflowOwnerRequest.Validate if the designated
// constraints aren't met.
type LoadBalancerWorkflowOwnerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoadBalancerWorkflowOwnerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoadBalancerWorkflowOwnerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoadBalancerWorkflowOwnerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoadBalancerWorkflowOwnerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoadBalancerWorkflowOwnerRequestValidationError) ErrorName() string {
	return "LoadBalancerWorkflowOwnerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LoadBalancerWorkflowOwnerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoadBalancerWorkflowOwnerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoadBalancerWorkflowOwnerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoadBalancerWorkflowOwnerRequestValidationError{}

// Validate checks the field values on TLSSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TLSSettings) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Sni

	// no validation rules for ServerCertificate

	// no validation rules for PrivateKey

	// no validation rules for CaCertificates

	// no validation rules for SecretName

	return nil
}

// TLSSettingsValidationError is the validation error returned by
// TLSSettings.Validate if the designated constraints aren't met.
type TLSSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TLSSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TLSSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TLSSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TLSSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TLSSettingsValidationError) ErrorName() string { return "TLSSettingsValidationError" }

// Error satisfies the builtin error interface
func (e TLSSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTLSSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TLSSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TLSSettingsValidationError{}

// Validate checks the field values on ListPendingResponse_PendingRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListPendingResponse_PendingRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Loadbalancer

	// no validation rules for Service

	// no validation rules for Operation

	return nil
}

// ListPendingResponse_PendingRequestValidationError is the validation error
// returned by ListPendingResponse_PendingRequest.Validate if the designated
// constraints aren't met.
type ListPendingResponse_PendingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPendingResponse_PendingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPendingResponse_PendingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPendingResponse_PendingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPendingResponse_PendingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPendingResponse_PendingRequestValidationError) ErrorName() string {
	return "ListPendingResponse_PendingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPendingResponse_PendingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPendingResponse_PendingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPendingResponse_PendingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPendingResponse_PendingRequestValidationError{}