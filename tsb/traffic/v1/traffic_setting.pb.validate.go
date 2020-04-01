// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: traffic_setting.proto

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

// Validate checks the field values on TrafficSetting with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TrafficSetting) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetReachability()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TrafficSettingValidationError{
					field:  "Reachability",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if !_TrafficSetting_EgressGateway_Pattern.MatchString(m.GetEgressGateway()) {
		return TrafficSettingValidationError{
			field:  "EgressGateway",
			reason: "value does not match regex pattern \"^[^/]+/[^/]+$\"",
		}
	}

	{
		tmp := m.GetResilience()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TrafficSettingValidationError{
					field:  "Resilience",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// TrafficSettingValidationError is the validation error returned by
// TrafficSetting.Validate if the designated constraints aren't met.
type TrafficSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrafficSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrafficSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrafficSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrafficSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrafficSettingValidationError) ErrorName() string { return "TrafficSettingValidationError" }

// Error satisfies the builtin error interface
func (e TrafficSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrafficSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrafficSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrafficSettingValidationError{}

var _TrafficSetting_EgressGateway_Pattern = regexp.MustCompile("^[^/]+/[^/]+$")

// Validate checks the field values on ResilienceSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResilienceSettings) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetHttpRequestTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ResilienceSettingsValidationError{
					field:  "HttpRequestTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetHttpRetries()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ResilienceSettingsValidationError{
					field:  "HttpRetries",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTcpKeepalive()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ResilienceSettingsValidationError{
					field:  "TcpKeepalive",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for CircuitBreakerSensitivity

	return nil
}

// ResilienceSettingsValidationError is the validation error returned by
// ResilienceSettings.Validate if the designated constraints aren't met.
type ResilienceSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResilienceSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResilienceSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResilienceSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResilienceSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResilienceSettingsValidationError) ErrorName() string {
	return "ResilienceSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e ResilienceSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResilienceSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResilienceSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResilienceSettingsValidationError{}

// Validate checks the field values on HTTPRetry with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HTTPRetry) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAttempts() < 1 {
		return HTTPRetryValidationError{
			field:  "Attempts",
			reason: "value must be greater than or equal to 1",
		}
	}

	{
		tmp := m.GetPerTryTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HTTPRetryValidationError{
					field:  "PerTryTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for RetryOn

	return nil
}

// HTTPRetryValidationError is the validation error returned by
// HTTPRetry.Validate if the designated constraints aren't met.
type HTTPRetryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPRetryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPRetryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPRetryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPRetryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPRetryValidationError) ErrorName() string { return "HTTPRetryValidationError" }

// Error satisfies the builtin error interface
func (e HTTPRetryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPRetry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPRetryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPRetryValidationError{}

// Validate checks the field values on ReachabilitySettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReachabilitySettings) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Mode

	return nil
}

// ReachabilitySettingsValidationError is the validation error returned by
// ReachabilitySettings.Validate if the designated constraints aren't met.
type ReachabilitySettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReachabilitySettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReachabilitySettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReachabilitySettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReachabilitySettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReachabilitySettingsValidationError) ErrorName() string {
	return "ReachabilitySettingsValidationError"
}

// Error satisfies the builtin error interface
func (e ReachabilitySettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReachabilitySettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReachabilitySettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReachabilitySettingsValidationError{}
