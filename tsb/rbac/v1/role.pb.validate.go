// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tsb/rbac/v1/role.proto

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
var _role_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Role) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 1 {
		return RoleValidationError{
			field:  "Rules",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RoleValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Subject) Validate() error {
	if m == nil {
		return nil
	}

	switch m.S.(type) {

	case *Subject_User:
		// no validation rules for User

	case *Subject_Team:
		// no validation rules for Team

	}

	return nil
}

// SubjectValidationError is the validation error returned by Subject.Validate
// if the designated constraints aren't met.
type SubjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectValidationError) ErrorName() string { return "SubjectValidationError" }

// Error satisfies the builtin error interface
func (e SubjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectValidationError{}

// Validate checks the field values on Role_ResourceType with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Role_ResourceType) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetApiGroup()) < 1 {
		return Role_ResourceTypeValidationError{
			field:  "ApiGroup",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// Role_ResourceTypeValidationError is the validation error returned by
// Role_ResourceType.Validate if the designated constraints aren't met.
type Role_ResourceTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Role_ResourceTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Role_ResourceTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Role_ResourceTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Role_ResourceTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Role_ResourceTypeValidationError) ErrorName() string {
	return "Role_ResourceTypeValidationError"
}

// Error satisfies the builtin error interface
func (e Role_ResourceTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole_ResourceType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Role_ResourceTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Role_ResourceTypeValidationError{}

// Validate checks the field values on Role_Rule with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Role_Rule) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetTypes()) < 1 {
		return Role_RuleValidationError{
			field:  "Types",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTypes() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Role_RuleValidationError{
						field:  fmt.Sprintf("Types[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	if len(m.GetPermissions()) < 1 {
		return Role_RuleValidationError{
			field:  "Permissions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// Role_RuleValidationError is the validation error returned by
// Role_Rule.Validate if the designated constraints aren't met.
type Role_RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Role_RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Role_RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Role_RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Role_RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Role_RuleValidationError) ErrorName() string { return "Role_RuleValidationError" }

// Error satisfies the builtin error interface
func (e Role_RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole_Rule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Role_RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Role_RuleValidationError{}