package config

import (
	"fmt"
	"time"
)

// SoupBinConfig holds all configuration parameters for a SoupBinTCP connection.
type SoupBinConfig struct {
	// Server address in "host:port" format.
	Addr string

	// Credentials (case-insensitive, right-padded with spaces to 6/10 bytes).
	Username string // max 6 chars
	Password string // max 10 chars

	// Session to connect to. Leave empty ("") to connect to the currently
	// active session (the server will left-pad with spaces to 10 bytes).
	Session string // max 10 chars

	// Next expected sequence number (ASCII numeric, 1-based).
	// Set to 0 to start from the most recently generated message.
	RequestedSeqNum uint64

	// HeartbeatInterval is how often a heartbeat is sent when no other
	// data has been sent. Must be ≤ 1 second per spec; defaults to 1s.
	HeartbeatInterval time.Duration

	// ReadTimeout is the maximum time the client waits for any data
	// (including heartbeats) before declaring the link dead and reconnecting.
	// Spec recommends several seconds; defaults to 15s.
	ReadTimeout time.Duration

	// LoginTimeout is the maximum time the server waits for a Login Request
	// after a new TCP connection. Spec recommends 30s.
	LoginTimeout time.Duration

	// MaxReconnectAttempts: 0 means retry forever.
	MaxReconnectAttempts int

	// ReconnectBackoff is the wait between reconnect attempts.
	ReconnectBackoff time.Duration

	// DialTimeout is the TCP connect timeout per attempt.
	DialTimeout time.Duration
}

func DefaultConfig(addr, username, password string) SoupBinConfig {
	return SoupBinConfig{
		Addr:                 addr,
		Username:             username,
		Password:             password,
		RequestedSeqNum:      1,
		HeartbeatInterval:    time.Second,
		ReadTimeout:          15 * time.Second,
		LoginTimeout:         30 * time.Second,
		MaxReconnectAttempts: 0, // infinite
		ReconnectBackoff:     2 * time.Second,
		DialTimeout:          5 * time.Second,
	}
}

// Validate checks that the config contains the minimum required fields and
// that field lengths are within protocol limits.
func (c *SoupBinConfig) Validate() error {
	if c.Addr == "" {
		return fmt.Errorf("soupbin: Addr is required")
	}
	if len(c.Username) > 6 {
		return fmt.Errorf("soupbin: Username exceeds 6 characters")
	}
	if len(c.Password) > 10 {
		return fmt.Errorf("soupbin: Password exceeds 10 characters")
	}
	if len(c.Session) > 10 {
		return fmt.Errorf("soupbin: Session exceeds 10 characters")
	}
	if c.HeartbeatInterval <= 0 || c.HeartbeatInterval > time.Second {
		c.HeartbeatInterval = time.Second
	}
	if c.ReadTimeout <= 0 {
		c.ReadTimeout = 15 * time.Second
	}
	if c.LoginTimeout <= 0 {
		c.LoginTimeout = 30 * time.Second
	}
	if c.DialTimeout <= 0 {
		c.DialTimeout = 5 * time.Second
	}
	return nil
}
