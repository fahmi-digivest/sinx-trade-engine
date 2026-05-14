package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type compactTextHandler struct {
	w      io.Writer
	opts   slog.HandlerOptions
	attrs  []slog.Attr
	groups []string
	mu     *sync.Mutex
}

func newCompactTextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	var copied slog.HandlerOptions
	if opts != nil {
		copied = *opts
	}

	return &compactTextHandler{
		w:    w,
		opts: copied,
		mu:   &sync.Mutex{},
	}
}

func (h *compactTextHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return level >= minLevel
}

func (h *compactTextHandler) Handle(_ context.Context, r slog.Record) error {
	attrs := make([]slog.Attr, 0, r.NumAttrs()+len(h.attrs)+1)
	attrs = append(attrs, h.attrs...)

	if h.opts.AddSource {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		if f.File != "" {
			attrs = append(attrs, slog.String("source", fmt.Sprintf("%s:%d", f.File, f.Line)))
		}
	}

	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	var b strings.Builder
	if !r.Time.IsZero() {
		b.WriteString(r.Time.Format("2006-01-02 15:04:05.000"))
		b.WriteByte(' ')
	}

	b.WriteString(strings.ToUpper(r.Level.String()))
	if r.Message != "" {
		b.WriteByte(' ')
		b.WriteString(`msg=`)
		b.WriteString(strconv.Quote(r.Message))
	}

	for _, attr := range attrs {
		attr = h.resolveAttr(attr)
		if attr.Equal(slog.Attr{}) {
			continue
		}

		b.WriteByte(' ')
		b.WriteString(h.attrKey(attr.Key))
		b.WriteByte('=')
		b.WriteString(formatAttrValue(attr.Value))
	}

	b.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := io.WriteString(h.w, b.String())
	return err
}

func (h *compactTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	cloned := *h
	cloned.attrs = append(append([]slog.Attr{}, h.attrs...), attrs...)
	return &cloned
}

func (h *compactTextHandler) WithGroup(name string) slog.Handler {
	cloned := *h
	cloned.groups = append(append([]string{}, h.groups...), name)
	return &cloned
}

func (h *compactTextHandler) resolveAttr(attr slog.Attr) slog.Attr {
	attr.Value = attr.Value.Resolve()
	if h.opts.ReplaceAttr != nil {
		attr = h.opts.ReplaceAttr(h.groups, attr)
	}
	return attr
}

func (h *compactTextHandler) attrKey(key string) string {
	if len(h.groups) == 0 {
		return key
	}
	return strings.Join(append(append([]string{}, h.groups...), key), ".")
}

func formatAttrValue(v slog.Value) string {
	switch v.Kind() {
	case slog.KindString:
		return strconv.Quote(v.String())
	case slog.KindTime:
		return strconv.Quote(v.Time().Format("2006-01-02 15:04:05.000"))
	case slog.KindBool:
		return strconv.FormatBool(v.Bool())
	case slog.KindInt64:
		return strconv.FormatInt(v.Int64(), 10)
	case slog.KindUint64:
		return strconv.FormatUint(v.Uint64(), 10)
	case slog.KindFloat64:
		return strconv.FormatFloat(v.Float64(), 'f', -1, 64)
	case slog.KindDuration:
		return strconv.Quote(v.Duration().String())
	case slog.KindAny:
		return strconv.Quote(fmt.Sprint(v.Any()))
	default:
		return strconv.Quote(v.String())
	}
}
