package codec

import (
	"fmt"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
)

// Encoder converts a message.Message into a frame.Frame ready to be written
// to the wire by a frame.Writer.
type Encoder struct{}

// NewEncoder returns a new Encoder.
func NewEncoder() *Encoder { return &Encoder{} }

// Encode serialises msg into a Frame.
func (e *Encoder) Encode(msg message.Message) (*frame.Frame, error) {
	payload, err := msg.Encode()
	if err != nil {
		return nil, fmt.Errorf("codec/encoder: %w", err)
	}
	return &frame.Frame{
		Type:    msg.Type(),
		Payload: payload,
	}, nil
}
