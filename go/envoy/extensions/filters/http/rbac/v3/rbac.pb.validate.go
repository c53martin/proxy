// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/rbac/v3/rbac.proto

package rbacv3

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

// Validate checks the field values on RBAC with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *RBAC) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RBAC with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RBACMultiError, or nil if none found.
func (m *RBAC) ValidateAll() error {
	return m.validate(true)
}

func (m *RBAC) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRules()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "Rules",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "Rules",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRules()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RBACValidationError{
				field:  "Rules",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RulesStatPrefix

	if all {
		switch v := interface{}(m.GetMatcher()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "Matcher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "Matcher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RBACValidationError{
				field:  "Matcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetShadowRules()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "ShadowRules",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "ShadowRules",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShadowRules()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RBACValidationError{
				field:  "ShadowRules",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetShadowMatcher()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "ShadowMatcher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RBACValidationError{
					field:  "ShadowMatcher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShadowMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RBACValidationError{
				field:  "ShadowMatcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ShadowRulesStatPrefix

	// no validation rules for TrackPerRuleStats

	if len(errors) > 0 {
		return RBACMultiError(errors)
	}

	return nil
}

// RBACMultiError is an error wrapping multiple validation errors returned by
// RBAC.ValidateAll() if the designated constraints aren't met.
type RBACMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RBACMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RBACMultiError) AllErrors() []error { return m }

// RBACValidationError is the validation error returned by RBAC.Validate if the
// designated constraints aren't met.
type RBACValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RBACValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RBACValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RBACValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RBACValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RBACValidationError) ErrorName() string { return "RBACValidationError" }

// Error satisfies the builtin error interface
func (e RBACValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRBAC.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RBACValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RBACValidationError{}

// Validate checks the field values on RBACPerRoute with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RBACPerRoute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RBACPerRoute with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RBACPerRouteMultiError, or
// nil if none found.
func (m *RBACPerRoute) ValidateAll() error {
	return m.validate(true)
}

func (m *RBACPerRoute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRbac()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RBACPerRouteValidationError{
					field:  "Rbac",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RBACPerRouteValidationError{
					field:  "Rbac",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRbac()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RBACPerRouteValidationError{
				field:  "Rbac",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RBACPerRouteMultiError(errors)
	}

	return nil
}

// RBACPerRouteMultiError is an error wrapping multiple validation errors
// returned by RBACPerRoute.ValidateAll() if the designated constraints aren't met.
type RBACPerRouteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RBACPerRouteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RBACPerRouteMultiError) AllErrors() []error { return m }

// RBACPerRouteValidationError is the validation error returned by
// RBACPerRoute.Validate if the designated constraints aren't met.
type RBACPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RBACPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RBACPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RBACPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RBACPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RBACPerRouteValidationError) ErrorName() string { return "RBACPerRouteValidationError" }

// Error satisfies the builtin error interface
func (e RBACPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRBACPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RBACPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RBACPerRouteValidationError{}
