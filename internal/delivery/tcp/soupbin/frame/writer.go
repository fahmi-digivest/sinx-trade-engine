package frame

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Writer writes SoupBinTCP frames to an io.Writer (typically a net.Conn).
type Writer struct {
	w io.Writer
}

// NewWriter creates a Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// WriteFrame encodes and writes a single SoupBinTCP frame.
//
// Wire format:
//
//	[0:2]  big-endian uint16  – len(payload) + 1  (the +1 accounts for the type byte)
//	[2]    byte               – packet type
//	[3:]   []byte             – payload
func (fw *Writer) WriteFrame(f *Frame) error {
	payloadLen := len(f.Payload)
	frameLen := payloadLen + 1 // +1 for the type byte

	if frameLen > 0xFFFF {
		return fmt.Errorf("frame: payload too large (%d bytes)", payloadLen)
	}

	// Build the complete wire buffer in one allocation to allow a single Write call,
	// which is important for atomicity on a shared net.Conn.
	buf := make([]byte, 2+frameLen)
	binary.BigEndian.PutUint16(buf[0:2], uint16(frameLen))
	buf[2] = byte(f.Type)
	copy(buf[3:], f.Payload)

	_, err := io.WriteString(fw.w, "") // noop – just ensuring interface is satisfied
	_, err = fw.w.Write(buf)
	if err != nil {
		return fmt.Errorf("frame: write: %w", err)
	}
	return nil
}
