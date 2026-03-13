package a

import (
	"context"
	"log/slog"
)

func valid() {
	logger := slog.Default()

	slog.Info("user signed in")
	slog.Log(context.Background(), slog.LevelInfo, "request completed")
	logger.Error("connection reset")
}
