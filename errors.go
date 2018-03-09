package totp

import "errors"

var (
	// ErrUndefinedAlgorithm is returned when an undefined algorithm type is found
	ErrUndefinedAlgorithm = errors.New("Undefined hash algorithm")

	// ErrUndefinedEncoding is returned when an undefined encoder type is encountered
	ErrUndefinedEncoding = errors.New("Undefined encoding")
)