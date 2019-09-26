// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: virtual_service.proto

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

// Validate checks the field values on VirtualService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *VirtualService) Validate() error {
	if m == nil {
		return nil
	}

	if !strings.HasPrefix(m.GetName(), "ingressGateways/") {
		return VirtualServiceValidationError{
			field:  "Name",
			reason: "value does not have prefix \"ingressGateways/\"",
		}
	}

	if !strings.HasSuffix(m.GetName(), "/virtualService") {
		return VirtualServiceValidationError{
			field:  "Name",
			reason: "value does not have suffix \"/virtualService\"",
		}
	}

	// no validation rules for DisplayName

	if v, ok := interface{}(m.GetVirtualService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VirtualServiceValidationError{
				field:  "VirtualService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// VirtualServiceValidationError is the validation error returned by
// VirtualService.Validate if the designated constraints aren't met.
type VirtualServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VirtualServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VirtualServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VirtualServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VirtualServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VirtualServiceValidationError) ErrorName() string { return "VirtualServiceValidationError" }

// Error satisfies the builtin error interface
func (e VirtualServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVirtualService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VirtualServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VirtualServiceValidationError{}

// Validate checks the field values on GetVirtualServiceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetVirtualServiceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// GetVirtualServiceRequestValidationError is the validation error returned by
// GetVirtualServiceRequest.Validate if the designated constraints aren't met.
type GetVirtualServiceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVirtualServiceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVirtualServiceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVirtualServiceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVirtualServiceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVirtualServiceRequestValidationError) ErrorName() string {
	return "GetVirtualServiceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetVirtualServiceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVirtualServiceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVirtualServiceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVirtualServiceRequestValidationError{}

// Validate checks the field values on UpdateVirtualServiceRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateVirtualServiceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetVirtualService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateVirtualServiceRequestValidationError{
				field:  "VirtualService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateVirtualServiceRequestValidationError is the validation error returned
// by UpdateVirtualServiceRequest.Validate if the designated constraints
// aren't met.
type UpdateVirtualServiceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateVirtualServiceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateVirtualServiceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateVirtualServiceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateVirtualServiceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateVirtualServiceRequestValidationError) ErrorName() string {
	return "UpdateVirtualServiceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateVirtualServiceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateVirtualServiceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateVirtualServiceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateVirtualServiceRequestValidationError{}
