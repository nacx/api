// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configproducer.proto

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

// Validate checks the field values on ConfigDownloadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ConfigDownloadRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Configtype

	// no validation rules for Cluster

	// no validation rules for Environment

	// no validation rules for Tenant

	return nil
}

// ConfigDownloadRequestValidationError is the validation error returned by
// ConfigDownloadRequest.Validate if the designated constraints aren't met.
type ConfigDownloadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigDownloadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigDownloadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigDownloadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigDownloadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigDownloadRequestValidationError) ErrorName() string {
	return "ConfigDownloadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigDownloadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigDownloadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigDownloadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigDownloadRequestValidationError{}

// Validate checks the field values on ConfigData with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ConfigData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Configtype

	// no validation rules for Cluster

	// no validation rules for Environment

	// no validation rules for Tenant

	// no validation rules for Payload

	return nil
}

// ConfigDataValidationError is the validation error returned by
// ConfigData.Validate if the designated constraints aren't met.
type ConfigDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigDataValidationError) ErrorName() string { return "ConfigDataValidationError" }

// Error satisfies the builtin error interface
func (e ConfigDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigDataValidationError{}
