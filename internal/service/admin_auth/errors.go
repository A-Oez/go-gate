package adminauth

import "errors"

var (
	ErrInvalidUser				= errors.New("authentication failed: user not found")
	ErrPasswordDoesNotMatch		= errors.New("authentication failed: incorrect password")
)