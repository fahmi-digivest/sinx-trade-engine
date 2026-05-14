package config

// AppConfig holds top-level app metadata and filesystem settings.
type AppConfig struct {
	Name        string `yaml:"name"`
	Desc        string `yaml:"desc"`
	Env         string `yaml:"env"`
	Version     string `yaml:"version"`
	StoragePath string `yaml:"storage_path"`
}
