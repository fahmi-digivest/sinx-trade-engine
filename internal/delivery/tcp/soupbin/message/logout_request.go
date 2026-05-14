package message

// LogoutRequest is sent by the client to ask the server to terminate the connection.
type LogoutRequest struct{}

func (m *LogoutRequest) Encode() ([]byte, error)           { return []byte{}, nil }
func DecodeLogoutRequest(_ []byte) (*LogoutRequest, error) { return &LogoutRequest{}, nil }
func (m *LogoutRequest) Type() PacketType                  { return PacketTypeLogoutRequest }
