package config

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

const (
	defaultConfigDir = "./config"
	homeConfigDir    = "ouch_mme/config"
)

type Config struct {
	App    AppConfig    `yaml:"app"`
	Logger LoggerConfig `yaml:"logger"`
	TCP    TCPConfig    `yaml:"tcp"`
}

// Load loads config with fallback priority:
//
// 1. $HOME/ouch_mme/config/<filename>
// 2. ./config/<filename>
func Load(fileName string) (*Config, error) {
	paths := buildConfigPaths(fileName)

	var data []byte
	var err error

	for _, path := range paths {
		data, err = os.ReadFile(path)
		if err == nil {
			var cfg Config

			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("config: unmarshal %s: %w", path, err)
			}

			return &cfg, nil
		}

		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("config: read %s: %w", path, err)
		}
	}

	return nil, fmt.Errorf("config: file %s not found", fileName)
}

func buildConfigPaths(fileName string) []string {
	paths := make([]string, 0, 2)

	homeDir, err := os.UserHomeDir()
	if err == nil {
		paths = append(paths,
			filepath.Join(homeDir, homeConfigDir, fileName),
		)
	}

	paths = append(paths,
		filepath.Join(defaultConfigDir, fileName),
	)

	return paths
}
