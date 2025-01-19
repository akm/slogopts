package slogopts

import (
	"io"
	"log/slog"
)

// NewHandler returns a new slog.Handler.
// The handler is created with the given io.Writer and options.
// The [Option] s are applied in the order they are given.
func NewHandler(w io.Writer, opts ...Option) slog.Handler {
	b := newBuilder()
	for _, opt := range opts {
		opt(b)
	}
	return b.build(w)
}

// New returns a new slog.Logger.
// The logger is created with the given io.Writer and options.
// The [Option] s are applied in the order they are given.
func New(w io.Writer, opts ...Option) *slog.Logger {
	return slog.New(NewHandler(w, opts...))
}
