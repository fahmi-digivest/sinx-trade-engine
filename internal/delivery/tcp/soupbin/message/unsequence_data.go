package message

// UnsequencedData is the client-to-server envelope for higher-level protocol messages.
// These are not sequenced and may be lost during TCP/IP failures.
type UnsequencedData struct {
	Message []byte
}

// Encode returns the raw message bytes as the payload.
func (m *UnsequencedData) Encode() ([]byte, error) {
	out := make([]byte, len(m.Message))
	copy(out, m.Message)
	return out, nil
}

// DecodeUnsequencedData wraps the raw payload bytes.
func DecodeUnsequencedData(p []byte) (*UnsequencedData, error) {
	msg := make([]byte, len(p))
	copy(msg, p)
	return &UnsequencedData{Message: msg}, nil
}

func (m *UnsequencedData) Type() PacketType { return PacketTypeUnsequencedData }
