// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tsb/traffic/v1/service_route.proto

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
var _service_route_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ServiceRoute with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ServiceRoute) Validate() error {
	if m == nil {
		return nil
	}

	if !_ServiceRoute_Service_Pattern.MatchString(m.GetService()) {
		return ServiceRouteValidationError{
			field:  "Service",
			reason: "value does not match regex pattern \"^[^/]+/[^/]+$\"",
		}
	}

	for idx, item := range m.GetSubsets() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ServiceRouteValidationError{
						field:  fmt.Sprintf("Subsets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ServiceRouteValidationError is the validation error returned by
// ServiceRoute.Validate if the designated constraints aren't met.
type ServiceRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceRouteValidationError) ErrorName() string { return "ServiceRouteValidationError" }

// Error satisfies the builtin error interface
func (e ServiceRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceRouteValidationError{}

var _ServiceRoute_Service_Pattern = regexp.MustCompile("^[^/]+/[^/]+$")

// Validate checks the field values on ServiceRoute_Subset with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServiceRoute_Subset) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Labels

	// no validation rules for Weight

	return nil
}

// ServiceRoute_SubsetValidationError is the validation error returned by
// ServiceRoute_Subset.Validate if the designated constraints aren't met.
type ServiceRoute_SubsetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceRoute_SubsetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceRoute_SubsetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceRoute_SubsetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceRoute_SubsetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceRoute_SubsetValidationError) ErrorName() string {
	return "ServiceRoute_SubsetValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceRoute_SubsetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceRoute_Subset.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceRoute_SubsetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceRoute_SubsetValidationError{}
