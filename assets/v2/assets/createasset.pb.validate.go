// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datatrails-common-api/assets/v2/assets/createasset.proto

package assets

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

// Validate checks the field values on CreateAssetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAssetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAssetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAssetRequestMultiError, or nil if none found.
func (m *CreateAssetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAssetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBehaviours() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) > 1024 {
			err := CreateAssetRequestValidationError{
				field:  fmt.Sprintf("Behaviours[%v]", idx),
				reason: "value length must be at most 1024 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	{
		sorted_keys := make([]string, len(m.GetAttributes()))
		i := 0
		for key := range m.GetAttributes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetAttributes()[key]
			_ = val

			if l := utf8.RuneCountInString(key); l < 1 || l > 1024 {
				err := CreateAssetRequestValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "value length must be between 1 and 1024 runes, inclusive",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if strings.Contains(key, ".") {
				err := CreateAssetRequestValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "value contains substring \".\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_CreateAssetRequest_Attributes_Pattern.MatchString(key) {
				err := CreateAssetRequestValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "value does not match regex pattern \"^[^[:cntrl:]]*?[^[[:cntrl:]]+?[^[:cntrl:]]$|^[^[:cntrl:]]$|^[^[:cntrl:]]*?[^][:cntrl:]]$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, CreateAssetRequestValidationError{
							field:  fmt.Sprintf("Attributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, CreateAssetRequestValidationError{
							field:  fmt.Sprintf("Attributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return CreateAssetRequestValidationError{
						field:  fmt.Sprintf("Attributes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if _, ok := ProofMechanism_name[int32(m.GetProofMechanism())]; !ok {
		err := CreateAssetRequestValidationError{
			field:  "ProofMechanism",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetChainId()) > 1024 {
		err := CreateAssetRequestValidationError{
			field:  "ChainId",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Public

	if len(errors) > 0 {
		return CreateAssetRequestMultiError(errors)
	}

	return nil
}

// CreateAssetRequestMultiError is an error wrapping multiple validation errors
// returned by CreateAssetRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateAssetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAssetRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAssetRequestMultiError) AllErrors() []error { return m }

// CreateAssetRequestValidationError is the validation error returned by
// CreateAssetRequest.Validate if the designated constraints aren't met.
type CreateAssetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAssetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAssetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAssetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAssetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAssetRequestValidationError) ErrorName() string {
	return "CreateAssetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAssetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAssetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAssetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAssetRequestValidationError{}

var _CreateAssetRequest_Attributes_Pattern = regexp.MustCompile("^[^[:cntrl:]]*?[^[[:cntrl:]]+?[^[:cntrl:]]$|^[^[:cntrl:]]$|^[^[:cntrl:]]*?[^][:cntrl:]]$")
