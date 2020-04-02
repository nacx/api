// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tcc/core/v1/tlssettings.proto

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
var _tlssettings_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TLSSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TLSSettings) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TlsMode

	// no validation rules for TlsEnabled

	// no validation rules for RedirectToHttps

	// no validation rules for ServerCertificate

	// no validation rules for PrivateKey

	// no validation rules for CaCertificates

	// no validation rules for SecretName

	return nil
}

// TLSSettingsValidationError is the validation error returned by
// TLSSettings.Validate if the designated constraints aren't met.
type TLSSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TLSSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TLSSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TLSSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TLSSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TLSSettingsValidationError) ErrorName() string { return "TLSSettingsValidationError" }

// Error satisfies the builtin error interface
func (e TLSSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTLSSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TLSSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TLSSettingsValidationError{}
