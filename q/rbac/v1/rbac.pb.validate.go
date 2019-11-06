// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rbac.proto

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

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Role) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for DisplayName

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

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Policy) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetBindings()) < 1 {
		return PolicyValidationError{
			field:  "Bindings",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetBindings() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return PolicyValidationError{
						field:  fmt.Sprintf("Bindings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// PolicyValidationError is the validation error returned by Policy.Validate if
// the designated constraints aren't met.
type PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyValidationError) ErrorName() string { return "PolicyValidationError" }

// Error satisfies the builtin error interface
func (e PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyValidationError{}

// Validate checks the field values on Binding with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Binding) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetRole()) < 1 {
		return BindingValidationError{
			field:  "Role",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetSubjects()) < 1 {
		return BindingValidationError{
			field:  "Subjects",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSubjects() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return BindingValidationError{
				field:  fmt.Sprintf("Subjects[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// BindingValidationError is the validation error returned by Binding.Validate
// if the designated constraints aren't met.
type BindingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BindingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BindingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BindingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BindingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BindingValidationError) ErrorName() string { return "BindingValidationError" }

// Error satisfies the builtin error interface
func (e BindingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBinding.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BindingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BindingValidationError{}

// Validate checks the field values on CreateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	if len(m.GetPermissions()) < 1 {
		return CreateRoleRequestValidationError{
			field:  "Permissions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetPermissions() {
		_, _ = idx, item

		if _, ok := Permission_name[int32(item)]; !ok {
			return CreateRoleRequestValidationError{
				field:  fmt.Sprintf("Permissions[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
		}

	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return CreateRoleRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for DisplayName

	return nil
}

// CreateRoleRequestValidationError is the validation error returned by
// CreateRoleRequest.Validate if the designated constraints aren't met.
type CreateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleRequestValidationError) ErrorName() string {
	return "CreateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleRequestValidationError{}

// Validate checks the field values on UpdateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return UpdateRoleRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Description

	if len(m.GetPermissions()) < 1 {
		return UpdateRoleRequestValidationError{
			field:  "Permissions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetPermissions() {
		_, _ = idx, item

		if _, ok := Permission_name[int32(item)]; !ok {
			return UpdateRoleRequestValidationError{
				field:  fmt.Sprintf("Permissions[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
		}

	}

	// no validation rules for Name

	// no validation rules for DisplayName

	return nil
}

// UpdateRoleRequestValidationError is the validation error returned by
// UpdateRoleRequest.Validate if the designated constraints aren't met.
type UpdateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleRequestValidationError) ErrorName() string {
	return "UpdateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleRequestValidationError{}

// Validate checks the field values on ListRolesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListRolesRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListRolesRequestValidationError is the validation error returned by
// ListRolesRequest.Validate if the designated constraints aren't met.
type ListRolesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesRequestValidationError) ErrorName() string { return "ListRolesRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRolesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesRequestValidationError{}

// Validate checks the field values on ListRolesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListRolesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListRolesResponseValidationError{
						field:  fmt.Sprintf("Roles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ListRolesResponseValidationError is the validation error returned by
// ListRolesResponse.Validate if the designated constraints aren't met.
type ListRolesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesResponseValidationError) ErrorName() string {
	return "ListRolesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRolesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesResponseValidationError{}

// Validate checks the field values on GetRoleRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return GetRoleRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Name

	return nil
}

// GetRoleRequestValidationError is the validation error returned by
// GetRoleRequest.Validate if the designated constraints aren't met.
type GetRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleRequestValidationError) ErrorName() string { return "GetRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleRequestValidationError{}

// Validate checks the field values on DeleteRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return DeleteRoleRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Name

	return nil
}

// DeleteRoleRequestValidationError is the validation error returned by
// DeleteRoleRequest.Validate if the designated constraints aren't met.
type DeleteRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleRequestValidationError) ErrorName() string {
	return "DeleteRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleRequestValidationError{}
