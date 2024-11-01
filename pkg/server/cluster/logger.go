package cluster

import (
	"github.com/shaj13/raft/raftlog"

	"darvaza.org/slog"
)

var (
	_ raftlog.Logger  = (*Logger)(nil)
	_ raftlog.Verbose = (*Logger)(nil)
)

// Logger implements a bridge between raftlog.Logger and darvaza.org/slog.Logger
type Logger struct {
	// Logger is the destination [slog.Logger] for the [raftlog.Logger] entries.
	Logger slog.Logger
	// Fields are fields that should be added to every entry.
	Fields slog.Fields
	// Verbosity is the maximum depth (V level) that should be logged.
	Verbosity int
	// VerbosityField is the label for the optional field to store
	// the V level.
	VerbosityField string

	v int
}

func (ll *Logger) getLogger(level slog.LogLevel) (slog.Logger, bool) {
	if !ll.Enabled() {
		return nil, false
	}

	l, ok := ll.Logger.WithLevel(level).WithEnabled()
	if !ok {
		return nil, false
	}

	if len(ll.Fields) > 0 {
		l = l.WithFields(ll.Fields)
	}

	if s := ll.VerbosityField; s != "" {
		l = l.WithField(s, ll.v)
	}

	return l, ok
}

func (ll *Logger) levelPrint(level slog.LogLevel, args ...any) {
	if l, ok := ll.getLogger(level); ok {
		l.Print(args...)
	}
}

func (ll *Logger) levelPrintf(level slog.LogLevel, format string, args ...any) {
	if l, ok := ll.getLogger(level); ok {
		l.Printf(format, args...)
	}
}

// Info implements the [raftlog.Logger] interface
func (ll *Logger) Info(args ...any) {
	ll.levelPrint(slog.Info, args...)
}

// Infof implements the [raftlog.Logger] interface
func (ll *Logger) Infof(format string, args ...any) {
	ll.levelPrintf(slog.Info, format, args...)
}

// Warning implements the [raftlog.Logger] interface
func (ll *Logger) Warning(args ...any) {
	ll.levelPrint(slog.Warn, args...)
}

// Warningf implements the [raftlog.Logger] interface
func (ll *Logger) Warningf(format string, args ...any) {
	ll.levelPrintf(slog.Warn, format, args...)
}

// Error implements the [raftlog.Logger] interface
func (ll *Logger) Error(args ...any) {
	ll.levelPrint(slog.Warn, args...)
}

// Errorf implements the [raftlog.Logger] interface
func (ll *Logger) Errorf(format string, args ...any) {
	ll.levelPrintf(slog.Error, format, args...)
}

// Fatal implements the [raftlog.Logger] interface
func (ll *Logger) Fatal(args ...any) {
	ll.levelPrint(slog.Fatal, args...)
}

// Fatalf implements the [raftlog.Logger] interface
func (ll *Logger) Fatalf(format string, args ...any) {
	ll.levelPrintf(slog.Fatal, format, args...)
}

// Panic implements the [raftlog.Logger] interface
func (ll *Logger) Panic(args ...any) {
	ll.levelPrint(slog.Panic, args...)
}

// Panicf implements the [raftlog.Logger] interface
func (ll *Logger) Panicf(format string, args ...any) {
	ll.levelPrintf(slog.Panic, format, args...)
}

// V implements the [raftlog.Logger] interface
func (ll *Logger) V(v int) raftlog.Verbose {
	if ll == nil {
		return &Logger{v: v}
	}

	out := *ll
	out.v = v
	return &out
}

// Enabled implements the [raftlog.Verbose] interface
func (ll *Logger) Enabled() bool {
	if ll == nil || ll.Logger == nil || ll.v > ll.Verbosity {
		return false
	}
	return true
}
