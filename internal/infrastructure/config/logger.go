package config

// LoggerConfig describes logging defaults and named log streams.
type LoggerConfig struct {
	Default LoggerDefaultConfig          `yaml:"default"`
	Logs    map[string]LoggerNamedConfig `yaml:"logs"`
}

type LoggerDefaultConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
	Format string `yaml:"format"`
	Dir    string `yaml:"dir"`
}

type LoggerNamedConfig struct {
	Enabled  bool     `yaml:"enabled"`
	Filename string   `yaml:"filename"`
	Output   []string `yaml:"output"`
	Level    string   `yaml:"level"`
}
