// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: alert.proto

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

// Validate checks the field values on Alert with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Alert) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Tenant

	// no validation rules for RuleName

	// no validation rules for Application

	// no validation rules for Message

	// no validation rules for Status

	// no validation rules for Acknowledged

	{
		tmp := m.GetTimestamp()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return AlertValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Severity

	// no validation rules for Properties

	return nil
}

// AlertValidationError is the validation error returned by Alert.Validate if
// the designated constraints aren't met.
type AlertValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlertValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlertValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlertValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlertValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlertValidationError) ErrorName() string { return "AlertValidationError" }

// Error satisfies the builtin error interface
func (e AlertValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlert.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlertValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlertValidationError{}

// Validate checks the field values on ListAlertsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListAlertsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return ListAlertsRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Status

	// no validation rules for Severity

	// no validation rules for PageSize

	// no validation rules for PageToken

	return nil
}

// ListAlertsRequestValidationError is the validation error returned by
// ListAlertsRequest.Validate if the designated constraints aren't met.
type ListAlertsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAlertsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAlertsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAlertsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAlertsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAlertsRequestValidationError) ErrorName() string {
	return "ListAlertsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAlertsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAlertsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAlertsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAlertsRequestValidationError{}

// Validate checks the field values on ListAlertsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAlertsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAlerts() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListAlertsResponseValidationError{
						field:  fmt.Sprintf("Alerts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for NextPageToken

	return nil
}

// ListAlertsResponseValidationError is the validation error returned by
// ListAlertsResponse.Validate if the designated constraints aren't met.
type ListAlertsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAlertsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAlertsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAlertsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAlertsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAlertsResponseValidationError) ErrorName() string {
	return "ListAlertsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAlertsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAlertsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAlertsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAlertsResponseValidationError{}

// Validate checks the field values on UpdateAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateAlertRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return UpdateAlertRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return UpdateAlertRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Status

	// no validation rules for Acknowledged

	return nil
}

// UpdateAlertRequestValidationError is the validation error returned by
// UpdateAlertRequest.Validate if the designated constraints aren't met.
type UpdateAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAlertRequestValidationError) ErrorName() string {
	return "UpdateAlertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAlertRequestValidationError{}

// Validate checks the field values on DeleteAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteAlertRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return DeleteAlertRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return DeleteAlertRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// DeleteAlertRequestValidationError is the validation error returned by
// DeleteAlertRequest.Validate if the designated constraints aren't met.
type DeleteAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAlertRequestValidationError) ErrorName() string {
	return "DeleteAlertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAlertRequestValidationError{}

// Validate checks the field values on GetAlertRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetAlertRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return GetAlertRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return GetAlertRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetAlertRequestValidationError is the validation error returned by
// GetAlertRequest.Validate if the designated constraints aren't met.
type GetAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlertRequestValidationError) ErrorName() string { return "GetAlertRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlertRequestValidationError{}
