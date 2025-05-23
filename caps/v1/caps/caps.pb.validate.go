// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datatrails-common-api/caps/v1/caps/caps.proto

package caps

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

// Validate checks the field values on GetCapsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCapsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCapsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCapsRequestMultiError,
// or nil if none found.
func (m *GetCapsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCapsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetCapsRequestMultiError(errors)
	}

	return nil
}

// GetCapsRequestMultiError is an error wrapping multiple validation errors
// returned by GetCapsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCapsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCapsRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCapsRequestMultiError) AllErrors() []error { return m }

// GetCapsRequestValidationError is the validation error returned by
// GetCapsRequest.Validate if the designated constraints aren't met.
type GetCapsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCapsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCapsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCapsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCapsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCapsRequestValidationError) ErrorName() string { return "GetCapsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCapsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCapsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCapsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCapsRequestValidationError{}

// Validate checks the field values on Cap with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Cap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cap with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CapMultiError, or nil if none found.
func (m *Cap) ValidateAll() error {
	return m.validate(true)
}

func (m *Cap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetResourceType()); l < 1 || l > 1024 {
		err := CapValidationError{
			field:  "ResourceType",
			reason: "value length must be between 1 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetResourceRemaining() < -1 {
		err := CapValidationError{
			field:  "ResourceRemaining",
			reason: "value must be greater than or equal to -1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetResourceCap() < -1 {
		err := CapValidationError{
			field:  "ResourceCap",
			reason: "value must be greater than or equal to -1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CapMultiError(errors)
	}

	return nil
}

// CapMultiError is an error wrapping multiple validation errors returned by
// Cap.ValidateAll() if the designated constraints aren't met.
type CapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CapMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CapMultiError) AllErrors() []error { return m }

// CapValidationError is the validation error returned by Cap.Validate if the
// designated constraints aren't met.
type CapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CapValidationError) ErrorName() string { return "CapValidationError" }

// Error satisfies the builtin error interface
func (e CapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CapValidationError{}

// Validate checks the field values on Caps with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Caps) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Caps with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CapsMultiError, or nil if none found.
func (m *Caps) ValidateAll() error {
	return m.validate(true)
}

func (m *Caps) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCaps() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CapsValidationError{
						field:  fmt.Sprintf("Caps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CapsValidationError{
						field:  fmt.Sprintf("Caps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CapsValidationError{
					field:  fmt.Sprintf("Caps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CapsMultiError(errors)
	}

	return nil
}

// CapsMultiError is an error wrapping multiple validation errors returned by
// Caps.ValidateAll() if the designated constraints aren't met.
type CapsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CapsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CapsMultiError) AllErrors() []error { return m }

// CapsValidationError is the validation error returned by Caps.Validate if the
// designated constraints aren't met.
type CapsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CapsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CapsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CapsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CapsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CapsValidationError) ErrorName() string { return "CapsValidationError" }

// Error satisfies the builtin error interface
func (e CapsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCaps.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CapsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CapsValidationError{}
