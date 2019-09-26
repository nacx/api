// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: team.proto

package invv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for DisplayName

	for idx, item := range m.GetGroups() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  fmt.Sprintf("Groups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Kind

	// no validation rules for Email

	// no validation rules for LoginName

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on Group with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Group) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for DisplayName

	for idx, item := range m.GetParents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GroupValidationError{
					field:  fmt.Sprintf("Parents[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GroupValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GroupValidationError is the validation error returned by Group.Validate if
// the designated constraints aren't met.
type GroupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupValidationError) ErrorName() string { return "GroupValidationError" }

// Error satisfies the builtin error interface
func (e GroupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroup.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupValidationError{}

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Role) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

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

// Validate checks the field values on RoleBinding with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RoleBinding) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleBindingValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleBindingValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPrincipals() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RoleBindingValidationError{
					field:  fmt.Sprintf("Principals[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RoleBindingValidationError is the validation error returned by
// RoleBinding.Validate if the designated constraints aren't met.
type RoleBindingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleBindingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleBindingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleBindingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleBindingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleBindingValidationError) ErrorName() string { return "RoleBindingValidationError" }

// Error satisfies the builtin error interface
func (e RoleBindingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoleBinding.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleBindingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleBindingValidationError{}

// Validate checks the field values on ResourceRef with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResourceRef) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for DisplayName

	// no validation rules for Type

	return nil
}

// ResourceRefValidationError is the validation error returned by
// ResourceRef.Validate if the designated constraints aren't met.
type ResourceRefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceRefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceRefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceRefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceRefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceRefValidationError) ErrorName() string { return "ResourceRefValidationError" }

// Error satisfies the builtin error interface
func (e ResourceRefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceRef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceRefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceRefValidationError{}

// Validate checks the field values on ListGroupsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListGroupsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListGroupsRequestValidationError is the validation error returned by
// ListGroupsRequest.Validate if the designated constraints aren't met.
type ListGroupsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGroupsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGroupsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGroupsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGroupsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGroupsRequestValidationError) ErrorName() string {
	return "ListGroupsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListGroupsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGroupsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGroupsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGroupsRequestValidationError{}

// Validate checks the field values on ListGroupsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListGroupsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetGroups() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListGroupsResponseValidationError{
					field:  fmt.Sprintf("Groups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListGroupsResponseValidationError is the validation error returned by
// ListGroupsResponse.Validate if the designated constraints aren't met.
type ListGroupsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGroupsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGroupsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGroupsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGroupsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGroupsResponseValidationError) ErrorName() string {
	return "ListGroupsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListGroupsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGroupsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGroupsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGroupsResponseValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on ListUsersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListUsersResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUsersResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListUsersResponseValidationError is the validation error returned by
// ListUsersResponse.Validate if the designated constraints aren't met.
type ListUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersResponseValidationError) ErrorName() string {
	return "ListUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersResponseValidationError{}

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

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRolesResponseValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
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

// Validate checks the field values on ListRoleBindingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRoleBindingsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// ListRoleBindingsRequestValidationError is the validation error returned by
// ListRoleBindingsRequest.Validate if the designated constraints aren't met.
type ListRoleBindingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleBindingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleBindingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleBindingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleBindingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleBindingsRequestValidationError) ErrorName() string {
	return "ListRoleBindingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoleBindingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleBindingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleBindingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleBindingsRequestValidationError{}

// Validate checks the field values on ListRoleBindingsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRoleBindingsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetBindings() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRoleBindingsResponseValidationError{
					field:  fmt.Sprintf("Bindings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRoleBindingsResponseValidationError is the validation error returned by
// ListRoleBindingsResponse.Validate if the designated constraints aren't met.
type ListRoleBindingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleBindingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleBindingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleBindingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleBindingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleBindingsResponseValidationError) ErrorName() string {
	return "ListRoleBindingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoleBindingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleBindingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleBindingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleBindingsResponseValidationError{}

// Validate checks the field values on ResetPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResetPasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// ResetPasswordRequestValidationError is the validation error returned by
// ResetPasswordRequest.Validate if the designated constraints aren't met.
type ResetPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordRequestValidationError) ErrorName() string {
	return "ResetPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordRequestValidationError{}

// Validate checks the field values on ResetPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResetPasswordResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Url

	return nil
}

// ResetPasswordResponseValidationError is the validation error returned by
// ResetPasswordResponse.Validate if the designated constraints aren't met.
type ResetPasswordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordResponseValidationError) ErrorName() string {
	return "ResetPasswordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordResponseValidationError{}

// Validate checks the field values on ChangePasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ChangePasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NewPassword

	return nil
}

// ChangePasswordRequestValidationError is the validation error returned by
// ChangePasswordRequest.Validate if the designated constraints aren't met.
type ChangePasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChangePasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChangePasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChangePasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChangePasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChangePasswordRequestValidationError) ErrorName() string {
	return "ChangePasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChangePasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChangePasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChangePasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChangePasswordRequestValidationError{}
