package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func SetupLogger() {
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		},
	}
	logger = slog.New(NewPrettyHandler(os.Stdout, opts))
}

func Logger() *slog.Logger {
	return logger
}
