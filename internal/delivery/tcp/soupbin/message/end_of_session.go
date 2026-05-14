package message

// EndOfSession is sent by the server to signal that the current session is complete.
// The server closes the connection shortly after.
type EndOfSession struct{}

func (m *EndOfSession) Encode() ([]byte, error)          { return []byte{}, nil }
func DecodeEndOfSession(_ []byte) (*EndOfSession, error) { return &EndOfSession{}, nil }
func (m *EndOfSession) Type() PacketType                 { return PacketTypeEndOfSession }
