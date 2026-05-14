package client

import "time"

// Config holds configuration for a SoupBinTCP client.
type Config struct {
	// ServerAddr is the upstream SoupBinTCP server address, e.g. "127.0.0.1:9000".
	ServerAddr string

	// Username and Password for authentication (case-insensitive, per spec).
	Username string
	Password string

	// RequestedSession is the session to log into.  Leave blank for the currently
	// active session.
	RequestedSession string

	// RequestedSequenceNumber is the first sequence number the client wants to
	// receive.  "0" means start from the most recent message.
	RequestedSequenceNumber string

	// HeartbeatInterval is how often to send a Client Heartbeat when idle.
	// Defaults to 1 second per the spec.
	HeartbeatInterval time.Duration

	// ServerTimeout is how long to wait without receiving anything before
	// treating the connection as dead.
	ServerTimeout time.Duration

	// DialTimeout is the timeout for the initial TCP connection.
	DialTimeout time.Duration
}

// DefaultConfig returns a Config with spec-recommended defaults.
func DefaultConfig(addr, username, password string) Config {
	return Config{
		ServerAddr:              addr,
		Username:                username,
		Password:                password,
		RequestedSequenceNumber: "1",
		HeartbeatInterval:       time.Second,
		ServerTimeout:           15 * time.Second,
		DialTimeout:             10 * time.Second,
	}
}
