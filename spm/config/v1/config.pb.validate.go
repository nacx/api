// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: spm/config/v1/config.proto

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _config_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ConfigurationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ConfigurationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetClusterName()) < 1 {
		return ConfigurationRequestValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// ConfigurationRequestValidationError is the validation error returned by
// ConfigurationRequest.Validate if the designated constraints aren't met.
type ConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationRequestValidationError) ErrorName() string {
	return "ConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationRequestValidationError{}

// Validate checks the field values on ConfigurationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ConfigurationResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetConfigTable() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ConfigurationResponseValidationError{
						field:  fmt.Sprintf("ConfigTable[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ConfigurationResponseValidationError is the validation error returned by
// ConfigurationResponse.Validate if the designated constraints aren't met.
type ConfigurationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationResponseValidationError) ErrorName() string {
	return "ConfigurationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigurationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigurationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationResponseValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Value

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}
