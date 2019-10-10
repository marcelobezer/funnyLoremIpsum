package ipsum

import "errors"

const (
	maxCap = 16777216
)

var (
	// ErrLenghtTooBig is returned whem maximum default size
	// (16777216 bytes) of ipsum is reach.
	ErrLenghtTooBig = errors.New("Required lorem ipsum size is too big")
)
