package slogopts

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(buf)
		logger.Info("hello")
		assert.Contains(t, buf.String(), "hello")
		assert.Contains(t, buf.String(), "INFO")
	})
	t.Run("LogLevel", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(buf, Level(slog.LevelInfo))
		logger.Info("hello")
		assert.Contains(t, buf.String(), "hello")
		assert.Contains(t, buf.String(), "INFO")
		buf.Reset()
		logger.Debug("hello")
		assert.Empty(t, buf.Bytes())
	})
}
