package slogopts

import (
	"bytes"
	"encoding/json"
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
	t.Run("Text", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(buf, Text())
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
	t.Run("JSON", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(buf, JSON())
		logger.Info("hello")
		var d map[string]interface{}
		assert.NoError(t, json.Unmarshal(buf.Bytes(), &d))
		assert.Equal(t, "hello", d["msg"])
		assert.Equal(t, "INFO", d["level"])
	})
	t.Run("Replace time and level", func(t *testing.T) {
		replTime := func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{Key: "t", Value: a.Value}
			}
			return a
		}
		replLevel := func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				return slog.Attr{Key: "lv", Value: a.Value}
			}
			return a
		}
		buf := bytes.NewBuffer(nil)
		logger := New(buf, JSON(), ReplaceAttr(MergeReplaceAttr(replTime, replLevel)))
		logger.Info("hello")
		var d map[string]interface{}
		assert.NoError(t, json.Unmarshal(buf.Bytes(), &d))
		assert.Equal(t, "hello", d["msg"])
		assert.Empty(t, d["time"])
		assert.Empty(t, d["level"])
		assert.NotEmpty(t, d["t"])
		assert.NotEmpty(t, d["lv"])
	})
	t.Run("AddSource", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		logger := New(buf, JSON(), AddSource(true))
		logger.Info("hello")
		var d map[string]interface{}
		assert.NoError(t, json.Unmarshal(buf.Bytes(), &d))
		assert.Equal(t, "hello", d["msg"])
		if assert.IsType(t, map[string]interface{}{}, d["source"]) {
			m := d["source"].(map[string]interface{})
			assert.Contains(t, m["file"], "new_test.go")
			assert.Contains(t, m["function"], "TestNew")
			assert.IsType(t, float64(0), m["line"])
		}
	})
}
