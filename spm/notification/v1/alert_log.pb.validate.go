// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: spm/notification/v1/alert_log.proto

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
var _alert_log_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AlertLog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AlertLog) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 0 {
		return AlertLogValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 0",
		}
	}

	if m.GetLastTriggered() == nil {
		return AlertLogValidationError{
			field:  "LastTriggered",
			reason: "value is required",
		}
	}

	if utf8.RuneCountInString(m.GetRuleName()) < 1 {
		return AlertLogValidationError{
			field:  "RuleName",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetApplication()) < 1 {
		return AlertLogValidationError{
			field:  "Application",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetMessage()) < 1 {
		return AlertLogValidationError{
			field:  "Message",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Read

	// no validation rules for Properties

	return nil
}

// AlertLogValidationError is the validation error returned by
// AlertLog.Validate if the designated constraints aren't met.
type AlertLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlertLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlertLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlertLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlertLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlertLogValidationError) ErrorName() string { return "AlertLogValidationError" }

// Error satisfies the builtin error interface
func (e AlertLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlertLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlertLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlertLogValidationError{}

// Validate checks the field values on ListAlertLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAlertLogsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// ListAlertLogsRequestValidationError is the validation error returned by
// ListAlertLogsRequest.Validate if the designated constraints aren't met.
type ListAlertLogsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAlertLogsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAlertLogsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAlertLogsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAlertLogsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAlertLogsRequestValidationError) ErrorName() string {
	return "ListAlertLogsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAlertLogsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAlertLogsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAlertLogsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAlertLogsRequestValidationError{}

// Validate checks the field values on ListAlertLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAlertLogsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAlertLogs() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListAlertLogsResponseValidationError{
						field:  fmt.Sprintf("AlertLogs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ListAlertLogsResponseValidationError is the validation error returned by
// ListAlertLogsResponse.Validate if the designated constraints aren't met.
type ListAlertLogsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAlertLogsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAlertLogsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAlertLogsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAlertLogsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAlertLogsResponseValidationError) ErrorName() string {
	return "ListAlertLogsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAlertLogsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAlertLogsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAlertLogsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAlertLogsResponseValidationError{}

// Validate checks the field values on DeleteAlertLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteAlertLogsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteAlertLogsRequestValidationError is the validation error returned by
// DeleteAlertLogsRequest.Validate if the designated constraints aren't met.
type DeleteAlertLogsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAlertLogsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAlertLogsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAlertLogsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAlertLogsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAlertLogsRequestValidationError) ErrorName() string {
	return "DeleteAlertLogsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAlertLogsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAlertLogsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAlertLogsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAlertLogsRequestValidationError{}

// Validate checks the field values on UpdateAlertLogReadStatusRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateAlertLogReadStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateAlertLogReadStatusRequestValidationError is the validation error
// returned by UpdateAlertLogReadStatusRequest.Validate if the designated
// constraints aren't met.
type UpdateAlertLogReadStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAlertLogReadStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAlertLogReadStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAlertLogReadStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAlertLogReadStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAlertLogReadStatusRequestValidationError) ErrorName() string {
	return "UpdateAlertLogReadStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAlertLogReadStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAlertLogReadStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAlertLogReadStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAlertLogReadStatusRequestValidationError{}
