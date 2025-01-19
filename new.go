package slogopts

import (
	"io"
	"log/slog"
)

func NewHandler(w io.Writer, opts ...Option) slog.Handler {
	b := &Builder{}
	for _, opt := range opts {
		opt(b)
	}
	return b.build(w)
}

func New(w io.Writer, opts ...Option) *slog.Logger {
	return slog.New(NewHandler(w, opts...))
}
