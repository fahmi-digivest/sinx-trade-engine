package message

// ServerHeartbeat is sent by the server when no data has been sent for > 1 second.
type ServerHeartbeat struct{}

func (m *ServerHeartbeat) Encode() ([]byte, error)             { return []byte{}, nil }
func DecodeServerHeartbeat(_ []byte) (*ServerHeartbeat, error) { return &ServerHeartbeat{}, nil }
func (m *ServerHeartbeat) Type() PacketType                    { return PacketTypeServerHB }

// ClientHeartbeat is sent by the client when no data has been sent for > 1 second.
type ClientHeartbeat struct{}

func (m *ClientHeartbeat) Encode() ([]byte, error)             { return []byte{}, nil }
func DecodeClientHeartbeat(_ []byte) (*ClientHeartbeat, error) { return &ClientHeartbeat{}, nil }
func (m *ClientHeartbeat) Type() PacketType                    { return PacketTypeClientHB }
