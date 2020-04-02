// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: audit/v1/audit.proto

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
var _audit_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AuditLog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AuditLog) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCreateTime() == nil {
		return AuditLogValidationError{
			field:  "CreateTime",
			reason: "value is required",
		}
	}

	if utf8.RuneCountInString(m.GetSeverity()) < 1 {
		return AuditLogValidationError{
			field:  "Severity",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetKind()) < 1 {
		return AuditLogValidationError{
			field:  "Kind",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetMessage()) < 1 {
		return AuditLogValidationError{
			field:  "Message",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetTriggeredBy()) < 1 {
		return AuditLogValidationError{
			field:  "TriggeredBy",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Properties

	return nil
}

// AuditLogValidationError is the validation error returned by
// AuditLog.Validate if the designated constraints aren't met.
type AuditLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuditLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuditLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuditLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuditLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuditLogValidationError) ErrorName() string { return "AuditLogValidationError" }

// Error satisfies the builtin error interface
func (e AuditLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuditLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuditLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuditLogValidationError{}

// Validate checks the field values on ListAuditLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAuditLogsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Count

	return nil
}

// ListAuditLogsRequestValidationError is the validation error returned by
// ListAuditLogsRequest.Validate if the designated constraints aren't met.
type ListAuditLogsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAuditLogsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAuditLogsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAuditLogsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAuditLogsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAuditLogsRequestValidationError) ErrorName() string {
	return "ListAuditLogsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAuditLogsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAuditLogsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAuditLogsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAuditLogsRequestValidationError{}

// Validate checks the field values on ListAuditLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAuditLogsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAuditLogs() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListAuditLogsResponseValidationError{
						field:  fmt.Sprintf("AuditLogs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ListAuditLogsResponseValidationError is the validation error returned by
// ListAuditLogsResponse.Validate if the designated constraints aren't met.
type ListAuditLogsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAuditLogsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAuditLogsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAuditLogsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAuditLogsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAuditLogsResponseValidationError) ErrorName() string {
	return "ListAuditLogsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAuditLogsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAuditLogsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAuditLogsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAuditLogsResponseValidationError{}
