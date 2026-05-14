package frame

import "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"

// Frame is a decoded SoupBinTCP logical packet.
// It contains the packet type and the raw variable-length payload
// (everything after the 1-byte type field).
type Frame struct {
	Type    message.PacketType
	Payload []byte
}
