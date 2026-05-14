package frame

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
)

// Reader reads SoupBinTCP frames from an io.Reader (typically a net.Conn).
//
// Wire format per frame:
//
//	[0:2]  big-endian uint16  – length of everything after these two bytes
//	[2]    byte               – packet type
//	[3:]   []byte             – payload (length - 1 bytes)
type Reader struct {
	r io.Reader
}

// NewReader creates a Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// ReadFrame reads exactly one SoupBinTCP frame from the underlying reader.
// It blocks until the complete frame is available or an error occurs.
func (fr *Reader) ReadFrame() (*Frame, error) {
	// Read the 2-byte length prefix.
	var lenBuf [2]byte
	if _, err := io.ReadFull(fr.r, lenBuf[:]); err != nil {
		return nil, fmt.Errorf("frame: reading length: %w", err)
	}
	frameLen := binary.BigEndian.Uint16(lenBuf[:])
	if frameLen < 1 {
		return nil, fmt.Errorf("frame: length %d is too short (must be >= 1 for type byte)", frameLen)
	}

	// Read type byte + payload.
	body := make([]byte, frameLen)
	if _, err := io.ReadFull(fr.r, body); err != nil {
		return nil, fmt.Errorf("frame: reading body: %w", err)
	}

	return &Frame{
		Type:    message.PacketType(body[0]),
		Payload: body[1:],
	}, nil
}
