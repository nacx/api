// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: webhook.proto

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

// Validate checks the field values on AlertMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AlertMessage) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetScopeId() < 0 {
		return AlertMessageValidationError{
			field:  "ScopeId",
			reason: "value must be greater than or equal to 0",
		}
	}

	if utf8.RuneCountInString(m.GetScope()) < 1 {
		return AlertMessageValidationError{
			field:  "Scope",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return AlertMessageValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetId0() < 0 {
		return AlertMessageValidationError{
			field:  "Id0",
			reason: "value must be greater than or equal to 0",
		}
	}

	// no validation rules for Id1

	if utf8.RuneCountInString(m.GetRuleName()) < 1 {
		return AlertMessageValidationError{
			field:  "RuleName",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetAlarmMessage()) < 1 {
		return AlertMessageValidationError{
			field:  "AlarmMessage",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetStartTime() < 0 {
		return AlertMessageValidationError{
			field:  "StartTime",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// AlertMessageValidationError is the validation error returned by
// AlertMessage.Validate if the designated constraints aren't met.
type AlertMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlertMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlertMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlertMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlertMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlertMessageValidationError) ErrorName() string { return "AlertMessageValidationError" }

// Error satisfies the builtin error interface
func (e AlertMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlertMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlertMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlertMessageValidationError{}
