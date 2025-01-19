package slogopts

import (
	"log/slog"
	"testing"
)

func TestBuilderBuild(t *testing.T) {
	t.Run("Text", func(t *testing.T) {
		b := newBuilder()
		b.kind = handlerKindTextHandler
		h := b.build(nil)
		if _, ok := h.(*slog.TextHandler); !ok {
			t.Error("expected TextHandler")
		}
	})
	t.Run("JSON", func(t *testing.T) {
		b := newBuilder()
		b.kind = handlerKindJSONHandler
		h := b.build(nil)
		if _, ok := h.(*slog.JSONHandler); !ok {
			t.Error("expected JSONHandler")
		}
	})
	t.Run("Unknown", func(t *testing.T) {
		b := newBuilder()
		b.kind = 2
		defer func() {
			if recover() == nil {
				t.Error("expected panic")
			}
		}()
		b.build(nil)
	})
}
