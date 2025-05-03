package adminauth

import "errors"

var (
	ErrInvalidUser				= errors.New("authentication failed: user not found")
	ErrPasswordDoesNotMatch		= errors.New("authentication failed: incorrect password")
	ErrReadBody					= errors.New("could not read request body")
	ErrUnmarshalJSON			= errors.New("could not parse JSON from body")
)