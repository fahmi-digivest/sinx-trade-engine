package message

import (
	"fmt"
)

// Message is the common interface all SoupBinTCP message types implement.
type Message interface {
	Type() PacketType
	Encode() ([]byte, error)
}

// writeAlpha writes s right-padded with spaces into buf.
// Returns an error if s is longer than len(buf).
func writeAlpha(buf []byte, s string) error {
	if len(s) > len(buf) {
		return fmt.Errorf("value %q too long for field of length %d", s, len(buf))
	}
	for i := range buf {
		buf[i] = ' '
	}
	copy(buf, s)
	return nil
}

// writeAlphaLeft writes s left-padded with spaces into buf.
func writeAlphaLeft(buf []byte, s string) error {
	if len(s) > len(buf) {
		return fmt.Errorf("value %q too long for field of length %d", s, len(buf))
	}
	for i := range buf {
		buf[i] = ' '
	}
	copy(buf[len(buf)-len(s):], s)
	return nil
}

// writeNumericASCIILeft writes a numeric ASCII string left-padded with spaces into buf.
func writeNumericASCIILeft(buf []byte, s string) error {
	return writeAlphaLeft(buf, s)
}
