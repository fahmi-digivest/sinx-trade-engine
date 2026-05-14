package message

// PacketType represents the single-byte SoupBinTCP packet type identifier.
type PacketType byte

const (
	// Server -> Client
	PacketTypeLoginAccepted PacketType = 'A'
	PacketTypeLoginRejected PacketType = 'J'
	PacketTypeSequencedData PacketType = 'S'
	PacketTypeServerHB      PacketType = 'H'
	PacketTypeEndOfSession  PacketType = 'Z'

	// Client -> Server
	PacketTypeLoginRequest    PacketType = 'L'
	PacketTypeUnsequencedData PacketType = 'U'
	PacketTypeClientHB        PacketType = 'R'
	PacketTypeLogoutRequest   PacketType = 'O'

	// Both directions
	PacketTypeDebug PacketType = '+'
)
