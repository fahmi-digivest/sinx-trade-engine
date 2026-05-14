package server

import (
	"fmt"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/codec"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
)

// Dispatcher sends sequenced and control packets to a single downstream client.
type Dispatcher struct {
	enc *codec.Encoder
	fw  *frame.Writer
}

// NewDispatcher creates a Dispatcher that writes to fw.
func NewDispatcher(fw *frame.Writer) *Dispatcher {
	return &Dispatcher{
		enc: codec.NewEncoder(),
		fw:  fw,
	}
}

// SendLoginAccepted sends a Login Accepted packet.
func (d *Dispatcher) SendLoginAccepted(session, seqNum string) error {
	return d.send(&message.LoginAccepted{Session: session, SequenceNumber: seqNum})
}

// SendLoginRejected sends a Login Rejected packet.
func (d *Dispatcher) SendLoginRejected(reason message.RejectReason) error {
	return d.send(&message.LoginRejected{RejectReasonCode: reason})
}

// SendSequencedData sends a Sequenced Data packet containing msg.
func (d *Dispatcher) SendSequencedData(msg []byte) error {
	return d.send(&message.SequencedData{Message: msg})
}

// SendHeartbeat sends a Server Heartbeat packet.
func (d *Dispatcher) SendHeartbeat() error {
	return d.send(&message.ServerHeartbeat{})
}

// SendEndOfSession sends an End of Session packet.
func (d *Dispatcher) SendEndOfSession() error {
	return d.send(&message.EndOfSession{})
}

func (d *Dispatcher) send(msg message.Message) error {
	f, err := d.enc.Encode(msg)
	if err != nil {
		return fmt.Errorf("dispatcher: encode %T: %w", msg, err)
	}
	return d.fw.WriteFrame(f)
}
