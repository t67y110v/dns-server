package logger

import (
	"log/slog"
	"os"
)

type Level int

type Leveler interface {
	Level() Level
}

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

func Get(level Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.Level(level)}))

}
