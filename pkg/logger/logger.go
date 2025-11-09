package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string, keysAndValues ...any)
	Error(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
	Debug(msg string, keysAndValues ...any)
}

type Config struct {
	Level string
}

type logger struct {
	log *slog.Logger
}

func New(config Config) Logger {
	level := slog.LevelInfo
	switch config.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	return &logger{
		log: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})),
	}
}

func (l *logger) Info(msg string, keysAndValues ...any) {
	l.log.Info(msg, keysAndValues...)
}

func (l *logger) Error(msg string, keysAndValues ...any) {
	l.log.Error(msg, keysAndValues...)
}

func (l *logger) Warn(msg string, keysAndValues ...any) {
	l.log.Warn(msg, keysAndValues...)
}

func (l *logger) Debug(msg string, keysAndValues ...any) {
	l.log.Debug(msg, keysAndValues...)
}
