package server

import (
	"sync"
	"sync/atomic"
)

// Session holds the per-session sequencing state managed by the server.
type Session struct {
	ID string

	mu       sync.RWMutex
	messages [][]byte // 0-indexed; messages[i] has implicit sequence number i+1
	nextSeq  uint64   // next sequence number to assign (1-based)
}

// NewSession creates a new empty session with the given ID.
func NewSession(id string) *Session {
	return &Session{
		ID:      id,
		nextSeq: 1,
	}
}

// Append adds a new message to the session and returns its sequence number.
func (s *Session) Append(msg []byte) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	seq := s.nextSeq
	cp := make([]byte, len(msg))
	copy(cp, msg)
	s.messages = append(s.messages, cp)
	atomic.AddUint64(&s.nextSeq, 1)
	return seq
}

// Get retrieves the message at the given 1-based sequence number.
// Returns nil if seq is out of range.
func (s *Session) Get(seq uint64) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	idx := int(seq) - 1
	if idx < 0 || idx >= len(s.messages) {
		return nil
	}
	return s.messages[idx]
}

// NextSeq returns the sequence number that will be assigned to the next message.
func (s *Session) NextSeq() uint64 {
	return atomic.LoadUint64(&s.nextSeq)
}

// Len returns the total number of messages stored in this session.
func (s *Session) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.messages)
}
