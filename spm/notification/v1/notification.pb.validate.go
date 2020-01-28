// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: notification.proto

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

// Validate checks the field values on Notification with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Notification) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for User

	// no validation rules for Tenant

	// no validation rules for Message

	// no validation rules for Type

	// no validation rules for Read

	// no validation rules for ResourceName

	{
		tmp := m.GetTimestamp()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return NotificationValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// NotificationValidationError is the validation error returned by
// Notification.Validate if the designated constraints aren't met.
type NotificationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationValidationError) ErrorName() string { return "NotificationValidationError" }

// Error satisfies the builtin error interface
func (e NotificationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotification.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationValidationError{}

// Validate checks the field values on ListNotificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListNotificationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return ListNotificationRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetUser()) < 1 {
		return ListNotificationRequestValidationError{
			field:  "User",
			reason: "value length must be at least 1 runes",
		}
	}

	if _, ok := NotificationType_name[int32(m.GetType())]; !ok {
		return ListNotificationRequestValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for Status

	// no validation rules for PageSize

	// no validation rules for PageToken

	return nil
}

// ListNotificationRequestValidationError is the validation error returned by
// ListNotificationRequest.Validate if the designated constraints aren't met.
type ListNotificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotificationRequestValidationError) ErrorName() string {
	return "ListNotificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotificationRequestValidationError{}

// Validate checks the field values on ListNotificationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListNotificationResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListNotificationResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
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

// ListNotificationResponseValidationError is the validation error returned by
// ListNotificationResponse.Validate if the designated constraints aren't met.
type ListNotificationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotificationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotificationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotificationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotificationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotificationResponseValidationError) ErrorName() string {
	return "ListNotificationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotificationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotificationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotificationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotificationResponseValidationError{}

// Validate checks the field values on UpdateNotificationReadStatusRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *UpdateNotificationReadStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return UpdateNotificationReadStatusRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetUser()) < 1 {
		return UpdateNotificationReadStatusRequestValidationError{
			field:  "User",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// UpdateNotificationReadStatusRequestValidationError is the validation error
// returned by UpdateNotificationReadStatusRequest.Validate if the designated
// constraints aren't met.
type UpdateNotificationReadStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNotificationReadStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNotificationReadStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNotificationReadStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNotificationReadStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNotificationReadStatusRequestValidationError) ErrorName() string {
	return "UpdateNotificationReadStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNotificationReadStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNotificationReadStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNotificationReadStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNotificationReadStatusRequestValidationError{}
