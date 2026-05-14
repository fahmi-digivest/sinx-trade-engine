package message

import (
	"fmt"
	"strings"
)

// LoginRequest is sent by the client immediately after opening a TCP connection.
// Username and Password are case-insensitive, right-padded with spaces.
type LoginRequest struct {
	Username                string // 6 bytes, right-padded
	Password                string // 10 bytes, right-padded
	RequestedSession        string // 10 bytes, blank = currently active session
	RequestedSequenceNumber string // 20 bytes ASCII numeric, "0" = most recent message
}

// Encode serialises the LoginRequest into its 46-byte wire payload.
func (m *LoginRequest) Encode() ([]byte, error) {
	buf := make([]byte, 46)
	if err := writeAlpha(buf[0:6], m.Username); err != nil {
		return nil, fmt.Errorf("login_request username: %w", err)
	}
	if err := writeAlpha(buf[6:16], m.Password); err != nil {
		return nil, fmt.Errorf("login_request password: %w", err)
	}
	if err := writeAlpha(buf[16:26], m.RequestedSession); err != nil {
		return nil, fmt.Errorf("login_request session: %w", err)
	}
	if err := writeNumericASCII(buf[26:46], m.RequestedSequenceNumber); err != nil {
		return nil, fmt.Errorf("login_request seq: %w", err)
	}
	return buf, nil
}

// DecodeLoginRequest parses a 46-byte payload into a LoginRequest.
func DecodeLoginRequest(p []byte) (*LoginRequest, error) {
	if len(p) < 46 {
		return nil, fmt.Errorf("login_request: payload too short (%d < 46)", len(p))
	}
	return &LoginRequest{
		Username:                strings.TrimRight(string(p[0:6]), " "),
		Password:                strings.TrimRight(string(p[6:16]), " "),
		RequestedSession:        strings.TrimRight(string(p[16:26]), " "),
		RequestedSequenceNumber: strings.TrimLeft(string(p[26:46]), " "),
	}, nil
}

func (m *LoginRequest) Type() PacketType { return PacketTypeLoginRequest }
