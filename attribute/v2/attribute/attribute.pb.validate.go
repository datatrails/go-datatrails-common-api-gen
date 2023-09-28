// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: attribute/v2/attribute/attribute.proto

package attribute

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

// Validate checks the field values on DictAttr with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DictAttr) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DictAttr with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DictAttrMultiError, or nil
// if none found.
func (m *DictAttr) ValidateAll() error {
	return m.validate(true)
}

func (m *DictAttr) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetValue()))
		i := 0
		for key := range m.GetValue() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetValue()[key]
			_ = val

			if utf8.RuneCountInString(key) > 1024 {
				err := DictAttrValidationError{
					field:  fmt.Sprintf("Value[%v]", key),
					reason: "value length must be at most 1024 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if strings.Contains(key, ".") {
				err := DictAttrValidationError{
					field:  fmt.Sprintf("Value[%v]", key),
					reason: "value contains substring \".\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if utf8.RuneCountInString(val) > 4096 {
				err := DictAttrValidationError{
					field:  fmt.Sprintf("Value[%v]", key),
					reason: "value length must be at most 4096 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return DictAttrMultiError(errors)
	}

	return nil
}

// DictAttrMultiError is an error wrapping multiple validation errors returned
// by DictAttr.ValidateAll() if the designated constraints aren't met.
type DictAttrMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DictAttrMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DictAttrMultiError) AllErrors() []error { return m }

// DictAttrValidationError is the validation error returned by
// DictAttr.Validate if the designated constraints aren't met.
type DictAttrValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DictAttrValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DictAttrValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DictAttrValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DictAttrValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DictAttrValidationError) ErrorName() string { return "DictAttrValidationError" }

// Error satisfies the builtin error interface
func (e DictAttrValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDictAttr.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DictAttrValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DictAttrValidationError{}

// Validate checks the field values on ListAttr with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListAttr) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAttr with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListAttrMultiError, or nil
// if none found.
func (m *ListAttr) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAttr) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetValue() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAttrValidationError{
						field:  fmt.Sprintf("Value[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAttrValidationError{
						field:  fmt.Sprintf("Value[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAttrValidationError{
					field:  fmt.Sprintf("Value[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAttrMultiError(errors)
	}

	return nil
}

// ListAttrMultiError is an error wrapping multiple validation errors returned
// by ListAttr.ValidateAll() if the designated constraints aren't met.
type ListAttrMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAttrMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAttrMultiError) AllErrors() []error { return m }

// ListAttrValidationError is the validation error returned by
// ListAttr.Validate if the designated constraints aren't met.
type ListAttrValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAttrValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAttrValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAttrValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAttrValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAttrValidationError) ErrorName() string { return "ListAttrValidationError" }

// Error satisfies the builtin error interface
func (e ListAttrValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAttr.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAttrValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAttrValidationError{}

// Validate checks the field values on Attribute with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attribute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttributeMultiError, or nil
// if none found.
func (m *Attribute) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Value.(type) {
	case *Attribute_StrVal:
		if v == nil {
			err := AttributeValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if utf8.RuneCountInString(m.GetStrVal()) > 4096 {
			err := AttributeValidationError{
				field:  "StrVal",
				reason: "value length must be at most 4096 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *Attribute_DictVal:
		if v == nil {
			err := AttributeValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetDictVal()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttributeValidationError{
						field:  "DictVal",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttributeValidationError{
						field:  "DictVal",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDictVal()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttributeValidationError{
					field:  "DictVal",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Attribute_ListVal:
		if v == nil {
			err := AttributeValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetListVal()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttributeValidationError{
						field:  "ListVal",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttributeValidationError{
						field:  "ListVal",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetListVal()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttributeValidationError{
					field:  "ListVal",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return AttributeMultiError(errors)
	}

	return nil
}

// AttributeMultiError is an error wrapping multiple validation errors returned
// by Attribute.ValidateAll() if the designated constraints aren't met.
type AttributeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeMultiError) AllErrors() []error { return m }

// AttributeValidationError is the validation error returned by
// Attribute.Validate if the designated constraints aren't met.
type AttributeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeValidationError) ErrorName() string { return "AttributeValidationError" }

// Error satisfies the builtin error interface
func (e AttributeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValidationError{}