// Package zerolog provides a dumb wrapper for logging to os.Stderr
// using github.com/rs/zerolog
package zerolog

import (
	"darvaza.org/slog"
	"darvaza.org/slog/handlers/filter"
	"darvaza.org/slog/handlers/zerolog"
)

// NewLogger wraps the zerolog.Logger for slog with a given filter level
func NewLogger(level slog.LogLevel) slog.Logger {
	return filter.New(zerolog.New(&zlog), level)
}
