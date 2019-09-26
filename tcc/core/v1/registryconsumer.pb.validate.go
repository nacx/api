// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: registryconsumer.proto

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

// Validate checks the field values on RegistryUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegistryUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Registrytype

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for Tenant

	// no validation rules for Payload

	return nil
}

// RegistryUpdateRequestValidationError is the validation error returned by
// RegistryUpdateRequest.Validate if the designated constraints aren't met.
type RegistryUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryUpdateRequestValidationError) ErrorName() string {
	return "RegistryUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegistryUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistryUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryUpdateRequestValidationError{}
