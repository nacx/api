// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ingress_gateway.proto

package v1alpha1

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

// Validate checks the field values on IngressGateway with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *IngressGateway) Validate() error {
	if m == nil {
		return nil
	}

	if !strings.HasPrefix(m.GetName(), "ingressGateways/") {
		return IngressGatewayValidationError{
			field:  "Name",
			reason: "value does not have prefix \"ingressGateways/\"",
		}
	}

	// no validation rules for DisplayName

	// no validation rules for Zone

	if v, ok := interface{}(m.GetGateway()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IngressGatewayValidationError{
				field:  "Gateway",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// IngressGatewayValidationError is the validation error returned by
// IngressGateway.Validate if the designated constraints aren't met.
type IngressGatewayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IngressGatewayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IngressGatewayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IngressGatewayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IngressGatewayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IngressGatewayValidationError) ErrorName() string { return "IngressGatewayValidationError" }

// Error satisfies the builtin error interface
func (e IngressGatewayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIngressGateway.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IngressGatewayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IngressGatewayValidationError{}

// Validate checks the field values on ListIngressGatewaysRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListIngressGatewaysRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListIngressGatewaysRequestValidationError is the validation error returned
// by ListIngressGatewaysRequest.Validate if the designated constraints aren't met.
type ListIngressGatewaysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListIngressGatewaysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListIngressGatewaysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListIngressGatewaysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListIngressGatewaysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListIngressGatewaysRequestValidationError) ErrorName() string {
	return "ListIngressGatewaysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListIngressGatewaysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListIngressGatewaysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListIngressGatewaysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListIngressGatewaysRequestValidationError{}

// Validate checks the field values on ListIngressGatewaysResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListIngressGatewaysResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetIngressGateways() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListIngressGatewaysResponseValidationError{
					field:  fmt.Sprintf("IngressGateways[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListIngressGatewaysResponseValidationError is the validation error returned
// by ListIngressGatewaysResponse.Validate if the designated constraints
// aren't met.
type ListIngressGatewaysResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListIngressGatewaysResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListIngressGatewaysResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListIngressGatewaysResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListIngressGatewaysResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListIngressGatewaysResponseValidationError) ErrorName() string {
	return "ListIngressGatewaysResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListIngressGatewaysResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListIngressGatewaysResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListIngressGatewaysResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListIngressGatewaysResponseValidationError{}

// Validate checks the field values on GetIngressGatewayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetIngressGatewayRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// GetIngressGatewayRequestValidationError is the validation error returned by
// GetIngressGatewayRequest.Validate if the designated constraints aren't met.
type GetIngressGatewayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIngressGatewayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIngressGatewayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIngressGatewayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIngressGatewayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIngressGatewayRequestValidationError) ErrorName() string {
	return "GetIngressGatewayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetIngressGatewayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIngressGatewayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIngressGatewayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIngressGatewayRequestValidationError{}

// Validate checks the field values on CreateIngressGatewayRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateIngressGatewayRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIngressGateway()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateIngressGatewayRequestValidationError{
				field:  "IngressGateway",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateIngressGatewayRequestValidationError is the validation error returned
// by CreateIngressGatewayRequest.Validate if the designated constraints
// aren't met.
type CreateIngressGatewayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIngressGatewayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIngressGatewayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIngressGatewayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIngressGatewayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIngressGatewayRequestValidationError) ErrorName() string {
	return "CreateIngressGatewayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIngressGatewayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIngressGatewayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIngressGatewayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIngressGatewayRequestValidationError{}

// Validate checks the field values on UpdateIngressGatewayRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateIngressGatewayRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIngressGateway()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateIngressGatewayRequestValidationError{
				field:  "IngressGateway",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateIngressGatewayRequestValidationError is the validation error returned
// by UpdateIngressGatewayRequest.Validate if the designated constraints
// aren't met.
type UpdateIngressGatewayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateIngressGatewayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateIngressGatewayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateIngressGatewayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateIngressGatewayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateIngressGatewayRequestValidationError) ErrorName() string {
	return "UpdateIngressGatewayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateIngressGatewayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateIngressGatewayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateIngressGatewayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateIngressGatewayRequestValidationError{}

// Validate checks the field values on DeleteIngressGatewayRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteIngressGatewayRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// DeleteIngressGatewayRequestValidationError is the validation error returned
// by DeleteIngressGatewayRequest.Validate if the designated constraints
// aren't met.
type DeleteIngressGatewayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteIngressGatewayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteIngressGatewayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteIngressGatewayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteIngressGatewayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteIngressGatewayRequestValidationError) ErrorName() string {
	return "DeleteIngressGatewayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteIngressGatewayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteIngressGatewayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteIngressGatewayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteIngressGatewayRequestValidationError{}
