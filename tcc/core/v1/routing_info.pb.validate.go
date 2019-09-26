// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: routing_info.proto

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

// Validate checks the field values on RoutingInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RoutingInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Labels

	for idx, item := range m.GetPorts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RoutingInfoValidationError{
					field:  fmt.Sprintf("Ports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSubsets() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RoutingInfoValidationError{
					field:  fmt.Sprintf("Subsets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetHttpSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoutingInfoValidationError{
				field:  "HttpSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTcpSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoutingInfoValidationError{
				field:  "TcpSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RoutingInfoValidationError is the validation error returned by
// RoutingInfo.Validate if the designated constraints aren't met.
type RoutingInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoutingInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoutingInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoutingInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoutingInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoutingInfoValidationError) ErrorName() string { return "RoutingInfoValidationError" }

// Error satisfies the builtin error interface
func (e RoutingInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoutingInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoutingInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoutingInfoValidationError{}

// Validate checks the field values on Port with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Port) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Number

	// no validation rules for Protocol

	// no validation rules for Name

	// no validation rules for EndpointPort

	return nil
}

// PortValidationError is the validation error returned by Port.Validate if the
// designated constraints aren't met.
type PortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PortValidationError) ErrorName() string { return "PortValidationError" }

// Error satisfies the builtin error interface
func (e PortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PortValidationError{}

// Validate checks the field values on Subset with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Subset) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Labels

	return nil
}

// SubsetValidationError is the validation error returned by Subset.Validate if
// the designated constraints aren't met.
type SubsetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubsetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubsetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubsetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubsetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubsetValidationError) ErrorName() string { return "SubsetValidationError" }

// Error satisfies the builtin error interface
func (e SubsetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubset.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubsetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubsetValidationError{}

// Validate checks the field values on HttpSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpSettings) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetStickySession()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpSettingsValidationError{
				field:  "StickySession",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCorsPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpSettingsValidationError{
				field:  "CorsPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetRouteRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpSettingsValidationError{
					field:  fmt.Sprintf("RouteRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpSettingsValidationError is the validation error returned by
// HttpSettings.Validate if the designated constraints aren't met.
type HttpSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpSettingsValidationError) ErrorName() string { return "HttpSettingsValidationError" }

// Error satisfies the builtin error interface
func (e HttpSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpSettingsValidationError{}

// Validate checks the field values on TcpSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TcpSettings) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRoute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpSettingsValidationError{
				field:  "Route",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TcpSettingsValidationError is the validation error returned by
// TcpSettings.Validate if the designated constraints aren't met.
type TcpSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpSettingsValidationError) ErrorName() string { return "TcpSettingsValidationError" }

// Error satisfies the builtin error interface
func (e TcpSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpSettingsValidationError{}

// Validate checks the field values on HttpRule with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpRule) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetMatch() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpRuleValidationError{
					field:  fmt.Sprintf("Match[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetModify()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpRuleValidationError{
				field:  "Modify",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.RouteOrRedirect.(type) {

	case *HttpRule_Route:

		if v, ok := interface{}(m.GetRoute()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpRuleValidationError{
					field:  "Route",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpRule_Redirect:

		if v, ok := interface{}(m.GetRedirect()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpRuleValidationError{
					field:  "Redirect",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpRuleValidationError is the validation error returned by
// HttpRule.Validate if the designated constraints aren't met.
type HttpRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpRuleValidationError) ErrorName() string { return "HttpRuleValidationError" }

// Error satisfies the builtin error interface
func (e HttpRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpRuleValidationError{}

// Validate checks the field values on HttpMatchCondition with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpMatchCondition) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUri()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpMatchConditionValidationError{
				field:  "Uri",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetScheme()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpMatchConditionValidationError{
				field:  "Scheme",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMethod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpMatchConditionValidationError{
				field:  "Method",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAuthority()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpMatchConditionValidationError{
				field:  "Authority",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Headers

	return nil
}

// HttpMatchConditionValidationError is the validation error returned by
// HttpMatchCondition.Validate if the designated constraints aren't met.
type HttpMatchConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpMatchConditionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpMatchConditionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpMatchConditionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpMatchConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpMatchConditionValidationError) ErrorName() string {
	return "HttpMatchConditionValidationError"
}

// Error satisfies the builtin error interface
func (e HttpMatchConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpMatchCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpMatchConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpMatchConditionValidationError{}

// Validate checks the field values on StringMatch with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StringMatch) Validate() error {
	if m == nil {
		return nil
	}

	switch m.MatchType.(type) {

	case *StringMatch_Exact:
		// no validation rules for Exact

	case *StringMatch_Prefix:
		// no validation rules for Prefix

	case *StringMatch_Regex:
		// no validation rules for Regex

	}

	return nil
}

// StringMatchValidationError is the validation error returned by
// StringMatch.Validate if the designated constraints aren't met.
type StringMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringMatchValidationError) ErrorName() string { return "StringMatchValidationError" }

// Error satisfies the builtin error interface
func (e StringMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringMatchValidationError{}

// Validate checks the field values on HTTPRewrite with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HTTPRewrite) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uri

	// no validation rules for Authority

	return nil
}

// HTTPRewriteValidationError is the validation error returned by
// HTTPRewrite.Validate if the designated constraints aren't met.
type HTTPRewriteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPRewriteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPRewriteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPRewriteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPRewriteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPRewriteValidationError) ErrorName() string { return "HTTPRewriteValidationError" }

// Error satisfies the builtin error interface
func (e HTTPRewriteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPRewrite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPRewriteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPRewriteValidationError{}

// Validate checks the field values on Headers with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Headers) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeadersValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeadersValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HeadersValidationError is the validation error returned by Headers.Validate
// if the designated constraints aren't met.
type HeadersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeadersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeadersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeadersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeadersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeadersValidationError) ErrorName() string { return "HeadersValidationError" }

// Error satisfies the builtin error interface
func (e HeadersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaders.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeadersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeadersValidationError{}

// Validate checks the field values on HttpModifyAction with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpModifyAction) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRewrite()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpModifyActionValidationError{
				field:  "Rewrite",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpModifyActionValidationError{
				field:  "Headers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpModifyActionValidationError is the validation error returned by
// HttpModifyAction.Validate if the designated constraints aren't met.
type HttpModifyActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpModifyActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpModifyActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpModifyActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpModifyActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpModifyActionValidationError) ErrorName() string { return "HttpModifyActionValidationError" }

// Error satisfies the builtin error interface
func (e HttpModifyActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpModifyAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpModifyActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpModifyActionValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetDestinations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteValidationError{
					field:  fmt.Sprintf("Destinations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on Redirect with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Redirect) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uri

	// no validation rules for Authority

	return nil
}

// RedirectValidationError is the validation error returned by
// Redirect.Validate if the designated constraints aren't met.
type RedirectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedirectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedirectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedirectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedirectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedirectValidationError) ErrorName() string { return "RedirectValidationError" }

// Error satisfies the builtin error interface
func (e RedirectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedirect.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedirectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedirectValidationError{}

// Validate checks the field values on CorsPolicy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CorsPolicy) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMaxAge()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CorsPolicyValidationError{
				field:  "MaxAge",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAllowCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CorsPolicyValidationError{
				field:  "AllowCredentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CorsPolicyValidationError is the validation error returned by
// CorsPolicy.Validate if the designated constraints aren't met.
type CorsPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CorsPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CorsPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CorsPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CorsPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CorsPolicyValidationError) ErrorName() string { return "CorsPolicyValidationError" }

// Error satisfies the builtin error interface
func (e CorsPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCorsPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CorsPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CorsPolicyValidationError{}

// Validate checks the field values on HttpSettings_HTTPCookie with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpSettings_HTTPCookie) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Path

	if v, ok := interface{}(m.GetTtl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpSettings_HTTPCookieValidationError{
				field:  "Ttl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpSettings_HTTPCookieValidationError is the validation error returned by
// HttpSettings_HTTPCookie.Validate if the designated constraints aren't met.
type HttpSettings_HTTPCookieValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpSettings_HTTPCookieValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpSettings_HTTPCookieValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpSettings_HTTPCookieValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpSettings_HTTPCookieValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpSettings_HTTPCookieValidationError) ErrorName() string {
	return "HttpSettings_HTTPCookieValidationError"
}

// Error satisfies the builtin error interface
func (e HttpSettings_HTTPCookieValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpSettings_HTTPCookie.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpSettings_HTTPCookieValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpSettings_HTTPCookieValidationError{}

// Validate checks the field values on HttpSettings_StickySession with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpSettings_StickySession) Validate() error {
	if m == nil {
		return nil
	}

	switch m.HashKey.(type) {

	case *HttpSettings_StickySession_Header:
		// no validation rules for Header

	case *HttpSettings_StickySession_Cookie:

		if v, ok := interface{}(m.GetCookie()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpSettings_StickySessionValidationError{
					field:  "Cookie",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpSettings_StickySession_UseSourceIp:
		// no validation rules for UseSourceIp

	}

	return nil
}

// HttpSettings_StickySessionValidationError is the validation error returned
// by HttpSettings_StickySession.Validate if the designated constraints aren't met.
type HttpSettings_StickySessionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpSettings_StickySessionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpSettings_StickySessionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpSettings_StickySessionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpSettings_StickySessionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpSettings_StickySessionValidationError) ErrorName() string {
	return "HttpSettings_StickySessionValidationError"
}

// Error satisfies the builtin error interface
func (e HttpSettings_StickySessionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpSettings_StickySession.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpSettings_StickySessionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpSettings_StickySessionValidationError{}

// Validate checks the field values on Headers_HeaderOperations with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Headers_HeaderOperations) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Set

	// no validation rules for Add

	return nil
}

// Headers_HeaderOperationsValidationError is the validation error returned by
// Headers_HeaderOperations.Validate if the designated constraints aren't met.
type Headers_HeaderOperationsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Headers_HeaderOperationsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Headers_HeaderOperationsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Headers_HeaderOperationsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Headers_HeaderOperationsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Headers_HeaderOperationsValidationError) ErrorName() string {
	return "Headers_HeaderOperationsValidationError"
}

// Error satisfies the builtin error interface
func (e Headers_HeaderOperationsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaders_HeaderOperations.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Headers_HeaderOperationsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Headers_HeaderOperationsValidationError{}

// Validate checks the field values on Route_Destination with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Route_Destination) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Host

	// no validation rules for Subset

	// no validation rules for Weight

	// no validation rules for Port

	return nil
}

// Route_DestinationValidationError is the validation error returned by
// Route_Destination.Validate if the designated constraints aren't met.
type Route_DestinationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Route_DestinationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Route_DestinationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Route_DestinationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Route_DestinationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Route_DestinationValidationError) ErrorName() string {
	return "Route_DestinationValidationError"
}

// Error satisfies the builtin error interface
func (e Route_DestinationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute_Destination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Route_DestinationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Route_DestinationValidationError{}
