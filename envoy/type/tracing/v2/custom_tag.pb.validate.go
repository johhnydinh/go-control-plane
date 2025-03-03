// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/tracing/v2/custom_tag.proto

package tracingv2

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

// Validate checks the field values on CustomTag with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CustomTag) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomTag with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CustomTagMultiError, or nil
// if none found.
func (m *CustomTag) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomTag) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTag()) < 1 {
		err := CustomTagValidationError{
			field:  "Tag",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch m.Type.(type) {

	case *CustomTag_Literal_:

		if all {
			switch v := interface{}(m.GetLiteral()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "Literal",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "Literal",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLiteral()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomTagValidationError{
					field:  "Literal",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomTag_Environment_:

		if all {
			switch v := interface{}(m.GetEnvironment()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "Environment",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "Environment",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEnvironment()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomTagValidationError{
					field:  "Environment",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomTag_RequestHeader:

		if all {
			switch v := interface{}(m.GetRequestHeader()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "RequestHeader",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "RequestHeader",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRequestHeader()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomTagValidationError{
					field:  "RequestHeader",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomTag_Metadata_:

		if all {
			switch v := interface{}(m.GetMetadata()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "Metadata",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomTagValidationError{
						field:  "Metadata",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomTagValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := CustomTagValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return CustomTagMultiError(errors)
	}
	return nil
}

// CustomTagMultiError is an error wrapping multiple validation errors returned
// by CustomTag.ValidateAll() if the designated constraints aren't met.
type CustomTagMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomTagMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomTagMultiError) AllErrors() []error { return m }

// CustomTagValidationError is the validation error returned by
// CustomTag.Validate if the designated constraints aren't met.
type CustomTagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomTagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomTagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomTagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomTagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomTagValidationError) ErrorName() string { return "CustomTagValidationError" }

// Error satisfies the builtin error interface
func (e CustomTagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomTagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomTagValidationError{}

// Validate checks the field values on CustomTag_Literal with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CustomTag_Literal) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomTag_Literal with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomTag_LiteralMultiError, or nil if none found.
func (m *CustomTag_Literal) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomTag_Literal) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetValue()) < 1 {
		err := CustomTag_LiteralValidationError{
			field:  "Value",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CustomTag_LiteralMultiError(errors)
	}
	return nil
}

// CustomTag_LiteralMultiError is an error wrapping multiple validation errors
// returned by CustomTag_Literal.ValidateAll() if the designated constraints
// aren't met.
type CustomTag_LiteralMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomTag_LiteralMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomTag_LiteralMultiError) AllErrors() []error { return m }

// CustomTag_LiteralValidationError is the validation error returned by
// CustomTag_Literal.Validate if the designated constraints aren't met.
type CustomTag_LiteralValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomTag_LiteralValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomTag_LiteralValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomTag_LiteralValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomTag_LiteralValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomTag_LiteralValidationError) ErrorName() string {
	return "CustomTag_LiteralValidationError"
}

// Error satisfies the builtin error interface
func (e CustomTag_LiteralValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomTag_Literal.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomTag_LiteralValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomTag_LiteralValidationError{}

// Validate checks the field values on CustomTag_Environment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CustomTag_Environment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomTag_Environment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomTag_EnvironmentMultiError, or nil if none found.
func (m *CustomTag_Environment) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomTag_Environment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) < 1 {
		err := CustomTag_EnvironmentValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for DefaultValue

	if len(errors) > 0 {
		return CustomTag_EnvironmentMultiError(errors)
	}
	return nil
}

// CustomTag_EnvironmentMultiError is an error wrapping multiple validation
// errors returned by CustomTag_Environment.ValidateAll() if the designated
// constraints aren't met.
type CustomTag_EnvironmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomTag_EnvironmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomTag_EnvironmentMultiError) AllErrors() []error { return m }

// CustomTag_EnvironmentValidationError is the validation error returned by
// CustomTag_Environment.Validate if the designated constraints aren't met.
type CustomTag_EnvironmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomTag_EnvironmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomTag_EnvironmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomTag_EnvironmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomTag_EnvironmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomTag_EnvironmentValidationError) ErrorName() string {
	return "CustomTag_EnvironmentValidationError"
}

// Error satisfies the builtin error interface
func (e CustomTag_EnvironmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomTag_Environment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomTag_EnvironmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomTag_EnvironmentValidationError{}

// Validate checks the field values on CustomTag_Header with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CustomTag_Header) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomTag_Header with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomTag_HeaderMultiError, or nil if none found.
func (m *CustomTag_Header) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomTag_Header) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) < 1 {
		err := CustomTag_HeaderValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CustomTag_Header_Name_Pattern.MatchString(m.GetName()) {
		err := CustomTag_HeaderValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for DefaultValue

	if len(errors) > 0 {
		return CustomTag_HeaderMultiError(errors)
	}
	return nil
}

// CustomTag_HeaderMultiError is an error wrapping multiple validation errors
// returned by CustomTag_Header.ValidateAll() if the designated constraints
// aren't met.
type CustomTag_HeaderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomTag_HeaderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomTag_HeaderMultiError) AllErrors() []error { return m }

// CustomTag_HeaderValidationError is the validation error returned by
// CustomTag_Header.Validate if the designated constraints aren't met.
type CustomTag_HeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomTag_HeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomTag_HeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomTag_HeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomTag_HeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomTag_HeaderValidationError) ErrorName() string { return "CustomTag_HeaderValidationError" }

// Error satisfies the builtin error interface
func (e CustomTag_HeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomTag_Header.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomTag_HeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomTag_HeaderValidationError{}

var _CustomTag_Header_Name_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

// Validate checks the field values on CustomTag_Metadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CustomTag_Metadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomTag_Metadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomTag_MetadataMultiError, or nil if none found.
func (m *CustomTag_Metadata) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomTag_Metadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetKind()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CustomTag_MetadataValidationError{
					field:  "Kind",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CustomTag_MetadataValidationError{
					field:  "Kind",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKind()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomTag_MetadataValidationError{
				field:  "Kind",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadataKey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CustomTag_MetadataValidationError{
					field:  "MetadataKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CustomTag_MetadataValidationError{
					field:  "MetadataKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadataKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomTag_MetadataValidationError{
				field:  "MetadataKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DefaultValue

	if len(errors) > 0 {
		return CustomTag_MetadataMultiError(errors)
	}
	return nil
}

// CustomTag_MetadataMultiError is an error wrapping multiple validation errors
// returned by CustomTag_Metadata.ValidateAll() if the designated constraints
// aren't met.
type CustomTag_MetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomTag_MetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomTag_MetadataMultiError) AllErrors() []error { return m }

// CustomTag_MetadataValidationError is the validation error returned by
// CustomTag_Metadata.Validate if the designated constraints aren't met.
type CustomTag_MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomTag_MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomTag_MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomTag_MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomTag_MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomTag_MetadataValidationError) ErrorName() string {
	return "CustomTag_MetadataValidationError"
}

// Error satisfies the builtin error interface
func (e CustomTag_MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomTag_Metadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomTag_MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomTag_MetadataValidationError{}
