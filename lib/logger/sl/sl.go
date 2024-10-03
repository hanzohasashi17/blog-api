package sl

import (
	"log/slog"
	"os"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key: "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Info(value int64) slog.Attr {
	return slog.Attr{
		Key: "info",
		Value: slog.Int64Value(value),
	}
}

func SetupLogger() *slog.Logger {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return log
}