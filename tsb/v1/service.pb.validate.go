// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service.proto

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

// Validate checks the field values on Object with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Object) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ApiVersion

	// no validation rules for Kind

	{
		tmp := m.GetMetadata()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ObjectValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Spec

	// no validation rules for Status

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

// Validate checks the field values on ObjectMeta with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ObjectMeta) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Namespace

	// no validation rules for Tenant

	// no validation rules for Workspace

	// no validation rules for Group

	// no validation rules for ResourceVersion

	// no validation rules for Labels

	// no validation rules for Annotations

	return nil
}

// ObjectMetaValidationError is the validation error returned by
// ObjectMeta.Validate if the designated constraints aren't met.
type ObjectMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectMetaValidationError) ErrorName() string { return "ObjectMetaValidationError" }

// Error satisfies the builtin error interface
func (e ObjectMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObjectMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectMetaValidationError{}
