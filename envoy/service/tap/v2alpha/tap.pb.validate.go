// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/tap.proto

package v2alpha

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

// Validate checks the field values on StreamTapsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StreamTapsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamTapsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamTapsRequestMultiError, or nil if none found.
func (m *StreamTapsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamTapsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIdentifier()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamTapsRequestValidationError{
					field:  "Identifier",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamTapsRequestValidationError{
					field:  "Identifier",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIdentifier()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamTapsRequestValidationError{
				field:  "Identifier",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TraceId

	if all {
		switch v := interface{}(m.GetTrace()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamTapsRequestValidationError{
					field:  "Trace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamTapsRequestValidationError{
					field:  "Trace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTrace()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamTapsRequestValidationError{
				field:  "Trace",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StreamTapsRequestMultiError(errors)
	}
	return nil
}

// StreamTapsRequestMultiError is an error wrapping multiple validation errors
// returned by StreamTapsRequest.ValidateAll() if the designated constraints
// aren't met.
type StreamTapsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamTapsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamTapsRequestMultiError) AllErrors() []error { return m }

// StreamTapsRequestValidationError is the validation error returned by
// StreamTapsRequest.Validate if the designated constraints aren't met.
type StreamTapsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamTapsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamTapsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamTapsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamTapsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamTapsRequestValidationError) ErrorName() string {
	return "StreamTapsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StreamTapsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTapsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamTapsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamTapsRequestValidationError{}

// Validate checks the field values on StreamTapsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamTapsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamTapsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamTapsResponseMultiError, or nil if none found.
func (m *StreamTapsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamTapsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StreamTapsResponseMultiError(errors)
	}
	return nil
}

// StreamTapsResponseMultiError is an error wrapping multiple validation errors
// returned by StreamTapsResponse.ValidateAll() if the designated constraints
// aren't met.
type StreamTapsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamTapsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamTapsResponseMultiError) AllErrors() []error { return m }

// StreamTapsResponseValidationError is the validation error returned by
// StreamTapsResponse.Validate if the designated constraints aren't met.
type StreamTapsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamTapsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamTapsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamTapsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamTapsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamTapsResponseValidationError) ErrorName() string {
	return "StreamTapsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StreamTapsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTapsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamTapsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamTapsResponseValidationError{}

// Validate checks the field values on StreamTapsRequest_Identifier with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamTapsRequest_Identifier) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamTapsRequest_Identifier with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamTapsRequest_IdentifierMultiError, or nil if none found.
func (m *StreamTapsRequest_Identifier) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamTapsRequest_Identifier) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNode() == nil {
		err := StreamTapsRequest_IdentifierValidationError{
			field:  "Node",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamTapsRequest_IdentifierValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamTapsRequest_IdentifierValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamTapsRequest_IdentifierValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TapId

	if len(errors) > 0 {
		return StreamTapsRequest_IdentifierMultiError(errors)
	}
	return nil
}

// StreamTapsRequest_IdentifierMultiError is an error wrapping multiple
// validation errors returned by StreamTapsRequest_Identifier.ValidateAll() if
// the designated constraints aren't met.
type StreamTapsRequest_IdentifierMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamTapsRequest_IdentifierMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamTapsRequest_IdentifierMultiError) AllErrors() []error { return m }

// StreamTapsRequest_IdentifierValidationError is the validation error returned
// by StreamTapsRequest_Identifier.Validate if the designated constraints
// aren't met.
type StreamTapsRequest_IdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamTapsRequest_IdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamTapsRequest_IdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamTapsRequest_IdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamTapsRequest_IdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamTapsRequest_IdentifierValidationError) ErrorName() string {
	return "StreamTapsRequest_IdentifierValidationError"
}

// Error satisfies the builtin error interface
func (e StreamTapsRequest_IdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTapsRequest_Identifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamTapsRequest_IdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamTapsRequest_IdentifierValidationError{}
