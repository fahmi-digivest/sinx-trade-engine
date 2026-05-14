package codec

import (
	"fmt"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
)

// Decoder converts a frame.Frame into the appropriate message.Message concrete type.
type Decoder struct{}

// NewDecoder returns a new Decoder.
func NewDecoder() *Decoder { return &Decoder{} }

// Decode inspects the frame's Type field and delegates to the appropriate
// message-level decoder.  Returns an error for unknown packet types.
func (d *Decoder) Decode(f *frame.Frame) (message.Message, error) {
	switch f.Type {
	// Server -> Client
	case message.PacketTypeLoginAccepted:
		return message.DecodeLoginAccepted(f.Payload)
	case message.PacketTypeLoginRejected:
		return message.DecodeLoginRejected(f.Payload)
	case message.PacketTypeSequencedData:
		return message.DecodeSequencedData(f.Payload)
	case message.PacketTypeServerHB:
		return message.DecodeServerHeartbeat(f.Payload)
	case message.PacketTypeEndOfSession:
		return message.DecodeEndOfSession(f.Payload)

	// Client -> Server
	case message.PacketTypeLoginRequest:
		return message.DecodeLoginRequest(f.Payload)
	case message.PacketTypeUnsequencedData:
		return message.DecodeUnsequencedData(f.Payload)
	case message.PacketTypeClientHB:
		return message.DecodeClientHeartbeat(f.Payload)
	case message.PacketTypeLogoutRequest:
		return message.DecodeLogoutRequest(f.Payload)

	// Debug (both directions)
	case message.PacketTypeDebug:
		// Return raw payload wrapped in a generic debug message.
		return &debugMessage{payload: f.Payload}, nil

	default:
		return nil, fmt.Errorf("codec/decoder: unknown packet type %q (0x%02X)", f.Type, byte(f.Type))
	}
}

// debugMessage is an internal type for debug packets.
type debugMessage struct {
	payload []byte
}

func (m *debugMessage) Type() message.PacketType { return message.PacketTypeDebug }
func (m *debugMessage) Encode() ([]byte, error)  { return m.payload, nil }
func (m *debugMessage) Text() string             { return string(m.payload) }
