package main

import (
	"log/slog"
)

// TODO: Implement newLogger(service, env string) *slog.Logger
// - Use slog.NewJSONHandler writing to os.Stdout
// - Set level to slog.LevelInfo
// - Attach base fields: "service" and "env"

func newLogger(service, env string) *slog.Logger {
	// TODO: implement
	return slog.Default()
}

func main() {
	logger := newLogger("token-service", "development")

	// TODO: Log service startup at Info level with a "version" field

	// TODO: Log a simulated incoming request at Info level
	// with fields: method, path, request_id

	// TODO: Log a simulated error at Error level
	// using errors.New("connection timeout") as the error field

	_ = logger
}
