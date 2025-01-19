package slogopts

import (
	"io"
	"log/slog"
)

type builder struct {
	kind    handerKind
	options *slog.HandlerOptions
}

func newBuilder() *builder {
	return &builder{
		kind:    handlerKindTextHandler,
		options: &slog.HandlerOptions{},
	}
}

type Option = func(*builder)

func (b *builder) build(w io.Writer) slog.Handler {
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

func Text() Option { return func(b *builder) { b.kind = handlerKindTextHandler } }
func JSON() Option { return func(b *builder) { b.kind = handlerKindJSONHandler } }
