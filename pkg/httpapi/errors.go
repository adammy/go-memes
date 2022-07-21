package httpapi

import (
	"errors"
)

var (
	// ErrInvalidServerType denotes an invalid ServerType was inputted.
	ErrInvalidServerType = errors.New("invalid server type defines")

	// ErrServerTypeNotImplemented denotes the ServerType is not implemented.
	ErrServerTypeNotImplemented = errors.New("server type not implemented")
)
