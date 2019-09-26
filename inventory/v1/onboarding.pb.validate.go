// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: onboarding.proto

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

// Validate checks the field values on CloudRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CloudRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Account

	return nil
}

// CloudRequestValidationError is the validation error returned by
// CloudRequest.Validate if the designated constraints aren't met.
type CloudRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloudRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloudRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloudRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloudRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloudRequestValidationError) ErrorName() string { return "CloudRequestValidationError" }

// Error satisfies the builtin error interface
func (e CloudRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloudRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloudRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloudRequestValidationError{}

// Validate checks the field values on ListCloudUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCloudUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCloud()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCloudUsersRequestValidationError{
				field:  "Cloud",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListCloudUsersRequestValidationError is the validation error returned by
// ListCloudUsersRequest.Validate if the designated constraints aren't met.
type ListCloudUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCloudUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCloudUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCloudUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCloudUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCloudUsersRequestValidationError) ErrorName() string {
	return "ListCloudUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCloudUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCloudUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCloudUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCloudUsersRequestValidationError{}

// Validate checks the field values on ListCloudGroupsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCloudGroupsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCloud()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCloudGroupsRequestValidationError{
				field:  "Cloud",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListCloudGroupsRequestValidationError is the validation error returned by
// ListCloudGroupsRequest.Validate if the designated constraints aren't met.
type ListCloudGroupsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCloudGroupsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCloudGroupsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCloudGroupsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCloudGroupsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCloudGroupsRequestValidationError) ErrorName() string {
	return "ListCloudGroupsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCloudGroupsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCloudGroupsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCloudGroupsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCloudGroupsRequestValidationError{}

// Validate checks the field values on ListCloudRolesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCloudRolesRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCloud()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCloudRolesRequestValidationError{
				field:  "Cloud",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListCloudRolesRequestValidationError is the validation error returned by
// ListCloudRolesRequest.Validate if the designated constraints aren't met.
type ListCloudRolesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCloudRolesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCloudRolesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCloudRolesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCloudRolesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCloudRolesRequestValidationError) ErrorName() string {
	return "ListCloudRolesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCloudRolesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCloudRolesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCloudRolesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCloudRolesRequestValidationError{}

// Validate checks the field values on ListCloudRoleBindingsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCloudRoleBindingsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCloud()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCloudRoleBindingsRequestValidationError{
				field:  "Cloud",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	return nil
}

// ListCloudRoleBindingsRequestValidationError is the validation error returned
// by ListCloudRoleBindingsRequest.Validate if the designated constraints
// aren't met.
type ListCloudRoleBindingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCloudRoleBindingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCloudRoleBindingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCloudRoleBindingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCloudRoleBindingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCloudRoleBindingsRequestValidationError) ErrorName() string {
	return "ListCloudRoleBindingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCloudRoleBindingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCloudRoleBindingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCloudRoleBindingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCloudRoleBindingsRequestValidationError{}

// Validate checks the field values on ImportGroupsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ImportGroupsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCloud()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImportGroupsRequestValidationError{
				field:  "Cloud",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ImportGroupsRequestValidationError is the validation error returned by
// ImportGroupsRequest.Validate if the designated constraints aren't met.
type ImportGroupsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportGroupsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportGroupsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportGroupsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportGroupsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportGroupsRequestValidationError) ErrorName() string {
	return "ImportGroupsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ImportGroupsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportGroupsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportGroupsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportGroupsRequestValidationError{}

// Validate checks the field values on ImportGroupsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ImportGroupsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetGroups() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImportGroupsResponseValidationError{
					field:  fmt.Sprintf("Groups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ImportGroupsResponseValidationError is the validation error returned by
// ImportGroupsResponse.Validate if the designated constraints aren't met.
type ImportGroupsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportGroupsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportGroupsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportGroupsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportGroupsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportGroupsResponseValidationError) ErrorName() string {
	return "ImportGroupsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ImportGroupsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportGroupsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportGroupsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportGroupsResponseValidationError{}

// Validate checks the field values on ImportUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ImportUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCloud()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImportUsersRequestValidationError{
				field:  "Cloud",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ImportUsersRequestValidationError is the validation error returned by
// ImportUsersRequest.Validate if the designated constraints aren't met.
type ImportUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportUsersRequestValidationError) ErrorName() string {
	return "ImportUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ImportUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportUsersRequestValidationError{}

// Validate checks the field values on ImportUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ImportUsersResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImportUsersResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ImportUsersResponseValidationError is the validation error returned by
// ImportUsersResponse.Validate if the designated constraints aren't met.
type ImportUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportUsersResponseValidationError) ErrorName() string {
	return "ImportUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ImportUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportUsersResponseValidationError{}