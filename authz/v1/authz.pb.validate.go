// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authz.proto

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

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Policy) Validate() error {
	if m == nil {
		return nil
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

	// no validation rules for Etag

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

	if utf8.RuneCountInString(m.GetDisplayName()) < 1 {
		return BindingValidationError{
			field:  "DisplayName",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Description

	for idx, item := range m.GetFrom() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return BindingValidationError{
						field:  fmt.Sprintf("From[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	switch m.Rule.(type) {

	case *Binding_AllowHttp:

		{
			tmp := m.GetAllowHttp()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return BindingValidationError{
						field:  "AllowHttp",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
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

// Validate checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Subject) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetJwt()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return SubjectValidationError{
					field:  "Jwt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for key, val := range m.GetRequestHeaders() {
		_ = val

		if utf8.RuneCountInString(key) < 1 {
			return SubjectValidationError{
				field:  fmt.Sprintf("RequestHeaders[%v]", key),
				reason: "value length must be at least 1 runes",
			}
		}

		// no validation rules for RequestHeaders[key]
	}

	switch m.TlsIdentity.(type) {

	case *Subject_ClusterLocalServiceAccount:

		if !_Subject_ClusterLocalServiceAccount_Pattern.MatchString(m.GetClusterLocalServiceAccount()) {
			return SubjectValidationError{
				field:  "ClusterLocalServiceAccount",
				reason: "value does not match regex pattern \"^$|^[^/]+/[^/]+$\"",
			}
		}

	case *Subject_SubjectAltName:
		// no validation rules for SubjectAltName

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

var _Subject_ClusterLocalServiceAccount_Pattern = regexp.MustCompile("^$|^[^/]+/[^/]+$")

// Validate checks the field values on Binding_HTTP_Rules with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Binding_HTTP_Rules) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetTo() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Binding_HTTP_RulesValidationError{
						field:  fmt.Sprintf("To[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// Binding_HTTP_RulesValidationError is the validation error returned by
// Binding_HTTP_Rules.Validate if the designated constraints aren't met.
type Binding_HTTP_RulesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Binding_HTTP_RulesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Binding_HTTP_RulesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Binding_HTTP_RulesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Binding_HTTP_RulesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Binding_HTTP_RulesValidationError) ErrorName() string {
	return "Binding_HTTP_RulesValidationError"
}

// Error satisfies the builtin error interface
func (e Binding_HTTP_RulesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBinding_HTTP_Rules.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Binding_HTTP_RulesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Binding_HTTP_RulesValidationError{}

// Validate checks the field values on Binding_HTTP_Rules_Http with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Binding_HTTP_Rules_Http) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHosts() {
		_, _ = idx, item

		if err := m._validateHostname(item); err != nil {
			return Binding_HTTP_Rules_HttpValidationError{
				field:  fmt.Sprintf("Hosts[%v]", idx),
				reason: "value must be a valid hostname",
				cause:  err,
			}
		}

	}

	for idx, item := range m.GetPaths() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return Binding_HTTP_Rules_HttpValidationError{
				field:  fmt.Sprintf("Paths[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	for idx, item := range m.GetMethods() {
		_, _ = idx, item

		if _, ok := _Binding_HTTP_Rules_Http_Methods_InLookup[item]; !ok {
			return Binding_HTTP_Rules_HttpValidationError{
				field:  fmt.Sprintf("Methods[%v]", idx),
				reason: "value must be in list [GET HEAD POST PUT PATCH DELETE OPTIONS]",
			}
		}

	}

	for key, val := range m.GetRequestHeaders() {
		_ = val

		if utf8.RuneCountInString(key) < 1 {
			return Binding_HTTP_Rules_HttpValidationError{
				field:  fmt.Sprintf("RequestHeaders[%v]", key),
				reason: "value length must be at least 1 runes",
			}
		}

		// no validation rules for RequestHeaders[key]
	}

	for idx, item := range m.GetPorts() {
		_, _ = idx, item

		if val := item; val < 0 || val > 65535 {
			return Binding_HTTP_Rules_HttpValidationError{
				field:  fmt.Sprintf("Ports[%v]", idx),
				reason: "value must be inside range [0, 65535]",
			}
		}

	}

	return nil
}

func (m *Binding_HTTP_Rules_Http) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// Binding_HTTP_Rules_HttpValidationError is the validation error returned by
// Binding_HTTP_Rules_Http.Validate if the designated constraints aren't met.
type Binding_HTTP_Rules_HttpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Binding_HTTP_Rules_HttpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Binding_HTTP_Rules_HttpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Binding_HTTP_Rules_HttpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Binding_HTTP_Rules_HttpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Binding_HTTP_Rules_HttpValidationError) ErrorName() string {
	return "Binding_HTTP_Rules_HttpValidationError"
}

// Error satisfies the builtin error interface
func (e Binding_HTTP_Rules_HttpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBinding_HTTP_Rules_Http.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Binding_HTTP_Rules_HttpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Binding_HTTP_Rules_HttpValidationError{}

var _Binding_HTTP_Rules_Http_Methods_InLookup = map[string]struct{}{
	"GET":     {},
	"HEAD":    {},
	"POST":    {},
	"PUT":     {},
	"PATCH":   {},
	"DELETE":  {},
	"OPTIONS": {},
}

// Validate checks the field values on Subject_JWT with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Subject_JWT) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Principal

	// no validation rules for Audience

	// no validation rules for Presenter

	for key, val := range m.GetClaims() {
		_ = val

		if utf8.RuneCountInString(key) < 1 {
			return Subject_JWTValidationError{
				field:  fmt.Sprintf("Claims[%v]", key),
				reason: "value length must be at least 1 runes",
			}
		}

		// no validation rules for Claims[key]
	}

	if m.GetValidation() == nil {
		return Subject_JWTValidationError{
			field:  "Validation",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetValidation()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return Subject_JWTValidationError{
					field:  "Validation",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// Subject_JWTValidationError is the validation error returned by
// Subject_JWT.Validate if the designated constraints aren't met.
type Subject_JWTValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Subject_JWTValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Subject_JWTValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Subject_JWTValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Subject_JWTValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Subject_JWTValidationError) ErrorName() string { return "Subject_JWTValidationError" }

// Error satisfies the builtin error interface
func (e Subject_JWTValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubject_JWT.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Subject_JWTValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Subject_JWTValidationError{}

// Validate checks the field values on Subject_JWT_Validation with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Subject_JWT_Validation) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetIssuer()) < 1 {
		return Subject_JWT_ValidationValidationError{
			field:  "Issuer",
			reason: "value length must be at least 1 runes",
		}
	}

	for idx, item := range m.GetAudiences() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return Subject_JWT_ValidationValidationError{
				field:  fmt.Sprintf("Audiences[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	switch m.Keys.(type) {

	case *Subject_JWT_Validation_JwksUri:

		if uri, err := url.Parse(m.GetJwksUri()); err != nil {
			return Subject_JWT_ValidationValidationError{
				field:  "JwksUri",
				reason: "value must be a valid URI",
				cause:  err,
			}
		} else if !uri.IsAbs() {
			return Subject_JWT_ValidationValidationError{
				field:  "JwksUri",
				reason: "value must be absolute",
			}
		}

	case *Subject_JWT_Validation_Jwks:

		if utf8.RuneCountInString(m.GetJwks()) < 1 {
			return Subject_JWT_ValidationValidationError{
				field:  "Jwks",
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// Subject_JWT_ValidationValidationError is the validation error returned by
// Subject_JWT_Validation.Validate if the designated constraints aren't met.
type Subject_JWT_ValidationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Subject_JWT_ValidationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Subject_JWT_ValidationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Subject_JWT_ValidationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Subject_JWT_ValidationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Subject_JWT_ValidationValidationError) ErrorName() string {
	return "Subject_JWT_ValidationValidationError"
}

// Error satisfies the builtin error interface
func (e Subject_JWT_ValidationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubject_JWT_Validation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Subject_JWT_ValidationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Subject_JWT_ValidationValidationError{}
