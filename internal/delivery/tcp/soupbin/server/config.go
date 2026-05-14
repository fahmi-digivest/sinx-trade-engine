package server

import "time"

// Config holds configuration for a SoupBinTCP server.
type Config struct {
	// ListenAddr is the TCP address to listen on, e.g. ":9000".
	ListenAddr string

	// HeartbeatInterval is how often the server sends a heartbeat when idle.
	// Defaults to 1 second per the spec.
	HeartbeatInterval time.Duration

	// ClientTimeout is how long the server waits before closing a connection
	// that has not sent anything.  Spec recommendation is 15 seconds.
	ClientTimeout time.Duration

	// LoginTimeout is how long to wait for a Login Request after a new TCP
	// connection is established.  Spec recommendation is 30 seconds.
	LoginTimeout time.Duration

	// Credentials holds the valid username/password pairs.
	// Keys are lower-cased usernames; values are lower-cased passwords.
	Credentials map[string]string
}

// DefaultConfig returns a Config with spec-recommended defaults.
func DefaultConfig(addr string) Config {
	return Config{
		ListenAddr:        addr,
		HeartbeatInterval: time.Second,
		ClientTimeout:     15 * time.Second,
		LoginTimeout:      30 * time.Second,
		Credentials:       make(map[string]string),
	}
}
