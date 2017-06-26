package pagerduty

import (
	"errors"
)

var (
	// ErrNoToken is returned by NewClient if a user
	// passed an empty/missing token.
	ErrNoToken = errors.New("an empty token was provided")

	// ErrAuthFailure is returned by NewClient if a user
	// passed an invalid token and failed validation against the PagerDuty API.
	ErrAuthFailure = errors.New("failed to authenticate using the provided token")
)

// ErrorResponse represents an error response from the PagerDuty API.
type ErrorResponse struct {
	ErrorResponse *ErrorResponse `json:"error,omitempty"`
	Code          int            `json:"code,omitempty"`
	Errors        interface{}    `json:"errors,omitempty"`
	Message       string         `json:"message,omitempty"`
	Response      *Response
}
