package a

import (
	"context"
	"log/slog"
)

func invalid() {
	logger := slog.Default()

	slog.Info("Uppercase message")                                         // want "message must start with a lowercase letter"
	slog.Info("message привет")                                            // want "message must be in english"
	slog.Info("message!")                                                  // want "message should not contain special symbols or emojis"
	slog.Info("password check failed")                                     // want "message may contain potentially sensitive data"
	slog.Info("Password!")                                                 // want "message must start with a lowercase letter" "message should not contain special symbols or emojis" "message may contain potentially sensitive data"
	slog.Log(context.Background(), slog.LevelInfo, "bearer token missing") // want "message may contain potentially sensitive data"
	logger.Error("message🙂")                                               // want "message should not contain special symbols or emojis"
}
