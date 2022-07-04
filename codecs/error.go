package codecs

import (
	"errors"
)

var (
	errInsufficientLengthForAHeader = errors.New("Insufficient length for a header")
	errT140NoExtensionAllowed       = errors.New("No extension is allowed for T140 packet")
	errT140CCNotZero                = errors.New("CC must be zero for T140 packet")
	errTooSmall                     = errors.New("Packet is too small")
)
