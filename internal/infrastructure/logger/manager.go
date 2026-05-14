package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/config"
)

// Manager holds all named loggers built from the YAML config.
type Manager struct {
	loggers map[string]*slog.Logger
	def     config.LoggerDefaultConfig
}

// NewManager builds a Manager from the parsed Config.
func NewManager(cfg config.LoggerConfig) (*Manager, error) {
	m := &Manager{
		loggers: make(map[string]*slog.Logger),
		def:     cfg.Default,
	}

	for name, lc := range cfg.Logs {
		if !lc.Enabled {
			continue
		}
		l, err := m.build(name, lc)
		if err != nil {
			return nil, fmt.Errorf("logger %q: %w", name, err)
		}
		m.loggers[name] = l
	}

	return m, nil
}

// Get returns the named logger. Returns slog.Default() if not found.
func (m *Manager) Get(name string) *slog.Logger {
	if l, ok := m.loggers[name]; ok {
		return l
	}
	return slog.Default()
}

// ---------------------------------------------------------------------------
// internal helpers
// ---------------------------------------------------------------------------

func (m *Manager) build(name string, lc config.LoggerNamedConfig) (*slog.Logger, error) {
	level, err := parseLevel(lc.Level, m.def.Level)
	if err != nil {
		return nil, err
	}

	writers := make([]io.Writer, 0, len(lc.Output))
	for _, out := range lc.Output {
		switch strings.ToLower(out) {
		case "stdout":
			writers = append(writers, os.Stdout)
		case "file":
			f, err := m.openFile(lc.Filename)
			if err != nil {
				return nil, err
			}
			writers = append(writers, f)
		default:
			return nil, fmt.Errorf("unknown output %q (valid: stdout, file)", out)
		}
	}

	if len(writers) == 0 {
		writers = append(writers, io.Discard)
	}

	w := io.MultiWriter(writers...)

	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if strings.ToLower(m.def.Format) == "json" {
		handler = slog.NewJSONHandler(w, opts)
	} else {
		handler = newCompactTextHandler(w, opts)
	}

	// Attach logger name as a fixed attribute.
	return slog.New(handler).With("logger", name), nil
}

// openFile creates or appends to a daily-rotated file: <filename>_YYYY-MM-DD.log
func (m *Manager) openFile(filename string) (*os.File, error) {
	if err := os.MkdirAll(m.def.Dir, 0o755); err != nil {
		return nil, fmt.Errorf("create log dir %q: %w", m.def.Dir, err)
	}
	dated := fmt.Sprintf("%s_%s.log", filename, time.Now().Format("2006-01-02"))
	path := filepath.Join(m.def.Dir, dated)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open log file %q: %w", path, err)
	}
	return f, nil
}

// parseLevel converts a string to slog.Level, falling back to fallback if empty.
func parseLevel(level, fallback string) (slog.Level, error) {
	if level == "" {
		level = fallback
	}
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid log level %q", level)
	}
}
