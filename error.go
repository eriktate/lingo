package lingo

import (
	"fmt"
	"strings"
)

// An Error is the structured error type that Linode returns on 4xx and 5xx status codes.
type Error struct {
	Field  string `json:"field,omitempty"`
	Reason string `json:"reason"`
}

// Error implements the go error interface for Errors.
func (e Error) Error() string {
	var errorText string
	if e.Field != "" {
		errorText = fmt.Sprintf("With field '%s', ", e.Field)
	}

	return errorText + e.Reason
}

// Errors is the full error response that Linode returns on 4xx and 5xx status codes. It
// aliases a slice of Error structs so that the go error interface can be fulfilled.
type Errors struct {
	Errors []Error `json:"errors"`
}

// Error implements the go error interface for Errors.
func (e Errors) Error() string {
	errorTexts := make([]string, len(e.Errors))

	for i, err := range e.Errors {
		errorTexts[i] = err.Error()
	}

	return strings.Join(errorTexts, "\n")
}
