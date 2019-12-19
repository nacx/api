// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: settings.proto

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

// Validate checks the field values on Settings with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Settings) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Rules

	return nil
}

// SettingsValidationError is the validation error returned by
// Settings.Validate if the designated constraints aren't met.
type SettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SettingsValidationError) ErrorName() string { return "SettingsValidationError" }

// Error satisfies the builtin error interface
func (e SettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SettingsValidationError{}

// Validate checks the field values on Rule with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Rule) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetMetricName()) < 6 {
		return RuleValidationError{
			field:  "MetricName",
			reason: "value length must be at least 6 runes",
		}
	}

	if !strings.HasSuffix(m.GetMetricName(), "_rule") {
		return RuleValidationError{
			field:  "MetricName",
			reason: "value does not have suffix \"_rule\"",
		}
	}

	if val := m.GetThreshold(); val < 0 || val > 100 {
		return RuleValidationError{
			field:  "Threshold",
			reason: "value must be inside range [0, 100]",
		}
	}

	// no validation rules for Operator

	if d := m.GetPeriod(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return RuleValidationError{
				field:  "Period",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		lt := time.Duration(1*time.Second + 0*time.Nanosecond)

		if dur >= lt {
			return RuleValidationError{
				field:  "Period",
				reason: "value must be less than 1s",
			}
		}

	}

	if m.GetCount() < 0 {
		return RuleValidationError{
			field:  "Count",
			reason: "value must be greater than or equal to 0",
		}
	}

	{
		tmp := m.GetSilencePeriod()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RuleValidationError{
					field:  "SilencePeriod",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for DisplayName

	// no validation rules for Enabled

	return nil
}

// RuleValidationError is the validation error returned by Rule.Validate if the
// designated constraints aren't met.
type RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuleValidationError) ErrorName() string { return "RuleValidationError" }

// Error satisfies the builtin error interface
func (e RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuleValidationError{}
