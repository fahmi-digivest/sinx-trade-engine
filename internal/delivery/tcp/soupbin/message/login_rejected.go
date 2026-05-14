package message

import "fmt"

// RejectReason is the single-byte reason code in a LoginRejected packet.
type RejectReason byte

const (
	RejectReasonNotAuthorized      RejectReason = 'A' // Invalid username/password
	RejectReasonSessionUnavailable RejectReason = 'S' // Session invalid or unavailable
)

// LoginRejected is sent by the server when the login request is invalid.
// The server closes the socket after sending this packet.
type LoginRejected struct {
	RejectReasonCode RejectReason
}

// Encode serialises LoginRejected into its 1-byte wire payload.
func (m *LoginRejected) Encode() ([]byte, error) {
	return []byte{byte(m.RejectReasonCode)}, nil
}

// DecodeLoginRejected parses a 1-byte payload.
func DecodeLoginRejected(p []byte) (*LoginRejected, error) {
	if len(p) < 1 {
		return nil, fmt.Errorf("login_rejected: payload too short")
	}
	return &LoginRejected{RejectReasonCode: RejectReason(p[0])}, nil
}

func (m *LoginRejected) Type() PacketType { return PacketTypeLoginRejected }

// Error implements the error interface for convenience.
func (m *LoginRejected) Error() string {
	switch m.RejectReasonCode {
	case RejectReasonNotAuthorized:
		return "login rejected: not authorized (invalid username/password)"
	case RejectReasonSessionUnavailable:
		return "login rejected: session not available"
	default:
		return fmt.Sprintf("login rejected: unknown reason %q", m.RejectReasonCode)
	}
}
