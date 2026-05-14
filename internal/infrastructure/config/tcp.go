package config

import "time"

// TCPConfig holds TCP client connection definitions.
type TCPConfig struct {
	Clients []TCPClientConfig `yaml:"clients"`
}

type TCPClientConfig struct {
	Name              string        `yaml:"name"`
	Run               bool          `yaml:"run"`
	Host              string        `yaml:"host"`
	Port              int           `yaml:"port"`
	Username          string        `yaml:"username"`
	Password          string        `yaml:"password"`
	RequestedSession  string        `yaml:"requested_session"`
	RequestedSequence string        `yaml:"requested_sequence"`
	HeartbeatInterval time.Duration `yaml:"heartbeat_interval"`
	DialTimeout       time.Duration `yaml:"dial_timeout"`
	ReconnectDelay    time.Duration `yaml:"reconnect_delay"`
	KeepAlive         time.Duration `yaml:"keep_alive"`
}
