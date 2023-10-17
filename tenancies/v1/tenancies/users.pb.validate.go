// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tenancies/v1/tenancies/users.proto

package tenancies

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListUserTenantsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUserTenantsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserTenantsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserTenantsRequestMultiError, or nil if none found.
func (m *ListUserTenantsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserTenantsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUserTenantsRequestMultiError(errors)
	}

	return nil
}

// ListUserTenantsRequestMultiError is an error wrapping multiple validation
// errors returned by ListUserTenantsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListUserTenantsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserTenantsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserTenantsRequestMultiError) AllErrors() []error { return m }

// ListUserTenantsRequestValidationError is the validation error returned by
// ListUserTenantsRequest.Validate if the designated constraints aren't met.
type ListUserTenantsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserTenantsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserTenantsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserTenantsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserTenantsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserTenantsRequestValidationError) ErrorName() string {
	return "ListUserTenantsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListUserTenantsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserTenantsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserTenantsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserTenantsRequestValidationError{}

// Validate checks the field values on UserTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserTenantResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserTenantResponseMultiError, or nil if none found.
func (m *UserTenantResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserTenantResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identity

	// no validation rules for DisplayName

	if len(errors) > 0 {
		return UserTenantResponseMultiError(errors)
	}

	return nil
}

// UserTenantResponseMultiError is an error wrapping multiple validation errors
// returned by UserTenantResponse.ValidateAll() if the designated constraints
// aren't met.
type UserTenantResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserTenantResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserTenantResponseMultiError) AllErrors() []error { return m }

// UserTenantResponseValidationError is the validation error returned by
// UserTenantResponse.Validate if the designated constraints aren't met.
type UserTenantResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserTenantResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserTenantResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserTenantResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserTenantResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserTenantResponseValidationError) ErrorName() string {
	return "UserTenantResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserTenantResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserTenantResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserTenantResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserTenantResponseValidationError{}

// Validate checks the field values on ListUserTenantsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUserTenantsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserTenantsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserTenantsResponseMultiError, or nil if none found.
func (m *ListUserTenantsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserTenantsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTenants() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserTenantsResponseValidationError{
						field:  fmt.Sprintf("Tenants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserTenantsResponseValidationError{
						field:  fmt.Sprintf("Tenants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserTenantsResponseValidationError{
					field:  fmt.Sprintf("Tenants[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListUserTenantsResponseMultiError(errors)
	}

	return nil
}

// ListUserTenantsResponseMultiError is an error wrapping multiple validation
// errors returned by ListUserTenantsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListUserTenantsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserTenantsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserTenantsResponseMultiError) AllErrors() []error { return m }

// ListUserTenantsResponseValidationError is the validation error returned by
// ListUserTenantsResponse.Validate if the designated constraints aren't met.
type ListUserTenantsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserTenantsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserTenantsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserTenantsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserTenantsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserTenantsResponseValidationError) ErrorName() string {
	return "ListUserTenantsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListUserTenantsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserTenantsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserTenantsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserTenantsResponseValidationError{}
