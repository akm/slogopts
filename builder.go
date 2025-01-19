package slogopts

import (
	"io"
	"log/slog"
)

type Builder struct {
	kind    handerKind
	options *slog.HandlerOptions
}

type Option = func(*Builder)

func (b *Builder) build(w io.Writer) slog.Handler {
	switch b.kind {
	case handlerKindTextHandler:
		return slog.NewTextHandler(w, b.options)
	case handlerKindJSONHandler:
		return slog.NewJSONHandler(w, b.options)
	default:
		panic("unknown handler kind")
	}
}

type handerKind int

const (
	handlerKindTextHandler handerKind = 0
	handlerKindJSONHandler handerKind = 1
)

func Text() Option { return func(b *Builder) { b.kind = handlerKindTextHandler } }
func JSON() Option { return func(b *Builder) { b.kind = handlerKindJSONHandler } }
