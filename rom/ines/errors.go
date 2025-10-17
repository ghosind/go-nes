package ines

import "errors"

var (
	ErrInvalidROM        = errors.New("invalid ROM")
	ErrUnsupportedMapper = errors.New("unsupported mapper")
)
