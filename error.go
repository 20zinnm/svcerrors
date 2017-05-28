package svcerrors

import "errors"

// Error represents a service error.
// It provides methods for an error status code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// ServiceError represents an error with an associated status code (for documentation lookups).
type ServiceError struct {
	Code int
	Err  error
}

// New returns a new Error with the given status code and text.
func New(code int, message string) Error {
	return ServiceError{Code: code, Err: errors.New(message)}
}

// Wrap returns a new Error with the given status code and error.
func Wrap(code int, err error) Error {
	return ServiceError{Code: code, Err: err}
}

// Allows StatusError to satisfy the error interface.
func (se ServiceError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se ServiceError) Status() int {
	return se.Code
}
