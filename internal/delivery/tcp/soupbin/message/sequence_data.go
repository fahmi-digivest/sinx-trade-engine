package message

// SequencedData is the server-to-client envelope for higher-level protocol messages.
// The sequence number is implicit (tracked by both sides).
type SequencedData struct {
	Message []byte
}

// Encode returns the raw message bytes as the payload.
func (m *SequencedData) Encode() ([]byte, error) {
	out := make([]byte, len(m.Message))
	copy(out, m.Message)
	return out, nil
}

// DecodeSequencedData wraps the raw payload bytes.
func DecodeSequencedData(p []byte) (*SequencedData, error) {
	msg := make([]byte, len(p))
	copy(msg, p)
	return &SequencedData{Message: msg}, nil
}

func (m *SequencedData) Type() PacketType { return PacketTypeSequencedData }
