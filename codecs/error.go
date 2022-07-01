package codecs

import (
	"errors"
)

var (
	errInsufficientLengthForAHeader = errors.New("Insufficient length for a header")
	errT140NoExtensionAllowed       = errors.New("No extension is allowed for T140 packet")
)
