package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

// Config is the root application configuration loaded from YAML.
type Config struct {
	App    AppConfig    `yaml:"app"`
	Logger LoggerConfig `yaml:"logger"`
	TCP    TCPConfig    `yaml:"tcp"`
}

// Load reads and decodes a YAML config file from disk.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config: read %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("config: unmarshal %s: %w", path, err)
	}

	return &cfg, nil
}
