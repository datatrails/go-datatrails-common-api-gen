// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tenancies/v1/tenancies/tenantusers_messages.proto

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

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetIdentity()) > 1024 {
		err := UserValidationError{
			field:  "Identity",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Email

	if utf8.RuneCountInString(m.GetIssuer()) > 1024 {
		err := UserValidationError{
			field:  "Issuer",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetSubject()) > 1024 {
		err := UserValidationError{
			field:  "Subject",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDisplayName()) > 1024 {
		err := UserValidationError{
			field:  "DisplayName",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

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

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersRequestMultiError, or nil if none found.
func (m *ListUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserUuid()) > 1024 {
		err := ListUsersRequestValidationError{
			field:  "UserUuid",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Email

	if utf8.RuneCountInString(m.GetDisplayName()) > 1024 {
		err := ListUsersRequestValidationError{
			field:  "DisplayName",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() < 0 {
		err := ListUsersRequestValidationError{
			field:  "PageSize",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListUsersRequestMultiError(errors)
	}

	return nil
}

// ListUsersRequestMultiError is an error wrapping multiple validation errors
// returned by ListUsersRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersRequestMultiError) AllErrors() []error { return m }

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
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersResponseMultiError, or nil if none found.
func (m *ListUsersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUsersResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListUsersResponseMultiError(errors)
	}

	return nil
}

// ListUsersResponseMultiError is an error wrapping multiple validation errors
// returned by ListUsersResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUsersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersResponseMultiError) AllErrors() []error { return m }

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

// Validate checks the field values on DeleteUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserRequestMultiError, or nil if none found.
func (m *DeleteUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserUuid()) > 1024 {
		err := DeleteUserRequestValidationError{
			field:  "UserUuid",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteUserRequestMultiError(errors)
	}

	return nil
}

// DeleteUserRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteUserRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserRequestMultiError) AllErrors() []error { return m }

// DeleteUserRequestValidationError is the validation error returned by
// DeleteUserRequest.Validate if the designated constraints aren't met.
type DeleteUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserRequestValidationError) ErrorName() string {
	return "DeleteUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserRequestValidationError{}

// Validate checks the field values on CountUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CountUsersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CountUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CountUsersResponseMultiError, or nil if none found.
func (m *CountUsersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CountUsersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return CountUsersResponseMultiError(errors)
	}

	return nil
}

// CountUsersResponseMultiError is an error wrapping multiple validation errors
// returned by CountUsersResponse.ValidateAll() if the designated constraints
// aren't met.
type CountUsersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountUsersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountUsersResponseMultiError) AllErrors() []error { return m }

// CountUsersResponseValidationError is the validation error returned by
// CountUsersResponse.Validate if the designated constraints aren't met.
type CountUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountUsersResponseValidationError) ErrorName() string {
	return "CountUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CountUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCountUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountUsersResponseValidationError{}

// Validate checks the field values on GetCapsTenantUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCapsTenantUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCapsTenantUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCapsTenantUserRequestMultiError, or nil if none found.
func (m *GetCapsTenantUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCapsTenantUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetCapsTenantUserRequestMultiError(errors)
	}

	return nil
}

// GetCapsTenantUserRequestMultiError is an error wrapping multiple validation
// errors returned by GetCapsTenantUserRequest.ValidateAll() if the designated
// constraints aren't met.
type GetCapsTenantUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCapsTenantUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCapsTenantUserRequestMultiError) AllErrors() []error { return m }

// GetCapsTenantUserRequestValidationError is the validation error returned by
// GetCapsTenantUserRequest.Validate if the designated constraints aren't met.
type GetCapsTenantUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCapsTenantUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCapsTenantUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCapsTenantUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCapsTenantUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCapsTenantUserRequestValidationError) ErrorName() string {
	return "GetCapsTenantUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCapsTenantUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCapsTenantUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCapsTenantUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCapsTenantUserRequestValidationError{}
