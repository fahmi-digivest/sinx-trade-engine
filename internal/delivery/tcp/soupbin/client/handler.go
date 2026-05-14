package client

import "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"

// Handler receives decoded SoupBinTCP messages from the server.
// Implementations must be safe for concurrent use.
type Handler interface {
	// OnLoginAccepted is called once the server accepts the login.
	OnLoginAccepted(session string, nextSeq uint64)

	// OnSequencedData is called for each Sequenced Data packet.
	// seq is the locally tracked sequence number (1-based).
	OnSequencedData(seq uint64, msg []byte)

	// OnEndOfSession is called when the server sends an End of Session packet.
	OnEndOfSession()

	// OnError is called when a non-recoverable error occurs.
	OnError(err error)
}

// UnsequencedHandler is an optional interface a Handler may also implement
// to receive Unsequenced messages echoed back from the server.
type UnsequencedHandler interface {
	OnUnsequencedData(msg *message.UnsequencedData)
}
