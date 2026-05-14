package message

import (
	"fmt"
	"strings"
)

// LoginAccepted is sent by the server after a successful login.
type LoginAccepted struct {
	Session        string // 10 bytes, left-padded with spaces
	SequenceNumber string // 20 bytes ASCII numeric, left-padded with spaces
}

// Encode serialises LoginAccepted into its 30-byte wire payload.
func (m *LoginAccepted) Encode() ([]byte, error) {
	buf := make([]byte, 30)
	if err := writeAlphaLeft(buf[0:10], m.Session); err != nil {
		return nil, fmt.Errorf("login_accepted session: %w", err)
	}
	if err := writeNumericASCIILeft(buf[10:30], m.SequenceNumber); err != nil {
		return nil, fmt.Errorf("login_accepted seq: %w", err)
	}
	return buf, nil
}

// DecodeLoginAccepted parses a 30-byte payload.
func DecodeLoginAccepted(p []byte) (*LoginAccepted, error) {
	if len(p) < 30 {
		return nil, fmt.Errorf("login_accepted: payload too short (%d < 30)", len(p))
	}
	return &LoginAccepted{
		Session:        strings.TrimLeft(string(p[0:10]), " "),
		SequenceNumber: strings.TrimLeft(string(p[10:30]), " "),
	}, nil
}

func (m *LoginAccepted) Type() PacketType { return PacketTypeLoginAccepted }
