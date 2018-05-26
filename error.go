package lingo

import (
	"fmt"
	"strings"
)

const busyText = "Linode busy."

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
	errorTexts := make([]string, len(e.Errors)+1)
	errorTexts[0] = "Linode API Error: "

	for i, err := range e.Errors {
		errorTexts[i+1] = err.Error()
	}

	return strings.Join(errorTexts, "\n\t")
}

func (e Error) IsBusy() bool {
	if e.Reason == busyText {
		return true
	}

	return false
}

func (e Errors) IsBusy() bool {
	for _, err := range e.Errors {
		if err.IsBusy() {
			return true
		}
	}

	return false
}
