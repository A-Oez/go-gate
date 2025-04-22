package routes

import "errors"

var (
	ErrInvalidID           = errors.New("invalid ID parameter")
	ErrReadBody		       = errors.New("could not read request body")
	ErrUnmarshalJSON	   = errors.New("could not parse JSON from body")
	ErrDBQueryFailed       = errors.New("database query failed")
)