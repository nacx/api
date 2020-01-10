// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: activity.proto

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

// Validate checks the field values on Activity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Activity) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Tenant

	{
		tmp := m.GetTimestamp()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ActivityValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetActor()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ActivityValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Verb

	{
		tmp := m.GetObject()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ActivityValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTarget()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ActivityValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Properties

	return nil
}

// ActivityValidationError is the validation error returned by
// Activity.Validate if the designated constraints aren't met.
type ActivityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActivityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActivityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActivityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActivityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActivityValidationError) ErrorName() string { return "ActivityValidationError" }

// Error satisfies the builtin error interface
func (e ActivityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActivity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActivityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActivityValidationError{}

// Validate checks the field values on Object with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Object) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	return nil
}

// ObjectValidationError is the validation error returned by Object.Validate if
// the designated constraints aren't met.
type ObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectValidationError) ErrorName() string { return "ObjectValidationError" }

// Error satisfies the builtin error interface
func (e ObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectValidationError{}

// Validate checks the field values on ListActivitiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListActivitiesRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return ListActivitiesRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for PageSize

	// no validation rules for PageToken

	switch m.Filter.(type) {

	case *ListActivitiesRequest_Actor:
		// no validation rules for Actor

	case *ListActivitiesRequest_ObjectOrTarget:
		// no validation rules for ObjectOrTarget

	}

	return nil
}

// ListActivitiesRequestValidationError is the validation error returned by
// ListActivitiesRequest.Validate if the designated constraints aren't met.
type ListActivitiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListActivitiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListActivitiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListActivitiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListActivitiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListActivitiesRequestValidationError) ErrorName() string {
	return "ListActivitiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListActivitiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListActivitiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListActivitiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListActivitiesRequestValidationError{}

// Validate checks the field values on ListActivitiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListActivitiesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetActivities() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListActivitiesResponseValidationError{
						field:  fmt.Sprintf("Activities[%v]", idx),
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

// ListActivitiesResponseValidationError is the validation error returned by
// ListActivitiesResponse.Validate if the designated constraints aren't met.
type ListActivitiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListActivitiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListActivitiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListActivitiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListActivitiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListActivitiesResponseValidationError) ErrorName() string {
	return "ListActivitiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListActivitiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListActivitiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListActivitiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListActivitiesResponseValidationError{}