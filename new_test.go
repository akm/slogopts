package slogopts

import (
	"bytes"
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
}
