package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	defaultTagName tagName = "default"
)

type tagName string

type fields struct {
	f map[string]string
}

type contextHandler struct {
	slog.Handler
}

func newContextHandler(h slog.Handler) contextHandler {
	return contextHandler{h}
}

func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	t, ok := ctx.Value(defaultTagName).(fields)
	if !ok {
		return h.Handler.Handle(ctx, r) // Если полей нет, просто продолжаем без них
	}

	attrs := make([]slog.Attr, 0, len(t.f)) // Инициализируем срез с capacity, но длиной 0
	for name, value := range t.f {
		attrs = append(attrs, slog.String(name, value))
	}
	r.AddAttrs(attrs...)

	return h.Handler.Handle(ctx, r)
}

type Logger struct {
	*slog.Logger
}

func New() *Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	contextHandler := newContextHandler(jsonHandler)
	return &Logger{slog.New(&contextHandler)}
}

func buildFields() fields {
	m := make(map[string]string, 0)
	return fields{f: m}
}

func (l *Logger) WithFields(ctx context.Context, f map[string]string) context.Context {
	t, ok := ctx.Value(defaultTagName).(fields)
	if !ok {
		t = buildFields() // Новый словарь, если полей еще нет
	} else {
		// Создаем копию существующего словаря
		newMap := make(map[string]string, len(t.f)+len(f))
		for k, v := range t.f {
			newMap[k] = v
		}
		t = fields{f: newMap}
	}

	// Добавляем новые поля
	for field, val := range f {
		t.f[field] = val
	}

	return context.WithValue(ctx, defaultTagName, t)
}

func (l *Logger) Info(ctx context.Context, message string, args ...any) {
	l.InfoContext(ctx, message, args...)
}

func (l *Logger) Error(ctx context.Context, err error, args ...any) {
	l.ErrorContext(ctx, err.Error(), args...)
}
